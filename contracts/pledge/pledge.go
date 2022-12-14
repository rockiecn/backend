// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pledge

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

// IAuthABI is the input ABI used to generate the binding from.
const IAuthABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"perm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IAuthFuncSigs maps the 4-byte function signature to its string representation.
var IAuthFuncSigs = map[string]string{
	"a96bba9d": "perm(bytes32,bytes[5])",
}

// IAuth is an auto generated Go binding around an Ethereum contract.
type IAuth struct {
	IAuthCaller     // Read-only binding to the contract
	IAuthTransactor // Write-only binding to the contract
	IAuthFilterer   // Log filterer for contract events
}

// IAuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAuthSession struct {
	Contract     *IAuth            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAuthCallerSession struct {
	Contract *IAuthCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IAuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAuthTransactorSession struct {
	Contract     *IAuthTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAuthRaw struct {
	Contract *IAuth // Generic contract binding to access the raw methods on
}

// IAuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAuthCallerRaw struct {
	Contract *IAuthCaller // Generic read-only contract binding to access the raw methods on
}

// IAuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAuthTransactorRaw struct {
	Contract *IAuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAuth creates a new instance of IAuth, bound to a specific deployed contract.
func NewIAuth(address common.Address, backend bind.ContractBackend) (*IAuth, error) {
	contract, err := bindIAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAuth{IAuthCaller: IAuthCaller{contract: contract}, IAuthTransactor: IAuthTransactor{contract: contract}, IAuthFilterer: IAuthFilterer{contract: contract}}, nil
}

// NewIAuthCaller creates a new read-only instance of IAuth, bound to a specific deployed contract.
func NewIAuthCaller(address common.Address, caller bind.ContractCaller) (*IAuthCaller, error) {
	contract, err := bindIAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthCaller{contract: contract}, nil
}

// NewIAuthTransactor creates a new write-only instance of IAuth, bound to a specific deployed contract.
func NewIAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*IAuthTransactor, error) {
	contract, err := bindIAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthTransactor{contract: contract}, nil
}

// NewIAuthFilterer creates a new log filterer instance of IAuth, bound to a specific deployed contract.
func NewIAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*IAuthFilterer, error) {
	contract, err := bindIAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAuthFilterer{contract: contract}, nil
}

// bindIAuth binds a generic wrapper to an already deployed contract.
func bindIAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAuthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuth *IAuthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuth.Contract.IAuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuth *IAuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuth.Contract.IAuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuth *IAuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuth.Contract.IAuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuth *IAuthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuth *IAuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuth *IAuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuth.Contract.contract.Transact(opts, method, params...)
}

// Perm is a paid mutator transaction binding the contract method 0xa96bba9d.
//
// Solidity: function perm(bytes32 _h, bytes[5] signs) returns()
func (_IAuth *IAuthTransactor) Perm(opts *bind.TransactOpts, _h [32]byte, signs [5][]byte) (*types.Transaction, error) {
	return _IAuth.contract.Transact(opts, "perm", _h, signs)
}

// Perm is a paid mutator transaction binding the contract method 0xa96bba9d.
//
// Solidity: function perm(bytes32 _h, bytes[5] signs) returns()
func (_IAuth *IAuthSession) Perm(_h [32]byte, signs [5][]byte) (*types.Transaction, error) {
	return _IAuth.Contract.Perm(&_IAuth.TransactOpts, _h, signs)
}

// Perm is a paid mutator transaction binding the contract method 0xa96bba9d.
//
// Solidity: function perm(bytes32 _h, bytes[5] signs) returns()
func (_IAuth *IAuthTransactorSession) Perm(_h [32]byte, signs [5][]byte) (*types.Transaction, error) {
	return _IAuth.Contract.Perm(&_IAuth.TransactOpts, _h, signs)
}

// IInstanceABI is the input ABI used to generate the binding from.
const IInstanceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Alter\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"instances\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IInstanceFuncSigs maps the 4-byte function signature to its string representation.
var IInstanceFuncSigs = map[string]string{
	"3ec7d5b9": "instances(uint8)",
}

// IInstance is an auto generated Go binding around an Ethereum contract.
type IInstance struct {
	IInstanceCaller     // Read-only binding to the contract
	IInstanceTransactor // Write-only binding to the contract
	IInstanceFilterer   // Log filterer for contract events
}

// IInstanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInstanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInstanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInstanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInstanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInstanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInstanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInstanceSession struct {
	Contract     *IInstance        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInstanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInstanceCallerSession struct {
	Contract *IInstanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IInstanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInstanceTransactorSession struct {
	Contract     *IInstanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IInstanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInstanceRaw struct {
	Contract *IInstance // Generic contract binding to access the raw methods on
}

// IInstanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInstanceCallerRaw struct {
	Contract *IInstanceCaller // Generic read-only contract binding to access the raw methods on
}

// IInstanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInstanceTransactorRaw struct {
	Contract *IInstanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInstance creates a new instance of IInstance, bound to a specific deployed contract.
func NewIInstance(address common.Address, backend bind.ContractBackend) (*IInstance, error) {
	contract, err := bindIInstance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInstance{IInstanceCaller: IInstanceCaller{contract: contract}, IInstanceTransactor: IInstanceTransactor{contract: contract}, IInstanceFilterer: IInstanceFilterer{contract: contract}}, nil
}

// NewIInstanceCaller creates a new read-only instance of IInstance, bound to a specific deployed contract.
func NewIInstanceCaller(address common.Address, caller bind.ContractCaller) (*IInstanceCaller, error) {
	contract, err := bindIInstance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInstanceCaller{contract: contract}, nil
}

// NewIInstanceTransactor creates a new write-only instance of IInstance, bound to a specific deployed contract.
func NewIInstanceTransactor(address common.Address, transactor bind.ContractTransactor) (*IInstanceTransactor, error) {
	contract, err := bindIInstance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInstanceTransactor{contract: contract}, nil
}

// NewIInstanceFilterer creates a new log filterer instance of IInstance, bound to a specific deployed contract.
func NewIInstanceFilterer(address common.Address, filterer bind.ContractFilterer) (*IInstanceFilterer, error) {
	contract, err := bindIInstance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInstanceFilterer{contract: contract}, nil
}

// bindIInstance binds a generic wrapper to an already deployed contract.
func bindIInstance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IInstanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInstance *IInstanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInstance.Contract.IInstanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInstance *IInstanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInstance.Contract.IInstanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInstance *IInstanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInstance.Contract.IInstanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInstance *IInstanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInstance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInstance *IInstanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInstance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInstance *IInstanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInstance.Contract.contract.Transact(opts, method, params...)
}

