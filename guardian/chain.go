package guardian

import (
	"context"
	"sync"
	"time"

	"github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-guardian/contracts"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

type OneChainConfig struct {
	ChainID                                             uint64
	Name, Gateway                                       string
	BlkInterval, BlkDelay, MaxBlkDelta, ForwardBlkDelay uint64
	GasLimit                                            uint64
	AddGasEstimateRatio                                 float64
	// Legacy gas price flags
	AddGasGwei, MinGasGwei, MaxGasGwei uint64
	ForceGasGwei                       string
	// EIP-1559 gas price flags
	MaxFeePerGasGwei, MaxPriorityFeePerGasGwei uint64

	MsgBus    string   // MessageBus contract address
	Receivers []string // Message receivers, i.e. dApp addresses
}

type Chain struct {
	*OneChainConfig
	MessageBus *contracts.MsgBusContract
	EthClient  *ethclient.Client
	monitor2   *mon2.Monitor
	pauser     *eth.Transactor
}

type Chains struct {
	// chainid => Chain
	chains map[uint64]*Chain
	lock   sync.RWMutex
	initWg sync.WaitGroup
}

var chains *Chains

func InitChains() *Chains {
	log.Infoln("Initializing chains")
	var configs []*OneChainConfig
	err := viper.UnmarshalKey(FlagMultiChain, &configs)
	if err != nil {
		log.Fatalf("failed to load multichain configs:%v", err)
		return nil
	}
	loadPauserConfig()
	cs := &Chains{chains: make(map[uint64]*Chain)}
	for _, config := range configs {
		cs.initWg.Add(1)
		go cs.addChain(config)
	}
	cs.initWg.Wait()
	log.Infoln("Finished initializing all chains")
	return cs

}

func (c *Chains) addChain(config *OneChainConfig) {
	ec := newEthClient(config)

	// init monitor
	chainConfig := mon2.PerChainCfg{
		BlkIntv:         time.Duration(config.BlkInterval) * time.Second,
		BlkDelay:        config.BlkDelay,
		MaxBlkDelta:     config.MaxBlkDelta,
		ForwardBlkDelay: config.ForwardBlkDelay,
	}
	dal := &MonitorDAL{kv: make(map[string]mon2.LogEventID)}
	mon, err := mon2.NewMonitor(ec, dal, chainConfig)
	if err != nil {
		log.Fatalln("failed to create monitor:", err)
	}

	chain := &Chain{
		OneChainConfig: config,
		EthClient:      ec,
		monitor2:       mon,
	}
	chain.MessageBus, err = contracts.NewMsgBusContract(Hex2Addr(config.MsgBus), ec)
	if err != nil {
		log.Fatalln("failed to get messageBus contract:", err)
	}
	chain.pauser, err = newPauserTransactor(chain)
	if err != nil {
		log.Fatalln("failed to get create pauser transactor:", err)
	}
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		c.initWg.Done()
	}()
	c.chains[chain.ChainID] = chain
}

func newEthClient(config *OneChainConfig) *ethclient.Client {
	// init eth client
	log.Infof("Dialing eth client for chain %d at %s", config.ChainID, config.Gateway)
	var ec *ethclient.Client
	var err error

	ec, err = ethclient.Dial(config.Gateway)
	if err != nil {
		log.Fatalln("dial", config.Gateway, "err:", err)
	}

	checkChainID(ec, config.ChainID)
	return ec
}

func checkChainID(ec *ethclient.Client, chainID uint64) {
	chid, err := ec.ChainID(context.Background())
	if err != nil {
		log.Fatalf("get chainid %d err: %s", chainID, err)
	}
	if chid.Uint64() != chainID {
		log.Fatalf("chainid mismatch! chainConf has %d but onchain has %d", chainID, chid.Uint64())
	}
}
