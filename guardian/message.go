package guardian

import (
	"context"
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-guardian/contracts"
	ec "github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

const (
	MsgType_MESSAGE_WITH_TRANSFER uint8 = 0
	MsgType_MESSAGE_ONLY          uint8 = 1
)

const (
	ResCode_Verified     = 0
	ResCode_MsgNotFound  = 1 // tx found, but no matching msg
	ResCode_UnkownChain  = 2 // src chain not configured
	ResCode_TxNotFound   = 3 // tx receipt not found by txHash
	ResCode_RpcFailed    = 4
	ResCode_NotSupported = 5
)

func (c *Chain) verifyMessage(ev *contracts.MessageBusExecuted) {
	resCode, err := checkSrcChainMsg(ev)
	if err != nil {
		log.Error(err)
		c.handleUnverifiedMessage(ev, resCode)
	}
}

// default action for invalid message: pause receiver contracts
// customized handling logic can be implemented here by the dApp community who run this guardian.
func (c *Chain) handleUnverifiedMessage(ev *contracts.MessageBusExecuted, resCode int) {
	c.pauseReceiverContracts(ev, resCode)
}

func checkSrcChainMsg(ev *contracts.MessageBusExecuted) (int, error) {
	if ev.MsgType != MsgType_MESSAGE_ONLY {
		// only support messageOnly for now
		return ResCode_NotSupported, nil
	}
	chains.lock.RLock()
	chain, ok := chains.chains[ev.SrcChainId]
	chains.lock.RUnlock()
	if !ok {
		return ResCode_UnkownChain, fmt.Errorf("src chain %d not configured", ev.SrcChainId)
	}
	txReceipt, err := chain.EthClient.TransactionReceipt(context.Background(), ev.SrcTxHash)
	if err != nil {
		return ResCode_TxNotFound, fmt.Errorf("get tx receipt err: %w", err)
	}
	elogs := contracts.FindMatchContractEvents("Message", chain.MessageBus.Address, txReceipt.Logs)
	for _, elog := range elogs {
		e, err := chain.MessageBus.ParseMessage(*elog)
		if err != nil {
			log.Errorf("parse event %v err: %s", e, err)
			continue
		}
		messageId := commputeMessageId(e, chain.ChainID)
		if messageId == ev.MsgId {
			log.Debugf("messageId %x found at src chain", ev.MsgId)
			return ResCode_Verified, nil
		}
	}
	return ResCode_MsgNotFound, fmt.Errorf("msg %x NOT found at src chain", ev.MsgId)
}

func commputeMessageId(ev *contracts.MessageBusMessage, chainId uint64) ec.Hash {
	var data []byte
	data = append(data, ToPadBytes(MsgType_MESSAGE_ONLY)...)
	data = append(data, ev.Sender.Bytes()...)
	data = append(data, ev.Receiver.Bytes()...)
	data = append(data, ToPadBytes(chainId)...)
	data = append(data, ev.Raw.TxHash.Bytes()...)
	data = append(data, ToPadBytes(ev.DstChainId.Uint64())...)
	data = append(data, ev.Message...)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return Bytes2Hash(hash.Sum(nil))
}
