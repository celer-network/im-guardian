// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MsgAppPauserMetaData contains all meta data concerning the MsgAppPauser contract.
var MsgAppPauserMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MsgAppPauserABI is the input ABI used to generate the binding from.
// Deprecated: Use MsgAppPauserMetaData.ABI instead.
var MsgAppPauserABI = MsgAppPauserMetaData.ABI

// MsgAppPauser is an auto generated Go binding around an Ethereum contract.
type MsgAppPauser struct {
	MsgAppPauserCaller     // Read-only binding to the contract
	MsgAppPauserTransactor // Write-only binding to the contract
	MsgAppPauserFilterer   // Log filterer for contract events
}

// MsgAppPauserCaller is an auto generated read-only Go binding around an Ethereum contract.
type MsgAppPauserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgAppPauserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MsgAppPauserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgAppPauserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MsgAppPauserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgAppPauserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MsgAppPauserSession struct {
	Contract     *MsgAppPauser     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MsgAppPauserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MsgAppPauserCallerSession struct {
	Contract *MsgAppPauserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MsgAppPauserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MsgAppPauserTransactorSession struct {
	Contract     *MsgAppPauserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MsgAppPauserRaw is an auto generated low-level Go binding around an Ethereum contract.
type MsgAppPauserRaw struct {
	Contract *MsgAppPauser // Generic contract binding to access the raw methods on
}

// MsgAppPauserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MsgAppPauserCallerRaw struct {
	Contract *MsgAppPauserCaller // Generic read-only contract binding to access the raw methods on
}

// MsgAppPauserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MsgAppPauserTransactorRaw struct {
	Contract *MsgAppPauserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMsgAppPauser creates a new instance of MsgAppPauser, bound to a specific deployed contract.
func NewMsgAppPauser(address common.Address, backend bind.ContractBackend) (*MsgAppPauser, error) {
	contract, err := bindMsgAppPauser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MsgAppPauser{MsgAppPauserCaller: MsgAppPauserCaller{contract: contract}, MsgAppPauserTransactor: MsgAppPauserTransactor{contract: contract}, MsgAppPauserFilterer: MsgAppPauserFilterer{contract: contract}}, nil
}

// NewMsgAppPauserCaller creates a new read-only instance of MsgAppPauser, bound to a specific deployed contract.
func NewMsgAppPauserCaller(address common.Address, caller bind.ContractCaller) (*MsgAppPauserCaller, error) {
	contract, err := bindMsgAppPauser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MsgAppPauserCaller{contract: contract}, nil
}

// NewMsgAppPauserTransactor creates a new write-only instance of MsgAppPauser, bound to a specific deployed contract.
func NewMsgAppPauserTransactor(address common.Address, transactor bind.ContractTransactor) (*MsgAppPauserTransactor, error) {
	contract, err := bindMsgAppPauser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MsgAppPauserTransactor{contract: contract}, nil
}

// NewMsgAppPauserFilterer creates a new log filterer instance of MsgAppPauser, bound to a specific deployed contract.
func NewMsgAppPauserFilterer(address common.Address, filterer bind.ContractFilterer) (*MsgAppPauserFilterer, error) {
	contract, err := bindMsgAppPauser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MsgAppPauserFilterer{contract: contract}, nil
}

// bindMsgAppPauser binds a generic wrapper to an already deployed contract.
func bindMsgAppPauser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MsgAppPauserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MsgAppPauser *MsgAppPauserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MsgAppPauser.Contract.MsgAppPauserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MsgAppPauser *MsgAppPauserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MsgAppPauser.Contract.MsgAppPauserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MsgAppPauser *MsgAppPauserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MsgAppPauser.Contract.MsgAppPauserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MsgAppPauser *MsgAppPauserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MsgAppPauser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MsgAppPauser *MsgAppPauserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MsgAppPauser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MsgAppPauser *MsgAppPauserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MsgAppPauser.Contract.contract.Transact(opts, method, params...)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_MsgAppPauser *MsgAppPauserCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _MsgAppPauser.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_MsgAppPauser *MsgAppPauserSession) IsPauser(account common.Address) (bool, error) {
	return _MsgAppPauser.Contract.IsPauser(&_MsgAppPauser.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_MsgAppPauser *MsgAppPauserCallerSession) IsPauser(account common.Address) (bool, error) {
	return _MsgAppPauser.Contract.IsPauser(&_MsgAppPauser.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MsgAppPauser *MsgAppPauserCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MsgAppPauser.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MsgAppPauser *MsgAppPauserSession) Owner() (common.Address, error) {
	return _MsgAppPauser.Contract.Owner(&_MsgAppPauser.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MsgAppPauser *MsgAppPauserCallerSession) Owner() (common.Address, error) {
	return _MsgAppPauser.Contract.Owner(&_MsgAppPauser.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MsgAppPauser *MsgAppPauserCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MsgAppPauser.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MsgAppPauser *MsgAppPauserSession) Paused() (bool, error) {
	return _MsgAppPauser.Contract.Paused(&_MsgAppPauser.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MsgAppPauser *MsgAppPauserCallerSession) Paused() (bool, error) {
	return _MsgAppPauser.Contract.Paused(&_MsgAppPauser.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_MsgAppPauser *MsgAppPauserCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MsgAppPauser.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_MsgAppPauser *MsgAppPauserSession) Pausers(arg0 common.Address) (bool, error) {
	return _MsgAppPauser.Contract.Pausers(&_MsgAppPauser.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_MsgAppPauser *MsgAppPauserCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _MsgAppPauser.Contract.Pausers(&_MsgAppPauser.CallOpts, arg0)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_MsgAppPauser *MsgAppPauserTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MsgAppPauser.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_MsgAppPauser *MsgAppPauserSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _MsgAppPauser.Contract.AddPauser(&_MsgAppPauser.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_MsgAppPauser *MsgAppPauserTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _MsgAppPauser.Contract.AddPauser(&_MsgAppPauser.TransactOpts, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MsgAppPauser *MsgAppPauserTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MsgAppPauser.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MsgAppPauser *MsgAppPauserSession) Pause() (*types.Transaction, error) {
	return _MsgAppPauser.Contract.Pause(&_MsgAppPauser.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MsgAppPauser *MsgAppPauserTransactorSession) Pause() (*types.Transaction, error) {
	return _MsgAppPauser.Contract.Pause(&_MsgAppPauser.TransactOpts)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_MsgAppPauser *MsgAppPauserTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MsgAppPauser.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_MsgAppPauser *MsgAppPauserSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _MsgAppPauser.Contract.RemovePauser(&_MsgAppPauser.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_MsgAppPauser *MsgAppPauserTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _MsgAppPauser.Contract.RemovePauser(&_MsgAppPauser.TransactOpts, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MsgAppPauser *MsgAppPauserTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MsgAppPauser.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MsgAppPauser *MsgAppPauserSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MsgAppPauser.Contract.TransferOwnership(&_MsgAppPauser.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MsgAppPauser *MsgAppPauserTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MsgAppPauser.Contract.TransferOwnership(&_MsgAppPauser.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MsgAppPauser *MsgAppPauserTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MsgAppPauser.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MsgAppPauser *MsgAppPauserSession) Unpause() (*types.Transaction, error) {
	return _MsgAppPauser.Contract.Unpause(&_MsgAppPauser.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MsgAppPauser *MsgAppPauserTransactorSession) Unpause() (*types.Transaction, error) {
	return _MsgAppPauser.Contract.Unpause(&_MsgAppPauser.TransactOpts)
}

// MsgAppPauserOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MsgAppPauser contract.
type MsgAppPauserOwnershipTransferredIterator struct {
	Event *MsgAppPauserOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MsgAppPauserOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MsgAppPauserOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MsgAppPauserOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MsgAppPauserOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MsgAppPauserOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MsgAppPauserOwnershipTransferred represents a OwnershipTransferred event raised by the MsgAppPauser contract.
type MsgAppPauserOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MsgAppPauser *MsgAppPauserFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MsgAppPauserOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MsgAppPauser.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MsgAppPauserOwnershipTransferredIterator{contract: _MsgAppPauser.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MsgAppPauser *MsgAppPauserFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MsgAppPauserOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MsgAppPauser.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MsgAppPauserOwnershipTransferred)
				if err := _MsgAppPauser.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MsgAppPauser *MsgAppPauserFilterer) ParseOwnershipTransferred(log types.Log) (*MsgAppPauserOwnershipTransferred, error) {
	event := new(MsgAppPauserOwnershipTransferred)
	if err := _MsgAppPauser.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MsgAppPauserPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MsgAppPauser contract.
type MsgAppPauserPausedIterator struct {
	Event *MsgAppPauserPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MsgAppPauserPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MsgAppPauserPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MsgAppPauserPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MsgAppPauserPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MsgAppPauserPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MsgAppPauserPaused represents a Paused event raised by the MsgAppPauser contract.
type MsgAppPauserPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MsgAppPauser *MsgAppPauserFilterer) FilterPaused(opts *bind.FilterOpts) (*MsgAppPauserPausedIterator, error) {

	logs, sub, err := _MsgAppPauser.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MsgAppPauserPausedIterator{contract: _MsgAppPauser.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MsgAppPauser *MsgAppPauserFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MsgAppPauserPaused) (event.Subscription, error) {

	logs, sub, err := _MsgAppPauser.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MsgAppPauserPaused)
				if err := _MsgAppPauser.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MsgAppPauser *MsgAppPauserFilterer) ParsePaused(log types.Log) (*MsgAppPauserPaused, error) {
	event := new(MsgAppPauserPaused)
	if err := _MsgAppPauser.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MsgAppPauserPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the MsgAppPauser contract.
type MsgAppPauserPauserAddedIterator struct {
	Event *MsgAppPauserPauserAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MsgAppPauserPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MsgAppPauserPauserAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MsgAppPauserPauserAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MsgAppPauserPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MsgAppPauserPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MsgAppPauserPauserAdded represents a PauserAdded event raised by the MsgAppPauser contract.
type MsgAppPauserPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_MsgAppPauser *MsgAppPauserFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*MsgAppPauserPauserAddedIterator, error) {

	logs, sub, err := _MsgAppPauser.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &MsgAppPauserPauserAddedIterator{contract: _MsgAppPauser.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_MsgAppPauser *MsgAppPauserFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *MsgAppPauserPauserAdded) (event.Subscription, error) {

	logs, sub, err := _MsgAppPauser.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MsgAppPauserPauserAdded)
				if err := _MsgAppPauser.contract.UnpackLog(event, "PauserAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_MsgAppPauser *MsgAppPauserFilterer) ParsePauserAdded(log types.Log) (*MsgAppPauserPauserAdded, error) {
	event := new(MsgAppPauserPauserAdded)
	if err := _MsgAppPauser.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MsgAppPauserPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the MsgAppPauser contract.
type MsgAppPauserPauserRemovedIterator struct {
	Event *MsgAppPauserPauserRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MsgAppPauserPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MsgAppPauserPauserRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MsgAppPauserPauserRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MsgAppPauserPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MsgAppPauserPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MsgAppPauserPauserRemoved represents a PauserRemoved event raised by the MsgAppPauser contract.
type MsgAppPauserPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_MsgAppPauser *MsgAppPauserFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*MsgAppPauserPauserRemovedIterator, error) {

	logs, sub, err := _MsgAppPauser.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &MsgAppPauserPauserRemovedIterator{contract: _MsgAppPauser.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_MsgAppPauser *MsgAppPauserFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *MsgAppPauserPauserRemoved) (event.Subscription, error) {

	logs, sub, err := _MsgAppPauser.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MsgAppPauserPauserRemoved)
				if err := _MsgAppPauser.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_MsgAppPauser *MsgAppPauserFilterer) ParsePauserRemoved(log types.Log) (*MsgAppPauserPauserRemoved, error) {
	event := new(MsgAppPauserPauserRemoved)
	if err := _MsgAppPauser.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MsgAppPauserUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MsgAppPauser contract.
type MsgAppPauserUnpausedIterator struct {
	Event *MsgAppPauserUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MsgAppPauserUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MsgAppPauserUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MsgAppPauserUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MsgAppPauserUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MsgAppPauserUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MsgAppPauserUnpaused represents a Unpaused event raised by the MsgAppPauser contract.
type MsgAppPauserUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MsgAppPauser *MsgAppPauserFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MsgAppPauserUnpausedIterator, error) {

	logs, sub, err := _MsgAppPauser.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MsgAppPauserUnpausedIterator{contract: _MsgAppPauser.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MsgAppPauser *MsgAppPauserFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MsgAppPauserUnpaused) (event.Subscription, error) {

	logs, sub, err := _MsgAppPauser.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MsgAppPauserUnpaused)
				if err := _MsgAppPauser.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MsgAppPauser *MsgAppPauserFilterer) ParseUnpaused(log types.Log) (*MsgAppPauserUnpaused, error) {
	event := new(MsgAppPauserUnpaused)
	if err := _MsgAppPauser.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