// Instances is a free data retrieval call binding the contract method 0x3ec7d5b9.
//
// Solidity: function instances(uint8 _type) view returns(address)
func (_IInstance *IInstanceCaller) Instances(opts *bind.CallOpts, _type uint8) (common.Address, error) {
	var out []interface{}
	err := _IInstance.contract.Call(opts, &out, "instances", _type)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Instances is a free data retrieval call binding the contract method 0x3ec7d5b9.
//
// Solidity: function instances(uint8 _type) view returns(address)
func (_IInstance *IInstanceSession) Instances(_type uint8) (common.Address, error) {
	return _IInstance.Contract.Instances(&_IInstance.CallOpts, _type)
}

// Instances is a free data retrieval call binding the contract method 0x3ec7d5b9.
//
// Solidity: function instances(uint8 _type) view returns(address)
func (_IInstance *IInstanceCallerSession) Instances(_type uint8) (common.Address, error) {
	return _IInstance.Contract.Instances(&_IInstance.CallOpts, _type)
}

// IInstanceAlterIterator is returned from FilterAlter and is used to iterate over the raw logs and unpacked data for Alter events raised by the IInstance contract.
type IInstanceAlterIterator struct {
	Event *IInstanceAlter // Event containing the contract specifics and raw log

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
func (it *IInstanceAlterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInstanceAlter)
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
		it.Event = new(IInstanceAlter)
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
func (it *IInstanceAlterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInstanceAlterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInstanceAlter represents a Alter event raised by the IInstance contract.
type IInstanceAlter struct {
	Type uint8
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAlter is a free log retrieval operation binding the contract event 0x69da8aaa18d0d64ebdd4d982e4d32ac4aaab1fe0c1034b19846e51a61fcc0f02.
//
// Solidity: event Alter(uint8 _type, address from, address to)
func (_IInstance *IInstanceFilterer) FilterAlter(opts *bind.FilterOpts) (*IInstanceAlterIterator, error) {

	logs, sub, err := _IInstance.contract.FilterLogs(opts, "Alter")
	if err != nil {
		return nil, err
	}
	return &IInstanceAlterIterator{contract: _IInstance.contract, event: "Alter", logs: logs, sub: sub}, nil
}

// WatchAlter is a free log subscription operation binding the contract event 0x69da8aaa18d0d64ebdd4d982e4d32ac4aaab1fe0c1034b19846e51a61fcc0f02.
//
// Solidity: event Alter(uint8 _type, address from, address to)
func (_IInstance *IInstanceFilterer) WatchAlter(opts *bind.WatchOpts, sink chan<- *IInstanceAlter) (event.Subscription, error) {

	logs, sub, err := _IInstance.contract.WatchLogs(opts, "Alter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInstanceAlter)
				if err := _IInstance.contract.UnpackLog(event, "Alter", log); err != nil {
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

// ParseAlter is a log parse operation binding the contract event 0x69da8aaa18d0d64ebdd4d982e4d32ac4aaab1fe0c1034b19846e51a61fcc0f02.
//
// Solidity: event Alter(uint8 _type, address from, address to)
func (_IInstance *IInstanceFilterer) ParseAlter(log types.Log) (*IInstanceAlter, error) {
	event := new(IInstanceAlter)
	if err := _IInstance.contract.UnpackLog(event, "Alter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOwnerABI is the input ABI used to generate the binding from.
const IOwnerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"}]"

// IOwner is an auto generated Go binding around an Ethereum contract.
type IOwner struct {
	IOwnerCaller     // Read-only binding to the contract
	IOwnerTransactor // Write-only binding to the contract
	IOwnerFilterer   // Log filterer for contract events
}

// IOwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOwnerSession struct {
	Contract     *IOwner           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOwnerCallerSession struct {
	Contract *IOwnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IOwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOwnerTransactorSession struct {
	Contract     *IOwnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOwnerRaw struct {
	Contract *IOwner // Generic contract binding to access the raw methods on
}

// IOwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOwnerCallerRaw struct {
	Contract *IOwnerCaller // Generic read-only contract binding to access the raw methods on
}

// IOwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOwnerTransactorRaw struct {
	Contract *IOwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOwner creates a new instance of IOwner, bound to a specific deployed contract.
func NewIOwner(address common.Address, backend bind.ContractBackend) (*IOwner, error) {
	contract, err := bindIOwner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOwner{IOwnerCaller: IOwnerCaller{contract: contract}, IOwnerTransactor: IOwnerTransactor{contract: contract}, IOwnerFilterer: IOwnerFilterer{contract: contract}}, nil
}

// NewIOwnerCaller creates a new read-only instance of IOwner, bound to a specific deployed contract.
func NewIOwnerCaller(address common.Address, caller bind.ContractCaller) (*IOwnerCaller, error) {
	contract, err := bindIOwner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOwnerCaller{contract: contract}, nil
}

// NewIOwnerTransactor creates a new write-only instance of IOwner, bound to a specific deployed contract.
func NewIOwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*IOwnerTransactor, error) {
	contract, err := bindIOwner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOwnerTransactor{contract: contract}, nil
}

// NewIOwnerFilterer creates a new log filterer instance of IOwner, bound to a specific deployed contract.
func NewIOwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*IOwnerFilterer, error) {
	contract, err := bindIOwner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOwnerFilterer{contract: contract}, nil
}

// bindIOwner binds a generic wrapper to an already deployed contract.
func bindIOwner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOwner *IOwnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOwner.Contract.IOwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOwner *IOwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOwner.Contract.IOwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOwner *IOwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOwner.Contract.IOwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOwner *IOwnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOwner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOwner *IOwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOwner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOwner *IOwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOwner.Contract.contract.Transact(opts, method, params...)
}

// IOwnerAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the IOwner contract.
type IOwnerAddOwnerIterator struct {
	Event *IOwnerAddOwner // Event containing the contract specifics and raw log

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
func (it *IOwnerAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOwnerAddOwner)
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
		it.Event = new(IOwnerAddOwner)
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
func (it *IOwnerAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOwnerAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOwnerAddOwner represents a AddOwner event raised by the IOwner contract.
type IOwnerAddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_IOwner *IOwnerFilterer) FilterAddOwner(opts *bind.FilterOpts) (*IOwnerAddOwnerIterator, error) {

	logs, sub, err := _IOwner.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &IOwnerAddOwnerIterator{contract: _IOwner.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_IOwner *IOwnerFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *IOwnerAddOwner) (event.Subscription, error) {

	logs, sub, err := _IOwner.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOwnerAddOwner)
				if err := _IOwner.contract.UnpackLog(event, "AddOwner", log); err != nil {
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

// ParseAddOwner is a log parse operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_IOwner *IOwnerFilterer) ParseAddOwner(log types.Log) (*IOwnerAddOwner, error) {
	event := new(IOwnerAddOwner)
	if err := _IOwner.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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

// ITokenABI is the input ABI used to generate the binding from.
const ITokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"}],\"name\":\"AddT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_mint\",\"type\":\"bool\"}],\"name\":\"addT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isBan\",\"type\":\"bool\"}],\"name\":\"banT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"checkT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getTA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTCnt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"}],\"name\":\"getTI\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ITokenFuncSigs maps the 4-byte function signature to its string representation.
var ITokenFuncSigs = map[string]string{
	"ff4268a4": "addT(address,bool)",
	"83373b36": "banT(address,bool)",
	"81abb8fe": "checkT(uint8)",
	"8bb4a637": "getTA(uint8)",
	"7600b86a": "getTCnt()",
	"2df2685f": "getTI(address)",
}

// IToken is an auto generated Go binding around an Ethereum contract.
type IToken struct {
	ITokenCaller     // Read-only binding to the contract
	ITokenTransactor // Write-only binding to the contract
	ITokenFilterer   // Log filterer for contract events
}

// ITokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenSession struct {
	Contract     *IToken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenCallerSession struct {
	Contract *ITokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ITokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenTransactorSession struct {
	Contract     *ITokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenRaw struct {
	Contract *IToken // Generic contract binding to access the raw methods on
}

// ITokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenCallerRaw struct {
	Contract *ITokenCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenTransactorRaw struct {
	Contract *ITokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIToken creates a new instance of IToken, bound to a specific deployed contract.
func NewIToken(address common.Address, backend bind.ContractBackend) (*IToken, error) {
	contract, err := bindIToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IToken{ITokenCaller: ITokenCaller{contract: contract}, ITokenTransactor: ITokenTransactor{contract: contract}, ITokenFilterer: ITokenFilterer{contract: contract}}, nil
}

// NewITokenCaller creates a new read-only instance of IToken, bound to a specific deployed contract.
func NewITokenCaller(address common.Address, caller bind.ContractCaller) (*ITokenCaller, error) {
	contract, err := bindIToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenCaller{contract: contract}, nil
}

// NewITokenTransactor creates a new write-only instance of IToken, bound to a specific deployed contract.
func NewITokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenTransactor, error) {
	contract, err := bindIToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenTransactor{contract: contract}, nil
}

// NewITokenFilterer creates a new log filterer instance of IToken, bound to a specific deployed contract.
func NewITokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenFilterer, error) {
	contract, err := bindIToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenFilterer{contract: contract}, nil
}

// bindIToken binds a generic wrapper to an already deployed contract.
func bindIToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IToken *ITokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IToken.Contract.ITokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IToken *ITokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.Contract.ITokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IToken *ITokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IToken.Contract.ITokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IToken *ITokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IToken *ITokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IToken *ITokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IToken.Contract.contract.Transact(opts, method, params...)
}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 _ti) view returns(address, bool)
func (_IToken *ITokenCaller) CheckT(opts *bind.CallOpts, _ti uint8) (common.Address, bool, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "checkT", _ti)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 _ti) view returns(address, bool)
func (_IToken *ITokenSession) CheckT(_ti uint8) (common.Address, bool, error) {
	return _IToken.Contract.CheckT(&_IToken.CallOpts, _ti)
}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 _ti) view returns(address, bool)
func (_IToken *ITokenCallerSession) CheckT(_ti uint8) (common.Address, bool, error) {
	return _IToken.Contract.CheckT(&_IToken.CallOpts, _ti)
}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 _ti) view returns(address)
func (_IToken *ITokenCaller) GetTA(opts *bind.CallOpts, _ti uint8) (common.Address, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "getTA", _ti)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 _ti) view returns(address)
func (_IToken *ITokenSession) GetTA(_ti uint8) (common.Address, error) {
	return _IToken.Contract.GetTA(&_IToken.CallOpts, _ti)
}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 _ti) view returns(address)
func (_IToken *ITokenCallerSession) GetTA(_ti uint8) (common.Address, error) {
	return _IToken.Contract.GetTA(&_IToken.CallOpts, _ti)
}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_IToken *ITokenCaller) GetTCnt(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "getTCnt")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_IToken *ITokenSession) GetTCnt() (uint8, error) {
	return _IToken.Contract.GetTCnt(&_IToken.CallOpts)
}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_IToken *ITokenCallerSession) GetTCnt() (uint8, error) {
	return _IToken.Contract.GetTCnt(&_IToken.CallOpts)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address _t) view returns(uint8, bool, bool)
func (_IToken *ITokenCaller) GetTI(opts *bind.CallOpts, _t common.Address) (uint8, bool, bool, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "getTI", _t)

	if err != nil {
		return *new(uint8), *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address _t) view returns(uint8, bool, bool)
func (_IToken *ITokenSession) GetTI(_t common.Address) (uint8, bool, bool, error) {
	return _IToken.Contract.GetTI(&_IToken.CallOpts, _t)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address _t) view returns(uint8, bool, bool)
func (_IToken *ITokenCallerSession) GetTI(_t common.Address) (uint8, bool, bool, error) {
	return _IToken.Contract.GetTI(&_IToken.CallOpts, _t)
}

// AddT is a paid mutator transaction binding the contract method 0xff4268a4.
//
// Solidity: function addT(address _t, bool _mint) returns(uint8)
func (_IToken *ITokenTransactor) AddT(opts *bind.TransactOpts, _t common.Address, _mint bool) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "addT", _t, _mint)
}

