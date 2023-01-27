package guardian

import (
	"strings"
	"sync"
	"time"

	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-guardian/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ec "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func StartMonitoring() {
	for _, chain := range chains.chains {
		go chain.monitorMsgExecuted()
	}
}

func (c *Chain) monitorMsgExecuted() {
	addrConfig := mon2.PerAddrCfg{
		Addr:    c.MessageBus.Address,
		ChkIntv: time.Minute,
		AbiStr:  contracts.MessageBusABI,
	}
	msgbusABI, err := abi.JSON(strings.NewReader(contracts.MessageBusABI))
	if err != nil {
		log.Fatalf("failed to parse MessageBus ABI")
	}
	evExecuted, ok := msgbusABI.Events["Executed"]
	if !ok {
		log.Fatalf("failed to get Executed event MessageBus ABI")
	}
	if len(c.Receivers) > 0 {
		addrs := []ec.Hash{}
		for _, addr := range c.Receivers {
			addrs = append(addrs, Hex2Addr(addr).Hash())
		}
		addrConfig.Topics = [][]ec.Hash{{evExecuted.ID}, addrs}
	} else {
		log.Infof("no address filter for chain %d, stop monitoring MessageBus Executed on this chain", c.ChainID)
		return
	}

	log.Infof("monitoring MessageBus Executed on chain %d with address filters %#s", c.ChainID, addrConfig.Topics[1])
	go c.monitor2.MonAddr(addrConfig, func(_ string, eLog ethtypes.Log) {
		ev, err := c.MessageBus.ParseExecuted(eLog)
		if err != nil {
			log.Errorln("cannot parse event Executed", err)
			return
		}
		log.Infof("monitorBusExecuted: got event Executed %v", ev)
		go c.verifyMessage(ev)
	})
}

type MonitorDAL struct {
	kv   map[string]mon2.LogEventID
	lock sync.RWMutex
}

func (d *MonitorDAL) GetMonitorBlock(key string) (uint64, int64, bool, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	le, ok := d.kv[key]
	return le.BlkNum, le.Index, ok, nil
}

func (d *MonitorDAL) SetMonitorBlock(key string, blockNum uint64, blockIdx int64) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.kv[key] = mon2.LogEventID{BlkNum: blockNum, Index: blockIdx}
	return nil
}
