package contracts

import (
	"strings"
	"sync"

	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ec "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	EvIdCache     = make(map[string]ec.Hash)
	EvIdCacheLock sync.RWMutex
	ZeroHash      ec.Hash
	ZeroAddr      ec.Address
)

type MsgBusContract struct {
	*MessageBus
	Address ec.Address
}

func NewMsgBusContract(address ec.Address, client *ethclient.Client) (*MsgBusContract, error) {
	bus, err := NewMessageBus(address, client)
	if err != nil {
		return nil, err
	}
	return &MsgBusContract{
		MessageBus: bus,
		Address:    address,
	}, nil
}

func (c *MsgBusContract) GetAddr() ec.Address {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func FindMatchContractEvents(evName string, expAddr ec.Address, logs []*ethtypes.Log) (matchLogs []*ethtypes.Log) {
	evID := GetContractEventID(evName)
	if evID == ZeroHash {
		return nil
	}
	for idx := len(logs) - 1; idx >= 0; idx-- {
		if len(logs[idx].Topics) > 0 && logs[idx].Topics[0] == evID {
			// event ID matches and from expected contract
			if logs[idx].Address == expAddr {
				matchLogs = append(matchLogs, logs[idx])
			} else {
				log.Warnln("topic match but contract addr mismatch, hack or misconfig. log has:", logs[idx].Address, "expect:", expAddr)
			}
		}
	}
	return matchLogs
}

func GetContractEventID(evname string) ec.Hash {
	if evId := GetEvIdCache(evname); evId != ZeroHash {
		return evId
	}
	contractAbi, _ := abi.JSON(strings.NewReader(MessageBusABI))
	evId := contractAbi.Events[evname].ID
	SetEvIdCache(evname, evId)
	return evId
}

func GetEvIdCache(evname string) ec.Hash {
	EvIdCacheLock.RLock()
	defer EvIdCacheLock.RUnlock()
	if evId, ok := EvIdCache[evname]; ok {
		return evId
	}
	return ZeroHash
}

func SetEvIdCache(evname string, evId ec.Hash) {
	EvIdCacheLock.Lock()
	defer EvIdCacheLock.Unlock()
	EvIdCache[evname] = evId
}