// AddT is a paid mutator transaction binding the contract method 0xff4268a4.
//
// Solidity: function addT(address _t, bool _mint) returns(uint8)
func (_IToken *ITokenSession) AddT(_t common.Address, _mint bool) (*types.Transaction, error) {
	return _IToken.Contract.AddT(&_IToken.TransactOpts, _t, _mint)
}

// AddT is a paid mutator transaction binding the contract method 0xff4268a4.
//
// Solidity: function addT(address _t, bool _mint) returns(uint8)
func (_IToken *ITokenTransactorSession) AddT(_t common.Address, _mint bool) (*types.Transaction, error) {
	return _IToken.Contract.AddT(&_IToken.TransactOpts, _t, _mint)
}

// BanT is a paid mutator transaction binding the contract method 0x83373b36.
//
// Solidity: function banT(address _t, bool _isBan) returns()
func (_IToken *ITokenTransactor) BanT(opts *bind.TransactOpts, _t common.Address, _isBan bool) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "banT", _t, _isBan)
}

// BanT is a paid mutator transaction binding the contract method 0x83373b36.
//
// Solidity: function banT(address _t, bool _isBan) returns()
func (_IToken *ITokenSession) BanT(_t common.Address, _isBan bool) (*types.Transaction, error) {
	return _IToken.Contract.BanT(&_IToken.TransactOpts, _t, _isBan)
}

// BanT is a paid mutator transaction binding the contract method 0x83373b36.
//
// Solidity: function banT(address _t, bool _isBan) returns()
func (_IToken *ITokenTransactorSession) BanT(_t common.Address, _isBan bool) (*types.Transaction, error) {
	return _IToken.Contract.BanT(&_IToken.TransactOpts, _t, _isBan)
}

// ITokenAddTIterator is returned from FilterAddT and is used to iterate over the raw logs and unpacked data for AddT events raised by the IToken contract.
type ITokenAddTIterator struct {
	Event *ITokenAddT // Event containing the contract specifics and raw log

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
func (it *ITokenAddTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenAddT)
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
		it.Event = new(ITokenAddT)
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
func (it *ITokenAddTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenAddTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenAddT represents a AddT event raised by the IToken contract.
type ITokenAddT struct {
	T      common.Address
	TIndex uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddT is a free log retrieval operation binding the contract event 0xf10191767d4ea3fc26b057d307336c7e8df8880bb07ddaebe4e853db6b5fd936.
//
// Solidity: event AddT(address t, uint8 tIndex)
func (_IToken *ITokenFilterer) FilterAddT(opts *bind.FilterOpts) (*ITokenAddTIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return &ITokenAddTIterator{contract: _IToken.contract, event: "AddT", logs: logs, sub: sub}, nil
}

// WatchAddT is a free log subscription operation binding the contract event 0xf10191767d4ea3fc26b057d307336c7e8df8880bb07ddaebe4e853db6b5fd936.
//
// Solidity: event AddT(address t, uint8 tIndex)
func (_IToken *ITokenFilterer) WatchAddT(opts *bind.WatchOpts, sink chan<- *ITokenAddT) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenAddT)
				if err := _IToken.contract.UnpackLog(event, "AddT", log); err != nil {
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

// ParseAddT is a log parse operation binding the contract event 0xf10191767d4ea3fc26b057d307336c7e8df8880bb07ddaebe4e853db6b5fd936.
//
// Solidity: event AddT(address t, uint8 tIndex)
func (_IToken *ITokenFilterer) ParseAddT(log types.Log) (*ITokenAddT, error) {
	event := new(ITokenAddT)
	if err := _IToken.contract.UnpackLog(event, "AddT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenGetterABI is the input ABI used to generate the binding from.
const ITokenGetterABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"checkT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getTA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTCnt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"}],\"name\":\"getTI\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ITokenGetterFuncSigs maps the 4-byte function signature to its string representation.
var ITokenGetterFuncSigs = map[string]string{
	"81abb8fe": "checkT(uint8)",
	"8bb4a637": "getTA(uint8)",
	"7600b86a": "getTCnt()",
	"2df2685f": "getTI(address)",
}

// ITokenGetter is an auto generated Go binding around an Ethereum contract.
type ITokenGetter struct {
	ITokenGetterCaller     // Read-only binding to the contract
	ITokenGetterTransactor // Write-only binding to the contract
	ITokenGetterFilterer   // Log filterer for contract events
}

// ITokenGetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenGetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenGetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenGetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenGetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenGetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenGetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenGetterSession struct {
	Contract     *ITokenGetter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenGetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenGetterCallerSession struct {
	Contract *ITokenGetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ITokenGetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenGetterTransactorSession struct {
	Contract     *ITokenGetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ITokenGetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenGetterRaw struct {
	Contract *ITokenGetter // Generic contract binding to access the raw methods on
}

// ITokenGetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenGetterCallerRaw struct {
	Contract *ITokenGetterCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenGetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenGetterTransactorRaw struct {
	Contract *ITokenGetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenGetter creates a new instance of ITokenGetter, bound to a specific deployed contract.
func NewITokenGetter(address common.Address, backend bind.ContractBackend) (*ITokenGetter, error) {
	contract, err := bindITokenGetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenGetter{ITokenGetterCaller: ITokenGetterCaller{contract: contract}, ITokenGetterTransactor: ITokenGetterTransactor{contract: contract}, ITokenGetterFilterer: ITokenGetterFilterer{contract: contract}}, nil
}

// NewITokenGetterCaller creates a new read-only instance of ITokenGetter, bound to a specific deployed contract.
func NewITokenGetterCaller(address common.Address, caller bind.ContractCaller) (*ITokenGetterCaller, error) {
	contract, err := bindITokenGetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenGetterCaller{contract: contract}, nil
}

// NewITokenGetterTransactor creates a new write-only instance of ITokenGetter, bound to a specific deployed contract.
func NewITokenGetterTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenGetterTransactor, error) {
	contract, err := bindITokenGetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenGetterTransactor{contract: contract}, nil
}

// NewITokenGetterFilterer creates a new log filterer instance of ITokenGetter, bound to a specific deployed contract.
func NewITokenGetterFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenGetterFilterer, error) {
	contract, err := bindITokenGetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenGetterFilterer{contract: contract}, nil
}

// bindITokenGetter binds a generic wrapper to an already deployed contract.
func bindITokenGetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenGetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenGetter *ITokenGetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenGetter.Contract.ITokenGetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenGetter *ITokenGetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenGetter.Contract.ITokenGetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenGetter *ITokenGetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenGetter.Contract.ITokenGetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenGetter *ITokenGetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenGetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenGetter *ITokenGetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenGetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenGetter *ITokenGetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenGetter.Contract.contract.Transact(opts, method, params...)
}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 _ti) view returns(address, bool)
func (_ITokenGetter *ITokenGetterCaller) CheckT(opts *bind.CallOpts, _ti uint8) (common.Address, bool, error) {
	var out []interface{}
	err := _ITokenGetter.contract.Call(opts, &out, "checkT", _ti)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 _ti) view returns(address, bool)
func (_ITokenGetter *ITokenGetterSession) CheckT(_ti uint8) (common.Address, bool, error) {
	return _ITokenGetter.Contract.CheckT(&_ITokenGetter.CallOpts, _ti)
}

// CheckT is a free data retrieval call binding the contract method 0x81abb8fe.
//
// Solidity: function checkT(uint8 _ti) view returns(address, bool)
func (_ITokenGetter *ITokenGetterCallerSession) CheckT(_ti uint8) (common.Address, bool, error) {
	return _ITokenGetter.Contract.CheckT(&_ITokenGetter.CallOpts, _ti)
}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 _ti) view returns(address)
func (_ITokenGetter *ITokenGetterCaller) GetTA(opts *bind.CallOpts, _ti uint8) (common.Address, error) {
	var out []interface{}
	err := _ITokenGetter.contract.Call(opts, &out, "getTA", _ti)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 _ti) view returns(address)
func (_ITokenGetter *ITokenGetterSession) GetTA(_ti uint8) (common.Address, error) {
	return _ITokenGetter.Contract.GetTA(&_ITokenGetter.CallOpts, _ti)
}

// GetTA is a free data retrieval call binding the contract method 0x8bb4a637.
//
// Solidity: function getTA(uint8 _ti) view returns(address)
func (_ITokenGetter *ITokenGetterCallerSession) GetTA(_ti uint8) (common.Address, error) {
	return _ITokenGetter.Contract.GetTA(&_ITokenGetter.CallOpts, _ti)
}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_ITokenGetter *ITokenGetterCaller) GetTCnt(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ITokenGetter.contract.Call(opts, &out, "getTCnt")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_ITokenGetter *ITokenGetterSession) GetTCnt() (uint8, error) {
	return _ITokenGetter.Contract.GetTCnt(&_ITokenGetter.CallOpts)
}

