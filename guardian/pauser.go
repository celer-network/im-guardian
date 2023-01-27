package guardian

import (
	"math/big"
	"time"

	"github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-guardian/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ec "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"
)

type PauserConfig struct {
	Keystore, Passphrase string
	PauseAll             bool
	PauseCode            []int
}

var pauserConfig PauserConfig

func loadPauserConfig() {
	err := viper.UnmarshalKey(FlagPauser, &pauserConfig)
	if err != nil {
		log.Fatalf("failed to load pauser config:%v", err)
	}
}

func (c *Chain) pauseReceiverContracts(ev *contracts.MessageBusExecuted, resCode int) {
	pauseCode := []int{ResCode_MsgNotFound, ResCode_UnkownChain, ResCode_TxNotFound} // default code
	if len(pauserConfig.PauseCode) > 0 {
		pauseCode = pauserConfig.PauseCode
	}
	pause := false
	for _, code := range pauseCode {
		if resCode == code {
			pause = true
			break
		}
	}
	if !pause {
		return
	}
	if pauserConfig.PauseAll {
		pauseAllContracts()
	} else {
		c.pauseContract(ev.Receiver)
	}
}

func (c *Chain) pauseContract(addr ec.Address) {
	tx, err := c.pauser.Transact(
		&eth.TransactionStateHandler{
			OnMined: func(receipt *ethtypes.Receipt) {
				if receipt.Status == ethtypes.ReceiptStatusSuccessful {
					log.Infof("Pause transaction %d %x succeeded", c.ChainID, receipt.TxHash)
				} else {
					log.Errorf("Pause transaction %d %x failed", c.ChainID, receipt.TxHash)
				}
			},
			OnError: func(tx *ethtypes.Transaction, err error) {
				log.Errorf("Pause transaction %d %x err: %s", c.ChainID, tx.Hash(), err)
			},
		},
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			contract, err2 := contracts.NewMsgAppPauserTransactor(addr, transactor)
			if err2 != nil {
				return nil, err2
			}
			return contract.Pause(opts)
		},
	)
	if err != nil {
		log.Errorf("Pause chain %d contract %x err: %s", c.ChainID, addr, err)
		return
	}
	log.Infof("Pause chain %d contract %x tx submitted: %x", c.ChainID, addr, tx.Hash())
}

func pauseAllContracts() {
	chains.lock.RLock()
	defer chains.lock.RUnlock()
	for _, chain := range chains.chains {
		for _, addr := range chain.Receivers {
			chain.pauseContract(Hex2Addr(addr))
		}
	}
}

func newPauserTransactor(chain *Chain) (*eth.Transactor, error) {
	return eth.NewTransactor(
		pauserConfig.Keystore,
		pauserConfig.Passphrase,
		chain.EthClient,
		big.NewInt(int64(chain.ChainID)),
		eth.WithBlockDelay(chain.BlkDelay),
		eth.WithPollingInterval(time.Duration(chain.BlkInterval)*time.Second),
		eth.WithAddGasEstimateRatio(chain.AddGasEstimateRatio),
		eth.WithGasLimit(chain.GasLimit),
		eth.WithAddGasGwei(chain.AddGasGwei),
		eth.WithMaxGasGwei(chain.MaxGasGwei),
		eth.WithMinGasGwei(chain.MinGasGwei),
		eth.WithMaxFeePerGasGwei(chain.MaxFeePerGasGwei),
		eth.WithMaxPriorityFeePerGasGwei(chain.MaxPriorityFeePerGasGwei),
	)
}
