// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tmp

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IPledgeABI is the input ABI used to generate the binding from.
const IPledgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"}],\"name\":\"AddT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"addT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lock\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IPledgeFuncSigs maps the 4-byte function signature to its string representation.
var IPledgeFuncSigs = map[string]string{
	"724a84a8": "addT(uint8)",
	"fc3ba0ad": "balanceOf(uint64,uint8)",
	"eb395fde": "getPledge(uint8)",
	"d23f7482": "pledge(uint64,uint256)",
	"14f732de": "rewards(uint64,uint8)",
	"c21a43e4": "totalPledge()",
	"5affa0f3": "withdraw(uint64,uint8,uint256,uint256)",
}

// IPledge is an auto generated Go binding around an Ethereum contract.
type IPledge struct {
	IPledgeCaller     // Read-only binding to the contract
	IPledgeTransactor // Write-only binding to the contract
	IPledgeFilterer   // Log filterer for contract events
}

// IPledgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPledgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPledgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPledgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPledgeSession struct {
	Contract     *IPledge          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPledgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPledgeCallerSession struct {
	Contract *IPledgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IPledgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPledgeTransactorSession struct {
	Contract     *IPledgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IPledgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPledgeRaw struct {
	Contract *IPledge // Generic contract binding to access the raw methods on
}

// IPledgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPledgeCallerRaw struct {
	Contract *IPledgeCaller // Generic read-only contract binding to access the raw methods on
}

// IPledgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPledgeTransactorRaw struct {
	Contract *IPledgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPledge creates a new instance of IPledge, bound to a specific deployed contract.
func NewIPledge(address common.Address, backend bind.ContractBackend) (*IPledge, error) {
	contract, err := bindIPledge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPledge{IPledgeCaller: IPledgeCaller{contract: contract}, IPledgeTransactor: IPledgeTransactor{contract: contract}, IPledgeFilterer: IPledgeFilterer{contract: contract}}, nil
}

// NewIPledgeCaller creates a new read-only instance of IPledge, bound to a specific deployed contract.
func NewIPledgeCaller(address common.Address, caller bind.ContractCaller) (*IPledgeCaller, error) {
	contract, err := bindIPledge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPledgeCaller{contract: contract}, nil
}

// NewIPledgeTransactor creates a new write-only instance of IPledge, bound to a specific deployed contract.
func NewIPledgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IPledgeTransactor, error) {
	contract, err := bindIPledge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPledgeTransactor{contract: contract}, nil
}

// NewIPledgeFilterer creates a new log filterer instance of IPledge, bound to a specific deployed contract.
func NewIPledgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IPledgeFilterer, error) {
	contract, err := bindIPledge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPledgeFilterer{contract: contract}, nil
}

// bindIPledge binds a generic wrapper to an already deployed contract.
func bindIPledge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPledgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPledge *IPledgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPledge.Contract.IPledgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPledge *IPledgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPledge.Contract.IPledgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPledge *IPledgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPledge.Contract.IPledgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPledge *IPledgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPledge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPledge *IPledgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPledge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPledge *IPledgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPledge.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256)
func (_IPledge *IPledgeCaller) BalanceOf(opts *bind.CallOpts, _i uint64, _ti uint8) (*big.Int, error) {
	var out []interface{}
	err := _IPledge.contract.Call(opts, &out, "balanceOf", _i, _ti)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256)