// GetTCnt is a free data retrieval call binding the contract method 0x7600b86a.
//
// Solidity: function getTCnt() view returns(uint8)
func (_ITokenGetter *ITokenGetterCallerSession) GetTCnt() (uint8, error) {
	return _ITokenGetter.Contract.GetTCnt(&_ITokenGetter.CallOpts)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address _t) view returns(uint8, bool, bool)
func (_ITokenGetter *ITokenGetterCaller) GetTI(opts *bind.CallOpts, _t common.Address) (uint8, bool, bool, error) {
	var out []interface{}
	err := _ITokenGetter.contract.Call(opts, &out, "getTI", _t)

	if err != nil {
		return *new(uint8), *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address _t) view returns(uint8, bool, bool)
func (_ITokenGetter *ITokenGetterSession) GetTI(_t common.Address) (uint8, bool, bool, error) {
	return _ITokenGetter.Contract.GetTI(&_ITokenGetter.CallOpts, _t)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address _t) view returns(uint8, bool, bool)
func (_ITokenGetter *ITokenGetterCallerSession) GetTI(_t common.Address) (uint8, bool, bool, error) {
	return _ITokenGetter.Contract.GetTI(&_ITokenGetter.CallOpts, _t)
}

// ITokenSetterABI is the input ABI used to generate the binding from.
const ITokenSetterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"}],\"name\":\"AddT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_mint\",\"type\":\"bool\"}],\"name\":\"addT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isBan\",\"type\":\"bool\"}],\"name\":\"banT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ITokenSetterFuncSigs maps the 4-byte function signature to its string representation.
var ITokenSetterFuncSigs = map[string]string{
	"ff4268a4": "addT(address,bool)",
	"83373b36": "banT(address,bool)",
}

// ITokenSetter is an auto generated Go binding around an Ethereum contract.
type ITokenSetter struct {
	ITokenSetterCaller     // Read-only binding to the contract
	ITokenSetterTransactor // Write-only binding to the contract
	ITokenSetterFilterer   // Log filterer for contract events
}

// ITokenSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenSetterSession struct {
	Contract     *ITokenSetter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenSetterCallerSession struct {
	Contract *ITokenSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ITokenSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenSetterTransactorSession struct {
	Contract     *ITokenSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ITokenSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenSetterRaw struct {
	Contract *ITokenSetter // Generic contract binding to access the raw methods on
}

// ITokenSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenSetterCallerRaw struct {
	Contract *ITokenSetterCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenSetterTransactorRaw struct {
	Contract *ITokenSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenSetter creates a new instance of ITokenSetter, bound to a specific deployed contract.
func NewITokenSetter(address common.Address, backend bind.ContractBackend) (*ITokenSetter, error) {
	contract, err := bindITokenSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenSetter{ITokenSetterCaller: ITokenSetterCaller{contract: contract}, ITokenSetterTransactor: ITokenSetterTransactor{contract: contract}, ITokenSetterFilterer: ITokenSetterFilterer{contract: contract}}, nil
}

// NewITokenSetterCaller creates a new read-only instance of ITokenSetter, bound to a specific deployed contract.
func NewITokenSetterCaller(address common.Address, caller bind.ContractCaller) (*ITokenSetterCaller, error) {
	contract, err := bindITokenSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenSetterCaller{contract: contract}, nil
}

// NewITokenSetterTransactor creates a new write-only instance of ITokenSetter, bound to a specific deployed contract.
func NewITokenSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenSetterTransactor, error) {
	contract, err := bindITokenSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenSetterTransactor{contract: contract}, nil
}

// NewITokenSetterFilterer creates a new log filterer instance of ITokenSetter, bound to a specific deployed contract.
func NewITokenSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenSetterFilterer, error) {
	contract, err := bindITokenSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenSetterFilterer{contract: contract}, nil
}

// bindITokenSetter binds a generic wrapper to an already deployed contract.
func bindITokenSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenSetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenSetter *ITokenSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenSetter.Contract.ITokenSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenSetter *ITokenSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenSetter.Contract.ITokenSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenSetter *ITokenSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenSetter.Contract.ITokenSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenSetter *ITokenSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenSetter *ITokenSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenSetter *ITokenSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenSetter.Contract.contract.Transact(opts, method, params...)
}

// AddT is a paid mutator transaction binding the contract method 0xff4268a4.
//
// Solidity: function addT(address _t, bool _mint) returns(uint8)
func (_ITokenSetter *ITokenSetterTransactor) AddT(opts *bind.TransactOpts, _t common.Address, _mint bool) (*types.Transaction, error) {
	return _ITokenSetter.contract.Transact(opts, "addT", _t, _mint)
}

// AddT is a paid mutator transaction binding the contract method 0xff4268a4.
//
// Solidity: function addT(address _t, bool _mint) returns(uint8)
func (_ITokenSetter *ITokenSetterSession) AddT(_t common.Address, _mint bool) (*types.Transaction, error) {
	return _ITokenSetter.Contract.AddT(&_ITokenSetter.TransactOpts, _t, _mint)
}

// AddT is a paid mutator transaction binding the contract method 0xff4268a4.
//
// Solidity: function addT(address _t, bool _mint) returns(uint8)
func (_ITokenSetter *ITokenSetterTransactorSession) AddT(_t common.Address, _mint bool) (*types.Transaction, error) {
	return _ITokenSetter.Contract.AddT(&_ITokenSetter.TransactOpts, _t, _mint)
}

// BanT is a paid mutator transaction binding the contract method 0x83373b36.
//
// Solidity: function banT(address _t, bool _isBan) returns()
func (_ITokenSetter *ITokenSetterTransactor) BanT(opts *bind.TransactOpts, _t common.Address, _isBan bool) (*types.Transaction, error) {
	return _ITokenSetter.contract.Transact(opts, "banT", _t, _isBan)
}

// BanT is a paid mutator transaction binding the contract method 0x83373b36.
//
// Solidity: function banT(address _t, bool _isBan) returns()
func (_ITokenSetter *ITokenSetterSession) BanT(_t common.Address, _isBan bool) (*types.Transaction, error) {
	return _ITokenSetter.Contract.BanT(&_ITokenSetter.TransactOpts, _t, _isBan)
}

// BanT is a paid mutator transaction binding the contract method 0x83373b36.
//
// Solidity: function banT(address _t, bool _isBan) returns()
func (_ITokenSetter *ITokenSetterTransactorSession) BanT(_t common.Address, _isBan bool) (*types.Transaction, error) {
	return _ITokenSetter.Contract.BanT(&_ITokenSetter.TransactOpts, _t, _isBan)
}

// ITokenSetterAddTIterator is returned from FilterAddT and is used to iterate over the raw logs and unpacked data for AddT events raised by the ITokenSetter contract.
type ITokenSetterAddTIterator struct {
	Event *ITokenSetterAddT // Event containing the contract specifics and raw log

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
func (it *ITokenSetterAddTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenSetterAddT)
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
		it.Event = new(ITokenSetterAddT)
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
func (it *ITokenSetterAddTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenSetterAddTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenSetterAddT represents a AddT event raised by the ITokenSetter contract.
type ITokenSetterAddT struct {
	T      common.Address
	TIndex uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddT is a free log retrieval operation binding the contract event 0xf10191767d4ea3fc26b057d307336c7e8df8880bb07ddaebe4e853db6b5fd936.
//
// Solidity: event AddT(address t, uint8 tIndex)
func (_ITokenSetter *ITokenSetterFilterer) FilterAddT(opts *bind.FilterOpts) (*ITokenSetterAddTIterator, error) {

	logs, sub, err := _ITokenSetter.contract.FilterLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return &ITokenSetterAddTIterator{contract: _ITokenSetter.contract, event: "AddT", logs: logs, sub: sub}, nil
}

// WatchAddT is a free log subscription operation binding the contract event 0xf10191767d4ea3fc26b057d307336c7e8df8880bb07ddaebe4e853db6b5fd936.
//
// Solidity: event AddT(address t, uint8 tIndex)
func (_ITokenSetter *ITokenSetterFilterer) WatchAddT(opts *bind.WatchOpts, sink chan<- *ITokenSetterAddT) (event.Subscription, error) {

	logs, sub, err := _ITokenSetter.contract.WatchLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenSetterAddT)
				if err := _ITokenSetter.contract.UnpackLog(event, "AddT", log); err != nil {
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

// ParseAddT is a log parse operation binding the contract event 0xf10191767d4ea3fc26b057d307336c7e8df8880bb07ddaebe4e853db6b5fd936.
//
// Solidity: event AddT(address t, uint8 tIndex)
func (_ITokenSetter *ITokenSetterFilterer) ParseAddT(log types.Log) (*ITokenSetterAddT, error) {
	event := new(ITokenSetterAddT)
	if err := _ITokenSetter.contract.UnpackLog(event, "AddT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerABI is the input ABI used to generate the binding from.
const OwnerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OwnerFuncSigs maps the 4-byte function signature to its string representation.
var OwnerFuncSigs = map[string]string{
	"4bf1b134": "add(address,bool,bytes[5])",
	"de9375f2": "auth()",
	"022914a7": "owners(address)",
}

// OwnerBin is the compiled bytecode used for deploying new contracts.
var OwnerBin = "0x608060405234801561001057600080fd5b5060405161054f38038061054f83398101604081905261002f91610054565b600180546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b6104bc806100936000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063022914a7146100465780634bf1b1341461007e578063de9375f214610093575b600080fd5b610069610054366004610257565b60006020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61009161008c3660046102e9565b6100be565b005b6001546100a6906001600160a01b031681565b6040516001600160a01b039091168152602001610075565b823b600081900361010a5760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b604482015260640160405180910390fd5b6040516bffffffffffffffffffffffff1930606090811b821660208401526218591960ea1b603484015286901b16603782015283151560f81b604b820152600090604c0160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d9061019790849087906004016103ff565b600060405180830381600087803b1580156101b157600080fd5b505af11580156101c5573d6000803e3d6000fd5b5050604080516001600160a01b038916815287151560208201527f938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807935001905060405180910390a15050506001600160a01b03919091166000908152602081905260409020805460ff1916911515919091179055565b80356001600160a01b038116811461025257600080fd5b919050565b60006020828403121561026957600080fd5b6102728261023b565b9392505050565b634e487b7160e01b600052604160045260246000fd5b60405160a0810167ffffffffffffffff811182821017156102b2576102b2610279565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156102e1576102e1610279565b604052919050565b6000806000606084860312156102fe57600080fd5b6103078461023b565b9250602080850135801515811461031d57600080fd5b9250604085013567ffffffffffffffff8082111561033a57600080fd5b8187019150601f888184011261034f57600080fd5b61035761028f565b8060a085018b81111561036957600080fd5b855b818110156103ed578035868111156103835760008081fd5b87018581018e136103945760008081fd5b8035878111156103a6576103a6610279565b6103b7818801601f19168b016102b8565b8181528f8b8385010111156103cc5760008081fd5b818b84018c83013760009181018b019190915285525092870192870161036b565b50508096505050505050509250925092565b8281526040602080830182905260009160e08401919084018584805b600581101561047857878603603f1901845282518051808852835b81811015610451578281018801518982018901528701610436565b508781018701849052601f01601f191690960185019550928401929184019160010161041b565b50939897505050505050505056fea264697066735822122051c6cfc7790ace9fc35fd7d24035222c18b0fbba24d8c520a60c25602823548b64736f6c63430008100033"

// DeployOwner deploys a new Ethereum contract, binding an instance of Owner to it.
func DeployOwner(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address) (common.Address, *types.Transaction, *Owner, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnerBin), backend, _a)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owner{OwnerCaller: OwnerCaller{contract: contract}, OwnerTransactor: OwnerTransactor{contract: contract}, OwnerFilterer: OwnerFilterer{contract: contract}}, nil
}

// Owner is an auto generated Go binding around an Ethereum contract.
type Owner struct {
	OwnerCaller     // Read-only binding to the contract
	OwnerTransactor // Write-only binding to the contract
	OwnerFilterer   // Log filterer for contract events
}

// OwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnerSession struct {
	Contract     *Owner            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnerCallerSession struct {
	Contract *OwnerCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnerTransactorSession struct {
	Contract     *OwnerTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnerRaw struct {
	Contract *Owner // Generic contract binding to access the raw methods on
}

// OwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnerCallerRaw struct {
	Contract *OwnerCaller // Generic read-only contract binding to access the raw methods on
}

// OwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnerTransactorRaw struct {
	Contract *OwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwner creates a new instance of Owner, bound to a specific deployed contract.
func NewOwner(address common.Address, backend bind.ContractBackend) (*Owner, error) {
	contract, err := bindOwner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owner{OwnerCaller: OwnerCaller{contract: contract}, OwnerTransactor: OwnerTransactor{contract: contract}, OwnerFilterer: OwnerFilterer{contract: contract}}, nil
}

// NewOwnerCaller creates a new read-only instance of Owner, bound to a specific deployed contract.
func NewOwnerCaller(address common.Address, caller bind.ContractCaller) (*OwnerCaller, error) {
	contract, err := bindOwner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerCaller{contract: contract}, nil
}

// NewOwnerTransactor creates a new write-only instance of Owner, bound to a specific deployed contract.
func NewOwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnerTransactor, error) {
	contract, err := bindOwner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerTransactor{contract: contract}, nil
}

// NewOwnerFilterer creates a new log filterer instance of Owner, bound to a specific deployed contract.
func NewOwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnerFilterer, error) {
	contract, err := bindOwner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnerFilterer{contract: contract}, nil
}

// bindOwner binds a generic wrapper to an already deployed contract.
func bindOwner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owner *OwnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owner.Contract.OwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owner *OwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owner.Contract.OwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owner *OwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owner.Contract.OwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owner *OwnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owner *OwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owner *OwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owner.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Owner *OwnerCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Owner.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Owner *OwnerSession) Auth() (common.Address, error) {
	return _Owner.Contract.Auth(&_Owner.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Owner *OwnerCallerSession) Auth() (common.Address, error) {
	return _Owner.Contract.Auth(&_Owner.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Owner *OwnerCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Owner.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Owner *OwnerSession) Owners(arg0 common.Address) (bool, error) {
	return _Owner.Contract.Owners(&_Owner.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Owner *OwnerCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _Owner.Contract.Owners(&_Owner.CallOpts, arg0)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Owner *OwnerTransactor) Add(opts *bind.TransactOpts, _a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Owner.contract.Transact(opts, "add", _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Owner *OwnerSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Owner.Contract.Add(&_Owner.TransactOpts, _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Owner *OwnerTransactorSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Owner.Contract.Add(&_Owner.TransactOpts, _a, isSet, signs)
}

// OwnerAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the Owner contract.
type OwnerAddOwnerIterator struct {
	Event *OwnerAddOwner // Event containing the contract specifics and raw log

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
func (it *OwnerAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerAddOwner)
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
		it.Event = new(OwnerAddOwner)
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
func (it *OwnerAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerAddOwner represents a AddOwner event raised by the Owner contract.
type OwnerAddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Owner *OwnerFilterer) FilterAddOwner(opts *bind.FilterOpts) (*OwnerAddOwnerIterator, error) {

	logs, sub, err := _Owner.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &OwnerAddOwnerIterator{contract: _Owner.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Owner *OwnerFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *OwnerAddOwner) (event.Subscription, error) {

	logs, sub, err := _Owner.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerAddOwner)
				if err := _Owner.contract.UnpackLog(event, "AddOwner", log); err != nil {
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

// ParseAddOwner is a log parse operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Owner *OwnerFilterer) ParseAddOwner(log types.Log) (*OwnerAddOwner, error) {
	event := new(OwnerAddOwner)
	if err := _Owner.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeABI is the input ABI used to generate the binding from.
const PledgeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_inst\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"}],\"name\":\"AddT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"addT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inst\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pledge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rs\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"tInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pledge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rs\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lock\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PledgeFuncSigs maps the 4-byte function signature to its string representation.
var PledgeFuncSigs = map[string]string{
	"4bf1b134": "add(address,bool,bytes[5])",
	"724a84a8": "addT(uint8)",
	"de9375f2": "auth()",
	"fc3ba0ad": "balanceOf(uint64,uint8)",
	"eb395fde": "getPledge(uint8)",
	"bd6b2222": "inst()",
	"022914a7": "owners(address)",
	"d23f7482": "pledge(uint64,uint256)",
	"14f732de": "rewards(uint64,uint8)",
	"041ac896": "tInfo(uint8)",
	"c21a43e4": "totalPledge()",
	"54fd4d50": "version()",
	"5affa0f3": "withdraw(uint64,uint8,uint256,uint256)",
}

// PledgeBin is the compiled bytecode used for deploying new contracts.
var PledgeBin = "0x60806040526001805461ffff60a81b1916600160a91b17905534801561002457600080fd5b506040516114a73803806114a783398101604081905261004391610090565b600180546001600160a01b039384166001600160a01b031991821617909155600280549290931691161790556100c3565b80516001600160a01b038116811461008b57600080fd5b919050565b600080604083850312156100a357600080fd5b6100ac83610074565b91506100ba60208401610074565b90509250929050565b6113d5806100d26000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063724a84a81161008c578063d23f748211610066578063d23f748214610246578063de9375f214610259578063eb395fde1461026c578063fc3ba0ad1461027f57600080fd5b8063724a84a8146101ff578063bd6b222214610212578063c21a43e41461023d57600080fd5b8063022914a7146100d4578063041ac8961461010c57806314f732de146101615780634bf1b134146101a157806354fd4d50146101b65780635affa0f3146101de575b600080fd5b6100f76100e2366004610f52565b60006020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61014161011a366004610f8c565b60046020526000908152604090208054600182015460028301546003909301549192909184565b604080519485526020850193909352918301526060820152608001610103565b61014161016f366004610fbe565b600560209081526000928352604080842090915290825290208054600182015460028301546003909301549192909184565b6101b46101af36600461105f565b610292565b005b6001546101cb90600160a81b900461ffff1681565b60405161ffff9091168152602001610103565b6101f16101ec366004611176565b610410565b604051908152602001610103565b6101b461020d366004610f8c565b610603565b600254610225906001600160a01b031681565b6040516001600160a01b039091168152602001610103565b6101f160035481565b6101b46102543660046111b8565b610755565b600154610225906001600160a01b031681565b6101f161027a366004610f8c565b6108ff565b6101f161028d366004610fbe565b610910565b823b60008190036102df5760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b60448201526064015b60405180910390fd5b6040516bffffffffffffffffffffffff1930606090811b821660208401526218591960ea1b603484015286901b16603782015283151560f81b604b820152600090604c0160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d9061036c9084908790600401611206565b600060405180830381600087803b15801561038657600080fd5b505af115801561039a573d6000803e3d6000fd5b5050604080516001600160a01b038916815287151560208201527f938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807935001905060405180910390a15050506001600160a01b03919091166000908152602081905260409020805460ff1916911515919091179055565b3360009081526020819052604081205460ff1661043f5760405162461bcd60e51b81526004016102d690611275565b600154600160a01b900460ff166104875760405162461bcd60e51b815260206004820152600c60248201526b3430b9b713ba103a37b5b2b760a11b60448201526064016102d6565b6001600160401b03851660009081526005602090815260408083208380529091529020600101548015806104bb5750600354155b156104ca5760009150506105fb565b8460ff166000036105135760005b60015460ff600160a01b9091048116908216101561050d576104fb828289610ab4565b80610505816112ae565b9150506104d8565b5061051e565b61051e818688610ab4565b6001600160401b038616600090815260056020908152604080832060ff891684529091528120600101546105539085906112cd565b905080851080156105645750600085115b1561056c5750835b60ff8616600090815260046020526040812060010180548392906105919084906112cd565b90915550506001600160401b038716600090815260056020908152604080832060ff8a168452909152812060010180548392906105cf9084906112cd565b909155505060ff86166000036105f75780600360008282546105f191906112cd565b90915550505b9150505b949350505050565b3360009081526020819052604090205460ff166106325760405162461bcd60e51b81526004016102d690611275565b60015460ff828116600160a01b909204161461067a5760405162461bcd60e51b8152602060048201526007602482015266696c6c6520746960c81b60448201526064016102d6565b600061068582610c9a565b60ff80841660009081527f05b8ccbb9d4d8fb16ea74ce3c29a41f1b461fbdaff4714a0d9a8eb05499746bc6020908152604080832060019081018690556004909252909120810183905580549293509182916014916106ed918491600160a01b9004166112e0565b92506101000a81548160ff021916908360ff1602179055508160ff166000036107165760038190555b6040805182815260ff841660208201527f8ce8baccab364cf4dd55bc29e9ca458faa5067327fb854f45973d4bdcc5b8cf9910160405180910390a15050565b3360009081526020819052604090205460ff166107845760405162461bcd60e51b81526004016102d690611275565b600154600160a01b900460ff166107cc5760405162461bcd60e51b815260206004820152600c60248201526b3430b9b713ba103a37b5b2b760a11b60448201526064016102d6565b6001600160401b0382166000908152600560209081526040808320838052909152812060010154905b60015460ff600160a01b9091048116908216101561082a57610818828286610ab4565b80610822816112ae565b9150506107f5565b50600080805260046020527f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ed80548492906108669084906112f9565b90915550506001600160401b0383166000908152600560209081526040808320838052909152812060010180548492906108a19084906112f9565b90915550506001600160401b0383166000908152600560209081526040808320838052909152812060020180548492906108dc9084906112f9565b9250508190555081600360008282546108f591906112f9565b9091555050505050565b600061090a82610c9a565b92915050565b6001600160401b03821660009081526005602090815260408083208380528252808320815160808101835281548152600182015493810184905260028201549281019290925260030154606082015290820361097057600091505061090a565b60208082015160ff85166000908152600483526040808220815160808101835281548082526001830154968201969096526002820154928101929092526003015460608201529192906109c287610c9a565b90508260200151811180156109d957506000600354115b15610a195760035460208401516109f090836112cd565b610a0290670de0b6b3a764000061130c565b610a0c919061132b565b610a1690836112f9565b91505b6001600160401b038816600090815260056020908152604080832060ff8b16845282529182902082516080810184528154808252600183015493820193909352600282015493810193909352600301546060830152670de0b6b3a7640000908690610a8490866112cd565b610a8e919061130c565b610a98919061132b565b8160200151610aa791906112f9565b9998505050505050505050565b6000610abf83610c9a565b60ff84166000908152600460205260409020600101549091508110801590610ae957506000600354115b15610c945760035460ff8416600090815260046020526040902060010154610b1190836112cd565b610b2390670de0b6b3a764000061130c565b610b2d919061132b565b60ff841660009081526004602052604081208054909190610b4f9084906112f9565b909155505060ff83166000818152600460208181526040808420600181018790556001600160401b03881685526005835281852095855294825283205491905291549091610b9c916112cd565b9050670de0b6b3a7640000610bb1868361130c565b610bbb919061132b565b6001600160401b038416600090815260056020908152604080832060ff89168452909152812060010180549293508392909190610bf99084906112f9565b90915550506001600160401b038316600090815260056020908152604080832060ff8816845290915281206003018054839290610c379084906112f9565b909155505060ff8416600003610c5f578060036000828254610c5991906112f9565b90915550505b5060ff83166000818152600460209081526040808320546001600160401b038716845260058352818420948452939091529020555b50505050565b600254604051633ec7d5b960e01b8152600760048201526000916001600160a01b03169082908290633ec7d5b990602401602060405180830381865afa158015610ce8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d0c919061134d565b604051638bb4a63760e01b815260ff861660048201526001600160a01b039190911690638bb4a63790602401602060405180830381865afa158015610d55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d79919061134d565b604080518082018252601281527162616c616e63654f6628616464726573732960701b60209091015251633ec7d5b960e01b81526005600482015290915060009081906001600160a01b03808516917f70a08231b98ef4ca268c9cc3f6b4590e4bfec28280db06bb5d45e689f2a360be91871690633ec7d5b990602401602060405180830381865afa158015610e13573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e37919061134d565b6040516001600160a01b03909116602482015260440160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051610e8b919061136a565b600060405180830381855afa9150503d8060008114610ec6576040519150601f19603f3d011682016040523d82523d6000602084013e610ecb565b606091505b5091509150818015610edf57506020815110155b610f1c5760405162461bcd60e51b815260206004820152600e60248201526d39ba30ba34b1b1b0b6361032b93960911b60448201526064016102d6565b80806020019051810190610f309190611386565b9695505050505050565b6001600160a01b0381168114610f4f57600080fd5b50565b600060208284031215610f6457600080fd5b8135610f6f81610f3a565b9392505050565b803560ff81168114610f8757600080fd5b919050565b600060208284031215610f9e57600080fd5b610f6f82610f76565b80356001600160401b0381168114610f8757600080fd5b60008060408385031215610fd157600080fd5b610fda83610fa7565b9150610fe860208401610f76565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b038111828210171561102957611029610ff1565b60405290565b604051601f8201601f191681016001600160401b038111828210171561105757611057610ff1565b604052919050565b60008060006060848603121561107457600080fd5b833561107f81610f3a565b9250602084810135801515811461109557600080fd5b925060408501356001600160401b03808211156110b157600080fd5b8187019150601f88818401126110c657600080fd5b6110ce611007565b8060a085018b8111156110e057600080fd5b855b81811015611164578035868111156110fa5760008081fd5b87018581018e1361110b5760008081fd5b80358781111561111d5761111d610ff1565b61112e818801601f19168b0161102f565b8181528f8b8385010111156111435760008081fd5b818b84018c83013760009181018b01919091528552509287019287016110e2565b50508096505050505050509250925092565b6000806000806080858703121561118c57600080fd5b61119585610fa7565b93506111a360208601610f76565b93969395505050506040820135916060013590565b600080604083850312156111cb57600080fd5b6111d483610fa7565b946020939093013593505050565b60005b838110156111fd5781810151838201526020016111e5565b50506000910152565b8281526040602080830182905260009160e084019190840185845b600581101561126857868503603f1901835281518051808752611249818789018885016111e2565b601f01601f191695909501840194509183019190830190600101611221565b5092979650505050505050565b6020808252600990820152683737ba1037bbb732b960b91b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b600060ff821660ff81036112c4576112c4611298565b60010192915050565b8181038181111561090a5761090a611298565b60ff818116838216019081111561090a5761090a611298565b8082018082111561090a5761090a611298565b600081600019048311821515161561132657611326611298565b500290565b60008261134857634e487b7160e01b600052601260045260246000fd5b500490565b60006020828403121561135f57600080fd5b8151610f6f81610f3a565b6000825161137c8184602087016111e2565b9190910192915050565b60006020828403121561139857600080fd5b505191905056fea264697066735822122085d6a31bd8b338f23b6affa067cd2d32c3a64625f5ad8a31348bd65226ddd14b64736f6c63430008100033"

// DeployPledge deploys a new Ethereum contract, binding an instance of Pledge to it.
func DeployPledge(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address, _inst common.Address) (common.Address, *types.Transaction, *Pledge, error) {
	parsed, err := abi.JSON(strings.NewReader(PledgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PledgeBin), backend, _a, _inst)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pledge{PledgeCaller: PledgeCaller{contract: contract}, PledgeTransactor: PledgeTransactor{contract: contract}, PledgeFilterer: PledgeFilterer{contract: contract}}, nil
}

// Pledge is an auto generated Go binding around an Ethereum contract.
type Pledge struct {
	PledgeCaller     // Read-only binding to the contract
	PledgeTransactor // Write-only binding to the contract
	PledgeFilterer   // Log filterer for contract events
}

// PledgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PledgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PledgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PledgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PledgeSession struct {
	Contract     *Pledge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PledgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PledgeCallerSession struct {
	Contract *PledgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PledgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PledgeTransactorSession struct {
	Contract     *PledgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PledgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PledgeRaw struct {
	Contract *Pledge // Generic contract binding to access the raw methods on
}

// PledgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PledgeCallerRaw struct {
	Contract *PledgeCaller // Generic read-only contract binding to access the raw methods on
}

// PledgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PledgeTransactorRaw struct {
	Contract *PledgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPledge creates a new instance of Pledge, bound to a specific deployed contract.
func NewPledge(address common.Address, backend bind.ContractBackend) (*Pledge, error) {
	contract, err := bindPledge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pledge{PledgeCaller: PledgeCaller{contract: contract}, PledgeTransactor: PledgeTransactor{contract: contract}, PledgeFilterer: PledgeFilterer{contract: contract}}, nil
}

// NewPledgeCaller creates a new read-only instance of Pledge, bound to a specific deployed contract.
func NewPledgeCaller(address common.Address, caller bind.ContractCaller) (*PledgeCaller, error) {
	contract, err := bindPledge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PledgeCaller{contract: contract}, nil
}

// NewPledgeTransactor creates a new write-only instance of Pledge, bound to a specific deployed contract.
func NewPledgeTransactor(address common.Address, transactor bind.ContractTransactor) (*PledgeTransactor, error) {
	contract, err := bindPledge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PledgeTransactor{contract: contract}, nil
}

// NewPledgeFilterer creates a new log filterer instance of Pledge, bound to a specific deployed contract.
func NewPledgeFilterer(address common.Address, filterer bind.ContractFilterer) (*PledgeFilterer, error) {
	contract, err := bindPledge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PledgeFilterer{contract: contract}, nil
}

// bindPledge binds a generic wrapper to an already deployed contract.
func bindPledge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PledgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pledge *PledgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pledge.Contract.PledgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pledge *PledgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge.Contract.PledgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pledge *PledgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pledge.Contract.PledgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pledge *PledgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pledge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pledge *PledgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pledge *PledgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pledge.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Pledge *PledgeCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Pledge *PledgeSession) Auth() (common.Address, error) {
	return _Pledge.Contract.Auth(&_Pledge.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Pledge *PledgeCallerSession) Auth() (common.Address, error) {
	return _Pledge.Contract.Auth(&_Pledge.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _index, uint8 _ti) view returns(uint256)
func (_Pledge *PledgeCaller) BalanceOf(opts *bind.CallOpts, _index uint64, _ti uint8) (*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "balanceOf", _index, _ti)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _index, uint8 _ti) view returns(uint256)
func (_Pledge *PledgeSession) BalanceOf(_index uint64, _ti uint8) (*big.Int, error) {
	return _Pledge.Contract.BalanceOf(&_Pledge.CallOpts, _index, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _index, uint8 _ti) view returns(uint256)
func (_Pledge *PledgeCallerSession) BalanceOf(_index uint64, _ti uint8) (*big.Int, error) {
	return _Pledge.Contract.BalanceOf(&_Pledge.CallOpts, _index, _ti)
}

// GetPledge is a free data retrieval call binding the contract method 0xeb395fde.
//
// Solidity: function getPledge(uint8 _ti) view returns(uint256)
func (_Pledge *PledgeCaller) GetPledge(opts *bind.CallOpts, _ti uint8) (*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getPledge", _ti)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPledge is a free data retrieval call binding the contract method 0xeb395fde.
//
// Solidity: function getPledge(uint8 _ti) view returns(uint256)
func (_Pledge *PledgeSession) GetPledge(_ti uint8) (*big.Int, error) {
	return _Pledge.Contract.GetPledge(&_Pledge.CallOpts, _ti)
}

// GetPledge is a free data retrieval call binding the contract method 0xeb395fde.
//
// Solidity: function getPledge(uint8 _ti) view returns(uint256)
func (_Pledge *PledgeCallerSession) GetPledge(_ti uint8) (*big.Int, error) {
	return _Pledge.Contract.GetPledge(&_Pledge.CallOpts, _ti)
}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Pledge *PledgeCaller) Inst(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "inst")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Pledge *PledgeSession) Inst() (common.Address, error) {
	return _Pledge.Contract.Inst(&_Pledge.CallOpts)
}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Pledge *PledgeCallerSession) Inst() (common.Address, error) {
	return _Pledge.Contract.Inst(&_Pledge.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Pledge *PledgeCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Pledge *PledgeSession) Owners(arg0 common.Address) (bool, error) {
	return _Pledge.Contract.Owners(&_Pledge.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Pledge *PledgeCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _Pledge.Contract.Owners(&_Pledge.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x14f732de.
//
// Solidity: function rewards(uint64 , uint8 ) view returns(uint256 accu, uint256 last, uint256 pledge, uint256 rs)
func (_Pledge *PledgeCaller) Rewards(opts *bind.CallOpts, arg0 uint64, arg1 uint8) (struct {
	Accu   *big.Int
	Last   *big.Int
	Pledge *big.Int
	Rs     *big.Int
}, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "rewards", arg0, arg1)

	outstruct := new(struct {
		Accu   *big.Int
		Last   *big.Int
		Pledge *big.Int
		Rs     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Accu = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Last = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Pledge = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Rs = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Rewards is a free data retrieval call binding the contract method 0x14f732de.
//
// Solidity: function rewards(uint64 , uint8 ) view returns(uint256 accu, uint256 last, uint256 pledge, uint256 rs)
func (_Pledge *PledgeSession) Rewards(arg0 uint64, arg1 uint8) (struct {
	Accu   *big.Int
	Last   *big.Int
	Pledge *big.Int
	Rs     *big.Int
}, error) {
	return _Pledge.Contract.Rewards(&_Pledge.CallOpts, arg0, arg1)
}

// Rewards is a free data retrieval call binding the contract method 0x14f732de.
//
// Solidity: function rewards(uint64 , uint8 ) view returns(uint256 accu, uint256 last, uint256 pledge, uint256 rs)
func (_Pledge *PledgeCallerSession) Rewards(arg0 uint64, arg1 uint8) (struct {
	Accu   *big.Int
	Last   *big.Int
	Pledge *big.Int
	Rs     *big.Int
}, error) {
	return _Pledge.Contract.Rewards(&_Pledge.CallOpts, arg0, arg1)
}

// TInfo is a free data retrieval call binding the contract method 0x041ac896.
//
// Solidity: function tInfo(uint8 ) view returns(uint256 accu, uint256 last, uint256 pledge, uint256 rs)
func (_Pledge *PledgeCaller) TInfo(opts *bind.CallOpts, arg0 uint8) (struct {
	Accu   *big.Int
	Last   *big.Int
	Pledge *big.Int
	Rs     *big.Int
}, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "tInfo", arg0)

	outstruct := new(struct {
		Accu   *big.Int
		Last   *big.Int
		Pledge *big.Int
		Rs     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Accu = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Last = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Pledge = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Rs = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TInfo is a free data retrieval call binding the contract method 0x041ac896.
//
// Solidity: function tInfo(uint8 ) view returns(uint256 accu, uint256 last, uint256 pledge, uint256 rs)
func (_Pledge *PledgeSession) TInfo(arg0 uint8) (struct {
	Accu   *big.Int
	Last   *big.Int
	Pledge *big.Int
	Rs     *big.Int
}, error) {
	return _Pledge.Contract.TInfo(&_Pledge.CallOpts, arg0)
}

// TInfo is a free data retrieval call binding the contract method 0x041ac896.
//
// Solidity: function tInfo(uint8 ) view returns(uint256 accu, uint256 last, uint256 pledge, uint256 rs)
func (_Pledge *PledgeCallerSession) TInfo(arg0 uint8) (struct {
	Accu   *big.Int
	Last   *big.Int
	Pledge *big.Int
	Rs     *big.Int
}, error) {
	return _Pledge.Contract.TInfo(&_Pledge.CallOpts, arg0)
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_Pledge *PledgeCaller) TotalPledge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "totalPledge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_Pledge *PledgeSession) TotalPledge() (*big.Int, error) {
	return _Pledge.Contract.TotalPledge(&_Pledge.CallOpts)
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_Pledge *PledgeCallerSession) TotalPledge() (*big.Int, error) {
	return _Pledge.Contract.TotalPledge(&_Pledge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Pledge *PledgeCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Pledge *PledgeSession) Version() (uint16, error) {
	return _Pledge.Contract.Version(&_Pledge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Pledge *PledgeCallerSession) Version() (uint16, error) {
	return _Pledge.Contract.Version(&_Pledge.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Pledge *PledgeTransactor) Add(opts *bind.TransactOpts, _a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "add", _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Pledge *PledgeSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Pledge.Contract.Add(&_Pledge.TransactOpts, _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Pledge *PledgeTransactorSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Pledge.Contract.Add(&_Pledge.TransactOpts, _a, isSet, signs)
}

// AddT is a paid mutator transaction binding the contract method 0x724a84a8.
//
// Solidity: function addT(uint8 _ti) returns()
func (_Pledge *PledgeTransactor) AddT(opts *bind.TransactOpts, _ti uint8) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "addT", _ti)
}

// AddT is a paid mutator transaction binding the contract method 0x724a84a8.
//
// Solidity: function addT(uint8 _ti) returns()
func (_Pledge *PledgeSession) AddT(_ti uint8) (*types.Transaction, error) {
	return _Pledge.Contract.AddT(&_Pledge.TransactOpts, _ti)
}

// AddT is a paid mutator transaction binding the contract method 0x724a84a8.
//
// Solidity: function addT(uint8 _ti) returns()
func (_Pledge *PledgeTransactorSession) AddT(_ti uint8) (*types.Transaction, error) {
	return _Pledge.Contract.AddT(&_Pledge.TransactOpts, _ti)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _index, uint256 money) returns()
func (_Pledge *PledgeTransactor) Pledge(opts *bind.TransactOpts, _index uint64, money *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "pledge", _index, money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _index, uint256 money) returns()
func (_Pledge *PledgeSession) Pledge(_index uint64, money *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.Pledge(&_Pledge.TransactOpts, _index, money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _index, uint256 money) returns()
func (_Pledge *PledgeTransactorSession) Pledge(_index uint64, money *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.Pledge(&_Pledge.TransactOpts, _index, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _index, uint8 _ti, uint256 money, uint256 _lock) returns(uint256)
func (_Pledge *PledgeTransactor) Withdraw(opts *bind.TransactOpts, _index uint64, _ti uint8, money *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "withdraw", _index, _ti, money, _lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _index, uint8 _ti, uint256 money, uint256 _lock) returns(uint256)
func (_Pledge *PledgeSession) Withdraw(_index uint64, _ti uint8, money *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.Withdraw(&_Pledge.TransactOpts, _index, _ti, money, _lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _index, uint8 _ti, uint256 money, uint256 _lock) returns(uint256)
func (_Pledge *PledgeTransactorSession) Withdraw(_index uint64, _ti uint8, money *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.Withdraw(&_Pledge.TransactOpts, _index, _ti, money, _lock)
}

// PledgeAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the Pledge contract.
type PledgeAddOwnerIterator struct {
	Event *PledgeAddOwner // Event containing the contract specifics and raw log

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
func (it *PledgeAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeAddOwner)
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
		it.Event = new(PledgeAddOwner)
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
func (it *PledgeAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeAddOwner represents a AddOwner event raised by the Pledge contract.
type PledgeAddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Pledge *PledgeFilterer) FilterAddOwner(opts *bind.FilterOpts) (*PledgeAddOwnerIterator, error) {

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &PledgeAddOwnerIterator{contract: _Pledge.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Pledge *PledgeFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *PledgeAddOwner) (event.Subscription, error) {

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeAddOwner)
				if err := _Pledge.contract.UnpackLog(event, "AddOwner", log); err != nil {
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

// ParseAddOwner is a log parse operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Pledge *PledgeFilterer) ParseAddOwner(log types.Log) (*PledgeAddOwner, error) {
	event := new(PledgeAddOwner)
	if err := _Pledge.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeAddTIterator is returned from FilterAddT and is used to iterate over the raw logs and unpacked data for AddT events raised by the Pledge contract.
type PledgeAddTIterator struct {
	Event *PledgeAddT // Event containing the contract specifics and raw log

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
func (it *PledgeAddTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeAddT)
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
		it.Event = new(PledgeAddT)
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
func (it *PledgeAddTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeAddTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeAddT represents a AddT event raised by the Pledge contract.
type PledgeAddT struct {
	Bal    *big.Int
	TIndex uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddT is a free log retrieval operation binding the contract event 0x8ce8baccab364cf4dd55bc29e9ca458faa5067327fb854f45973d4bdcc5b8cf9.
//
// Solidity: event AddT(uint256 bal, uint8 tIndex)
func (_Pledge *PledgeFilterer) FilterAddT(opts *bind.FilterOpts) (*PledgeAddTIterator, error) {

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return &PledgeAddTIterator{contract: _Pledge.contract, event: "AddT", logs: logs, sub: sub}, nil
}

// WatchAddT is a free log subscription operation binding the contract event 0x8ce8baccab364cf4dd55bc29e9ca458faa5067327fb854f45973d4bdcc5b8cf9.
//
// Solidity: event AddT(uint256 bal, uint8 tIndex)
func (_Pledge *PledgeFilterer) WatchAddT(opts *bind.WatchOpts, sink chan<- *PledgeAddT) (event.Subscription, error) {

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeAddT)
				if err := _Pledge.contract.UnpackLog(event, "AddT", log); err != nil {
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
func (_Pledge *PledgeFilterer) ParseAddT(log types.Log) (*PledgeAddT, error) {
	event := new(PledgeAddT)
	if err := _Pledge.contract.UnpackLog(event, "AddT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