func (_IPledge *IPledgeSession) BalanceOf(_i uint64, _ti uint8) (*big.Int, error) {
	return _IPledge.Contract.BalanceOf(&_IPledge.CallOpts, _i, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256)
func (_IPledge *IPledgeCallerSession) BalanceOf(_i uint64, _ti uint8) (*big.Int, error) {
	return _IPledge.Contract.BalanceOf(&_IPledge.CallOpts, _i, _ti)
}

// GetPledge is a free data retrieval call binding the contract method 0xeb395fde.
//
// Solidity: function getPledge(uint8 _ti) view returns(uint256)
func (_IPledge *IPledgeCaller) GetPledge(opts *bind.CallOpts, _ti uint8) (*big.Int, error) {
	var out []interface{}
	err := _IPledge.contract.Call(opts, &out, "getPledge", _ti)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPledge is a free data retrieval call binding the contract method 0xeb395fde.
//
// Solidity: function getPledge(uint8 _ti) view returns(uint256)
func (_IPledge *IPledgeSession) GetPledge(_ti uint8) (*big.Int, error) {
	return _IPledge.Contract.GetPledge(&_IPledge.CallOpts, _ti)
}

// GetPledge is a free data retrieval call binding the contract method 0xeb395fde.
//
// Solidity: function getPledge(uint8 _ti) view returns(uint256)
func (_IPledge *IPledgeCallerSession) GetPledge(_ti uint8) (*big.Int, error) {
	return _IPledge.Contract.GetPledge(&_IPledge.CallOpts, _ti)
}

// Rewards is a free data retrieval call binding the contract method 0x14f732de.
//
// Solidity: function rewards(uint64 _index, uint8 _ti) view returns(uint256, uint256, uint256, uint256)
func (_IPledge *IPledgeCaller) Rewards(opts *bind.CallOpts, _index uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _IPledge.contract.Call(opts, &out, "rewards", _index, _ti)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// Rewards is a free data retrieval call binding the contract method 0x14f732de.
//
// Solidity: function rewards(uint64 _index, uint8 _ti) view returns(uint256, uint256, uint256, uint256)
func (_IPledge *IPledgeSession) Rewards(_index uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IPledge.Contract.Rewards(&_IPledge.CallOpts, _index, _ti)
}

// Rewards is a free data retrieval call binding the contract method 0x14f732de.
//
// Solidity: function rewards(uint64 _index, uint8 _ti) view returns(uint256, uint256, uint256, uint256)
func (_IPledge *IPledgeCallerSession) Rewards(_index uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IPledge.Contract.Rewards(&_IPledge.CallOpts, _index, _ti)
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_IPledge *IPledgeCaller) TotalPledge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPledge.contract.Call(opts, &out, "totalPledge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_IPledge *IPledgeSession) TotalPledge() (*big.Int, error) {
	return _IPledge.Contract.TotalPledge(&_IPledge.CallOpts)
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_IPledge *IPledgeCallerSession) TotalPledge() (*big.Int, error) {
	return _IPledge.Contract.TotalPledge(&_IPledge.CallOpts)
}

// AddT is a paid mutator transaction binding the contract method 0x724a84a8.
//
// Solidity: function addT(uint8 _ti) returns()
func (_IPledge *IPledgeTransactor) AddT(opts *bind.TransactOpts, _ti uint8) (*types.Transaction, error) {
	return _IPledge.contract.Transact(opts, "addT", _ti)
}

// AddT is a paid mutator transaction binding the contract method 0x724a84a8.
//
// Solidity: function addT(uint8 _ti) returns()
func (_IPledge *IPledgeSession) AddT(_ti uint8) (*types.Transaction, error) {
	return _IPledge.Contract.AddT(&_IPledge.TransactOpts, _ti)
}

// AddT is a paid mutator transaction binding the contract method 0x724a84a8.
//
// Solidity: function addT(uint8 _ti) returns()
func (_IPledge *IPledgeTransactorSession) AddT(_ti uint8) (*types.Transaction, error) {
	return _IPledge.Contract.AddT(&_IPledge.TransactOpts, _ti)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 money) returns()
func (_IPledge *IPledgeTransactor) Pledge(opts *bind.TransactOpts, _i uint64, money *big.Int) (*types.Transaction, error) {
	return _IPledge.contract.Transact(opts, "pledge", _i, money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 money) returns()
func (_IPledge *IPledgeSession) Pledge(_i uint64, money *big.Int) (*types.Transaction, error) {
	return _IPledge.Contract.Pledge(&_IPledge.TransactOpts, _i, money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 money) returns()
func (_IPledge *IPledgeTransactorSession) Pledge(_i uint64, money *big.Int) (*types.Transaction, error) {
	return _IPledge.Contract.Pledge(&_IPledge.TransactOpts, _i, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 lock) returns(uint256)
func (_IPledge *IPledgeTransactor) Withdraw(opts *bind.TransactOpts, _i uint64, _ti uint8, money *big.Int, lock *big.Int) (*types.Transaction, error) {
	return _IPledge.contract.Transact(opts, "withdraw", _i, _ti, money, lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 lock) returns(uint256)
func (_IPledge *IPledgeSession) Withdraw(_i uint64, _ti uint8, money *big.Int, lock *big.Int) (*types.Transaction, error) {
	return _IPledge.Contract.Withdraw(&_IPledge.TransactOpts, _i, _ti, money, lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 lock) returns(uint256)
func (_IPledge *IPledgeTransactorSession) Withdraw(_i uint64, _ti uint8, money *big.Int, lock *big.Int) (*types.Transaction, error) {
	return _IPledge.Contract.Withdraw(&_IPledge.TransactOpts, _i, _ti, money, lock)
}

// IPledgeAddTIterator is returned from FilterAddT and is used to iterate over the raw logs and unpacked data for AddT events raised by the IPledge contract.
type IPledgeAddTIterator struct {
	Event *IPledgeAddT // Event containing the contract specifics and raw log

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
func (it *IPledgeAddTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPledgeAddT)
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
		it.Event = new(IPledgeAddT)
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
func (it *IPledgeAddTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPledgeAddTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPledgeAddT represents a AddT event raised by the IPledge contract.
type IPledgeAddT struct {
	Bal    *big.Int
	TIndex uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddT is a free log retrieval operation binding the contract event 0x8ce8baccab364cf4dd55bc29e9ca458faa5067327fb854f45973d4bdcc5b8cf9.
//
// Solidity: event AddT(uint256 bal, uint8 tIndex)
func (_IPledge *IPledgeFilterer) FilterAddT(opts *bind.FilterOpts) (*IPledgeAddTIterator, error) {

	logs, sub, err := _IPledge.contract.FilterLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return &IPledgeAddTIterator{contract: _IPledge.contract, event: "AddT", logs: logs, sub: sub}, nil
}

// WatchAddT is a free log subscription operation binding the contract event 0x8ce8baccab364cf4dd55bc29e9ca458faa5067327fb854f45973d4bdcc5b8cf9.
//
// Solidity: event AddT(uint256 bal, uint8 tIndex)
func (_IPledge *IPledgeFilterer) WatchAddT(opts *bind.WatchOpts, sink chan<- *IPledgeAddT) (event.Subscription, error) {

	logs, sub, err := _IPledge.contract.WatchLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPledgeAddT)
				if err := _IPledge.contract.UnpackLog(event, "AddT", log); err != nil {
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

// ParseAddT is a log parse operation binding the contract event 0x8ce8baccab364cf4dd55bc29e9ca458faa5067327fb854f45973d4bdcc5b8cf9.
//
// Solidity: event AddT(uint256 bal, uint8 tIndex)
func (_IPledge *IPledgeFilterer) ParseAddT(log types.Log) (*IPledgeAddT, error) {
	event := new(IPledgeAddT)
	if err := _IPledge.contract.UnpackLog(event, "AddT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPledgeGetterABI is the input ABI used to generate the binding from.
const IPledgeGetterABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IPledgeGetterFuncSigs maps the 4-byte function signature to its string representation.
var IPledgeGetterFuncSigs = map[string]string{
	"fc3ba0ad": "balanceOf(uint64,uint8)",
	"eb395fde": "getPledge(uint8)",
	"14f732de": "rewards(uint64,uint8)",
	"c21a43e4": "totalPledge()",
}

// IPledgeGetter is an auto generated Go binding around an Ethereum contract.
type IPledgeGetter struct {
	IPledgeGetterCaller     // Read-only binding to the contract
	IPledgeGetterTransactor // Write-only binding to the contract
	IPledgeGetterFilterer   // Log filterer for contract events
}

// IPledgeGetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPledgeGetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeGetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPledgeGetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeGetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPledgeGetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeGetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPledgeGetterSession struct {
	Contract     *IPledgeGetter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPledgeGetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPledgeGetterCallerSession struct {
	Contract *IPledgeGetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IPledgeGetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPledgeGetterTransactorSession struct {
	Contract     *IPledgeGetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IPledgeGetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPledgeGetterRaw struct {
	Contract *IPledgeGetter // Generic contract binding to access the raw methods on
}

// IPledgeGetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPledgeGetterCallerRaw struct {
	Contract *IPledgeGetterCaller // Generic read-only contract binding to access the raw methods on
}

// IPledgeGetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPledgeGetterTransactorRaw struct {
	Contract *IPledgeGetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPledgeGetter creates a new instance of IPledgeGetter, bound to a specific deployed contract.
func NewIPledgeGetter(address common.Address, backend bind.ContractBackend) (*IPledgeGetter, error) {
	contract, err := bindIPledgeGetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPledgeGetter{IPledgeGetterCaller: IPledgeGetterCaller{contract: contract}, IPledgeGetterTransactor: IPledgeGetterTransactor{contract: contract}, IPledgeGetterFilterer: IPledgeGetterFilterer{contract: contract}}, nil
}

// NewIPledgeGetterCaller creates a new read-only instance of IPledgeGetter, bound to a specific deployed contract.
func NewIPledgeGetterCaller(address common.Address, caller bind.ContractCaller) (*IPledgeGetterCaller, error) {
	contract, err := bindIPledgeGetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPledgeGetterCaller{contract: contract}, nil
}

// NewIPledgeGetterTransactor creates a new write-only instance of IPledgeGetter, bound to a specific deployed contract.
func NewIPledgeGetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IPledgeGetterTransactor, error) {
	contract, err := bindIPledgeGetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPledgeGetterTransactor{contract: contract}, nil
}

// NewIPledgeGetterFilterer creates a new log filterer instance of IPledgeGetter, bound to a specific deployed contract.
func NewIPledgeGetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IPledgeGetterFilterer, error) {
	contract, err := bindIPledgeGetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPledgeGetterFilterer{contract: contract}, nil
}

// bindIPledgeGetter binds a generic wrapper to an already deployed contract.
func bindIPledgeGetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPledgeGetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPledgeGetter *IPledgeGetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPledgeGetter.Contract.IPledgeGetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPledgeGetter *IPledgeGetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPledgeGetter.Contract.IPledgeGetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPledgeGetter *IPledgeGetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPledgeGetter.Contract.IPledgeGetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPledgeGetter *IPledgeGetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPledgeGetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPledgeGetter *IPledgeGetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPledgeGetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPledgeGetter *IPledgeGetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPledgeGetter.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256)
func (_IPledgeGetter *IPledgeGetterCaller) BalanceOf(opts *bind.CallOpts, _i uint64, _ti uint8) (*big.Int, error) {
	var out []interface{}
	err := _IPledgeGetter.contract.Call(opts, &out, "balanceOf", _i, _ti)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256)
func (_IPledgeGetter *IPledgeGetterSession) BalanceOf(_i uint64, _ti uint8) (*big.Int, error) {
	return _IPledgeGetter.Contract.BalanceOf(&_IPledgeGetter.CallOpts, _i, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _i, uint8 _ti) view returns(uint256)
func (_IPledgeGetter *IPledgeGetterCallerSession) BalanceOf(_i uint64, _ti uint8) (*big.Int, error) {
	return _IPledgeGetter.Contract.BalanceOf(&_IPledgeGetter.CallOpts, _i, _ti)
}

// GetPledge is a free data retrieval call binding the contract method 0xeb395fde.
//
// Solidity: function getPledge(uint8 _ti) view returns(uint256)
func (_IPledgeGetter *IPledgeGetterCaller) GetPledge(opts *bind.CallOpts, _ti uint8) (*big.Int, error) {
	var out []interface{}
	err := _IPledgeGetter.contract.Call(opts, &out, "getPledge", _ti)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPledge is a free data retrieval call binding the contract method 0xeb395fde.
//
// Solidity: function getPledge(uint8 _ti) view returns(uint256)
func (_IPledgeGetter *IPledgeGetterSession) GetPledge(_ti uint8) (*big.Int, error) {
	return _IPledgeGetter.Contract.GetPledge(&_IPledgeGetter.CallOpts, _ti)
}

// GetPledge is a free data retrieval call binding the contract method 0xeb395fde.
//
// Solidity: function getPledge(uint8 _ti) view returns(uint256)
func (_IPledgeGetter *IPledgeGetterCallerSession) GetPledge(_ti uint8) (*big.Int, error) {
	return _IPledgeGetter.Contract.GetPledge(&_IPledgeGetter.CallOpts, _ti)
}

// Rewards is a free data retrieval call binding the contract method 0x14f732de.
//
// Solidity: function rewards(uint64 _index, uint8 _ti) view returns(uint256, uint256, uint256, uint256)
func (_IPledgeGetter *IPledgeGetterCaller) Rewards(opts *bind.CallOpts, _index uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _IPledgeGetter.contract.Call(opts, &out, "rewards", _index, _ti)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// Rewards is a free data retrieval call binding the contract method 0x14f732de.
//
// Solidity: function rewards(uint64 _index, uint8 _ti) view returns(uint256, uint256, uint256, uint256)
func (_IPledgeGetter *IPledgeGetterSession) Rewards(_index uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IPledgeGetter.Contract.Rewards(&_IPledgeGetter.CallOpts, _index, _ti)
}

// Rewards is a free data retrieval call binding the contract method 0x14f732de.
//
// Solidity: function rewards(uint64 _index, uint8 _ti) view returns(uint256, uint256, uint256, uint256)
func (_IPledgeGetter *IPledgeGetterCallerSession) Rewards(_index uint64, _ti uint8) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IPledgeGetter.Contract.Rewards(&_IPledgeGetter.CallOpts, _index, _ti)
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_IPledgeGetter *IPledgeGetterCaller) TotalPledge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPledgeGetter.contract.Call(opts, &out, "totalPledge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_IPledgeGetter *IPledgeGetterSession) TotalPledge() (*big.Int, error) {
	return _IPledgeGetter.Contract.TotalPledge(&_IPledgeGetter.CallOpts)
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_IPledgeGetter *IPledgeGetterCallerSession) TotalPledge() (*big.Int, error) {
	return _IPledgeGetter.Contract.TotalPledge(&_IPledgeGetter.CallOpts)
}

// IPledgeSetterABI is the input ABI used to generate the binding from.
const IPledgeSetterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"}],\"name\":\"AddT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"addT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lock\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IPledgeSetterFuncSigs maps the 4-byte function signature to its string representation.
var IPledgeSetterFuncSigs = map[string]string{
	"724a84a8": "addT(uint8)",
	"d23f7482": "pledge(uint64,uint256)",
	"5affa0f3": "withdraw(uint64,uint8,uint256,uint256)",
}

// IPledgeSetter is an auto generated Go binding around an Ethereum contract.
type IPledgeSetter struct {
	IPledgeSetterCaller     // Read-only binding to the contract
	IPledgeSetterTransactor // Write-only binding to the contract
	IPledgeSetterFilterer   // Log filterer for contract events
}

// IPledgeSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPledgeSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPledgeSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPledgeSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgeSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPledgeSetterSession struct {
	Contract     *IPledgeSetter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPledgeSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPledgeSetterCallerSession struct {
	Contract *IPledgeSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IPledgeSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPledgeSetterTransactorSession struct {
	Contract     *IPledgeSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IPledgeSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPledgeSetterRaw struct {
	Contract *IPledgeSetter // Generic contract binding to access the raw methods on
}

// IPledgeSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPledgeSetterCallerRaw struct {
	Contract *IPledgeSetterCaller // Generic read-only contract binding to access the raw methods on
}

// IPledgeSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPledgeSetterTransactorRaw struct {
	Contract *IPledgeSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPledgeSetter creates a new instance of IPledgeSetter, bound to a specific deployed contract.
func NewIPledgeSetter(address common.Address, backend bind.ContractBackend) (*IPledgeSetter, error) {
	contract, err := bindIPledgeSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPledgeSetter{IPledgeSetterCaller: IPledgeSetterCaller{contract: contract}, IPledgeSetterTransactor: IPledgeSetterTransactor{contract: contract}, IPledgeSetterFilterer: IPledgeSetterFilterer{contract: contract}}, nil
}

// NewIPledgeSetterCaller creates a new read-only instance of IPledgeSetter, bound to a specific deployed contract.
func NewIPledgeSetterCaller(address common.Address, caller bind.ContractCaller) (*IPledgeSetterCaller, error) {
	contract, err := bindIPledgeSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPledgeSetterCaller{contract: contract}, nil
}

// NewIPledgeSetterTransactor creates a new write-only instance of IPledgeSetter, bound to a specific deployed contract.
func NewIPledgeSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IPledgeSetterTransactor, error) {
	contract, err := bindIPledgeSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPledgeSetterTransactor{contract: contract}, nil
}

// NewIPledgeSetterFilterer creates a new log filterer instance of IPledgeSetter, bound to a specific deployed contract.
func NewIPledgeSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IPledgeSetterFilterer, error) {
	contract, err := bindIPledgeSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPledgeSetterFilterer{contract: contract}, nil
}

// bindIPledgeSetter binds a generic wrapper to an already deployed contract.
func bindIPledgeSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPledgeSetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPledgeSetter *IPledgeSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPledgeSetter.Contract.IPledgeSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPledgeSetter *IPledgeSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.IPledgeSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPledgeSetter *IPledgeSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.IPledgeSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPledgeSetter *IPledgeSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPledgeSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPledgeSetter *IPledgeSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPledgeSetter *IPledgeSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.contract.Transact(opts, method, params...)
}

// AddT is a paid mutator transaction binding the contract method 0x724a84a8.
//
// Solidity: function addT(uint8 _ti) returns()
func (_IPledgeSetter *IPledgeSetterTransactor) AddT(opts *bind.TransactOpts, _ti uint8) (*types.Transaction, error) {
	return _IPledgeSetter.contract.Transact(opts, "addT", _ti)
}

// AddT is a paid mutator transaction binding the contract method 0x724a84a8.
//
// Solidity: function addT(uint8 _ti) returns()
func (_IPledgeSetter *IPledgeSetterSession) AddT(_ti uint8) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.AddT(&_IPledgeSetter.TransactOpts, _ti)
}

// AddT is a paid mutator transaction binding the contract method 0x724a84a8.
//
// Solidity: function addT(uint8 _ti) returns()
func (_IPledgeSetter *IPledgeSetterTransactorSession) AddT(_ti uint8) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.AddT(&_IPledgeSetter.TransactOpts, _ti)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 money) returns()
func (_IPledgeSetter *IPledgeSetterTransactor) Pledge(opts *bind.TransactOpts, _i uint64, money *big.Int) (*types.Transaction, error) {
	return _IPledgeSetter.contract.Transact(opts, "pledge", _i, money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 money) returns()
func (_IPledgeSetter *IPledgeSetterSession) Pledge(_i uint64, money *big.Int) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.Pledge(&_IPledgeSetter.TransactOpts, _i, money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 money) returns()
func (_IPledgeSetter *IPledgeSetterTransactorSession) Pledge(_i uint64, money *big.Int) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.Pledge(&_IPledgeSetter.TransactOpts, _i, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 lock) returns(uint256)
func (_IPledgeSetter *IPledgeSetterTransactor) Withdraw(opts *bind.TransactOpts, _i uint64, _ti uint8, money *big.Int, lock *big.Int) (*types.Transaction, error) {
	return _IPledgeSetter.contract.Transact(opts, "withdraw", _i, _ti, money, lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 lock) returns(uint256)
func (_IPledgeSetter *IPledgeSetterSession) Withdraw(_i uint64, _ti uint8, money *big.Int, lock *big.Int) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.Withdraw(&_IPledgeSetter.TransactOpts, _i, _ti, money, lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 lock) returns(uint256)
func (_IPledgeSetter *IPledgeSetterTransactorSession) Withdraw(_i uint64, _ti uint8, money *big.Int, lock *big.Int) (*types.Transaction, error) {
	return _IPledgeSetter.Contract.Withdraw(&_IPledgeSetter.TransactOpts, _i, _ti, money, lock)
}

// IPledgeSetterAddTIterator is returned from FilterAddT and is used to iterate over the raw logs and unpacked data for AddT events raised by the IPledgeSetter contract.
type IPledgeSetterAddTIterator struct {
	Event *IPledgeSetterAddT // Event containing the contract specifics and raw log

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
func (it *IPledgeSetterAddTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPledgeSetterAddT)
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
		it.Event = new(IPledgeSetterAddT)
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
func (it *IPledgeSetterAddTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPledgeSetterAddTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPledgeSetterAddT represents a AddT event raised by the IPledgeSetter contract.
type IPledgeSetterAddT struct {
	Bal    *big.Int
	TIndex uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddT is a free log retrieval operation binding the contract event 0x8ce8baccab364cf4dd55bc29e9ca458faa5067327fb854f45973d4bdcc5b8cf9.
//
// Solidity: event AddT(uint256 bal, uint8 tIndex)
func (_IPledgeSetter *IPledgeSetterFilterer) FilterAddT(opts *bind.FilterOpts) (*IPledgeSetterAddTIterator, error) {

	logs, sub, err := _IPledgeSetter.contract.FilterLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return &IPledgeSetterAddTIterator{contract: _IPledgeSetter.contract, event: "AddT", logs: logs, sub: sub}, nil
}

// WatchAddT is a free log subscription operation binding the contract event 0x8ce8baccab364cf4dd55bc29e9ca458faa5067327fb854f45973d4bdcc5b8cf9.
//
// Solidity: event AddT(uint256 bal, uint8 tIndex)
func (_IPledgeSetter *IPledgeSetterFilterer) WatchAddT(opts *bind.WatchOpts, sink chan<- *IPledgeSetterAddT) (event.Subscription, error) {

	logs, sub, err := _IPledgeSetter.contract.WatchLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPledgeSetterAddT)
				if err := _IPledgeSetter.contract.UnpackLog(event, "AddT", log); err != nil {
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

// ParseAddT is a log parse operation binding the contract event 0x8ce8baccab364cf4dd55bc29e9ca458faa5067327fb854f45973d4bdcc5b8cf9.
//
// Solidity: event AddT(uint256 bal, uint8 tIndex)
func (_IPledgeSetter *IPledgeSetterFilterer) ParseAddT(log types.Log) (*IPledgeSetterAddT, error) {
	event := new(IPledgeSetterAddT)
	if err := _IPledgeSetter.contract.UnpackLog(event, "AddT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeAddTABI is the input ABI used to generate the binding from.
const PledgeAddTABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"pledgeAddr\",\"type\":\"address\"}],\"name\":\"addT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PledgeAddTFuncSigs maps the 4-byte function signature to its string representation.
var PledgeAddTFuncSigs = map[string]string{
	"c1046ee4": "addT(uint8,address)",
}

// PledgeAddTBin is the compiled bytecode used for deploying new contracts.
var PledgeAddTBin = "0x608060405234801561001057600080fd5b5061011a806100206000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063c1046ee414602d575b600080fd5b603c6038366004609c565b603e565b005b604051630e49509560e31b815260ff831660048201526001600160a01b0382169063724a84a890602401600060405180830381600087803b158015608157600080fd5b505af11580156094573d6000803e3d6000fd5b505050505050565b6000806040838503121560ae57600080fd5b823560ff8116811460be57600080fd5b915060208301356001600160a01b038116811460d957600080fd5b80915050925092905056fea264697066735822122025f00c17b07553bb41069feb750d50edb4d8b1e61ba244cfc3f643ae677192c864736f6c634300080e0033"

// DeployPledgeAddT deploys a new Ethereum contract, binding an instance of PledgeAddT to it.
func DeployPledgeAddT(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PledgeAddT, error) {
	parsed, err := abi.JSON(strings.NewReader(PledgeAddTABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PledgeAddTBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PledgeAddT{PledgeAddTCaller: PledgeAddTCaller{contract: contract}, PledgeAddTTransactor: PledgeAddTTransactor{contract: contract}, PledgeAddTFilterer: PledgeAddTFilterer{contract: contract}}, nil
}

// PledgeAddT is an auto generated Go binding around an Ethereum contract.
type PledgeAddT struct {
	PledgeAddTCaller     // Read-only binding to the contract
	PledgeAddTTransactor // Write-only binding to the contract
	PledgeAddTFilterer   // Log filterer for contract events
}

// PledgeAddTCaller is an auto generated read-only Go binding around an Ethereum contract.
type PledgeAddTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeAddTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PledgeAddTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeAddTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PledgeAddTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeAddTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PledgeAddTSession struct {
	Contract     *PledgeAddT       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PledgeAddTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PledgeAddTCallerSession struct {
	Contract *PledgeAddTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PledgeAddTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PledgeAddTTransactorSession struct {
	Contract     *PledgeAddTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PledgeAddTRaw is an auto generated low-level Go binding around an Ethereum contract.
type PledgeAddTRaw struct {
	Contract *PledgeAddT // Generic contract binding to access the raw methods on
}

// PledgeAddTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PledgeAddTCallerRaw struct {
	Contract *PledgeAddTCaller // Generic read-only contract binding to access the raw methods on
}

// PledgeAddTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PledgeAddTTransactorRaw struct {
	Contract *PledgeAddTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPledgeAddT creates a new instance of PledgeAddT, bound to a specific deployed contract.
func NewPledgeAddT(address common.Address, backend bind.ContractBackend) (*PledgeAddT, error) {
	contract, err := bindPledgeAddT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PledgeAddT{PledgeAddTCaller: PledgeAddTCaller{contract: contract}, PledgeAddTTransactor: PledgeAddTTransactor{contract: contract}, PledgeAddTFilterer: PledgeAddTFilterer{contract: contract}}, nil
}

// NewPledgeAddTCaller creates a new read-only instance of PledgeAddT, bound to a specific deployed contract.
func NewPledgeAddTCaller(address common.Address, caller bind.ContractCaller) (*PledgeAddTCaller, error) {
	contract, err := bindPledgeAddT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PledgeAddTCaller{contract: contract}, nil
}

// NewPledgeAddTTransactor creates a new write-only instance of PledgeAddT, bound to a specific deployed contract.
func NewPledgeAddTTransactor(address common.Address, transactor bind.ContractTransactor) (*PledgeAddTTransactor, error) {
	contract, err := bindPledgeAddT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PledgeAddTTransactor{contract: contract}, nil
}

// NewPledgeAddTFilterer creates a new log filterer instance of PledgeAddT, bound to a specific deployed contract.
func NewPledgeAddTFilterer(address common.Address, filterer bind.ContractFilterer) (*PledgeAddTFilterer, error) {
	contract, err := bindPledgeAddT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PledgeAddTFilterer{contract: contract}, nil
}

// bindPledgeAddT binds a generic wrapper to an already deployed contract.
func bindPledgeAddT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PledgeAddTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PledgeAddT *PledgeAddTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PledgeAddT.Contract.PledgeAddTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PledgeAddT *PledgeAddTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PledgeAddT.Contract.PledgeAddTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PledgeAddT *PledgeAddTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PledgeAddT.Contract.PledgeAddTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PledgeAddT *PledgeAddTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PledgeAddT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PledgeAddT *PledgeAddTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PledgeAddT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PledgeAddT *PledgeAddTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PledgeAddT.Contract.contract.Transact(opts, method, params...)
}

// AddT is a paid mutator transaction binding the contract method 0xc1046ee4.
//
// Solidity: function addT(uint8 _ti, address pledgeAddr) returns()
func (_PledgeAddT *PledgeAddTTransactor) AddT(opts *bind.TransactOpts, _ti uint8, pledgeAddr common.Address) (*types.Transaction, error) {
	return _PledgeAddT.contract.Transact(opts, "addT", _ti, pledgeAddr)
}

// AddT is a paid mutator transaction binding the contract method 0xc1046ee4.
//
// Solidity: function addT(uint8 _ti, address pledgeAddr) returns()
func (_PledgeAddT *PledgeAddTSession) AddT(_ti uint8, pledgeAddr common.Address) (*types.Transaction, error) {
	return _PledgeAddT.Contract.AddT(&_PledgeAddT.TransactOpts, _ti, pledgeAddr)
}

// AddT is a paid mutator transaction binding the contract method 0xc1046ee4.
//
// Solidity: function addT(uint8 _ti, address pledgeAddr) returns()
func (_PledgeAddT *PledgeAddTTransactorSession) AddT(_ti uint8, pledgeAddr common.Address) (*types.Transaction, error) {
	return _PledgeAddT.Contract.AddT(&_PledgeAddT.TransactOpts, _ti, pledgeAddr)
}
