// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxy

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

// OrderIn is an auto generated low-level Go binding around an user-defined struct.
type OrderIn struct {
	UIndex uint64
	PIndex uint64
	Start  uint64
	End    uint64
	Size   uint64
	Nonce  uint64
	TIndex uint8
	Sprice *big.Int
}

// PWIn is an auto generated low-level Go binding around an user-defined struct.
type PWIn struct {
	PIndex uint64
	TIndex uint8
	Pay    *big.Int
	Lost   *big.Int
}

// IControl1ABI is the input ABI used to generate the binding from.
const IControl1ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"Pledge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lock\",\"type\":\"uint256\"}],\"name\":\"Unpledge\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_mint\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"addT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_p\",\"type\":\"address\"}],\"name\":\"alterPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_ban\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"banT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"confirmPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_level\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tr\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_kpPerB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ppPerB\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"}],\"name\":\"createGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"reAcc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_rtype\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_extra\",\"type\":\"bytes\"}],\"name\":\"reRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"}],\"name\":\"setDesc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setGDesc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_s\",\"type\":\"uint8\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setGS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_kp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pp\",\"type\":\"uint256\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_s\",\"type\":\"uint8\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setRS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tr\",\"type\":\"uint8\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"unpledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IControl1FuncSigs maps the 4-byte function signature to its string representation.
var IControl1FuncSigs = map[string]string{
	"843d1b49": "activate(uint64,bytes[5])",
	"68612348": "addT(address,bool,bytes[5])",
	"d9637231": "addToGroup(address,uint64)",
	"4a593dd4": "alterPayee(address,uint64,address)",
	"7a743984": "banT(address,bool,bytes[5])",
	"51c10b28": "confirmPayee(address,uint64)",
	"8aae670a": "createGroup(uint8,uint8,uint8,uint256,uint256,uint256,uint256,bytes)",
	"1a22e62e": "pledge(address,uint64,uint256)",
	"effcafce": "reAcc(address)",
	"9ccfe4ff": "reRole(address,uint8,bytes)",
	"722cc4a4": "setDesc(address,bytes)",
	"aa756343": "setGDesc(uint64,bytes,bytes[5])",
	"f62ea8d1": "setGS(uint64,uint8,bytes[5])",
	"5192e18a": "setP(uint64,uint256,uint256,bytes[5])",
	"88602357": "setRS(uint64,uint8,bytes[5])",
	"4d525421": "setRate(uint64,uint8,uint8,bytes[5])",
	"0a866aed": "unpledge(address,uint64,uint8,uint256)",
}

// IControl1 is an auto generated Go binding around an Ethereum contract.
type IControl1 struct {
	IControl1Caller     // Read-only binding to the contract
	IControl1Transactor // Write-only binding to the contract
	IControl1Filterer   // Log filterer for contract events
}

// IControl1Caller is an auto generated read-only Go binding around an Ethereum contract.
type IControl1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IControl1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IControl1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IControl1Session struct {
	Contract     *IControl1        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IControl1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IControl1CallerSession struct {
	Contract *IControl1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IControl1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IControl1TransactorSession struct {
	Contract     *IControl1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IControl1Raw is an auto generated low-level Go binding around an Ethereum contract.
type IControl1Raw struct {
	Contract *IControl1 // Generic contract binding to access the raw methods on
}

// IControl1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IControl1CallerRaw struct {
	Contract *IControl1Caller // Generic read-only contract binding to access the raw methods on
}

// IControl1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IControl1TransactorRaw struct {
	Contract *IControl1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIControl1 creates a new instance of IControl1, bound to a specific deployed contract.
func NewIControl1(address common.Address, backend bind.ContractBackend) (*IControl1, error) {
	contract, err := bindIControl1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IControl1{IControl1Caller: IControl1Caller{contract: contract}, IControl1Transactor: IControl1Transactor{contract: contract}, IControl1Filterer: IControl1Filterer{contract: contract}}, nil
}

// NewIControl1Caller creates a new read-only instance of IControl1, bound to a specific deployed contract.
func NewIControl1Caller(address common.Address, caller bind.ContractCaller) (*IControl1Caller, error) {
	contract, err := bindIControl1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IControl1Caller{contract: contract}, nil
}

// NewIControl1Transactor creates a new write-only instance of IControl1, bound to a specific deployed contract.
func NewIControl1Transactor(address common.Address, transactor bind.ContractTransactor) (*IControl1Transactor, error) {
	contract, err := bindIControl1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IControl1Transactor{contract: contract}, nil
}

// NewIControl1Filterer creates a new log filterer instance of IControl1, bound to a specific deployed contract.
func NewIControl1Filterer(address common.Address, filterer bind.ContractFilterer) (*IControl1Filterer, error) {
	contract, err := bindIControl1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IControl1Filterer{contract: contract}, nil
}

// bindIControl1 binds a generic wrapper to an already deployed contract.
func bindIControl1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IControl1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl1 *IControl1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl1.Contract.IControl1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl1 *IControl1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl1.Contract.IControl1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl1 *IControl1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl1.Contract.IControl1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl1 *IControl1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl1 *IControl1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl1 *IControl1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl1.Contract.contract.Transact(opts, method, params...)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_IControl1 *IControl1Transactor) Activate(opts *bind.TransactOpts, _i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "activate", _i, signs)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_IControl1 *IControl1Session) Activate(_i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.Activate(&_IControl1.TransactOpts, _i, signs)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_IControl1 *IControl1TransactorSession) Activate(_i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.Activate(&_IControl1.TransactOpts, _i, signs)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_IControl1 *IControl1Transactor) AddT(opts *bind.TransactOpts, _t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "addT", _t, _mint, signs)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_IControl1 *IControl1Session) AddT(_t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.AddT(&_IControl1.TransactOpts, _t, _mint, signs)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_IControl1 *IControl1TransactorSession) AddT(_t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.AddT(&_IControl1.TransactOpts, _t, _mint, signs)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xd9637231.
//
// Solidity: function addToGroup(address _a, uint64 _gi) returns()
func (_IControl1 *IControl1Transactor) AddToGroup(opts *bind.TransactOpts, _a common.Address, _gi uint64) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "addToGroup", _a, _gi)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xd9637231.
//
// Solidity: function addToGroup(address _a, uint64 _gi) returns()
func (_IControl1 *IControl1Session) AddToGroup(_a common.Address, _gi uint64) (*types.Transaction, error) {
	return _IControl1.Contract.AddToGroup(&_IControl1.TransactOpts, _a, _gi)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xd9637231.
//
// Solidity: function addToGroup(address _a, uint64 _gi) returns()
func (_IControl1 *IControl1TransactorSession) AddToGroup(_a common.Address, _gi uint64) (*types.Transaction, error) {
	return _IControl1.Contract.AddToGroup(&_IControl1.TransactOpts, _a, _gi)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x4a593dd4.
//
// Solidity: function alterPayee(address _a, uint64 _i, address _p) returns()
func (_IControl1 *IControl1Transactor) AlterPayee(opts *bind.TransactOpts, _a common.Address, _i uint64, _p common.Address) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "alterPayee", _a, _i, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x4a593dd4.
//
// Solidity: function alterPayee(address _a, uint64 _i, address _p) returns()
func (_IControl1 *IControl1Session) AlterPayee(_a common.Address, _i uint64, _p common.Address) (*types.Transaction, error) {
	return _IControl1.Contract.AlterPayee(&_IControl1.TransactOpts, _a, _i, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x4a593dd4.
//
// Solidity: function alterPayee(address _a, uint64 _i, address _p) returns()
func (_IControl1 *IControl1TransactorSession) AlterPayee(_a common.Address, _i uint64, _p common.Address) (*types.Transaction, error) {
	return _IControl1.Contract.AlterPayee(&_IControl1.TransactOpts, _a, _i, _p)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _ban, bytes[5] signs) returns()
func (_IControl1 *IControl1Transactor) BanT(opts *bind.TransactOpts, _t common.Address, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "banT", _t, _ban, signs)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _ban, bytes[5] signs) returns()
func (_IControl1 *IControl1Session) BanT(_t common.Address, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.BanT(&_IControl1.TransactOpts, _t, _ban, signs)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _ban, bytes[5] signs) returns()
func (_IControl1 *IControl1TransactorSession) BanT(_t common.Address, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.BanT(&_IControl1.TransactOpts, _t, _ban, signs)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x51c10b28.
//
// Solidity: function confirmPayee(address _a, uint64 _i) returns()
func (_IControl1 *IControl1Transactor) ConfirmPayee(opts *bind.TransactOpts, _a common.Address, _i uint64) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "confirmPayee", _a, _i)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x51c10b28.
//
// Solidity: function confirmPayee(address _a, uint64 _i) returns()
func (_IControl1 *IControl1Session) ConfirmPayee(_a common.Address, _i uint64) (*types.Transaction, error) {
	return _IControl1.Contract.ConfirmPayee(&_IControl1.TransactOpts, _a, _i)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x51c10b28.
//
// Solidity: function confirmPayee(address _a, uint64 _i) returns()
func (_IControl1 *IControl1TransactorSession) ConfirmPayee(_a common.Address, _i uint64) (*types.Transaction, error) {
	return _IControl1.Contract.ConfirmPayee(&_IControl1.TransactOpts, _a, _i)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x8aae670a.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kpPerB, uint256 _ppPerB, bytes _desc) returns()
func (_IControl1 *IControl1Transactor) CreateGroup(opts *bind.TransactOpts, _level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kpPerB *big.Int, _ppPerB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "createGroup", _level, _mr, _tr, _kr, _pr, _kpPerB, _ppPerB, _desc)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x8aae670a.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kpPerB, uint256 _ppPerB, bytes _desc) returns()
func (_IControl1 *IControl1Session) CreateGroup(_level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kpPerB *big.Int, _ppPerB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _IControl1.Contract.CreateGroup(&_IControl1.TransactOpts, _level, _mr, _tr, _kr, _pr, _kpPerB, _ppPerB, _desc)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x8aae670a.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kpPerB, uint256 _ppPerB, bytes _desc) returns()
func (_IControl1 *IControl1TransactorSession) CreateGroup(_level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kpPerB *big.Int, _ppPerB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _IControl1.Contract.CreateGroup(&_IControl1.TransactOpts, _level, _mr, _tr, _kr, _pr, _kpPerB, _ppPerB, _desc)
}

// Pledge is a paid mutator transaction binding the contract method 0x1a22e62e.
//
// Solidity: function pledge(address _a, uint64 _i, uint256 _money) returns()
func (_IControl1 *IControl1Transactor) Pledge(opts *bind.TransactOpts, _a common.Address, _i uint64, _money *big.Int) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "pledge", _a, _i, _money)
}

// Pledge is a paid mutator transaction binding the contract method 0x1a22e62e.
//
// Solidity: function pledge(address _a, uint64 _i, uint256 _money) returns()
func (_IControl1 *IControl1Session) Pledge(_a common.Address, _i uint64, _money *big.Int) (*types.Transaction, error) {
	return _IControl1.Contract.Pledge(&_IControl1.TransactOpts, _a, _i, _money)
}

// Pledge is a paid mutator transaction binding the contract method 0x1a22e62e.
//
// Solidity: function pledge(address _a, uint64 _i, uint256 _money) returns()
func (_IControl1 *IControl1TransactorSession) Pledge(_a common.Address, _i uint64, _money *big.Int) (*types.Transaction, error) {
	return _IControl1.Contract.Pledge(&_IControl1.TransactOpts, _a, _i, _money)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address _a) returns()
func (_IControl1 *IControl1Transactor) ReAcc(opts *bind.TransactOpts, _a common.Address) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "reAcc", _a)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address _a) returns()
func (_IControl1 *IControl1Session) ReAcc(_a common.Address) (*types.Transaction, error) {
	return _IControl1.Contract.ReAcc(&_IControl1.TransactOpts, _a)
}

// ReAcc is a paid mutator transaction binding the contract method 0xeffcafce.
//
// Solidity: function reAcc(address _a) returns()
func (_IControl1 *IControl1TransactorSession) ReAcc(_a common.Address) (*types.Transaction, error) {
	return _IControl1.Contract.ReAcc(&_IControl1.TransactOpts, _a)
}

// ReRole is a paid mutator transaction binding the contract method 0x9ccfe4ff.
//
// Solidity: function reRole(address _a, uint8 _rtype, bytes _extra) returns()
func (_IControl1 *IControl1Transactor) ReRole(opts *bind.TransactOpts, _a common.Address, _rtype uint8, _extra []byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "reRole", _a, _rtype, _extra)
}

// ReRole is a paid mutator transaction binding the contract method 0x9ccfe4ff.
//
// Solidity: function reRole(address _a, uint8 _rtype, bytes _extra) returns()
func (_IControl1 *IControl1Session) ReRole(_a common.Address, _rtype uint8, _extra []byte) (*types.Transaction, error) {
	return _IControl1.Contract.ReRole(&_IControl1.TransactOpts, _a, _rtype, _extra)
}

// ReRole is a paid mutator transaction binding the contract method 0x9ccfe4ff.
//
// Solidity: function reRole(address _a, uint8 _rtype, bytes _extra) returns()
func (_IControl1 *IControl1TransactorSession) ReRole(_a common.Address, _rtype uint8, _extra []byte) (*types.Transaction, error) {
	return _IControl1.Contract.ReRole(&_IControl1.TransactOpts, _a, _rtype, _extra)
}

// SetDesc is a paid mutator transaction binding the contract method 0x722cc4a4.
//
// Solidity: function setDesc(address _a, bytes _desc) returns()
func (_IControl1 *IControl1Transactor) SetDesc(opts *bind.TransactOpts, _a common.Address, _desc []byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "setDesc", _a, _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x722cc4a4.
//
// Solidity: function setDesc(address _a, bytes _desc) returns()
func (_IControl1 *IControl1Session) SetDesc(_a common.Address, _desc []byte) (*types.Transaction, error) {
	return _IControl1.Contract.SetDesc(&_IControl1.TransactOpts, _a, _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x722cc4a4.
//
// Solidity: function setDesc(address _a, bytes _desc) returns()
func (_IControl1 *IControl1TransactorSession) SetDesc(_a common.Address, _desc []byte) (*types.Transaction, error) {
	return _IControl1.Contract.SetDesc(&_IControl1.TransactOpts, _a, _desc)
}

// SetGDesc is a paid mutator transaction binding the contract method 0xaa756343.
//
// Solidity: function setGDesc(uint64 _gi, bytes _desc, bytes[5] signs) returns()
func (_IControl1 *IControl1Transactor) SetGDesc(opts *bind.TransactOpts, _gi uint64, _desc []byte, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "setGDesc", _gi, _desc, signs)
}

// SetGDesc is a paid mutator transaction binding the contract method 0xaa756343.
//
// Solidity: function setGDesc(uint64 _gi, bytes _desc, bytes[5] signs) returns()
func (_IControl1 *IControl1Session) SetGDesc(_gi uint64, _desc []byte, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.SetGDesc(&_IControl1.TransactOpts, _gi, _desc, signs)
}

// SetGDesc is a paid mutator transaction binding the contract method 0xaa756343.
//
// Solidity: function setGDesc(uint64 _gi, bytes _desc, bytes[5] signs) returns()
func (_IControl1 *IControl1TransactorSession) SetGDesc(_gi uint64, _desc []byte, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.SetGDesc(&_IControl1.TransactOpts, _gi, _desc, signs)
}

// SetGS is a paid mutator transaction binding the contract method 0xf62ea8d1.
//
// Solidity: function setGS(uint64 _gi, uint8 _s, bytes[5] signs) returns()
func (_IControl1 *IControl1Transactor) SetGS(opts *bind.TransactOpts, _gi uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "setGS", _gi, _s, signs)
}

// SetGS is a paid mutator transaction binding the contract method 0xf62ea8d1.
//
// Solidity: function setGS(uint64 _gi, uint8 _s, bytes[5] signs) returns()
func (_IControl1 *IControl1Session) SetGS(_gi uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.SetGS(&_IControl1.TransactOpts, _gi, _s, signs)
}

// SetGS is a paid mutator transaction binding the contract method 0xf62ea8d1.
//
// Solidity: function setGS(uint64 _gi, uint8 _s, bytes[5] signs) returns()
func (_IControl1 *IControl1TransactorSession) SetGS(_gi uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.SetGS(&_IControl1.TransactOpts, _gi, _s, signs)
}

// SetP is a paid mutator transaction binding the contract method 0x5192e18a.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp, bytes[5] signs) returns()
func (_IControl1 *IControl1Transactor) SetP(opts *bind.TransactOpts, _gi uint64, _kp *big.Int, _pp *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "setP", _gi, _kp, _pp, signs)
}

// SetP is a paid mutator transaction binding the contract method 0x5192e18a.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp, bytes[5] signs) returns()
func (_IControl1 *IControl1Session) SetP(_gi uint64, _kp *big.Int, _pp *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.SetP(&_IControl1.TransactOpts, _gi, _kp, _pp, signs)
}

// SetP is a paid mutator transaction binding the contract method 0x5192e18a.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp, bytes[5] signs) returns()
func (_IControl1 *IControl1TransactorSession) SetP(_gi uint64, _kp *big.Int, _pp *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.SetP(&_IControl1.TransactOpts, _gi, _kp, _pp, signs)
}

// SetRS is a paid mutator transaction binding the contract method 0x88602357.
//
// Solidity: function setRS(uint64 _i, uint8 _s, bytes[5] signs) returns()
func (_IControl1 *IControl1Transactor) SetRS(opts *bind.TransactOpts, _i uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "setRS", _i, _s, signs)
}

// SetRS is a paid mutator transaction binding the contract method 0x88602357.
//
// Solidity: function setRS(uint64 _i, uint8 _s, bytes[5] signs) returns()
func (_IControl1 *IControl1Session) SetRS(_i uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.SetRS(&_IControl1.TransactOpts, _i, _s, signs)
}

// SetRS is a paid mutator transaction binding the contract method 0x88602357.
//
// Solidity: function setRS(uint64 _i, uint8 _s, bytes[5] signs) returns()
func (_IControl1 *IControl1TransactorSession) SetRS(_i uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.SetRS(&_IControl1.TransactOpts, _i, _s, signs)
}

// SetRate is a paid mutator transaction binding the contract method 0x4d525421.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr, bytes[5] signs) returns()
func (_IControl1 *IControl1Transactor) SetRate(opts *bind.TransactOpts, _gi uint64, _mr uint8, _tr uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "setRate", _gi, _mr, _tr, signs)
}

// SetRate is a paid mutator transaction binding the contract method 0x4d525421.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr, bytes[5] signs) returns()
func (_IControl1 *IControl1Session) SetRate(_gi uint64, _mr uint8, _tr uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.SetRate(&_IControl1.TransactOpts, _gi, _mr, _tr, signs)
}

// SetRate is a paid mutator transaction binding the contract method 0x4d525421.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr, bytes[5] signs) returns()
func (_IControl1 *IControl1TransactorSession) SetRate(_gi uint64, _mr uint8, _tr uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IControl1.Contract.SetRate(&_IControl1.TransactOpts, _gi, _mr, _tr, signs)
}

// Unpledge is a paid mutator transaction binding the contract method 0x0a866aed.
//
// Solidity: function unpledge(address _a, uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IControl1 *IControl1Transactor) Unpledge(opts *bind.TransactOpts, _a common.Address, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IControl1.contract.Transact(opts, "unpledge", _a, _i, _ti, _money)
}

// Unpledge is a paid mutator transaction binding the contract method 0x0a866aed.
//
// Solidity: function unpledge(address _a, uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IControl1 *IControl1Session) Unpledge(_a common.Address, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IControl1.Contract.Unpledge(&_IControl1.TransactOpts, _a, _i, _ti, _money)
}

// Unpledge is a paid mutator transaction binding the contract method 0x0a866aed.
//
// Solidity: function unpledge(address _a, uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IControl1 *IControl1TransactorSession) Unpledge(_a common.Address, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IControl1.Contract.Unpledge(&_IControl1.TransactOpts, _a, _i, _ti, _money)
}

// IControl1PledgeIterator is returned from FilterPledge and is used to iterate over the raw logs and unpacked data for Pledge events raised by the IControl1 contract.
type IControl1PledgeIterator struct {
	Event *IControl1Pledge // Event containing the contract specifics and raw log

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
func (it *IControl1PledgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IControl1Pledge)
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
		it.Event = new(IControl1Pledge)
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
func (it *IControl1PledgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IControl1PledgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IControl1Pledge represents a Pledge event raised by the IControl1 contract.
type IControl1Pledge struct {
	Payer common.Address
	I     uint64
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPledge is a free log retrieval operation binding the contract event 0x09b6ee23f9f6fb394aa2974c56e3d3bef9d90a07cb6ed69bea1b1ca84a958ad9.
//
// Solidity: event Pledge(address indexed payer, uint64 indexed i, uint256 money)
func (_IControl1 *IControl1Filterer) FilterPledge(opts *bind.FilterOpts, payer []common.Address, i []uint64) (*IControl1PledgeIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _IControl1.contract.FilterLogs(opts, "Pledge", payerRule, iRule)
	if err != nil {
		return nil, err
	}
	return &IControl1PledgeIterator{contract: _IControl1.contract, event: "Pledge", logs: logs, sub: sub}, nil
}

// WatchPledge is a free log subscription operation binding the contract event 0x09b6ee23f9f6fb394aa2974c56e3d3bef9d90a07cb6ed69bea1b1ca84a958ad9.
//
// Solidity: event Pledge(address indexed payer, uint64 indexed i, uint256 money)
func (_IControl1 *IControl1Filterer) WatchPledge(opts *bind.WatchOpts, sink chan<- *IControl1Pledge, payer []common.Address, i []uint64) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _IControl1.contract.WatchLogs(opts, "Pledge", payerRule, iRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IControl1Pledge)
				if err := _IControl1.contract.UnpackLog(event, "Pledge", log); err != nil {
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

// ParsePledge is a log parse operation binding the contract event 0x09b6ee23f9f6fb394aa2974c56e3d3bef9d90a07cb6ed69bea1b1ca84a958ad9.
//
// Solidity: event Pledge(address indexed payer, uint64 indexed i, uint256 money)
func (_IControl1 *IControl1Filterer) ParsePledge(log types.Log) (*IControl1Pledge, error) {
	event := new(IControl1Pledge)
	if err := _IControl1.contract.UnpackLog(event, "Pledge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IControl1UnpledgeIterator is returned from FilterUnpledge and is used to iterate over the raw logs and unpacked data for Unpledge events raised by the IControl1 contract.
type IControl1UnpledgeIterator struct {
	Event *IControl1Unpledge // Event containing the contract specifics and raw log

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
func (it *IControl1UnpledgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IControl1Unpledge)
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
		it.Event = new(IControl1Unpledge)
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
func (it *IControl1UnpledgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IControl1UnpledgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IControl1Unpledge represents a Unpledge event raised by the IControl1 contract.
type IControl1Unpledge struct {
	Payee common.Address
	I     uint64
	Ti    uint8
	Money *big.Int
	Lock  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnpledge is a free log retrieval operation binding the contract event 0xbccb4889fc3c63ae9fc95055fbaa744519cd0a8117b6e6602bac77fdffe62456.
//
// Solidity: event Unpledge(address indexed payee, uint64 indexed i, uint8 indexed ti, uint256 money, uint256 lock)
func (_IControl1 *IControl1Filterer) FilterUnpledge(opts *bind.FilterOpts, payee []common.Address, i []uint64, ti []uint8) (*IControl1UnpledgeIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}
	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _IControl1.contract.FilterLogs(opts, "Unpledge", payeeRule, iRule, tiRule)
	if err != nil {
		return nil, err
	}
	return &IControl1UnpledgeIterator{contract: _IControl1.contract, event: "Unpledge", logs: logs, sub: sub}, nil
}

// WatchUnpledge is a free log subscription operation binding the contract event 0xbccb4889fc3c63ae9fc95055fbaa744519cd0a8117b6e6602bac77fdffe62456.
//
// Solidity: event Unpledge(address indexed payee, uint64 indexed i, uint8 indexed ti, uint256 money, uint256 lock)
func (_IControl1 *IControl1Filterer) WatchUnpledge(opts *bind.WatchOpts, sink chan<- *IControl1Unpledge, payee []common.Address, i []uint64, ti []uint8) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}
	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _IControl1.contract.WatchLogs(opts, "Unpledge", payeeRule, iRule, tiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IControl1Unpledge)
				if err := _IControl1.contract.UnpackLog(event, "Unpledge", log); err != nil {
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

// ParseUnpledge is a log parse operation binding the contract event 0xbccb4889fc3c63ae9fc95055fbaa744519cd0a8117b6e6602bac77fdffe62456.
//
// Solidity: event Unpledge(address indexed payee, uint64 indexed i, uint8 indexed ti, uint256 money, uint256 lock)
func (_IControl1 *IControl1Filterer) ParseUnpledge(log types.Log) (*IControl1Unpledge, error) {
	event := new(IControl1Unpledge)
	if err := _IControl1.contract.UnpackLog(event, "Unpledge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IControl2ABI is the input ABI used to generate the binding from.
const IControl2ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"name\":\"ProWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"Recharge\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"uSign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pSign\",\"type\":\"bytes\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"uint64[]\",\"name\":\"_kis\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structPWIn\",\"name\":\"_ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64[]\",\"name\":\"_kis\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"name\":\"proWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"uSign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pSign\",\"type\":\"bytes\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IControl2FuncSigs maps the 4-byte function signature to its string representation.
var IControl2FuncSigs = map[string]string{
	"ae9d0b40": "addOrder(address,(uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),bytes,bytes)",
	"af99c59a": "addRepair(address,(uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64[],bytes[])",
	"54aa6642": "proWithdraw(address,(uint64,uint8,uint256,uint256),uint64[],bytes[])",
	"f661f9e3": "recharge(address,uint64,uint8,bool,uint256)",
	"42f45166": "subOrder(address,(uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),bytes,bytes)",
}

// IControl2 is an auto generated Go binding around an Ethereum contract.
type IControl2 struct {
	IControl2Caller     // Read-only binding to the contract
	IControl2Transactor // Write-only binding to the contract
	IControl2Filterer   // Log filterer for contract events
}

// IControl2Caller is an auto generated read-only Go binding around an Ethereum contract.
type IControl2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IControl2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IControl2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IControl2Session struct {
	Contract     *IControl2        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IControl2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IControl2CallerSession struct {
	Contract *IControl2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IControl2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IControl2TransactorSession struct {
	Contract     *IControl2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IControl2Raw is an auto generated low-level Go binding around an Ethereum contract.
type IControl2Raw struct {
	Contract *IControl2 // Generic contract binding to access the raw methods on
}

// IControl2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IControl2CallerRaw struct {
	Contract *IControl2Caller // Generic read-only contract binding to access the raw methods on
}

// IControl2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IControl2TransactorRaw struct {
	Contract *IControl2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIControl2 creates a new instance of IControl2, bound to a specific deployed contract.
func NewIControl2(address common.Address, backend bind.ContractBackend) (*IControl2, error) {
	contract, err := bindIControl2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IControl2{IControl2Caller: IControl2Caller{contract: contract}, IControl2Transactor: IControl2Transactor{contract: contract}, IControl2Filterer: IControl2Filterer{contract: contract}}, nil
}

// NewIControl2Caller creates a new read-only instance of IControl2, bound to a specific deployed contract.
func NewIControl2Caller(address common.Address, caller bind.ContractCaller) (*IControl2Caller, error) {
	contract, err := bindIControl2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IControl2Caller{contract: contract}, nil
}

// NewIControl2Transactor creates a new write-only instance of IControl2, bound to a specific deployed contract.
func NewIControl2Transactor(address common.Address, transactor bind.ContractTransactor) (*IControl2Transactor, error) {
	contract, err := bindIControl2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IControl2Transactor{contract: contract}, nil
}

// NewIControl2Filterer creates a new log filterer instance of IControl2, bound to a specific deployed contract.
func NewIControl2Filterer(address common.Address, filterer bind.ContractFilterer) (*IControl2Filterer, error) {
	contract, err := bindIControl2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IControl2Filterer{contract: contract}, nil
}

// bindIControl2 binds a generic wrapper to an already deployed contract.
func bindIControl2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IControl2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl2 *IControl2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl2.Contract.IControl2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl2 *IControl2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl2.Contract.IControl2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl2 *IControl2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl2.Contract.IControl2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl2 *IControl2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl2 *IControl2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl2 *IControl2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl2.Contract.contract.Transact(opts, method, params...)
}

// AddOrder is a paid mutator transaction binding the contract method 0xae9d0b40.
//
// Solidity: function addOrder(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IControl2 *IControl2Transactor) AddOrder(opts *bind.TransactOpts, _a common.Address, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IControl2.contract.Transact(opts, "addOrder", _a, _oi, uSign, pSign)
}

// AddOrder is a paid mutator transaction binding the contract method 0xae9d0b40.
//
// Solidity: function addOrder(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IControl2 *IControl2Session) AddOrder(_a common.Address, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IControl2.Contract.AddOrder(&_IControl2.TransactOpts, _a, _oi, uSign, pSign)
}

// AddOrder is a paid mutator transaction binding the contract method 0xae9d0b40.
//
// Solidity: function addOrder(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IControl2 *IControl2TransactorSession) AddOrder(_a common.Address, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IControl2.Contract.AddOrder(&_IControl2.TransactOpts, _a, _oi, uSign, pSign)
}

// AddRepair is a paid mutator transaction binding the contract method 0xaf99c59a.
//
// Solidity: function addRepair(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_IControl2 *IControl2Transactor) AddRepair(opts *bind.TransactOpts, _a common.Address, _oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IControl2.contract.Transact(opts, "addRepair", _a, _oi, _kis, ksigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0xaf99c59a.
//
// Solidity: function addRepair(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_IControl2 *IControl2Session) AddRepair(_a common.Address, _oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IControl2.Contract.AddRepair(&_IControl2.TransactOpts, _a, _oi, _kis, ksigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0xaf99c59a.
//
// Solidity: function addRepair(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_IControl2 *IControl2TransactorSession) AddRepair(_a common.Address, _oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IControl2.Contract.AddRepair(&_IControl2.TransactOpts, _a, _oi, _kis, ksigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0x54aa6642.
//
// Solidity: function proWithdraw(address _a, (uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_IControl2 *IControl2Transactor) ProWithdraw(opts *bind.TransactOpts, _a common.Address, _ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IControl2.contract.Transact(opts, "proWithdraw", _a, _ps, _kis, ksigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0x54aa6642.
//
// Solidity: function proWithdraw(address _a, (uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_IControl2 *IControl2Session) ProWithdraw(_a common.Address, _ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IControl2.Contract.ProWithdraw(&_IControl2.TransactOpts, _a, _ps, _kis, ksigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0x54aa6642.
//
// Solidity: function proWithdraw(address _a, (uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_IControl2 *IControl2TransactorSession) ProWithdraw(_a common.Address, _ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IControl2.Contract.ProWithdraw(&_IControl2.TransactOpts, _a, _ps, _kis, ksigns)
}

// Recharge is a paid mutator transaction binding the contract method 0xf661f9e3.
//
// Solidity: function recharge(address _a, uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_IControl2 *IControl2Transactor) Recharge(opts *bind.TransactOpts, _a common.Address, _i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _IControl2.contract.Transact(opts, "recharge", _a, _i, _ti, isLock, _money)
}

// Recharge is a paid mutator transaction binding the contract method 0xf661f9e3.
//
// Solidity: function recharge(address _a, uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_IControl2 *IControl2Session) Recharge(_a common.Address, _i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _IControl2.Contract.Recharge(&_IControl2.TransactOpts, _a, _i, _ti, isLock, _money)
}

// Recharge is a paid mutator transaction binding the contract method 0xf661f9e3.
//
// Solidity: function recharge(address _a, uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_IControl2 *IControl2TransactorSession) Recharge(_a common.Address, _i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _IControl2.Contract.Recharge(&_IControl2.TransactOpts, _a, _i, _ti, isLock, _money)
}

// SubOrder is a paid mutator transaction binding the contract method 0x42f45166.
//
// Solidity: function subOrder(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IControl2 *IControl2Transactor) SubOrder(opts *bind.TransactOpts, _a common.Address, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IControl2.contract.Transact(opts, "subOrder", _a, _oi, uSign, pSign)
}

// SubOrder is a paid mutator transaction binding the contract method 0x42f45166.
//
// Solidity: function subOrder(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IControl2 *IControl2Session) SubOrder(_a common.Address, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IControl2.Contract.SubOrder(&_IControl2.TransactOpts, _a, _oi, uSign, pSign)
}

// SubOrder is a paid mutator transaction binding the contract method 0x42f45166.
//
// Solidity: function subOrder(address _a, (uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IControl2 *IControl2TransactorSession) SubOrder(_a common.Address, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IControl2.Contract.SubOrder(&_IControl2.TransactOpts, _a, _oi, uSign, pSign)
}

// IControl2ProWithdrawIterator is returned from FilterProWithdraw and is used to iterate over the raw logs and unpacked data for ProWithdraw events raised by the IControl2 contract.
type IControl2ProWithdrawIterator struct {
	Event *IControl2ProWithdraw // Event containing the contract specifics and raw log

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
func (it *IControl2ProWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IControl2ProWithdraw)
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
		it.Event = new(IControl2ProWithdraw)
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
func (it *IControl2ProWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IControl2ProWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IControl2ProWithdraw represents a ProWithdraw event raised by the IControl2 contract.
type IControl2ProWithdraw struct {
	A    common.Address
	Pi   uint64
	Ti   uint8
	Pay  *big.Int
	Lost *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterProWithdraw is a free log retrieval operation binding the contract event 0x3ea20fac35aae529e59d08cd5e2b0481108ccd275b50215084c05039245766db.
//
// Solidity: event ProWithdraw(address indexed a, uint64 indexed pi, uint8 ti, uint256 pay, uint256 lost)
func (_IControl2 *IControl2Filterer) FilterProWithdraw(opts *bind.FilterOpts, a []common.Address, pi []uint64) (*IControl2ProWithdrawIterator, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IControl2.contract.FilterLogs(opts, "ProWithdraw", aRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IControl2ProWithdrawIterator{contract: _IControl2.contract, event: "ProWithdraw", logs: logs, sub: sub}, nil
}

// WatchProWithdraw is a free log subscription operation binding the contract event 0x3ea20fac35aae529e59d08cd5e2b0481108ccd275b50215084c05039245766db.
//
// Solidity: event ProWithdraw(address indexed a, uint64 indexed pi, uint8 ti, uint256 pay, uint256 lost)
func (_IControl2 *IControl2Filterer) WatchProWithdraw(opts *bind.WatchOpts, sink chan<- *IControl2ProWithdraw, a []common.Address, pi []uint64) (event.Subscription, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IControl2.contract.WatchLogs(opts, "ProWithdraw", aRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IControl2ProWithdraw)
				if err := _IControl2.contract.UnpackLog(event, "ProWithdraw", log); err != nil {
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

// ParseProWithdraw is a log parse operation binding the contract event 0x3ea20fac35aae529e59d08cd5e2b0481108ccd275b50215084c05039245766db.
//
// Solidity: event ProWithdraw(address indexed a, uint64 indexed pi, uint8 ti, uint256 pay, uint256 lost)
func (_IControl2 *IControl2Filterer) ParseProWithdraw(log types.Log) (*IControl2ProWithdraw, error) {
	event := new(IControl2ProWithdraw)
	if err := _IControl2.contract.UnpackLog(event, "ProWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IControl2RechargeIterator is returned from FilterRecharge and is used to iterate over the raw logs and unpacked data for Recharge events raised by the IControl2 contract.
type IControl2RechargeIterator struct {
	Event *IControl2Recharge // Event containing the contract specifics and raw log

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
func (it *IControl2RechargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IControl2Recharge)
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
		it.Event = new(IControl2Recharge)
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
func (it *IControl2RechargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IControl2RechargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IControl2Recharge represents a Recharge event raised by the IControl2 contract.
type IControl2Recharge struct {
	Payer  common.Address
	I      uint64
	IsLock bool
	Ti     uint8
	Money  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecharge is a free log retrieval operation binding the contract event 0x544832b39fc44fc05cca7aa2c9f113bada9a6316b096b6cefcb302e433454a17.
//
// Solidity: event Recharge(address indexed payer, uint64 indexed i, bool indexed isLock, uint8 ti, uint256 money)
func (_IControl2 *IControl2Filterer) FilterRecharge(opts *bind.FilterOpts, payer []common.Address, i []uint64, isLock []bool) (*IControl2RechargeIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}
	var isLockRule []interface{}
	for _, isLockItem := range isLock {
		isLockRule = append(isLockRule, isLockItem)
	}

	logs, sub, err := _IControl2.contract.FilterLogs(opts, "Recharge", payerRule, iRule, isLockRule)
	if err != nil {
		return nil, err
	}
	return &IControl2RechargeIterator{contract: _IControl2.contract, event: "Recharge", logs: logs, sub: sub}, nil
}

// WatchRecharge is a free log subscription operation binding the contract event 0x544832b39fc44fc05cca7aa2c9f113bada9a6316b096b6cefcb302e433454a17.
//
// Solidity: event Recharge(address indexed payer, uint64 indexed i, bool indexed isLock, uint8 ti, uint256 money)
func (_IControl2 *IControl2Filterer) WatchRecharge(opts *bind.WatchOpts, sink chan<- *IControl2Recharge, payer []common.Address, i []uint64, isLock []bool) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}
	var isLockRule []interface{}
	for _, isLockItem := range isLock {
		isLockRule = append(isLockRule, isLockItem)
	}

	logs, sub, err := _IControl2.contract.WatchLogs(opts, "Recharge", payerRule, iRule, isLockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IControl2Recharge)
				if err := _IControl2.contract.UnpackLog(event, "Recharge", log); err != nil {
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

// ParseRecharge is a log parse operation binding the contract event 0x544832b39fc44fc05cca7aa2c9f113bada9a6316b096b6cefcb302e433454a17.
//
// Solidity: event Recharge(address indexed payer, uint64 indexed i, bool indexed isLock, uint8 ti, uint256 money)
func (_IControl2 *IControl2Filterer) ParseRecharge(log types.Log) (*IControl2Recharge, error) {
	event := new(IControl2Recharge)
	if err := _IControl2.contract.UnpackLog(event, "Recharge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IControl3ABI is the input ABI used to generate the binding from.
const IControl3ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualMoney\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"quitRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_le\",\"type\":\"uint8\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setLevel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pe\",\"type\":\"uint64\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_kpPerB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ppPerB\",\"type\":\"uint256\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setPlePerB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IControl3FuncSigs maps the 4-byte function signature to its string representation.
var IControl3FuncSigs = map[string]string{
	"d6985699": "quitRole(address,uint64)",
	"713d3b95": "setLevel(uint64,uint8,bytes[5])",
	"bd5d54b7": "setPeriod(uint64,bytes[5])",
	"130b89b1": "setPlePerB(uint64,uint256,uint256,bytes[5])",
	"9b4c757a": "withdraw(address,uint64,uint8,uint256)",
}

// IControl3 is an auto generated Go binding around an Ethereum contract.
type IControl3 struct {
	IControl3Caller     // Read-only binding to the contract
	IControl3Transactor // Write-only binding to the contract
	IControl3Filterer   // Log filterer for contract events
}

// IControl3Caller is an auto generated read-only Go binding around an Ethereum contract.
type IControl3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IControl3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IControl3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControl3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IControl3Session struct {
	Contract     *IControl3        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IControl3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IControl3CallerSession struct {
	Contract *IControl3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IControl3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IControl3TransactorSession struct {
	Contract     *IControl3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IControl3Raw is an auto generated low-level Go binding around an Ethereum contract.
type IControl3Raw struct {
	Contract *IControl3 // Generic contract binding to access the raw methods on
}

// IControl3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IControl3CallerRaw struct {
	Contract *IControl3Caller // Generic read-only contract binding to access the raw methods on
}

// IControl3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IControl3TransactorRaw struct {
	Contract *IControl3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIControl3 creates a new instance of IControl3, bound to a specific deployed contract.
func NewIControl3(address common.Address, backend bind.ContractBackend) (*IControl3, error) {
	contract, err := bindIControl3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IControl3{IControl3Caller: IControl3Caller{contract: contract}, IControl3Transactor: IControl3Transactor{contract: contract}, IControl3Filterer: IControl3Filterer{contract: contract}}, nil
}

// NewIControl3Caller creates a new read-only instance of IControl3, bound to a specific deployed contract.
func NewIControl3Caller(address common.Address, caller bind.ContractCaller) (*IControl3Caller, error) {
	contract, err := bindIControl3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IControl3Caller{contract: contract}, nil
}

// NewIControl3Transactor creates a new write-only instance of IControl3, bound to a specific deployed contract.
func NewIControl3Transactor(address common.Address, transactor bind.ContractTransactor) (*IControl3Transactor, error) {
	contract, err := bindIControl3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IControl3Transactor{contract: contract}, nil
}

// NewIControl3Filterer creates a new log filterer instance of IControl3, bound to a specific deployed contract.
func NewIControl3Filterer(address common.Address, filterer bind.ContractFilterer) (*IControl3Filterer, error) {
	contract, err := bindIControl3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IControl3Filterer{contract: contract}, nil
}

// bindIControl3 binds a generic wrapper to an already deployed contract.
func bindIControl3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IControl3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl3 *IControl3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl3.Contract.IControl3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl3 *IControl3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl3.Contract.IControl3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl3 *IControl3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl3.Contract.IControl3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl3 *IControl3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl3 *IControl3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl3 *IControl3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl3.Contract.contract.Transact(opts, method, params...)
}

// QuitRole is a paid mutator transaction binding the contract method 0xd6985699.
//
// Solidity: function quitRole(address _a, uint64 _i) returns()
func (_IControl3 *IControl3Transactor) QuitRole(opts *bind.TransactOpts, _a common.Address, _i uint64) (*types.Transaction, error) {
	return _IControl3.contract.Transact(opts, "quitRole", _a, _i)
}

// QuitRole is a paid mutator transaction binding the contract method 0xd6985699.
//
// Solidity: function quitRole(address _a, uint64 _i) returns()
func (_IControl3 *IControl3Session) QuitRole(_a common.Address, _i uint64) (*types.Transaction, error) {
	return _IControl3.Contract.QuitRole(&_IControl3.TransactOpts, _a, _i)
}

// QuitRole is a paid mutator transaction binding the contract method 0xd6985699.
//
// Solidity: function quitRole(address _a, uint64 _i) returns()
func (_IControl3 *IControl3TransactorSession) QuitRole(_a common.Address, _i uint64) (*types.Transaction, error) {
	return _IControl3.Contract.QuitRole(&_IControl3.TransactOpts, _a, _i)
}

// SetLevel is a paid mutator transaction binding the contract method 0x713d3b95.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le, bytes[5] signs) returns()
func (_IControl3 *IControl3Transactor) SetLevel(opts *bind.TransactOpts, _gi uint64, _le uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.contract.Transact(opts, "setLevel", _gi, _le, signs)
}

// SetLevel is a paid mutator transaction binding the contract method 0x713d3b95.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le, bytes[5] signs) returns()
func (_IControl3 *IControl3Session) SetLevel(_gi uint64, _le uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.Contract.SetLevel(&_IControl3.TransactOpts, _gi, _le, signs)
}

// SetLevel is a paid mutator transaction binding the contract method 0x713d3b95.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le, bytes[5] signs) returns()
func (_IControl3 *IControl3TransactorSession) SetLevel(_gi uint64, _le uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.Contract.SetLevel(&_IControl3.TransactOpts, _gi, _le, signs)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xbd5d54b7.
//
// Solidity: function setPeriod(uint64 _pe, bytes[5] signs) returns()
func (_IControl3 *IControl3Transactor) SetPeriod(opts *bind.TransactOpts, _pe uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.contract.Transact(opts, "setPeriod", _pe, signs)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xbd5d54b7.
//
// Solidity: function setPeriod(uint64 _pe, bytes[5] signs) returns()
func (_IControl3 *IControl3Session) SetPeriod(_pe uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.Contract.SetPeriod(&_IControl3.TransactOpts, _pe, signs)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xbd5d54b7.
//
// Solidity: function setPeriod(uint64 _pe, bytes[5] signs) returns()
func (_IControl3 *IControl3TransactorSession) SetPeriod(_pe uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.Contract.SetPeriod(&_IControl3.TransactOpts, _pe, signs)
}

// SetPlePerB is a paid mutator transaction binding the contract method 0x130b89b1.
//
// Solidity: function setPlePerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB, bytes[5] signs) returns()
func (_IControl3 *IControl3Transactor) SetPlePerB(opts *bind.TransactOpts, _gi uint64, _kpPerB *big.Int, _ppPerB *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.contract.Transact(opts, "setPlePerB", _gi, _kpPerB, _ppPerB, signs)
}

// SetPlePerB is a paid mutator transaction binding the contract method 0x130b89b1.
//
// Solidity: function setPlePerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB, bytes[5] signs) returns()
func (_IControl3 *IControl3Session) SetPlePerB(_gi uint64, _kpPerB *big.Int, _ppPerB *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.Contract.SetPlePerB(&_IControl3.TransactOpts, _gi, _kpPerB, _ppPerB, signs)
}

// SetPlePerB is a paid mutator transaction binding the contract method 0x130b89b1.
//
// Solidity: function setPlePerB(uint64 _gi, uint256 _kpPerB, uint256 _ppPerB, bytes[5] signs) returns()
func (_IControl3 *IControl3TransactorSession) SetPlePerB(_gi uint64, _kpPerB *big.Int, _ppPerB *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IControl3.Contract.SetPlePerB(&_IControl3.TransactOpts, _gi, _kpPerB, _ppPerB, signs)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b4c757a.
//
// Solidity: function withdraw(address _a, uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IControl3 *IControl3Transactor) Withdraw(opts *bind.TransactOpts, _a common.Address, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IControl3.contract.Transact(opts, "withdraw", _a, _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b4c757a.
//
// Solidity: function withdraw(address _a, uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IControl3 *IControl3Session) Withdraw(_a common.Address, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IControl3.Contract.Withdraw(&_IControl3.TransactOpts, _a, _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b4c757a.
//
// Solidity: function withdraw(address _a, uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IControl3 *IControl3TransactorSession) Withdraw(_a common.Address, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IControl3.Contract.Withdraw(&_IControl3.TransactOpts, _a, _i, _ti, _money)
}

// IControl3WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the IControl3 contract.
type IControl3WithdrawIterator struct {
	Event *IControl3Withdraw // Event containing the contract specifics and raw log

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
func (it *IControl3WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IControl3Withdraw)
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
		it.Event = new(IControl3Withdraw)
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
func (it *IControl3WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IControl3WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IControl3Withdraw represents a Withdraw event raised by the IControl3 contract.
type IControl3Withdraw struct {
	Addr        common.Address
	I           uint64
	Ti          uint8
	Money       *big.Int
	ActualMoney *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x8cf9f6665801a7a41feb765d72a4472b5f7d327ed18c77c6e81d2bb5622cd8c0.
//
// Solidity: event Withdraw(address indexed addr, uint64 indexed i, uint8 ti, uint256 money, uint256 actualMoney)
func (_IControl3 *IControl3Filterer) FilterWithdraw(opts *bind.FilterOpts, addr []common.Address, i []uint64) (*IControl3WithdrawIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _IControl3.contract.FilterLogs(opts, "Withdraw", addrRule, iRule)
	if err != nil {
		return nil, err
	}
	return &IControl3WithdrawIterator{contract: _IControl3.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x8cf9f6665801a7a41feb765d72a4472b5f7d327ed18c77c6e81d2bb5622cd8c0.
//
// Solidity: event Withdraw(address indexed addr, uint64 indexed i, uint8 ti, uint256 money, uint256 actualMoney)
func (_IControl3 *IControl3Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *IControl3Withdraw, addr []common.Address, i []uint64) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _IControl3.contract.WatchLogs(opts, "Withdraw", addrRule, iRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IControl3Withdraw)
				if err := _IControl3.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x8cf9f6665801a7a41feb765d72a4472b5f7d327ed18c77c6e81d2bb5622cd8c0.
//
// Solidity: event Withdraw(address indexed addr, uint64 indexed i, uint8 ti, uint256 money, uint256 actualMoney)
func (_IControl3 *IControl3Filterer) ParseWithdraw(log types.Log) (*IControl3Withdraw, error) {
	event := new(IControl3Withdraw)
	if err := _IControl3.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileSysSetterABI is the input ABI used to generate the binding from.
const IFileSysSetterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"AddOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"AddRepair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ui\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"SubOrder\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tr\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_rtype\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kPB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pPB\",\"type\":\"uint256\"}],\"name\":\"addPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structPWIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"proWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lock\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IFileSysSetterFuncSigs maps the 4-byte function signature to its string representation.
var IFileSysSetterFuncSigs = map[string]string{
	"9058d8e7": "addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint8,uint8,uint64)",
	"2449a168": "addPenalty(uint64,uint64,uint8,uint8,uint256,uint256)",
	"80faaf88": "addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64)",
	"a4703e16": "proWithdraw((uint64,uint8,uint256,uint256),uint64)",
	"24d11d40": "recharge(uint64,uint8,bool,uint256)",
	"248d02a0": "subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64)",
	"5affa0f3": "withdraw(uint64,uint8,uint256,uint256)",
}

// IFileSysSetter is an auto generated Go binding around an Ethereum contract.
type IFileSysSetter struct {
	IFileSysSetterCaller     // Read-only binding to the contract
	IFileSysSetterTransactor // Write-only binding to the contract
	IFileSysSetterFilterer   // Log filterer for contract events
}

// IFileSysSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFileSysSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFileSysSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFileSysSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFileSysSetterSession struct {
	Contract     *IFileSysSetter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFileSysSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFileSysSetterCallerSession struct {
	Contract *IFileSysSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IFileSysSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFileSysSetterTransactorSession struct {
	Contract     *IFileSysSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IFileSysSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFileSysSetterRaw struct {
	Contract *IFileSysSetter // Generic contract binding to access the raw methods on
}

// IFileSysSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFileSysSetterCallerRaw struct {
	Contract *IFileSysSetterCaller // Generic read-only contract binding to access the raw methods on
}

// IFileSysSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFileSysSetterTransactorRaw struct {
	Contract *IFileSysSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFileSysSetter creates a new instance of IFileSysSetter, bound to a specific deployed contract.
func NewIFileSysSetter(address common.Address, backend bind.ContractBackend) (*IFileSysSetter, error) {
	contract, err := bindIFileSysSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetter{IFileSysSetterCaller: IFileSysSetterCaller{contract: contract}, IFileSysSetterTransactor: IFileSysSetterTransactor{contract: contract}, IFileSysSetterFilterer: IFileSysSetterFilterer{contract: contract}}, nil
}

// NewIFileSysSetterCaller creates a new read-only instance of IFileSysSetter, bound to a specific deployed contract.
func NewIFileSysSetterCaller(address common.Address, caller bind.ContractCaller) (*IFileSysSetterCaller, error) {
	contract, err := bindIFileSysSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterCaller{contract: contract}, nil
}

// NewIFileSysSetterTransactor creates a new write-only instance of IFileSysSetter, bound to a specific deployed contract.
func NewIFileSysSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IFileSysSetterTransactor, error) {
	contract, err := bindIFileSysSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterTransactor{contract: contract}, nil
}

// NewIFileSysSetterFilterer creates a new log filterer instance of IFileSysSetter, bound to a specific deployed contract.
func NewIFileSysSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IFileSysSetterFilterer, error) {
	contract, err := bindIFileSysSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterFilterer{contract: contract}, nil
}

// bindIFileSysSetter binds a generic wrapper to an already deployed contract.
func bindIFileSysSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFileSysSetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSysSetter *IFileSysSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSysSetter.Contract.IFileSysSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSysSetter *IFileSysSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.IFileSysSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSysSetter *IFileSysSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.IFileSysSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSysSetter *IFileSysSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSysSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSysSetter *IFileSysSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSysSetter *IFileSysSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.contract.Transact(opts, method, params...)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9058d8e7.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint8 _mr, uint8 _tr, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) AddOrder(opts *bind.TransactOpts, ps OrderIn, _mr uint8, _tr uint8, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "addOrder", ps, _mr, _tr, _gi)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9058d8e7.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint8 _mr, uint8 _tr, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterSession) AddOrder(ps OrderIn, _mr uint8, _tr uint8, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddOrder(&_IFileSysSetter.TransactOpts, ps, _mr, _tr, _gi)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9058d8e7.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint8 _mr, uint8 _tr, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) AddOrder(ps OrderIn, _mr uint8, _tr uint8, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddOrder(&_IFileSysSetter.TransactOpts, ps, _mr, _tr, _gi)
}

// AddPenalty is a paid mutator transaction binding the contract method 0x2449a168.
//
// Solidity: function addPenalty(uint64 _gi, uint64 _i, uint8 _ti, uint8 _rtype, uint256 _kPB, uint256 _pPB) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) AddPenalty(opts *bind.TransactOpts, _gi uint64, _i uint64, _ti uint8, _rtype uint8, _kPB *big.Int, _pPB *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "addPenalty", _gi, _i, _ti, _rtype, _kPB, _pPB)
}

// AddPenalty is a paid mutator transaction binding the contract method 0x2449a168.
//
// Solidity: function addPenalty(uint64 _gi, uint64 _i, uint8 _ti, uint8 _rtype, uint256 _kPB, uint256 _pPB) returns()
func (_IFileSysSetter *IFileSysSetterSession) AddPenalty(_gi uint64, _i uint64, _ti uint8, _rtype uint8, _kPB *big.Int, _pPB *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddPenalty(&_IFileSysSetter.TransactOpts, _gi, _i, _ti, _rtype, _kPB, _pPB)
}

// AddPenalty is a paid mutator transaction binding the contract method 0x2449a168.
//
// Solidity: function addPenalty(uint64 _gi, uint64 _i, uint8 _ti, uint8 _rtype, uint256 _kPB, uint256 _pPB) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) AddPenalty(_gi uint64, _i uint64, _ti uint8, _rtype uint8, _kPB *big.Int, _pPB *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddPenalty(&_IFileSysSetter.TransactOpts, _gi, _i, _ti, _rtype, _kPB, _pPB)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) AddRepair(opts *bind.TransactOpts, ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "addRepair", ps, _gi)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterSession) AddRepair(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddRepair(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// AddRepair is a paid mutator transaction binding the contract method 0x80faaf88.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) AddRepair(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.AddRepair(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_IFileSysSetter *IFileSysSetterTransactor) ProWithdraw(opts *bind.TransactOpts, ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "proWithdraw", ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_IFileSysSetter *IFileSysSetterSession) ProWithdraw(ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.ProWithdraw(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xa4703e16.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) ps, uint64 _gi) returns(uint256)
func (_IFileSysSetter *IFileSysSetterTransactorSession) ProWithdraw(ps PWIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.ProWithdraw(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) Recharge(opts *bind.TransactOpts, _i uint64, _ti uint8, isLock bool, money *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "recharge", _i, _ti, isLock, money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) returns()
func (_IFileSysSetter *IFileSysSetterSession) Recharge(_i uint64, _ti uint8, isLock bool, money *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.Recharge(&_IFileSysSetter.TransactOpts, _i, _ti, isLock, money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) Recharge(_i uint64, _ti uint8, isLock bool, money *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.Recharge(&_IFileSysSetter.TransactOpts, _i, _ti, isLock, money)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactor) SubOrder(opts *bind.TransactOpts, ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "subOrder", ps, _gi)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterSession) SubOrder(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.SubOrder(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// SubOrder is a paid mutator transaction binding the contract method 0x248d02a0.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) ps, uint64 _gi) returns()
func (_IFileSysSetter *IFileSysSetterTransactorSession) SubOrder(ps OrderIn, _gi uint64) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.SubOrder(&_IFileSysSetter.TransactOpts, ps, _gi)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 _lock) returns(uint256)
func (_IFileSysSetter *IFileSysSetterTransactor) Withdraw(opts *bind.TransactOpts, _i uint64, _ti uint8, money *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.contract.Transact(opts, "withdraw", _i, _ti, money, _lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 _lock) returns(uint256)
func (_IFileSysSetter *IFileSysSetterSession) Withdraw(_i uint64, _ti uint8, money *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.Withdraw(&_IFileSysSetter.TransactOpts, _i, _ti, money, _lock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5affa0f3.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 _lock) returns(uint256)
func (_IFileSysSetter *IFileSysSetterTransactorSession) Withdraw(_i uint64, _ti uint8, money *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _IFileSysSetter.Contract.Withdraw(&_IFileSysSetter.TransactOpts, _i, _ti, money, _lock)
}

// IFileSysSetterAddOrderIterator is returned from FilterAddOrder and is used to iterate over the raw logs and unpacked data for AddOrder events raised by the IFileSysSetter contract.
type IFileSysSetterAddOrderIterator struct {
	Event *IFileSysSetterAddOrder // Event containing the contract specifics and raw log

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
func (it *IFileSysSetterAddOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileSysSetterAddOrder)
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
		it.Event = new(IFileSysSetterAddOrder)
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
func (it *IFileSysSetterAddOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileSysSetterAddOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileSysSetterAddOrder represents a AddOrder event raised by the IFileSysSetter contract.
type IFileSysSetterAddOrder struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddOrder is a free log retrieval operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) FilterAddOrder(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*IFileSysSetterAddOrderIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.FilterLogs(opts, "AddOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterAddOrderIterator{contract: _IFileSysSetter.contract, event: "AddOrder", logs: logs, sub: sub}, nil
}

// WatchAddOrder is a free log subscription operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) WatchAddOrder(opts *bind.WatchOpts, sink chan<- *IFileSysSetterAddOrder, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.WatchLogs(opts, "AddOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileSysSetterAddOrder)
				if err := _IFileSysSetter.contract.UnpackLog(event, "AddOrder", log); err != nil {
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

// ParseAddOrder is a log parse operation binding the contract event 0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314.
//
// Solidity: event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) ParseAddOrder(log types.Log) (*IFileSysSetterAddOrder, error) {
	event := new(IFileSysSetterAddOrder)
	if err := _IFileSysSetter.contract.UnpackLog(event, "AddOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileSysSetterAddRepairIterator is returned from FilterAddRepair and is used to iterate over the raw logs and unpacked data for AddRepair events raised by the IFileSysSetter contract.
type IFileSysSetterAddRepairIterator struct {
	Event *IFileSysSetterAddRepair // Event containing the contract specifics and raw log

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
func (it *IFileSysSetterAddRepairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileSysSetterAddRepair)
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
		it.Event = new(IFileSysSetterAddRepair)
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
func (it *IFileSysSetterAddRepairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileSysSetterAddRepairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileSysSetterAddRepair represents a AddRepair event raised by the IFileSysSetter contract.
type IFileSysSetterAddRepair struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddRepair is a free log retrieval operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) FilterAddRepair(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*IFileSysSetterAddRepairIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.FilterLogs(opts, "AddRepair", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterAddRepairIterator{contract: _IFileSysSetter.contract, event: "AddRepair", logs: logs, sub: sub}, nil
}

// WatchAddRepair is a free log subscription operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) WatchAddRepair(opts *bind.WatchOpts, sink chan<- *IFileSysSetterAddRepair, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.WatchLogs(opts, "AddRepair", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileSysSetterAddRepair)
				if err := _IFileSysSetter.contract.UnpackLog(event, "AddRepair", log); err != nil {
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

// ParseAddRepair is a log parse operation binding the contract event 0x82f944cf0134663ff19eea526226f1cd8d5746a71cc837983659b8cbb2674e48.
//
// Solidity: event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) ParseAddRepair(log types.Log) (*IFileSysSetterAddRepair, error) {
	event := new(IFileSysSetterAddRepair)
	if err := _IFileSysSetter.contract.UnpackLog(event, "AddRepair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileSysSetterSubOrderIterator is returned from FilterSubOrder and is used to iterate over the raw logs and unpacked data for SubOrder events raised by the IFileSysSetter contract.
type IFileSysSetterSubOrderIterator struct {
	Event *IFileSysSetterSubOrder // Event containing the contract specifics and raw log

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
func (it *IFileSysSetterSubOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileSysSetterSubOrder)
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
		it.Event = new(IFileSysSetterSubOrder)
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
func (it *IFileSysSetterSubOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileSysSetterSubOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileSysSetterSubOrder represents a SubOrder event raised by the IFileSysSetter contract.
type IFileSysSetterSubOrder struct {
	Ti     uint8
	Ui     uint64
	Pi     uint64
	Start  uint64
	End    uint64
	Size   uint64
	Sprice *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubOrder is a free log retrieval operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) FilterSubOrder(opts *bind.FilterOpts, ti []uint8, ui []uint64, pi []uint64) (*IFileSysSetterSubOrderIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.FilterLogs(opts, "SubOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return &IFileSysSetterSubOrderIterator{contract: _IFileSysSetter.contract, event: "SubOrder", logs: logs, sub: sub}, nil
}

// WatchSubOrder is a free log subscription operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) WatchSubOrder(opts *bind.WatchOpts, sink chan<- *IFileSysSetterSubOrder, ti []uint8, ui []uint64, pi []uint64) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}
	var uiRule []interface{}
	for _, uiItem := range ui {
		uiRule = append(uiRule, uiItem)
	}
	var piRule []interface{}
	for _, piItem := range pi {
		piRule = append(piRule, piItem)
	}

	logs, sub, err := _IFileSysSetter.contract.WatchLogs(opts, "SubOrder", tiRule, uiRule, piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileSysSetterSubOrder)
				if err := _IFileSysSetter.contract.UnpackLog(event, "SubOrder", log); err != nil {
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

// ParseSubOrder is a log parse operation binding the contract event 0xe42ef64dc7dfe9121f13026a70df328cf6687348fdf8f71745f8e40fa5e52b82.
//
// Solidity: event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice)
func (_IFileSysSetter *IFileSysSetterFilterer) ParseSubOrder(log types.Log) (*IFileSysSetterSubOrder, error) {
	event := new(IFileSysSetterSubOrder)
	if err := _IFileSysSetter.contract.UnpackLog(event, "SubOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// IProxyABI is the input ABI used to generate the binding from.
const IProxyABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"Activate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"mint\",\"type\":\"bool\"}],\"name\":\"AddT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"ban\",\"type\":\"bool\"}],\"name\":\"BanT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"desc\",\"type\":\"bytes\"}],\"name\":\"SetGDesc\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_s\",\"type\":\"uint8\"}],\"name\":\"SetGS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"le\",\"type\":\"uint8\"}],\"name\":\"SetLevel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"pe\",\"type\":\"uint64\"}],\"name\":\"SetPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kpPerB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ppPerB\",\"type\":\"uint256\"}],\"name\":\"SetPlePerB\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_s\",\"type\":\"uint8\"}],\"name\":\"SetRS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"mr\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kr\",\"type\":\"uint8\"}],\"name\":\"SetRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pp\",\"type\":\"uint256\"}],\"name\":\"SetkpP\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"uSign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pSign\",\"type\":\"bytes\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"uint64[]\",\"name\":\"_kis\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_mint\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"addT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_p\",\"type\":\"address\"}],\"name\":\"alterPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_ban\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"banT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"confirmPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_level\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tr\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_kpPerB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ppPerB\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"}],\"name\":\"createGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structPWIn\",\"name\":\"_ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64[]\",\"name\":\"_kis\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"name\":\"proWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"quitRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reAcc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_rtype\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_extra\",\"type\":\"bytes\"}],\"name\":\"reRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"}],\"name\":\"setDesc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setGDesc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_s\",\"type\":\"uint8\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setGS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_le\",\"type\":\"uint8\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setLevel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_kp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pp\",\"type\":\"uint256\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pe\",\"type\":\"uint64\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_kpPerB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ppPerB\",\"type\":\"uint256\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setPlePerB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_s\",\"type\":\"uint8\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setRS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tr\",\"type\":\"uint8\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"uSign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pSign\",\"type\":\"bytes\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"unpledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IProxyFuncSigs maps the 4-byte function signature to its string representation.
var IProxyFuncSigs = map[string]string{
	"843d1b49": "activate(uint64,bytes[5])",
	"c05bdcd0": "addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),bytes,bytes)",
	"97f2352c": "addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64[],bytes[])",
	"68612348": "addT(address,bool,bytes[5])",
	"a246de42": "addToGroup(uint64)",
	"2915824a": "alterPayee(uint64,address)",
	"7a743984": "banT(address,bool,bytes[5])",
	"310aa0be": "confirmPayee(uint64)",
	"8aae670a": "createGroup(uint8,uint8,uint8,uint256,uint256,uint256,uint256,bytes)",
	"d23f7482": "pledge(uint64,uint256)",
	"e182d54c": "proWithdraw((uint64,uint8,uint256,uint256),uint64[],bytes[])",
	"adad6d3a": "quitRole(uint64)",
	"6526250f": "reAcc()",
	"898e00f8": "reRole(uint8,bytes)",
	"24d11d40": "recharge(uint64,uint8,bool,uint256)",
	"0fc6fdd7": "setDesc(bytes)",
	"aa756343": "setGDesc(uint64,bytes,bytes[5])",
	"f62ea8d1": "setGS(uint64,uint8,bytes[5])",
	"713d3b95": "setLevel(uint64,uint8,bytes[5])",
	"5192e18a": "setP(uint64,uint256,uint256,bytes[5])",
	"bd5d54b7": "setPeriod(uint64,bytes[5])",
	"130b89b1": "setPlePerB(uint64,uint256,uint256,bytes[5])",
	"88602357": "setRS(uint64,uint8,bytes[5])",
	"4d525421": "setRate(uint64,uint8,uint8,bytes[5])",
	"153ede07": "subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),bytes,bytes)",
	"60985756": "unpledge(uint64,uint8,uint256)",
	"259c6d5e": "withdraw(uint64,uint8,uint256)",
}

// IProxy is an auto generated Go binding around an Ethereum contract.
type IProxy struct {
	IProxyCaller     // Read-only binding to the contract
	IProxyTransactor // Write-only binding to the contract
	IProxyFilterer   // Log filterer for contract events
}

// IProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IProxySession struct {
	Contract     *IProxy           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IProxyCallerSession struct {
	Contract *IProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IProxyTransactorSession struct {
	Contract     *IProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IProxyRaw struct {
	Contract *IProxy // Generic contract binding to access the raw methods on
}

// IProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IProxyCallerRaw struct {
	Contract *IProxyCaller // Generic read-only contract binding to access the raw methods on
}

// IProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IProxyTransactorRaw struct {
	Contract *IProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIProxy creates a new instance of IProxy, bound to a specific deployed contract.
func NewIProxy(address common.Address, backend bind.ContractBackend) (*IProxy, error) {
	contract, err := bindIProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IProxy{IProxyCaller: IProxyCaller{contract: contract}, IProxyTransactor: IProxyTransactor{contract: contract}, IProxyFilterer: IProxyFilterer{contract: contract}}, nil
}

// NewIProxyCaller creates a new read-only instance of IProxy, bound to a specific deployed contract.
func NewIProxyCaller(address common.Address, caller bind.ContractCaller) (*IProxyCaller, error) {
	contract, err := bindIProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IProxyCaller{contract: contract}, nil
}

// NewIProxyTransactor creates a new write-only instance of IProxy, bound to a specific deployed contract.
func NewIProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IProxyTransactor, error) {
	contract, err := bindIProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IProxyTransactor{contract: contract}, nil
}

// NewIProxyFilterer creates a new log filterer instance of IProxy, bound to a specific deployed contract.
func NewIProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IProxyFilterer, error) {
	contract, err := bindIProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IProxyFilterer{contract: contract}, nil
}

// bindIProxy binds a generic wrapper to an already deployed contract.
func bindIProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProxy *IProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProxy.Contract.IProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProxy *IProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProxy.Contract.IProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProxy *IProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProxy.Contract.IProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProxy *IProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProxy *IProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProxy *IProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProxy.Contract.contract.Transact(opts, method, params...)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_IProxy *IProxyTransactor) Activate(opts *bind.TransactOpts, _i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "activate", _i, signs)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_IProxy *IProxySession) Activate(_i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.Activate(&_IProxy.TransactOpts, _i, signs)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_IProxy *IProxyTransactorSession) Activate(_i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.Activate(&_IProxy.TransactOpts, _i, signs)
}

// AddOrder is a paid mutator transaction binding the contract method 0xc05bdcd0.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IProxy *IProxyTransactor) AddOrder(opts *bind.TransactOpts, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "addOrder", _oi, uSign, pSign)
}

// AddOrder is a paid mutator transaction binding the contract method 0xc05bdcd0.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IProxy *IProxySession) AddOrder(_oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IProxy.Contract.AddOrder(&_IProxy.TransactOpts, _oi, uSign, pSign)
}

// AddOrder is a paid mutator transaction binding the contract method 0xc05bdcd0.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IProxy *IProxyTransactorSession) AddOrder(_oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IProxy.Contract.AddOrder(&_IProxy.TransactOpts, _oi, uSign, pSign)
}

// AddRepair is a paid mutator transaction binding the contract method 0x97f2352c.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_IProxy *IProxyTransactor) AddRepair(opts *bind.TransactOpts, _oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "addRepair", _oi, _kis, ksigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0x97f2352c.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_IProxy *IProxySession) AddRepair(_oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.AddRepair(&_IProxy.TransactOpts, _oi, _kis, ksigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0x97f2352c.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_IProxy *IProxyTransactorSession) AddRepair(_oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.AddRepair(&_IProxy.TransactOpts, _oi, _kis, ksigns)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_IProxy *IProxyTransactor) AddT(opts *bind.TransactOpts, _t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "addT", _t, _mint, signs)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_IProxy *IProxySession) AddT(_t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.AddT(&_IProxy.TransactOpts, _t, _mint, signs)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_IProxy *IProxyTransactorSession) AddT(_t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.AddT(&_IProxy.TransactOpts, _t, _mint, signs)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xa246de42.
//
// Solidity: function addToGroup(uint64 _gi) returns()
func (_IProxy *IProxyTransactor) AddToGroup(opts *bind.TransactOpts, _gi uint64) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "addToGroup", _gi)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xa246de42.
//
// Solidity: function addToGroup(uint64 _gi) returns()
func (_IProxy *IProxySession) AddToGroup(_gi uint64) (*types.Transaction, error) {
	return _IProxy.Contract.AddToGroup(&_IProxy.TransactOpts, _gi)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xa246de42.
//
// Solidity: function addToGroup(uint64 _gi) returns()
func (_IProxy *IProxyTransactorSession) AddToGroup(_gi uint64) (*types.Transaction, error) {
	return _IProxy.Contract.AddToGroup(&_IProxy.TransactOpts, _gi)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _i, address _p) returns()
func (_IProxy *IProxyTransactor) AlterPayee(opts *bind.TransactOpts, _i uint64, _p common.Address) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "alterPayee", _i, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _i, address _p) returns()
func (_IProxy *IProxySession) AlterPayee(_i uint64, _p common.Address) (*types.Transaction, error) {
	return _IProxy.Contract.AlterPayee(&_IProxy.TransactOpts, _i, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _i, address _p) returns()
func (_IProxy *IProxyTransactorSession) AlterPayee(_i uint64, _p common.Address) (*types.Transaction, error) {
	return _IProxy.Contract.AlterPayee(&_IProxy.TransactOpts, _i, _p)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _ban, bytes[5] signs) returns()
func (_IProxy *IProxyTransactor) BanT(opts *bind.TransactOpts, _t common.Address, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "banT", _t, _ban, signs)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _ban, bytes[5] signs) returns()
func (_IProxy *IProxySession) BanT(_t common.Address, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.BanT(&_IProxy.TransactOpts, _t, _ban, signs)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _ban, bytes[5] signs) returns()
func (_IProxy *IProxyTransactorSession) BanT(_t common.Address, _ban bool, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.BanT(&_IProxy.TransactOpts, _t, _ban, signs)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x310aa0be.
//
// Solidity: function confirmPayee(uint64 _i) returns()
func (_IProxy *IProxyTransactor) ConfirmPayee(opts *bind.TransactOpts, _i uint64) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "confirmPayee", _i)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x310aa0be.
//
// Solidity: function confirmPayee(uint64 _i) returns()
func (_IProxy *IProxySession) ConfirmPayee(_i uint64) (*types.Transaction, error) {
	return _IProxy.Contract.ConfirmPayee(&_IProxy.TransactOpts, _i)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x310aa0be.
//
// Solidity: function confirmPayee(uint64 _i) returns()
func (_IProxy *IProxyTransactorSession) ConfirmPayee(_i uint64) (*types.Transaction, error) {
	return _IProxy.Contract.ConfirmPayee(&_IProxy.TransactOpts, _i)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x8aae670a.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kpPerB, uint256 _ppPerB, bytes _desc) returns()
func (_IProxy *IProxyTransactor) CreateGroup(opts *bind.TransactOpts, _level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kpPerB *big.Int, _ppPerB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "createGroup", _level, _mr, _tr, _kr, _pr, _kpPerB, _ppPerB, _desc)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x8aae670a.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kpPerB, uint256 _ppPerB, bytes _desc) returns()
func (_IProxy *IProxySession) CreateGroup(_level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kpPerB *big.Int, _ppPerB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _IProxy.Contract.CreateGroup(&_IProxy.TransactOpts, _level, _mr, _tr, _kr, _pr, _kpPerB, _ppPerB, _desc)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x8aae670a.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kpPerB, uint256 _ppPerB, bytes _desc) returns()
func (_IProxy *IProxyTransactorSession) CreateGroup(_level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kpPerB *big.Int, _ppPerB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _IProxy.Contract.CreateGroup(&_IProxy.TransactOpts, _level, _mr, _tr, _kr, _pr, _kpPerB, _ppPerB, _desc)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 _money) returns()
func (_IProxy *IProxyTransactor) Pledge(opts *bind.TransactOpts, _i uint64, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "pledge", _i, _money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 _money) returns()
func (_IProxy *IProxySession) Pledge(_i uint64, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.Pledge(&_IProxy.TransactOpts, _i, _money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 _money) returns()
func (_IProxy *IProxyTransactorSession) Pledge(_i uint64, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.Pledge(&_IProxy.TransactOpts, _i, _money)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xe182d54c.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_IProxy *IProxyTransactor) ProWithdraw(opts *bind.TransactOpts, _ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "proWithdraw", _ps, _kis, ksigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xe182d54c.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_IProxy *IProxySession) ProWithdraw(_ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.ProWithdraw(&_IProxy.TransactOpts, _ps, _kis, ksigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xe182d54c.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_IProxy *IProxyTransactorSession) ProWithdraw(_ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.ProWithdraw(&_IProxy.TransactOpts, _ps, _kis, ksigns)
}

// QuitRole is a paid mutator transaction binding the contract method 0xadad6d3a.
//
// Solidity: function quitRole(uint64 _i) returns()
func (_IProxy *IProxyTransactor) QuitRole(opts *bind.TransactOpts, _i uint64) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "quitRole", _i)
}

// QuitRole is a paid mutator transaction binding the contract method 0xadad6d3a.
//
// Solidity: function quitRole(uint64 _i) returns()
func (_IProxy *IProxySession) QuitRole(_i uint64) (*types.Transaction, error) {
	return _IProxy.Contract.QuitRole(&_IProxy.TransactOpts, _i)
}

// QuitRole is a paid mutator transaction binding the contract method 0xadad6d3a.
//
// Solidity: function quitRole(uint64 _i) returns()
func (_IProxy *IProxyTransactorSession) QuitRole(_i uint64) (*types.Transaction, error) {
	return _IProxy.Contract.QuitRole(&_IProxy.TransactOpts, _i)
}

// ReAcc is a paid mutator transaction binding the contract method 0x6526250f.
//
// Solidity: function reAcc() returns()
func (_IProxy *IProxyTransactor) ReAcc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "reAcc")
}

// ReAcc is a paid mutator transaction binding the contract method 0x6526250f.
//
// Solidity: function reAcc() returns()
func (_IProxy *IProxySession) ReAcc() (*types.Transaction, error) {
	return _IProxy.Contract.ReAcc(&_IProxy.TransactOpts)
}

// ReAcc is a paid mutator transaction binding the contract method 0x6526250f.
//
// Solidity: function reAcc() returns()
func (_IProxy *IProxyTransactorSession) ReAcc() (*types.Transaction, error) {
	return _IProxy.Contract.ReAcc(&_IProxy.TransactOpts)
}

// ReRole is a paid mutator transaction binding the contract method 0x898e00f8.
//
// Solidity: function reRole(uint8 _rtype, bytes _extra) returns()
func (_IProxy *IProxyTransactor) ReRole(opts *bind.TransactOpts, _rtype uint8, _extra []byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "reRole", _rtype, _extra)
}

// ReRole is a paid mutator transaction binding the contract method 0x898e00f8.
//
// Solidity: function reRole(uint8 _rtype, bytes _extra) returns()
func (_IProxy *IProxySession) ReRole(_rtype uint8, _extra []byte) (*types.Transaction, error) {
	return _IProxy.Contract.ReRole(&_IProxy.TransactOpts, _rtype, _extra)
}

// ReRole is a paid mutator transaction binding the contract method 0x898e00f8.
//
// Solidity: function reRole(uint8 _rtype, bytes _extra) returns()
func (_IProxy *IProxyTransactorSession) ReRole(_rtype uint8, _extra []byte) (*types.Transaction, error) {
	return _IProxy.Contract.ReRole(&_IProxy.TransactOpts, _rtype, _extra)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_IProxy *IProxyTransactor) Recharge(opts *bind.TransactOpts, _i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "recharge", _i, _ti, isLock, _money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_IProxy *IProxySession) Recharge(_i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.Recharge(&_IProxy.TransactOpts, _i, _ti, isLock, _money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_IProxy *IProxyTransactorSession) Recharge(_i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.Recharge(&_IProxy.TransactOpts, _i, _ti, isLock, _money)
}

// SetDesc is a paid mutator transaction binding the contract method 0x0fc6fdd7.
//
// Solidity: function setDesc(bytes _desc) returns()
func (_IProxy *IProxyTransactor) SetDesc(opts *bind.TransactOpts, _desc []byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "setDesc", _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x0fc6fdd7.
//
// Solidity: function setDesc(bytes _desc) returns()
func (_IProxy *IProxySession) SetDesc(_desc []byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetDesc(&_IProxy.TransactOpts, _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x0fc6fdd7.
//
// Solidity: function setDesc(bytes _desc) returns()
func (_IProxy *IProxyTransactorSession) SetDesc(_desc []byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetDesc(&_IProxy.TransactOpts, _desc)
}

// SetGDesc is a paid mutator transaction binding the contract method 0xaa756343.
//
// Solidity: function setGDesc(uint64 _gi, bytes _desc, bytes[5] signs) returns()
func (_IProxy *IProxyTransactor) SetGDesc(opts *bind.TransactOpts, _gi uint64, _desc []byte, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "setGDesc", _gi, _desc, signs)
}

// SetGDesc is a paid mutator transaction binding the contract method 0xaa756343.
//
// Solidity: function setGDesc(uint64 _gi, bytes _desc, bytes[5] signs) returns()
func (_IProxy *IProxySession) SetGDesc(_gi uint64, _desc []byte, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetGDesc(&_IProxy.TransactOpts, _gi, _desc, signs)
}

// SetGDesc is a paid mutator transaction binding the contract method 0xaa756343.
//
// Solidity: function setGDesc(uint64 _gi, bytes _desc, bytes[5] signs) returns()
func (_IProxy *IProxyTransactorSession) SetGDesc(_gi uint64, _desc []byte, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetGDesc(&_IProxy.TransactOpts, _gi, _desc, signs)
}

// SetGS is a paid mutator transaction binding the contract method 0xf62ea8d1.
//
// Solidity: function setGS(uint64 _gi, uint8 _s, bytes[5] signs) returns()
func (_IProxy *IProxyTransactor) SetGS(opts *bind.TransactOpts, _gi uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "setGS", _gi, _s, signs)
}

// SetGS is a paid mutator transaction binding the contract method 0xf62ea8d1.
//
// Solidity: function setGS(uint64 _gi, uint8 _s, bytes[5] signs) returns()
func (_IProxy *IProxySession) SetGS(_gi uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetGS(&_IProxy.TransactOpts, _gi, _s, signs)
}

// SetGS is a paid mutator transaction binding the contract method 0xf62ea8d1.
//
// Solidity: function setGS(uint64 _gi, uint8 _s, bytes[5] signs) returns()
func (_IProxy *IProxyTransactorSession) SetGS(_gi uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetGS(&_IProxy.TransactOpts, _gi, _s, signs)
}

// SetLevel is a paid mutator transaction binding the contract method 0x713d3b95.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le, bytes[5] signs) returns()
func (_IProxy *IProxyTransactor) SetLevel(opts *bind.TransactOpts, _gi uint64, _le uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "setLevel", _gi, _le, signs)
}

// SetLevel is a paid mutator transaction binding the contract method 0x713d3b95.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le, bytes[5] signs) returns()
func (_IProxy *IProxySession) SetLevel(_gi uint64, _le uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetLevel(&_IProxy.TransactOpts, _gi, _le, signs)
}

// SetLevel is a paid mutator transaction binding the contract method 0x713d3b95.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le, bytes[5] signs) returns()
func (_IProxy *IProxyTransactorSession) SetLevel(_gi uint64, _le uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetLevel(&_IProxy.TransactOpts, _gi, _le, signs)
}

// SetP is a paid mutator transaction binding the contract method 0x5192e18a.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp, bytes[5] signs) returns()
func (_IProxy *IProxyTransactor) SetP(opts *bind.TransactOpts, _gi uint64, _kp *big.Int, _pp *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "setP", _gi, _kp, _pp, signs)
}

// SetP is a paid mutator transaction binding the contract method 0x5192e18a.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp, bytes[5] signs) returns()
func (_IProxy *IProxySession) SetP(_gi uint64, _kp *big.Int, _pp *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetP(&_IProxy.TransactOpts, _gi, _kp, _pp, signs)
}

// SetP is a paid mutator transaction binding the contract method 0x5192e18a.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp, bytes[5] signs) returns()
func (_IProxy *IProxyTransactorSession) SetP(_gi uint64, _kp *big.Int, _pp *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetP(&_IProxy.TransactOpts, _gi, _kp, _pp, signs)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xbd5d54b7.
//
// Solidity: function setPeriod(uint64 _pe, bytes[5] signs) returns()
func (_IProxy *IProxyTransactor) SetPeriod(opts *bind.TransactOpts, _pe uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "setPeriod", _pe, signs)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xbd5d54b7.
//
// Solidity: function setPeriod(uint64 _pe, bytes[5] signs) returns()
func (_IProxy *IProxySession) SetPeriod(_pe uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetPeriod(&_IProxy.TransactOpts, _pe, signs)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xbd5d54b7.
//
// Solidity: function setPeriod(uint64 _pe, bytes[5] signs) returns()
func (_IProxy *IProxyTransactorSession) SetPeriod(_pe uint64, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetPeriod(&_IProxy.TransactOpts, _pe, signs)
}

// SetPlePerB is a paid mutator transaction binding the contract method 0x130b89b1.
//
// Solidity: function setPlePerB(uint64 _gi, uint256 _kpPerB, uint256 ppPerB, bytes[5] signs) returns()
func (_IProxy *IProxyTransactor) SetPlePerB(opts *bind.TransactOpts, _gi uint64, _kpPerB *big.Int, ppPerB *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "setPlePerB", _gi, _kpPerB, ppPerB, signs)
}

// SetPlePerB is a paid mutator transaction binding the contract method 0x130b89b1.
//
// Solidity: function setPlePerB(uint64 _gi, uint256 _kpPerB, uint256 ppPerB, bytes[5] signs) returns()
func (_IProxy *IProxySession) SetPlePerB(_gi uint64, _kpPerB *big.Int, ppPerB *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetPlePerB(&_IProxy.TransactOpts, _gi, _kpPerB, ppPerB, signs)
}

// SetPlePerB is a paid mutator transaction binding the contract method 0x130b89b1.
//
// Solidity: function setPlePerB(uint64 _gi, uint256 _kpPerB, uint256 ppPerB, bytes[5] signs) returns()
func (_IProxy *IProxyTransactorSession) SetPlePerB(_gi uint64, _kpPerB *big.Int, ppPerB *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetPlePerB(&_IProxy.TransactOpts, _gi, _kpPerB, ppPerB, signs)
}

// SetRS is a paid mutator transaction binding the contract method 0x88602357.
//
// Solidity: function setRS(uint64 _i, uint8 _s, bytes[5] signs) returns()
func (_IProxy *IProxyTransactor) SetRS(opts *bind.TransactOpts, _i uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "setRS", _i, _s, signs)
}

// SetRS is a paid mutator transaction binding the contract method 0x88602357.
//
// Solidity: function setRS(uint64 _i, uint8 _s, bytes[5] signs) returns()
func (_IProxy *IProxySession) SetRS(_i uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetRS(&_IProxy.TransactOpts, _i, _s, signs)
}

// SetRS is a paid mutator transaction binding the contract method 0x88602357.
//
// Solidity: function setRS(uint64 _i, uint8 _s, bytes[5] signs) returns()
func (_IProxy *IProxyTransactorSession) SetRS(_i uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetRS(&_IProxy.TransactOpts, _i, _s, signs)
}

// SetRate is a paid mutator transaction binding the contract method 0x4d525421.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr, bytes[5] signs) returns()
func (_IProxy *IProxyTransactor) SetRate(opts *bind.TransactOpts, _gi uint64, _mr uint8, _tr uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "setRate", _gi, _mr, _tr, signs)
}

// SetRate is a paid mutator transaction binding the contract method 0x4d525421.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr, bytes[5] signs) returns()
func (_IProxy *IProxySession) SetRate(_gi uint64, _mr uint8, _tr uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetRate(&_IProxy.TransactOpts, _gi, _mr, _tr, signs)
}

// SetRate is a paid mutator transaction binding the contract method 0x4d525421.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr, bytes[5] signs) returns()
func (_IProxy *IProxyTransactorSession) SetRate(_gi uint64, _mr uint8, _tr uint8, signs [5][]byte) (*types.Transaction, error) {
	return _IProxy.Contract.SetRate(&_IProxy.TransactOpts, _gi, _mr, _tr, signs)
}

// SubOrder is a paid mutator transaction binding the contract method 0x153ede07.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IProxy *IProxyTransactor) SubOrder(opts *bind.TransactOpts, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "subOrder", _oi, uSign, pSign)
}

// SubOrder is a paid mutator transaction binding the contract method 0x153ede07.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IProxy *IProxySession) SubOrder(_oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IProxy.Contract.SubOrder(&_IProxy.TransactOpts, _oi, uSign, pSign)
}

// SubOrder is a paid mutator transaction binding the contract method 0x153ede07.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_IProxy *IProxyTransactorSession) SubOrder(_oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _IProxy.Contract.SubOrder(&_IProxy.TransactOpts, _oi, uSign, pSign)
}

// Unpledge is a paid mutator transaction binding the contract method 0x60985756.
//
// Solidity: function unpledge(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IProxy *IProxyTransactor) Unpledge(opts *bind.TransactOpts, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "unpledge", _i, _ti, _money)
}

// Unpledge is a paid mutator transaction binding the contract method 0x60985756.
//
// Solidity: function unpledge(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IProxy *IProxySession) Unpledge(_i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.Unpledge(&_IProxy.TransactOpts, _i, _ti, _money)
}

// Unpledge is a paid mutator transaction binding the contract method 0x60985756.
//
// Solidity: function unpledge(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IProxy *IProxyTransactorSession) Unpledge(_i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.Unpledge(&_IProxy.TransactOpts, _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IProxy *IProxyTransactor) Withdraw(opts *bind.TransactOpts, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.contract.Transact(opts, "withdraw", _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IProxy *IProxySession) Withdraw(_i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.Withdraw(&_IProxy.TransactOpts, _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_IProxy *IProxyTransactorSession) Withdraw(_i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IProxy.Contract.Withdraw(&_IProxy.TransactOpts, _i, _ti, _money)
}

// IProxyActivateIterator is returned from FilterActivate and is used to iterate over the raw logs and unpacked data for Activate events raised by the IProxy contract.
type IProxyActivateIterator struct {
	Event *IProxyActivate // Event containing the contract specifics and raw log

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
func (it *IProxyActivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProxyActivate)
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
		it.Event = new(IProxyActivate)
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
func (it *IProxyActivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProxyActivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProxyActivate represents a Activate event raised by the IProxy contract.
type IProxyActivate struct {
	I   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterActivate is a free log retrieval operation binding the contract event 0x452a42fb51ebab9c8ebcbb3d536371397d861478788df33e9d2c9ae57765431c.
//
// Solidity: event Activate(uint64 i)
func (_IProxy *IProxyFilterer) FilterActivate(opts *bind.FilterOpts) (*IProxyActivateIterator, error) {

	logs, sub, err := _IProxy.contract.FilterLogs(opts, "Activate")
	if err != nil {
		return nil, err
	}
	return &IProxyActivateIterator{contract: _IProxy.contract, event: "Activate", logs: logs, sub: sub}, nil
}

// WatchActivate is a free log subscription operation binding the contract event 0x452a42fb51ebab9c8ebcbb3d536371397d861478788df33e9d2c9ae57765431c.
//
// Solidity: event Activate(uint64 i)
func (_IProxy *IProxyFilterer) WatchActivate(opts *bind.WatchOpts, sink chan<- *IProxyActivate) (event.Subscription, error) {

	logs, sub, err := _IProxy.contract.WatchLogs(opts, "Activate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProxyActivate)
				if err := _IProxy.contract.UnpackLog(event, "Activate", log); err != nil {
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

// ParseActivate is a log parse operation binding the contract event 0x452a42fb51ebab9c8ebcbb3d536371397d861478788df33e9d2c9ae57765431c.
//
// Solidity: event Activate(uint64 i)
func (_IProxy *IProxyFilterer) ParseActivate(log types.Log) (*IProxyActivate, error) {
	event := new(IProxyActivate)
	if err := _IProxy.contract.UnpackLog(event, "Activate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProxyAddTIterator is returned from FilterAddT and is used to iterate over the raw logs and unpacked data for AddT events raised by the IProxy contract.
type IProxyAddTIterator struct {
	Event *IProxyAddT // Event containing the contract specifics and raw log

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
func (it *IProxyAddTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProxyAddT)
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
		it.Event = new(IProxyAddT)
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
func (it *IProxyAddTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProxyAddTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProxyAddT represents a AddT event raised by the IProxy contract.
type IProxyAddT struct {
	T    common.Address
	Mint bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddT is a free log retrieval operation binding the contract event 0x4954a00722fe171d5c0079b423ebc3b3984d83c24170fdb068dfbd979d250412.
//
// Solidity: event AddT(address t, bool mint)
func (_IProxy *IProxyFilterer) FilterAddT(opts *bind.FilterOpts) (*IProxyAddTIterator, error) {

	logs, sub, err := _IProxy.contract.FilterLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return &IProxyAddTIterator{contract: _IProxy.contract, event: "AddT", logs: logs, sub: sub}, nil
}

// WatchAddT is a free log subscription operation binding the contract event 0x4954a00722fe171d5c0079b423ebc3b3984d83c24170fdb068dfbd979d250412.
//
// Solidity: event AddT(address t, bool mint)
func (_IProxy *IProxyFilterer) WatchAddT(opts *bind.WatchOpts, sink chan<- *IProxyAddT) (event.Subscription, error) {

	logs, sub, err := _IProxy.contract.WatchLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProxyAddT)
				if err := _IProxy.contract.UnpackLog(event, "AddT", log); err != nil {
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

// ParseAddT is a log parse operation binding the contract event 0x4954a00722fe171d5c0079b423ebc3b3984d83c24170fdb068dfbd979d250412.
//
// Solidity: event AddT(address t, bool mint)
func (_IProxy *IProxyFilterer) ParseAddT(log types.Log) (*IProxyAddT, error) {
	event := new(IProxyAddT)
	if err := _IProxy.contract.UnpackLog(event, "AddT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProxyBanTIterator is returned from FilterBanT and is used to iterate over the raw logs and unpacked data for BanT events raised by the IProxy contract.
type IProxyBanTIterator struct {
	Event *IProxyBanT // Event containing the contract specifics and raw log

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
func (it *IProxyBanTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProxyBanT)
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
		it.Event = new(IProxyBanT)
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
func (it *IProxyBanTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProxyBanTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProxyBanT represents a BanT event raised by the IProxy contract.
type IProxyBanT struct {
	T   common.Address
	Ban bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBanT is a free log retrieval operation binding the contract event 0xffaa091f00488e7d9934b8efc3cbe385e3bb89957862cb891e428474b69b6bc9.
//
// Solidity: event BanT(address t, bool ban)
func (_IProxy *IProxyFilterer) FilterBanT(opts *bind.FilterOpts) (*IProxyBanTIterator, error) {

	logs, sub, err := _IProxy.contract.FilterLogs(opts, "BanT")
	if err != nil {
		return nil, err
	}
	return &IProxyBanTIterator{contract: _IProxy.contract, event: "BanT", logs: logs, sub: sub}, nil
}

// WatchBanT is a free log subscription operation binding the contract event 0xffaa091f00488e7d9934b8efc3cbe385e3bb89957862cb891e428474b69b6bc9.
//
// Solidity: event BanT(address t, bool ban)
func (_IProxy *IProxyFilterer) WatchBanT(opts *bind.WatchOpts, sink chan<- *IProxyBanT) (event.Subscription, error) {

	logs, sub, err := _IProxy.contract.WatchLogs(opts, "BanT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProxyBanT)
				if err := _IProxy.contract.UnpackLog(event, "BanT", log); err != nil {
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

// ParseBanT is a log parse operation binding the contract event 0xffaa091f00488e7d9934b8efc3cbe385e3bb89957862cb891e428474b69b6bc9.
//
// Solidity: event BanT(address t, bool ban)
func (_IProxy *IProxyFilterer) ParseBanT(log types.Log) (*IProxyBanT, error) {
	event := new(IProxyBanT)
	if err := _IProxy.contract.UnpackLog(event, "BanT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProxySetGDescIterator is returned from FilterSetGDesc and is used to iterate over the raw logs and unpacked data for SetGDesc events raised by the IProxy contract.
type IProxySetGDescIterator struct {
	Event *IProxySetGDesc // Event containing the contract specifics and raw log

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
func (it *IProxySetGDescIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProxySetGDesc)
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
		it.Event = new(IProxySetGDesc)
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
func (it *IProxySetGDescIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProxySetGDescIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProxySetGDesc represents a SetGDesc event raised by the IProxy contract.
type IProxySetGDesc struct {
	Gi   uint64
	Desc []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetGDesc is a free log retrieval operation binding the contract event 0x402035b2534df57a3c7490784b37ac991fa59096f8e251b49d5364967424b85d.
//
// Solidity: event SetGDesc(uint64 indexed gi, bytes desc)
func (_IProxy *IProxyFilterer) FilterSetGDesc(opts *bind.FilterOpts, gi []uint64) (*IProxySetGDescIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IProxy.contract.FilterLogs(opts, "SetGDesc", giRule)
	if err != nil {
		return nil, err
	}
	return &IProxySetGDescIterator{contract: _IProxy.contract, event: "SetGDesc", logs: logs, sub: sub}, nil
}

// WatchSetGDesc is a free log subscription operation binding the contract event 0x402035b2534df57a3c7490784b37ac991fa59096f8e251b49d5364967424b85d.
//
// Solidity: event SetGDesc(uint64 indexed gi, bytes desc)
func (_IProxy *IProxyFilterer) WatchSetGDesc(opts *bind.WatchOpts, sink chan<- *IProxySetGDesc, gi []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IProxy.contract.WatchLogs(opts, "SetGDesc", giRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProxySetGDesc)
				if err := _IProxy.contract.UnpackLog(event, "SetGDesc", log); err != nil {
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

// ParseSetGDesc is a log parse operation binding the contract event 0x402035b2534df57a3c7490784b37ac991fa59096f8e251b49d5364967424b85d.
//
// Solidity: event SetGDesc(uint64 indexed gi, bytes desc)
func (_IProxy *IProxyFilterer) ParseSetGDesc(log types.Log) (*IProxySetGDesc, error) {
	event := new(IProxySetGDesc)
	if err := _IProxy.contract.UnpackLog(event, "SetGDesc", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProxySetGSIterator is returned from FilterSetGS and is used to iterate over the raw logs and unpacked data for SetGS events raised by the IProxy contract.
type IProxySetGSIterator struct {
	Event *IProxySetGS // Event containing the contract specifics and raw log

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
func (it *IProxySetGSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProxySetGS)
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
		it.Event = new(IProxySetGS)
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
func (it *IProxySetGSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProxySetGSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProxySetGS represents a SetGS event raised by the IProxy contract.
type IProxySetGS struct {
	Gi  uint64
	S   uint8
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetGS is a free log retrieval operation binding the contract event 0x6868b79de8578929833c8106714d184d883afab70f4161a0d447c4f2d899417a.
//
// Solidity: event SetGS(uint64 indexed gi, uint8 _s)
func (_IProxy *IProxyFilterer) FilterSetGS(opts *bind.FilterOpts, gi []uint64) (*IProxySetGSIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IProxy.contract.FilterLogs(opts, "SetGS", giRule)
	if err != nil {
		return nil, err
	}
	return &IProxySetGSIterator{contract: _IProxy.contract, event: "SetGS", logs: logs, sub: sub}, nil
}

// WatchSetGS is a free log subscription operation binding the contract event 0x6868b79de8578929833c8106714d184d883afab70f4161a0d447c4f2d899417a.
//
// Solidity: event SetGS(uint64 indexed gi, uint8 _s)
func (_IProxy *IProxyFilterer) WatchSetGS(opts *bind.WatchOpts, sink chan<- *IProxySetGS, gi []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IProxy.contract.WatchLogs(opts, "SetGS", giRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProxySetGS)
				if err := _IProxy.contract.UnpackLog(event, "SetGS", log); err != nil {
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

// ParseSetGS is a log parse operation binding the contract event 0x6868b79de8578929833c8106714d184d883afab70f4161a0d447c4f2d899417a.
//
// Solidity: event SetGS(uint64 indexed gi, uint8 _s)
func (_IProxy *IProxyFilterer) ParseSetGS(log types.Log) (*IProxySetGS, error) {
	event := new(IProxySetGS)
	if err := _IProxy.contract.UnpackLog(event, "SetGS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProxySetLevelIterator is returned from FilterSetLevel and is used to iterate over the raw logs and unpacked data for SetLevel events raised by the IProxy contract.
type IProxySetLevelIterator struct {
	Event *IProxySetLevel // Event containing the contract specifics and raw log

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
func (it *IProxySetLevelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProxySetLevel)
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
		it.Event = new(IProxySetLevel)
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
func (it *IProxySetLevelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProxySetLevelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProxySetLevel represents a SetLevel event raised by the IProxy contract.
type IProxySetLevel struct {
	Gi  uint64
	Le  uint8
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetLevel is a free log retrieval operation binding the contract event 0x5aee2b03a7ccab0db1135bf51f379d3b74e615f96cb0e7d43db42767638cf899.
//
// Solidity: event SetLevel(uint64 indexed gi, uint8 le)
func (_IProxy *IProxyFilterer) FilterSetLevel(opts *bind.FilterOpts, gi []uint64) (*IProxySetLevelIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IProxy.contract.FilterLogs(opts, "SetLevel", giRule)
	if err != nil {
		return nil, err
	}
	return &IProxySetLevelIterator{contract: _IProxy.contract, event: "SetLevel", logs: logs, sub: sub}, nil
}

// WatchSetLevel is a free log subscription operation binding the contract event 0x5aee2b03a7ccab0db1135bf51f379d3b74e615f96cb0e7d43db42767638cf899.
//
// Solidity: event SetLevel(uint64 indexed gi, uint8 le)
func (_IProxy *IProxyFilterer) WatchSetLevel(opts *bind.WatchOpts, sink chan<- *IProxySetLevel, gi []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IProxy.contract.WatchLogs(opts, "SetLevel", giRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProxySetLevel)
				if err := _IProxy.contract.UnpackLog(event, "SetLevel", log); err != nil {
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

// ParseSetLevel is a log parse operation binding the contract event 0x5aee2b03a7ccab0db1135bf51f379d3b74e615f96cb0e7d43db42767638cf899.
//
// Solidity: event SetLevel(uint64 indexed gi, uint8 le)
func (_IProxy *IProxyFilterer) ParseSetLevel(log types.Log) (*IProxySetLevel, error) {
	event := new(IProxySetLevel)
	if err := _IProxy.contract.UnpackLog(event, "SetLevel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProxySetPeriodIterator is returned from FilterSetPeriod and is used to iterate over the raw logs and unpacked data for SetPeriod events raised by the IProxy contract.
type IProxySetPeriodIterator struct {
	Event *IProxySetPeriod // Event containing the contract specifics and raw log

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
func (it *IProxySetPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProxySetPeriod)
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
		it.Event = new(IProxySetPeriod)
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
func (it *IProxySetPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProxySetPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProxySetPeriod represents a SetPeriod event raised by the IProxy contract.
type IProxySetPeriod struct {
	Pe  uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetPeriod is a free log retrieval operation binding the contract event 0x11718e279b88c35d197671951471b347900bac4b36c87d6f5a34ede95b3823d0.
//
// Solidity: event SetPeriod(uint64 pe)
func (_IProxy *IProxyFilterer) FilterSetPeriod(opts *bind.FilterOpts) (*IProxySetPeriodIterator, error) {

	logs, sub, err := _IProxy.contract.FilterLogs(opts, "SetPeriod")
	if err != nil {
		return nil, err
	}
	return &IProxySetPeriodIterator{contract: _IProxy.contract, event: "SetPeriod", logs: logs, sub: sub}, nil
}

// WatchSetPeriod is a free log subscription operation binding the contract event 0x11718e279b88c35d197671951471b347900bac4b36c87d6f5a34ede95b3823d0.
//
// Solidity: event SetPeriod(uint64 pe)
func (_IProxy *IProxyFilterer) WatchSetPeriod(opts *bind.WatchOpts, sink chan<- *IProxySetPeriod) (event.Subscription, error) {

	logs, sub, err := _IProxy.contract.WatchLogs(opts, "SetPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProxySetPeriod)
				if err := _IProxy.contract.UnpackLog(event, "SetPeriod", log); err != nil {
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

// ParseSetPeriod is a log parse operation binding the contract event 0x11718e279b88c35d197671951471b347900bac4b36c87d6f5a34ede95b3823d0.
//
// Solidity: event SetPeriod(uint64 pe)
func (_IProxy *IProxyFilterer) ParseSetPeriod(log types.Log) (*IProxySetPeriod, error) {
	event := new(IProxySetPeriod)
	if err := _IProxy.contract.UnpackLog(event, "SetPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProxySetPlePerBIterator is returned from FilterSetPlePerB and is used to iterate over the raw logs and unpacked data for SetPlePerB events raised by the IProxy contract.
type IProxySetPlePerBIterator struct {
	Event *IProxySetPlePerB // Event containing the contract specifics and raw log

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
func (it *IProxySetPlePerBIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProxySetPlePerB)
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
		it.Event = new(IProxySetPlePerB)
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
func (it *IProxySetPlePerBIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProxySetPlePerBIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProxySetPlePerB represents a SetPlePerB event raised by the IProxy contract.
type IProxySetPlePerB struct {
	Gi     uint64
	KpPerB *big.Int
	PpPerB *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetPlePerB is a free log retrieval operation binding the contract event 0x4f8fb948a05297fa2147f00955cede56207c0eb67e57201c6a3fdc53588f27b8.
//
// Solidity: event SetPlePerB(uint64 indexed gi, uint256 kpPerB, uint256 ppPerB)
func (_IProxy *IProxyFilterer) FilterSetPlePerB(opts *bind.FilterOpts, gi []uint64) (*IProxySetPlePerBIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IProxy.contract.FilterLogs(opts, "SetPlePerB", giRule)
	if err != nil {
		return nil, err
	}
	return &IProxySetPlePerBIterator{contract: _IProxy.contract, event: "SetPlePerB", logs: logs, sub: sub}, nil
}

// WatchSetPlePerB is a free log subscription operation binding the contract event 0x4f8fb948a05297fa2147f00955cede56207c0eb67e57201c6a3fdc53588f27b8.
//
// Solidity: event SetPlePerB(uint64 indexed gi, uint256 kpPerB, uint256 ppPerB)
func (_IProxy *IProxyFilterer) WatchSetPlePerB(opts *bind.WatchOpts, sink chan<- *IProxySetPlePerB, gi []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IProxy.contract.WatchLogs(opts, "SetPlePerB", giRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProxySetPlePerB)
				if err := _IProxy.contract.UnpackLog(event, "SetPlePerB", log); err != nil {
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

// ParseSetPlePerB is a log parse operation binding the contract event 0x4f8fb948a05297fa2147f00955cede56207c0eb67e57201c6a3fdc53588f27b8.
//
// Solidity: event SetPlePerB(uint64 indexed gi, uint256 kpPerB, uint256 ppPerB)
func (_IProxy *IProxyFilterer) ParseSetPlePerB(log types.Log) (*IProxySetPlePerB, error) {
	event := new(IProxySetPlePerB)
	if err := _IProxy.contract.UnpackLog(event, "SetPlePerB", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProxySetRSIterator is returned from FilterSetRS and is used to iterate over the raw logs and unpacked data for SetRS events raised by the IProxy contract.
type IProxySetRSIterator struct {
	Event *IProxySetRS // Event containing the contract specifics and raw log

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
func (it *IProxySetRSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProxySetRS)
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
		it.Event = new(IProxySetRS)
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
func (it *IProxySetRSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProxySetRSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProxySetRS represents a SetRS event raised by the IProxy contract.
type IProxySetRS struct {
	I   uint64
	S   uint8
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetRS is a free log retrieval operation binding the contract event 0xf101cf457bf0c852aed48bf40478c9eb079d6b5a23f9dc8c00c467712b3d293e.
//
// Solidity: event SetRS(uint64 i, uint8 _s)
func (_IProxy *IProxyFilterer) FilterSetRS(opts *bind.FilterOpts) (*IProxySetRSIterator, error) {

	logs, sub, err := _IProxy.contract.FilterLogs(opts, "SetRS")
	if err != nil {
		return nil, err
	}
	return &IProxySetRSIterator{contract: _IProxy.contract, event: "SetRS", logs: logs, sub: sub}, nil
}

// WatchSetRS is a free log subscription operation binding the contract event 0xf101cf457bf0c852aed48bf40478c9eb079d6b5a23f9dc8c00c467712b3d293e.
//
// Solidity: event SetRS(uint64 i, uint8 _s)
func (_IProxy *IProxyFilterer) WatchSetRS(opts *bind.WatchOpts, sink chan<- *IProxySetRS) (event.Subscription, error) {

	logs, sub, err := _IProxy.contract.WatchLogs(opts, "SetRS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProxySetRS)
				if err := _IProxy.contract.UnpackLog(event, "SetRS", log); err != nil {
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

// ParseSetRS is a log parse operation binding the contract event 0xf101cf457bf0c852aed48bf40478c9eb079d6b5a23f9dc8c00c467712b3d293e.
//
// Solidity: event SetRS(uint64 i, uint8 _s)
func (_IProxy *IProxyFilterer) ParseSetRS(log types.Log) (*IProxySetRS, error) {
	event := new(IProxySetRS)
	if err := _IProxy.contract.UnpackLog(event, "SetRS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProxySetRateIterator is returned from FilterSetRate and is used to iterate over the raw logs and unpacked data for SetRate events raised by the IProxy contract.
type IProxySetRateIterator struct {
	Event *IProxySetRate // Event containing the contract specifics and raw log

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
func (it *IProxySetRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProxySetRate)
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
		it.Event = new(IProxySetRate)
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
func (it *IProxySetRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProxySetRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProxySetRate represents a SetRate event raised by the IProxy contract.
type IProxySetRate struct {
	Gi  uint64
	Mr  uint8
	Kr  uint8
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetRate is a free log retrieval operation binding the contract event 0x1ee4799609f0de30722e40275c40d1dd80883a00e6d3ceda4f5d4328968c9a4a.
//
// Solidity: event SetRate(uint64 indexed gi, uint8 mr, uint8 kr)
func (_IProxy *IProxyFilterer) FilterSetRate(opts *bind.FilterOpts, gi []uint64) (*IProxySetRateIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IProxy.contract.FilterLogs(opts, "SetRate", giRule)
	if err != nil {
		return nil, err
	}
	return &IProxySetRateIterator{contract: _IProxy.contract, event: "SetRate", logs: logs, sub: sub}, nil
}

// WatchSetRate is a free log subscription operation binding the contract event 0x1ee4799609f0de30722e40275c40d1dd80883a00e6d3ceda4f5d4328968c9a4a.
//
// Solidity: event SetRate(uint64 indexed gi, uint8 mr, uint8 kr)
func (_IProxy *IProxyFilterer) WatchSetRate(opts *bind.WatchOpts, sink chan<- *IProxySetRate, gi []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IProxy.contract.WatchLogs(opts, "SetRate", giRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProxySetRate)
				if err := _IProxy.contract.UnpackLog(event, "SetRate", log); err != nil {
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

// ParseSetRate is a log parse operation binding the contract event 0x1ee4799609f0de30722e40275c40d1dd80883a00e6d3ceda4f5d4328968c9a4a.
//
// Solidity: event SetRate(uint64 indexed gi, uint8 mr, uint8 kr)
func (_IProxy *IProxyFilterer) ParseSetRate(log types.Log) (*IProxySetRate, error) {
	event := new(IProxySetRate)
	if err := _IProxy.contract.UnpackLog(event, "SetRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProxySetkpPIterator is returned from FilterSetkpP and is used to iterate over the raw logs and unpacked data for SetkpP events raised by the IProxy contract.
type IProxySetkpPIterator struct {
	Event *IProxySetkpP // Event containing the contract specifics and raw log

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
func (it *IProxySetkpPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProxySetkpP)
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
		it.Event = new(IProxySetkpP)
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
func (it *IProxySetkpPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProxySetkpPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProxySetkpP represents a SetkpP event raised by the IProxy contract.
type IProxySetkpP struct {
	Gi  uint64
	Kp  *big.Int
	Pp  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetkpP is a free log retrieval operation binding the contract event 0xaa3e90b31b32985da56d5c241157e7caec89ed412da1ba26278f59b165a21585.
//
// Solidity: event SetkpP(uint64 indexed gi, uint256 kp, uint256 pp)
func (_IProxy *IProxyFilterer) FilterSetkpP(opts *bind.FilterOpts, gi []uint64) (*IProxySetkpPIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IProxy.contract.FilterLogs(opts, "SetkpP", giRule)
	if err != nil {
		return nil, err
	}
	return &IProxySetkpPIterator{contract: _IProxy.contract, event: "SetkpP", logs: logs, sub: sub}, nil
}

// WatchSetkpP is a free log subscription operation binding the contract event 0xaa3e90b31b32985da56d5c241157e7caec89ed412da1ba26278f59b165a21585.
//
// Solidity: event SetkpP(uint64 indexed gi, uint256 kp, uint256 pp)
func (_IProxy *IProxyFilterer) WatchSetkpP(opts *bind.WatchOpts, sink chan<- *IProxySetkpP, gi []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _IProxy.contract.WatchLogs(opts, "SetkpP", giRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProxySetkpP)
				if err := _IProxy.contract.UnpackLog(event, "SetkpP", log); err != nil {
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

// ParseSetkpP is a log parse operation binding the contract event 0xaa3e90b31b32985da56d5c241157e7caec89ed412da1ba26278f59b165a21585.
//
// Solidity: event SetkpP(uint64 indexed gi, uint256 kp, uint256 pp)
func (_IProxy *IProxyFilterer) ParseSetkpP(log types.Log) (*IProxySetkpP, error) {
	event := new(IProxySetkpP)
	if err := _IProxy.contract.UnpackLog(event, "SetkpP", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyABI is the input ABI used to generate the binding from.
const ProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inst\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"Activate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"mint\",\"type\":\"bool\"}],\"name\":\"AddT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"ban\",\"type\":\"bool\"}],\"name\":\"BanT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"desc\",\"type\":\"bytes\"}],\"name\":\"SetGDesc\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_s\",\"type\":\"uint8\"}],\"name\":\"SetGS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"le\",\"type\":\"uint8\"}],\"name\":\"SetLevel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"pe\",\"type\":\"uint64\"}],\"name\":\"SetPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kpPerB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ppPerB\",\"type\":\"uint256\"}],\"name\":\"SetPlePerB\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_s\",\"type\":\"uint8\"}],\"name\":\"SetRS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"mr\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kr\",\"type\":\"uint8\"}],\"name\":\"SetRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pp\",\"type\":\"uint256\"}],\"name\":\"SetkpP\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"uSign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pSign\",\"type\":\"bytes\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"uint64[]\",\"name\":\"_kis\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_mint\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"addT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"addToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_p\",\"type\":\"address\"}],\"name\":\"alterPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isBan\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"banT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"confirmPayee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_level\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tr\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_kpPerB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ppPerB\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"}],\"name\":\"createGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inst\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structPWIn\",\"name\":\"_ps\",\"type\":\"tuple\"},{\"internalType\":\"uint64[]\",\"name\":\"_kis\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"name\":\"proWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"quitRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reAcc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_rtype\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_vk\",\"type\":\"bytes\"}],\"name\":\"reRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"}],\"name\":\"setDesc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_desc\",\"type\":\"bytes\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setGDesc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_s\",\"type\":\"uint8\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setGS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_le\",\"type\":\"uint8\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setLevel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_kp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pp\",\"type\":\"uint256\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pe\",\"type\":\"uint64\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_kPB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pPB\",\"type\":\"uint256\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setPlePerB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_s\",\"type\":\"uint8\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setRS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tr\",\"type\":\"uint8\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"tIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"internalType\":\"structOrderIn\",\"name\":\"_oi\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"uSign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pSign\",\"type\":\"bytes\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"unpledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ProxyFuncSigs maps the 4-byte function signature to its string representation.
var ProxyFuncSigs = map[string]string{
	"843d1b49": "activate(uint64,bytes[5])",
	"c05bdcd0": "addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),bytes,bytes)",
	"97f2352c": "addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),uint64[],bytes[])",
	"68612348": "addT(address,bool,bytes[5])",
	"a246de42": "addToGroup(uint64)",
	"2915824a": "alterPayee(uint64,address)",
	"7a743984": "banT(address,bool,bytes[5])",
	"310aa0be": "confirmPayee(uint64)",
	"8aae670a": "createGroup(uint8,uint8,uint8,uint256,uint256,uint256,uint256,bytes)",
	"bd6b2222": "inst()",
	"d23f7482": "pledge(uint64,uint256)",
	"e182d54c": "proWithdraw((uint64,uint8,uint256,uint256),uint64[],bytes[])",
	"adad6d3a": "quitRole(uint64)",
	"6526250f": "reAcc()",
	"898e00f8": "reRole(uint8,bytes)",
	"24d11d40": "recharge(uint64,uint8,bool,uint256)",
	"0fc6fdd7": "setDesc(bytes)",
	"aa756343": "setGDesc(uint64,bytes,bytes[5])",
	"f62ea8d1": "setGS(uint64,uint8,bytes[5])",
	"713d3b95": "setLevel(uint64,uint8,bytes[5])",
	"5192e18a": "setP(uint64,uint256,uint256,bytes[5])",
	"bd5d54b7": "setPeriod(uint64,bytes[5])",
	"130b89b1": "setPlePerB(uint64,uint256,uint256,bytes[5])",
	"88602357": "setRS(uint64,uint8,bytes[5])",
	"4d525421": "setRate(uint64,uint8,uint8,bytes[5])",
	"153ede07": "subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256),bytes,bytes)",
	"60985756": "unpledge(uint64,uint8,uint256)",
	"54fd4d50": "version()",
	"259c6d5e": "withdraw(uint64,uint8,uint256)",
}

// ProxyBin is the compiled bytecode used for deploying new contracts.
var ProxyBin = "0x60806040526000805461ffff191660021790553480156200001f57600080fd5b5060405162002aa538038062002aa5833981016040819052620000429162000070565b600080546001600160a01b03909216620100000262010000600160b01b0319909216919091179055620000a2565b6000602082840312156200008357600080fd5b81516001600160a01b03811681146200009b57600080fd5b9392505050565b6129f380620000b26000396000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c80637a74398411610104578063aa756343116100a2578063c05bdcd011610071578063c05bdcd0146103d7578063d23f7482146103ea578063e182d54c146103fd578063f62ea8d11461041057600080fd5b8063aa7563431461036d578063adad6d3a14610380578063bd5d54b714610393578063bd6b2222146103a657600080fd5b8063898e00f8116100de578063898e00f8146103215780638aae670a1461033457806397f2352c14610347578063a246de421461035a57600080fd5b80637a743984146102e8578063843d1b49146102fb578063886023571461030e57600080fd5b80634d52542111610171578063609857561161014b57806360985756146102a75780636526250f146102ba57806368612348146102c2578063713d3b95146102d557600080fd5b80634d5254211461025b5780635192e18a1461026e57806354fd4d501461028157600080fd5b806324d11d40116101ad57806324d11d401461020f578063259c6d5e146102225780632915824a14610235578063310aa0be1461024857600080fd5b80630fc6fdd7146101d4578063130b89b1146101e9578063153ede07146101fc575b600080fd5b6101e76101e2366004611cfd565b610423565b005b6101e76101f7366004611de2565b6104f7565b6101e761020a366004611f13565b610617565b6101e761021d366004611f9a565b6106f1565b6101e7610230366004611fe5565b6107eb565b6101e7610243366004612039565b6108aa565b6101e7610256366004612070565b610993565b6101e7610269366004612092565b610a42565b6101e761027c366004611de2565b610b5c565b60005461028f9061ffff1681565b60405161ffff90911681526020015b60405180910390f35b6101e76102b5366004611fe5565b610c72565b6101e7610d31565b6101e76102d03660046120e8565b610e00565b6101e76102e336600461213d565b610f1d565b6101e76102f63660046120e8565b611037565b6101e7610309366004612169565b61114b565b6101e761031c36600461213d565b61125e565b6101e761032f3660046121b6565b611373565b6101e76103423660046121f9565b611414565b6101e76103553660046123a2565b6114fb565b6101e7610368366004612070565b61159e565b6101e761037b36600461240f565b61164d565b6101e761038e366004612070565b611759565b6101e76103a1366004612169565b611808565b6000546103bf906201000090046001600160a01b031681565b6040516001600160a01b03909116815260200161029e565b6101e76103e5366004611f13565b611913565b6101e76103f8366004612478565b6119b6565b6101e761040b3660046124a2565b611a6c565b6101e761041e36600461213d565b611b0f565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015610471573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610495919061253f565b6001600160a01b031663722cc4a433836040518363ffffffff1660e01b81526004016104c29291906125a2565b600060405180830381600087803b1580156104dc57600080fd5b505af11580156104f0573d6000803e3d6000fd5b5050505050565b600054604051633ec7d5b960e01b815260676004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015610545573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610569919061253f565b6001600160a01b031663130b89b1858585856040518563ffffffff1660e01b815260040161059a949392919061260b565b600060405180830381600087803b1580156105b457600080fd5b505af11580156105c8573d6000803e3d6000fd5b505060408051868152602081018690526001600160401b03881693507f4f8fb948a05297fa2147f00955cede56207c0eb67e57201c6a3fdc53588f27b89250015b60405180910390a250505050565b600054604051633ec7d5b960e01b815260666004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015610665573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610689919061253f565b6001600160a01b03166342f45166338585856040518563ffffffff1660e01b81526004016106ba94939291906126a9565b600060405180830381600087803b1580156106d457600080fd5b505af11580156106e8573d6000803e3d6000fd5b50505050505050565b600054604051633ec7d5b960e01b815260666004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa15801561073f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610763919061253f565b60405163f661f9e360e01b81523360048201526001600160401b038616602482015260ff851660448201528315156064820152608481018390526001600160a01b03919091169063f661f9e39060a401600060405180830381600087803b1580156107cd57600080fd5b505af11580156107e1573d6000803e3d6000fd5b5050505050505050565b600054604051633ec7d5b960e01b815260676004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015610839573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085d919061253f565b604051634da63abd60e11b81523360048201526001600160401b038516602482015260ff84166044820152606481018390526001600160a01b039190911690639b4c757a906084016106ba565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa1580156108f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061091c919061253f565b6040516312964f7560e21b81523360048201526001600160401b03841660248201526001600160a01b0383811660448301529190911690634a593dd4906064015b600060405180830381600087803b15801561097757600080fd5b505af115801561098b573d6000803e3d6000fd5b505050505050565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa1580156109e1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a05919061253f565b604051630a38216560e31b81523360048201526001600160401b03831660248201526001600160a01b0391909116906351c10b28906044016104c2565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015610a90573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab4919061253f565b6001600160a01b0316634d525421858585856040518563ffffffff1660e01b8152600401610ae594939291906126fa565b600060405180830381600087803b158015610aff57600080fd5b505af1158015610b13573d6000803e3d6000fd5b50506040805160ff8088168252861660208201526001600160401b03881693507f1ee4799609f0de30722e40275c40d1dd80883a00e6d3ceda4f5d4328968c9a4a925001610609565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015610baa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bce919061253f565b6001600160a01b0316635192e18a858585856040518563ffffffff1660e01b8152600401610bff949392919061260b565b600060405180830381600087803b158015610c1957600080fd5b505af1158015610c2d573d6000803e3d6000fd5b505060408051868152602081018690526001600160401b03881693507faa3e90b31b32985da56d5c241157e7caec89ed412da1ba26278f59b165a21585925001610609565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015610cc0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce4919061253f565b604051630a866aed60e01b81523360048201526001600160401b038516602482015260ff84166044820152606481018390526001600160a01b039190911690630a866aed906084016106ba565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015610d7f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610da3919061253f565b6040516377fe57e760e11b81523360048201526001600160a01b03919091169063effcafce90602401600060405180830381600087803b158015610de657600080fd5b505af1158015610dfa573d6000803e3d6000fd5b50505050565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015610e4e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e72919061253f565b6001600160a01b031663686123488484846040518463ffffffff1660e01b8152600401610ea19392919061272e565b600060405180830381600087803b158015610ebb57600080fd5b505af1158015610ecf573d6000803e3d6000fd5b5050604080516001600160a01b038716815285151560208201527f4954a00722fe171d5c0079b423ebc3b3984d83c24170fdb068dfbd979d25041293500190505b60405180910390a1505050565b600054604051633ec7d5b960e01b815260676004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015610f6b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f8f919061253f565b6001600160a01b031663713d3b958484846040518463ffffffff1660e01b8152600401610fbe93929190612763565b600060405180830381600087803b158015610fd857600080fd5b505af1158015610fec573d6000803e3d6000fd5b505060405160ff851681526001600160401b03861692507f5aee2b03a7ccab0db1135bf51f379d3b74e615f96cb0e7d43db42767638cf89991506020015b60405180910390a2505050565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015611085573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110a9919061253f565b6001600160a01b0316637a7439848484846040518463ffffffff1660e01b81526004016110d89392919061272e565b600060405180830381600087803b1580156110f257600080fd5b505af1158015611106573d6000803e3d6000fd5b5050604080516001600160a01b038716815285151560208201527fffaa091f00488e7d9934b8efc3cbe385e3bb89957862cb891e428474b69b6bc99350019050610f10565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015611199573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111bd919061253f565b6001600160a01b031663843d1b4983836040518363ffffffff1660e01b81526004016111ea92919061278e565b600060405180830381600087803b15801561120457600080fd5b505af1158015611218573d6000803e3d6000fd5b50506040516001600160401b03851681527f452a42fb51ebab9c8ebcbb3d536371397d861478788df33e9d2c9ae57765431c925060200190505b60405180910390a15050565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa1580156112ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112d0919061253f565b6001600160a01b031663886023578484846040518463ffffffff1660e01b81526004016112ff93929190612763565b600060405180830381600087803b15801561131957600080fd5b505af115801561132d573d6000803e3d6000fd5b5050604080516001600160401b038716815260ff861660208201527ff101cf457bf0c852aed48bf40478c9eb079d6b5a23f9dc8c00c467712b3d293e9350019050610f10565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa1580156113c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113e5919061253f565b6001600160a01b0316639ccfe4ff3384846040518463ffffffff1660e01b815260040161095d939291906127b0565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015611462573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611486919061253f565b6001600160a01b0316638aae670a89898989898989896040518963ffffffff1660e01b81526004016114bf9897969594939291906127dd565b600060405180830381600087803b1580156114d957600080fd5b505af11580156114ed573d6000803e3d6000fd5b505050505050505050505050565b600054604051633ec7d5b960e01b815260666004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015611549573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061156d919061253f565b6001600160a01b031663af99c59a338585856040518563ffffffff1660e01b81526004016106ba94939291906128cc565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa1580156115ec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611610919061253f565b60405163d963723160e01b81523360048201526001600160401b03831660248201526001600160a01b03919091169063d9637231906044016104c2565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa15801561169b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116bf919061253f565b6001600160a01b031663aa7563438484846040518463ffffffff1660e01b81526004016116ee93929190612912565b600060405180830381600087803b15801561170857600080fd5b505af115801561171c573d6000803e3d6000fd5b50505050826001600160401b03167f402035b2534df57a3c7490784b37ac991fa59096f8e251b49d5364967424b85d8360405161102a9190612946565b600054604051633ec7d5b960e01b815260676004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa1580156117a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117cb919061253f565b60405163d698569960e01b81523360048201526001600160401b03831660248201526001600160a01b03919091169063d6985699906044016104c2565b600054604051633ec7d5b960e01b815260676004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015611856573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061187a919061253f565b6001600160a01b031663bd5d54b783836040518363ffffffff1660e01b81526004016118a792919061278e565b600060405180830381600087803b1580156118c157600080fd5b505af11580156118d5573d6000803e3d6000fd5b50506040516001600160401b03851681527f11718e279b88c35d197671951471b347900bac4b36c87d6f5a34ede95b3823d092506020019050611252565b600054604051633ec7d5b960e01b815260666004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015611961573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611985919061253f565b6001600160a01b031663ae9d0b40338585856040518563ffffffff1660e01b81526004016106ba94939291906126a9565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015611a04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a28919061253f565b604051630d11731760e11b81523360048201526001600160401b0384166024820152604481018390526001600160a01b039190911690631a22e62e9060640161095d565b600054604051633ec7d5b960e01b815260666004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015611aba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ade919061253f565b6001600160a01b03166354aa6642338585856040518563ffffffff1660e01b81526004016106ba9493929190612959565b600054604051633ec7d5b960e01b815260656004820152620100009091046001600160a01b031690633ec7d5b990602401602060405180830381865afa158015611b5d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b81919061253f565b6001600160a01b031663f62ea8d18484846040518463ffffffff1660e01b8152600401611bb093929190612763565b600060405180830381600087803b158015611bca57600080fd5b505af1158015611bde573d6000803e3d6000fd5b505060405160ff851681526001600160401b03861692507f6868b79de8578929833c8106714d184d883afab70f4161a0d447c4f2d899417a915060200161102a565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b0381118282101715611c5857611c58611c20565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611c8657611c86611c20565b604052919050565b600082601f830112611c9f57600080fd5b81356001600160401b03811115611cb857611cb8611c20565b611ccb601f8201601f1916602001611c5e565b818152846020838601011115611ce057600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215611d0f57600080fd5b81356001600160401b03811115611d2557600080fd5b611d3184828501611c8e565b949350505050565b80356001600160401b0381168114611d5057600080fd5b919050565b600082601f830112611d6657600080fd5b60405160a081016001600160401b038282108183111715611d8957611d89611c20565b8160405282915060a0850186811115611da157600080fd5b855b81811015611dd657803583811115611dbb5760008081fd5b611dc789828a01611c8e565b85525060209384019301611da3565b50929695505050505050565b60008060008060808587031215611df857600080fd5b611e0185611d39565b9350602085013592506040850135915060608501356001600160401b03811115611e2a57600080fd5b611e3687828801611d55565b91505092959194509250565b803560ff81168114611d5057600080fd5b6000610100808385031215611e6757600080fd5b604051908101906001600160401b0382118183101715611e8957611e89611c20565b81604052809250611e9984611d39565b8152611ea760208501611d39565b6020820152611eb860408501611d39565b6040820152611ec960608501611d39565b6060820152611eda60808501611d39565b6080820152611eeb60a08501611d39565b60a0820152611efc60c08501611e42565b60c082015260e084013560e0820152505092915050565b60008060006101408486031215611f2957600080fd5b611f338585611e53565b92506101008401356001600160401b0380821115611f5057600080fd5b611f5c87838801611c8e565b9350610120860135915080821115611f7357600080fd5b50611f8086828701611c8e565b9150509250925092565b80358015158114611d5057600080fd5b60008060008060808587031215611fb057600080fd5b611fb985611d39565b9350611fc760208601611e42565b9250611fd560408601611f8a565b9396929550929360600135925050565b600080600060608486031215611ffa57600080fd5b61200384611d39565b925061201160208501611e42565b9150604084013590509250925092565b6001600160a01b038116811461203657600080fd5b50565b6000806040838503121561204c57600080fd5b61205583611d39565b9150602083013561206581612021565b809150509250929050565b60006020828403121561208257600080fd5b61208b82611d39565b9392505050565b600080600080608085870312156120a857600080fd5b6120b185611d39565b93506120bf60208601611e42565b92506120cd60408601611e42565b915060608501356001600160401b03811115611e2a57600080fd5b6000806000606084860312156120fd57600080fd5b833561210881612021565b925061211660208501611f8a565b915060408401356001600160401b0381111561213157600080fd5b611f8086828701611d55565b60008060006060848603121561215257600080fd5b61215b84611d39565b925061211660208501611e42565b6000806040838503121561217c57600080fd5b61218583611d39565b915060208301356001600160401b038111156121a057600080fd5b6121ac85828601611d55565b9150509250929050565b600080604083850312156121c957600080fd5b6121d283611e42565b915060208301356001600160401b038111156121ed57600080fd5b6121ac85828601611c8e565b600080600080600080600080610100898b03121561221657600080fd5b61221f89611e42565b975061222d60208a01611e42565b965061223b60408a01611e42565b9550606089013594506080890135935060a0890135925060c0890135915060e08901356001600160401b0381111561227257600080fd5b61227e8b828c01611c8e565b9150509295985092959890939650565b60006001600160401b038211156122a7576122a7611c20565b5060051b60200190565b600082601f8301126122c257600080fd5b813560206122d76122d28361228e565b611c5e565b82815260059290921b840181019181810190868411156122f657600080fd5b8286015b848110156123185761230b81611d39565b83529183019183016122fa565b509695505050505050565b600082601f83011261233457600080fd5b813560206123446122d28361228e565b82815260059290921b8401810191818101908684111561236357600080fd5b8286015b848110156123185780356001600160401b038111156123865760008081fd5b6123948986838b0101611c8e565b845250918301918301612367565b600080600061014084860312156123b857600080fd5b6123c28585611e53565b92506101008401356001600160401b03808211156123df57600080fd5b6123eb878388016122b1565b935061012086013591508082111561240257600080fd5b50611f8086828701612323565b60008060006060848603121561242457600080fd5b61242d84611d39565b925060208401356001600160401b038082111561244957600080fd5b61245587838801611c8e565b9350604086013591508082111561246b57600080fd5b50611f8086828701611d55565b6000806040838503121561248b57600080fd5b61249483611d39565b946020939093013593505050565b600080600083850360c08112156124b857600080fd5b60808112156124c657600080fd5b506124cf611c36565b6124d885611d39565b81526124e660208601611e42565b602082015260408501356040820152606085013560608201528093505060808401356001600160401b038082111561251d57600080fd5b612529878388016122b1565b935060a086013591508082111561240257600080fd5b60006020828403121561255157600080fd5b815161208b81612021565b6000815180845260005b8181101561258257602081850181015186830182015201612566565b506000602082860101526020601f19601f83011685010191505092915050565b6001600160a01b0383168152604060208201819052600090611d319083018461255c565b60008260a081018360005b60058110156126005783830387526125ea83835161255c565b60209788019790935091909101906001016125d1565b509095945050505050565b6001600160401b038516815283602082015282604082015260806060820152600061263960808301846125c6565b9695505050505050565b6001600160401b038082511683528060208301511660208401528060408301511660408401528060608301511660608401528060808301511660808401528060a08301511660a08401525060ff60c08201511660c083015260e081015160e08301525050565b6001600160a01b038516815260006101606126c76020840187612643565b806101208401526126da8184018661255c565b90508281036101408401526126ef818561255c565b979650505050505050565b6001600160401b038516815260ff8416602082015260ff8316604082015260806060820152600061263960808301846125c6565b6001600160a01b0384168152821515602082015260606040820181905260009061275a908301846125c6565b95945050505050565b6001600160401b038416815260ff8316602082015260606040820152600061275a60608301846125c6565b6001600160401b0383168152604060208201526000611d3160408301846125c6565b6001600160a01b038416815260ff8316602082015260606040820181905260009061275a9083018461255c565b600061010060ff8b16835260ff8a16602084015260ff891660408401528760608401528660808401528560a08401528460c08401528060e08401526128248184018561255c565b9b9a5050505050505050505050565b600081518084526020808501945080840160005b8381101561286c5781516001600160401b031687529582019590820190600101612847565b509495945050505050565b600081518084526020808501808196508360051b8101915082860160005b858110156128bf5782840389526128ad84835161255c565b98850198935090840190600101612895565b5091979650505050505050565b6001600160a01b038516815260006101606128ea6020840187612643565b806101208401526128fd81840186612833565b90508281036101408401526126ef8185612877565b6001600160401b0384168152606060208201526000612934606083018561255c565b828103604084015261263981856125c6565b60208152600061208b602083018461255c565b60018060a01b03851681526001600160401b03845116602082015260ff6020850151166040820152604084015160608201526060840151608082015260e060a082015260006129ab60e0830185612833565b82810360c08401526126ef818561287756fea26469706673582212208ab6a17828cc3e97722d65dd2913568c4c85213f9378e3defbbf0a2905d2f3fb64736f6c63430008100033"

// DeployProxy deploys a new Ethereum contract, binding an instance of Proxy to it.
func DeployProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _inst common.Address) (common.Address, *types.Transaction, *Proxy, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProxyBin), backend, _inst)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Proxy *ProxyCaller) Inst(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "inst")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Proxy *ProxySession) Inst() (common.Address, error) {
	return _Proxy.Contract.Inst(&_Proxy.CallOpts)
}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Proxy *ProxyCallerSession) Inst() (common.Address, error) {
	return _Proxy.Contract.Inst(&_Proxy.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Proxy *ProxyCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Proxy *ProxySession) Version() (uint16, error) {
	return _Proxy.Contract.Version(&_Proxy.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Proxy *ProxyCallerSession) Version() (uint16, error) {
	return _Proxy.Contract.Version(&_Proxy.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) Activate(opts *bind.TransactOpts, _i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "activate", _i, signs)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_Proxy *ProxySession) Activate(_i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.Activate(&_Proxy.TransactOpts, _i, signs)
}

// Activate is a paid mutator transaction binding the contract method 0x843d1b49.
//
// Solidity: function activate(uint64 _i, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) Activate(_i uint64, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.Activate(&_Proxy.TransactOpts, _i, signs)
}

// AddOrder is a paid mutator transaction binding the contract method 0xc05bdcd0.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_Proxy *ProxyTransactor) AddOrder(opts *bind.TransactOpts, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "addOrder", _oi, uSign, pSign)
}

// AddOrder is a paid mutator transaction binding the contract method 0xc05bdcd0.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_Proxy *ProxySession) AddOrder(_oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddOrder(&_Proxy.TransactOpts, _oi, uSign, pSign)
}

// AddOrder is a paid mutator transaction binding the contract method 0xc05bdcd0.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_Proxy *ProxyTransactorSession) AddOrder(_oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddOrder(&_Proxy.TransactOpts, _oi, uSign, pSign)
}

// AddRepair is a paid mutator transaction binding the contract method 0x97f2352c.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_Proxy *ProxyTransactor) AddRepair(opts *bind.TransactOpts, _oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "addRepair", _oi, _kis, ksigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0x97f2352c.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_Proxy *ProxySession) AddRepair(_oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddRepair(&_Proxy.TransactOpts, _oi, _kis, ksigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0x97f2352c.
//
// Solidity: function addRepair((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, uint64[] _kis, bytes[] ksigns) returns()
func (_Proxy *ProxyTransactorSession) AddRepair(_oi OrderIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddRepair(&_Proxy.TransactOpts, _oi, _kis, ksigns)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) AddT(opts *bind.TransactOpts, _t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "addT", _t, _mint, signs)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_Proxy *ProxySession) AddT(_t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddT(&_Proxy.TransactOpts, _t, _mint, signs)
}

// AddT is a paid mutator transaction binding the contract method 0x68612348.
//
// Solidity: function addT(address _t, bool _mint, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) AddT(_t common.Address, _mint bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddT(&_Proxy.TransactOpts, _t, _mint, signs)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xa246de42.
//
// Solidity: function addToGroup(uint64 _gi) returns()
func (_Proxy *ProxyTransactor) AddToGroup(opts *bind.TransactOpts, _gi uint64) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "addToGroup", _gi)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xa246de42.
//
// Solidity: function addToGroup(uint64 _gi) returns()
func (_Proxy *ProxySession) AddToGroup(_gi uint64) (*types.Transaction, error) {
	return _Proxy.Contract.AddToGroup(&_Proxy.TransactOpts, _gi)
}

// AddToGroup is a paid mutator transaction binding the contract method 0xa246de42.
//
// Solidity: function addToGroup(uint64 _gi) returns()
func (_Proxy *ProxyTransactorSession) AddToGroup(_gi uint64) (*types.Transaction, error) {
	return _Proxy.Contract.AddToGroup(&_Proxy.TransactOpts, _gi)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _i, address _p) returns()
func (_Proxy *ProxyTransactor) AlterPayee(opts *bind.TransactOpts, _i uint64, _p common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "alterPayee", _i, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _i, address _p) returns()
func (_Proxy *ProxySession) AlterPayee(_i uint64, _p common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.AlterPayee(&_Proxy.TransactOpts, _i, _p)
}

// AlterPayee is a paid mutator transaction binding the contract method 0x2915824a.
//
// Solidity: function alterPayee(uint64 _i, address _p) returns()
func (_Proxy *ProxyTransactorSession) AlterPayee(_i uint64, _p common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.AlterPayee(&_Proxy.TransactOpts, _i, _p)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _isBan, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) BanT(opts *bind.TransactOpts, _t common.Address, _isBan bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "banT", _t, _isBan, signs)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _isBan, bytes[5] signs) returns()
func (_Proxy *ProxySession) BanT(_t common.Address, _isBan bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.BanT(&_Proxy.TransactOpts, _t, _isBan, signs)
}

// BanT is a paid mutator transaction binding the contract method 0x7a743984.
//
// Solidity: function banT(address _t, bool _isBan, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) BanT(_t common.Address, _isBan bool, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.BanT(&_Proxy.TransactOpts, _t, _isBan, signs)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x310aa0be.
//
// Solidity: function confirmPayee(uint64 _i) returns()
func (_Proxy *ProxyTransactor) ConfirmPayee(opts *bind.TransactOpts, _i uint64) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "confirmPayee", _i)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x310aa0be.
//
// Solidity: function confirmPayee(uint64 _i) returns()
func (_Proxy *ProxySession) ConfirmPayee(_i uint64) (*types.Transaction, error) {
	return _Proxy.Contract.ConfirmPayee(&_Proxy.TransactOpts, _i)
}

// ConfirmPayee is a paid mutator transaction binding the contract method 0x310aa0be.
//
// Solidity: function confirmPayee(uint64 _i) returns()
func (_Proxy *ProxyTransactorSession) ConfirmPayee(_i uint64) (*types.Transaction, error) {
	return _Proxy.Contract.ConfirmPayee(&_Proxy.TransactOpts, _i)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x8aae670a.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kpPerB, uint256 _ppPerB, bytes _desc) returns()
func (_Proxy *ProxyTransactor) CreateGroup(opts *bind.TransactOpts, _level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kpPerB *big.Int, _ppPerB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "createGroup", _level, _mr, _tr, _kr, _pr, _kpPerB, _ppPerB, _desc)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x8aae670a.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kpPerB, uint256 _ppPerB, bytes _desc) returns()
func (_Proxy *ProxySession) CreateGroup(_level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kpPerB *big.Int, _ppPerB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _Proxy.Contract.CreateGroup(&_Proxy.TransactOpts, _level, _mr, _tr, _kr, _pr, _kpPerB, _ppPerB, _desc)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x8aae670a.
//
// Solidity: function createGroup(uint8 _level, uint8 _mr, uint8 _tr, uint256 _kr, uint256 _pr, uint256 _kpPerB, uint256 _ppPerB, bytes _desc) returns()
func (_Proxy *ProxyTransactorSession) CreateGroup(_level uint8, _mr uint8, _tr uint8, _kr *big.Int, _pr *big.Int, _kpPerB *big.Int, _ppPerB *big.Int, _desc []byte) (*types.Transaction, error) {
	return _Proxy.Contract.CreateGroup(&_Proxy.TransactOpts, _level, _mr, _tr, _kr, _pr, _kpPerB, _ppPerB, _desc)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 _money) returns()
func (_Proxy *ProxyTransactor) Pledge(opts *bind.TransactOpts, _i uint64, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "pledge", _i, _money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 _money) returns()
func (_Proxy *ProxySession) Pledge(_i uint64, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Pledge(&_Proxy.TransactOpts, _i, _money)
}

// Pledge is a paid mutator transaction binding the contract method 0xd23f7482.
//
// Solidity: function pledge(uint64 _i, uint256 _money) returns()
func (_Proxy *ProxyTransactorSession) Pledge(_i uint64, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Pledge(&_Proxy.TransactOpts, _i, _money)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xe182d54c.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_Proxy *ProxyTransactor) ProWithdraw(opts *bind.TransactOpts, _ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "proWithdraw", _ps, _kis, ksigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xe182d54c.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_Proxy *ProxySession) ProWithdraw(_ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.ProWithdraw(&_Proxy.TransactOpts, _ps, _kis, ksigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xe182d54c.
//
// Solidity: function proWithdraw((uint64,uint8,uint256,uint256) _ps, uint64[] _kis, bytes[] ksigns) returns()
func (_Proxy *ProxyTransactorSession) ProWithdraw(_ps PWIn, _kis []uint64, ksigns [][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.ProWithdraw(&_Proxy.TransactOpts, _ps, _kis, ksigns)
}

// QuitRole is a paid mutator transaction binding the contract method 0xadad6d3a.
//
// Solidity: function quitRole(uint64 _i) returns()
func (_Proxy *ProxyTransactor) QuitRole(opts *bind.TransactOpts, _i uint64) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "quitRole", _i)
}

// QuitRole is a paid mutator transaction binding the contract method 0xadad6d3a.
//
// Solidity: function quitRole(uint64 _i) returns()
func (_Proxy *ProxySession) QuitRole(_i uint64) (*types.Transaction, error) {
	return _Proxy.Contract.QuitRole(&_Proxy.TransactOpts, _i)
}

// QuitRole is a paid mutator transaction binding the contract method 0xadad6d3a.
//
// Solidity: function quitRole(uint64 _i) returns()
func (_Proxy *ProxyTransactorSession) QuitRole(_i uint64) (*types.Transaction, error) {
	return _Proxy.Contract.QuitRole(&_Proxy.TransactOpts, _i)
}

// ReAcc is a paid mutator transaction binding the contract method 0x6526250f.
//
// Solidity: function reAcc() returns()
func (_Proxy *ProxyTransactor) ReAcc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "reAcc")
}

// ReAcc is a paid mutator transaction binding the contract method 0x6526250f.
//
// Solidity: function reAcc() returns()
func (_Proxy *ProxySession) ReAcc() (*types.Transaction, error) {
	return _Proxy.Contract.ReAcc(&_Proxy.TransactOpts)
}

// ReAcc is a paid mutator transaction binding the contract method 0x6526250f.
//
// Solidity: function reAcc() returns()
func (_Proxy *ProxyTransactorSession) ReAcc() (*types.Transaction, error) {
	return _Proxy.Contract.ReAcc(&_Proxy.TransactOpts)
}

// ReRole is a paid mutator transaction binding the contract method 0x898e00f8.
//
// Solidity: function reRole(uint8 _rtype, bytes _vk) returns()
func (_Proxy *ProxyTransactor) ReRole(opts *bind.TransactOpts, _rtype uint8, _vk []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "reRole", _rtype, _vk)
}

// ReRole is a paid mutator transaction binding the contract method 0x898e00f8.
//
// Solidity: function reRole(uint8 _rtype, bytes _vk) returns()
func (_Proxy *ProxySession) ReRole(_rtype uint8, _vk []byte) (*types.Transaction, error) {
	return _Proxy.Contract.ReRole(&_Proxy.TransactOpts, _rtype, _vk)
}

// ReRole is a paid mutator transaction binding the contract method 0x898e00f8.
//
// Solidity: function reRole(uint8 _rtype, bytes _vk) returns()
func (_Proxy *ProxyTransactorSession) ReRole(_rtype uint8, _vk []byte) (*types.Transaction, error) {
	return _Proxy.Contract.ReRole(&_Proxy.TransactOpts, _rtype, _vk)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_Proxy *ProxyTransactor) Recharge(opts *bind.TransactOpts, _i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "recharge", _i, _ti, isLock, _money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_Proxy *ProxySession) Recharge(_i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Recharge(&_Proxy.TransactOpts, _i, _ti, isLock, _money)
}

// Recharge is a paid mutator transaction binding the contract method 0x24d11d40.
//
// Solidity: function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 _money) returns()
func (_Proxy *ProxyTransactorSession) Recharge(_i uint64, _ti uint8, isLock bool, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Recharge(&_Proxy.TransactOpts, _i, _ti, isLock, _money)
}

// SetDesc is a paid mutator transaction binding the contract method 0x0fc6fdd7.
//
// Solidity: function setDesc(bytes _desc) returns()
func (_Proxy *ProxyTransactor) SetDesc(opts *bind.TransactOpts, _desc []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setDesc", _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x0fc6fdd7.
//
// Solidity: function setDesc(bytes _desc) returns()
func (_Proxy *ProxySession) SetDesc(_desc []byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetDesc(&_Proxy.TransactOpts, _desc)
}

// SetDesc is a paid mutator transaction binding the contract method 0x0fc6fdd7.
//
// Solidity: function setDesc(bytes _desc) returns()
func (_Proxy *ProxyTransactorSession) SetDesc(_desc []byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetDesc(&_Proxy.TransactOpts, _desc)
}

// SetGDesc is a paid mutator transaction binding the contract method 0xaa756343.
//
// Solidity: function setGDesc(uint64 _gi, bytes _desc, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) SetGDesc(opts *bind.TransactOpts, _gi uint64, _desc []byte, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setGDesc", _gi, _desc, signs)
}

// SetGDesc is a paid mutator transaction binding the contract method 0xaa756343.
//
// Solidity: function setGDesc(uint64 _gi, bytes _desc, bytes[5] signs) returns()
func (_Proxy *ProxySession) SetGDesc(_gi uint64, _desc []byte, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetGDesc(&_Proxy.TransactOpts, _gi, _desc, signs)
}

// SetGDesc is a paid mutator transaction binding the contract method 0xaa756343.
//
// Solidity: function setGDesc(uint64 _gi, bytes _desc, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) SetGDesc(_gi uint64, _desc []byte, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetGDesc(&_Proxy.TransactOpts, _gi, _desc, signs)
}

// SetGS is a paid mutator transaction binding the contract method 0xf62ea8d1.
//
// Solidity: function setGS(uint64 _gi, uint8 _s, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) SetGS(opts *bind.TransactOpts, _gi uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setGS", _gi, _s, signs)
}

// SetGS is a paid mutator transaction binding the contract method 0xf62ea8d1.
//
// Solidity: function setGS(uint64 _gi, uint8 _s, bytes[5] signs) returns()
func (_Proxy *ProxySession) SetGS(_gi uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetGS(&_Proxy.TransactOpts, _gi, _s, signs)
}

// SetGS is a paid mutator transaction binding the contract method 0xf62ea8d1.
//
// Solidity: function setGS(uint64 _gi, uint8 _s, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) SetGS(_gi uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetGS(&_Proxy.TransactOpts, _gi, _s, signs)
}

// SetLevel is a paid mutator transaction binding the contract method 0x713d3b95.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) SetLevel(opts *bind.TransactOpts, _gi uint64, _le uint8, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setLevel", _gi, _le, signs)
}

// SetLevel is a paid mutator transaction binding the contract method 0x713d3b95.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le, bytes[5] signs) returns()
func (_Proxy *ProxySession) SetLevel(_gi uint64, _le uint8, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetLevel(&_Proxy.TransactOpts, _gi, _le, signs)
}

// SetLevel is a paid mutator transaction binding the contract method 0x713d3b95.
//
// Solidity: function setLevel(uint64 _gi, uint8 _le, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) SetLevel(_gi uint64, _le uint8, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetLevel(&_Proxy.TransactOpts, _gi, _le, signs)
}

// SetP is a paid mutator transaction binding the contract method 0x5192e18a.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) SetP(opts *bind.TransactOpts, _gi uint64, _kp *big.Int, _pp *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setP", _gi, _kp, _pp, signs)
}

// SetP is a paid mutator transaction binding the contract method 0x5192e18a.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp, bytes[5] signs) returns()
func (_Proxy *ProxySession) SetP(_gi uint64, _kp *big.Int, _pp *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetP(&_Proxy.TransactOpts, _gi, _kp, _pp, signs)
}

// SetP is a paid mutator transaction binding the contract method 0x5192e18a.
//
// Solidity: function setP(uint64 _gi, uint256 _kp, uint256 _pp, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) SetP(_gi uint64, _kp *big.Int, _pp *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetP(&_Proxy.TransactOpts, _gi, _kp, _pp, signs)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xbd5d54b7.
//
// Solidity: function setPeriod(uint64 _pe, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) SetPeriod(opts *bind.TransactOpts, _pe uint64, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setPeriod", _pe, signs)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xbd5d54b7.
//
// Solidity: function setPeriod(uint64 _pe, bytes[5] signs) returns()
func (_Proxy *ProxySession) SetPeriod(_pe uint64, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetPeriod(&_Proxy.TransactOpts, _pe, signs)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xbd5d54b7.
//
// Solidity: function setPeriod(uint64 _pe, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) SetPeriod(_pe uint64, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetPeriod(&_Proxy.TransactOpts, _pe, signs)
}

// SetPlePerB is a paid mutator transaction binding the contract method 0x130b89b1.
//
// Solidity: function setPlePerB(uint64 _gi, uint256 _kPB, uint256 _pPB, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) SetPlePerB(opts *bind.TransactOpts, _gi uint64, _kPB *big.Int, _pPB *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setPlePerB", _gi, _kPB, _pPB, signs)
}

// SetPlePerB is a paid mutator transaction binding the contract method 0x130b89b1.
//
// Solidity: function setPlePerB(uint64 _gi, uint256 _kPB, uint256 _pPB, bytes[5] signs) returns()
func (_Proxy *ProxySession) SetPlePerB(_gi uint64, _kPB *big.Int, _pPB *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetPlePerB(&_Proxy.TransactOpts, _gi, _kPB, _pPB, signs)
}

// SetPlePerB is a paid mutator transaction binding the contract method 0x130b89b1.
//
// Solidity: function setPlePerB(uint64 _gi, uint256 _kPB, uint256 _pPB, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) SetPlePerB(_gi uint64, _kPB *big.Int, _pPB *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetPlePerB(&_Proxy.TransactOpts, _gi, _kPB, _pPB, signs)
}

// SetRS is a paid mutator transaction binding the contract method 0x88602357.
//
// Solidity: function setRS(uint64 _i, uint8 _s, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) SetRS(opts *bind.TransactOpts, _i uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setRS", _i, _s, signs)
}

// SetRS is a paid mutator transaction binding the contract method 0x88602357.
//
// Solidity: function setRS(uint64 _i, uint8 _s, bytes[5] signs) returns()
func (_Proxy *ProxySession) SetRS(_i uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetRS(&_Proxy.TransactOpts, _i, _s, signs)
}

// SetRS is a paid mutator transaction binding the contract method 0x88602357.
//
// Solidity: function setRS(uint64 _i, uint8 _s, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) SetRS(_i uint64, _s uint8, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetRS(&_Proxy.TransactOpts, _i, _s, signs)
}

// SetRate is a paid mutator transaction binding the contract method 0x4d525421.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr, bytes[5] signs) returns()
func (_Proxy *ProxyTransactor) SetRate(opts *bind.TransactOpts, _gi uint64, _mr uint8, _tr uint8, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setRate", _gi, _mr, _tr, signs)
}

// SetRate is a paid mutator transaction binding the contract method 0x4d525421.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr, bytes[5] signs) returns()
func (_Proxy *ProxySession) SetRate(_gi uint64, _mr uint8, _tr uint8, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetRate(&_Proxy.TransactOpts, _gi, _mr, _tr, signs)
}

// SetRate is a paid mutator transaction binding the contract method 0x4d525421.
//
// Solidity: function setRate(uint64 _gi, uint8 _mr, uint8 _tr, bytes[5] signs) returns()
func (_Proxy *ProxyTransactorSession) SetRate(_gi uint64, _mr uint8, _tr uint8, signs [5][]byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetRate(&_Proxy.TransactOpts, _gi, _mr, _tr, signs)
}

// SubOrder is a paid mutator transaction binding the contract method 0x153ede07.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_Proxy *ProxyTransactor) SubOrder(opts *bind.TransactOpts, _oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "subOrder", _oi, uSign, pSign)
}

// SubOrder is a paid mutator transaction binding the contract method 0x153ede07.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_Proxy *ProxySession) SubOrder(_oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _Proxy.Contract.SubOrder(&_Proxy.TransactOpts, _oi, uSign, pSign)
}

// SubOrder is a paid mutator transaction binding the contract method 0x153ede07.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint8,uint256) _oi, bytes uSign, bytes pSign) returns()
func (_Proxy *ProxyTransactorSession) SubOrder(_oi OrderIn, uSign []byte, pSign []byte) (*types.Transaction, error) {
	return _Proxy.Contract.SubOrder(&_Proxy.TransactOpts, _oi, uSign, pSign)
}

// Unpledge is a paid mutator transaction binding the contract method 0x60985756.
//
// Solidity: function unpledge(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_Proxy *ProxyTransactor) Unpledge(opts *bind.TransactOpts, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "unpledge", _i, _ti, _money)
}

// Unpledge is a paid mutator transaction binding the contract method 0x60985756.
//
// Solidity: function unpledge(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_Proxy *ProxySession) Unpledge(_i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Unpledge(&_Proxy.TransactOpts, _i, _ti, _money)
}

// Unpledge is a paid mutator transaction binding the contract method 0x60985756.
//
// Solidity: function unpledge(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_Proxy *ProxyTransactorSession) Unpledge(_i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Unpledge(&_Proxy.TransactOpts, _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_Proxy *ProxyTransactor) Withdraw(opts *bind.TransactOpts, _i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "withdraw", _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_Proxy *ProxySession) Withdraw(_i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Withdraw(&_Proxy.TransactOpts, _i, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _i, uint8 _ti, uint256 _money) returns()
func (_Proxy *ProxyTransactorSession) Withdraw(_i uint64, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Withdraw(&_Proxy.TransactOpts, _i, _ti, _money)
}

// ProxyActivateIterator is returned from FilterActivate and is used to iterate over the raw logs and unpacked data for Activate events raised by the Proxy contract.
type ProxyActivateIterator struct {
	Event *ProxyActivate // Event containing the contract specifics and raw log

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
func (it *ProxyActivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyActivate)
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
		it.Event = new(ProxyActivate)
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
func (it *ProxyActivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyActivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyActivate represents a Activate event raised by the Proxy contract.
type ProxyActivate struct {
	I   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterActivate is a free log retrieval operation binding the contract event 0x452a42fb51ebab9c8ebcbb3d536371397d861478788df33e9d2c9ae57765431c.
//
// Solidity: event Activate(uint64 i)
func (_Proxy *ProxyFilterer) FilterActivate(opts *bind.FilterOpts) (*ProxyActivateIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "Activate")
	if err != nil {
		return nil, err
	}
	return &ProxyActivateIterator{contract: _Proxy.contract, event: "Activate", logs: logs, sub: sub}, nil
}

// WatchActivate is a free log subscription operation binding the contract event 0x452a42fb51ebab9c8ebcbb3d536371397d861478788df33e9d2c9ae57765431c.
//
// Solidity: event Activate(uint64 i)
func (_Proxy *ProxyFilterer) WatchActivate(opts *bind.WatchOpts, sink chan<- *ProxyActivate) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "Activate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyActivate)
				if err := _Proxy.contract.UnpackLog(event, "Activate", log); err != nil {
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

// ParseActivate is a log parse operation binding the contract event 0x452a42fb51ebab9c8ebcbb3d536371397d861478788df33e9d2c9ae57765431c.
//
// Solidity: event Activate(uint64 i)
func (_Proxy *ProxyFilterer) ParseActivate(log types.Log) (*ProxyActivate, error) {
	event := new(ProxyActivate)
	if err := _Proxy.contract.UnpackLog(event, "Activate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyAddTIterator is returned from FilterAddT and is used to iterate over the raw logs and unpacked data for AddT events raised by the Proxy contract.
type ProxyAddTIterator struct {
	Event *ProxyAddT // Event containing the contract specifics and raw log

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
func (it *ProxyAddTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyAddT)
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
		it.Event = new(ProxyAddT)
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
func (it *ProxyAddTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyAddTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyAddT represents a AddT event raised by the Proxy contract.
type ProxyAddT struct {
	T    common.Address
	Mint bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddT is a free log retrieval operation binding the contract event 0x4954a00722fe171d5c0079b423ebc3b3984d83c24170fdb068dfbd979d250412.
//
// Solidity: event AddT(address t, bool mint)
func (_Proxy *ProxyFilterer) FilterAddT(opts *bind.FilterOpts) (*ProxyAddTIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return &ProxyAddTIterator{contract: _Proxy.contract, event: "AddT", logs: logs, sub: sub}, nil
}

// WatchAddT is a free log subscription operation binding the contract event 0x4954a00722fe171d5c0079b423ebc3b3984d83c24170fdb068dfbd979d250412.
//
// Solidity: event AddT(address t, bool mint)
func (_Proxy *ProxyFilterer) WatchAddT(opts *bind.WatchOpts, sink chan<- *ProxyAddT) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyAddT)
				if err := _Proxy.contract.UnpackLog(event, "AddT", log); err != nil {
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

// ParseAddT is a log parse operation binding the contract event 0x4954a00722fe171d5c0079b423ebc3b3984d83c24170fdb068dfbd979d250412.
//
// Solidity: event AddT(address t, bool mint)
func (_Proxy *ProxyFilterer) ParseAddT(log types.Log) (*ProxyAddT, error) {
	event := new(ProxyAddT)
	if err := _Proxy.contract.UnpackLog(event, "AddT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyBanTIterator is returned from FilterBanT and is used to iterate over the raw logs and unpacked data for BanT events raised by the Proxy contract.
type ProxyBanTIterator struct {
	Event *ProxyBanT // Event containing the contract specifics and raw log

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
func (it *ProxyBanTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyBanT)
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
		it.Event = new(ProxyBanT)
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
func (it *ProxyBanTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyBanTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyBanT represents a BanT event raised by the Proxy contract.
type ProxyBanT struct {
	T   common.Address
	Ban bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBanT is a free log retrieval operation binding the contract event 0xffaa091f00488e7d9934b8efc3cbe385e3bb89957862cb891e428474b69b6bc9.
//
// Solidity: event BanT(address t, bool ban)
func (_Proxy *ProxyFilterer) FilterBanT(opts *bind.FilterOpts) (*ProxyBanTIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "BanT")
	if err != nil {
		return nil, err
	}
	return &ProxyBanTIterator{contract: _Proxy.contract, event: "BanT", logs: logs, sub: sub}, nil
}

// WatchBanT is a free log subscription operation binding the contract event 0xffaa091f00488e7d9934b8efc3cbe385e3bb89957862cb891e428474b69b6bc9.
//
// Solidity: event BanT(address t, bool ban)
func (_Proxy *ProxyFilterer) WatchBanT(opts *bind.WatchOpts, sink chan<- *ProxyBanT) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "BanT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyBanT)
				if err := _Proxy.contract.UnpackLog(event, "BanT", log); err != nil {
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

// ParseBanT is a log parse operation binding the contract event 0xffaa091f00488e7d9934b8efc3cbe385e3bb89957862cb891e428474b69b6bc9.
//
// Solidity: event BanT(address t, bool ban)
func (_Proxy *ProxyFilterer) ParseBanT(log types.Log) (*ProxyBanT, error) {
	event := new(ProxyBanT)
	if err := _Proxy.contract.UnpackLog(event, "BanT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxySetGDescIterator is returned from FilterSetGDesc and is used to iterate over the raw logs and unpacked data for SetGDesc events raised by the Proxy contract.
type ProxySetGDescIterator struct {
	Event *ProxySetGDesc // Event containing the contract specifics and raw log

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
func (it *ProxySetGDescIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxySetGDesc)
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
		it.Event = new(ProxySetGDesc)
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
func (it *ProxySetGDescIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxySetGDescIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxySetGDesc represents a SetGDesc event raised by the Proxy contract.
type ProxySetGDesc struct {
	Gi   uint64
	Desc []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetGDesc is a free log retrieval operation binding the contract event 0x402035b2534df57a3c7490784b37ac991fa59096f8e251b49d5364967424b85d.
//
// Solidity: event SetGDesc(uint64 indexed gi, bytes desc)
func (_Proxy *ProxyFilterer) FilterSetGDesc(opts *bind.FilterOpts, gi []uint64) (*ProxySetGDescIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "SetGDesc", giRule)
	if err != nil {
		return nil, err
	}
	return &ProxySetGDescIterator{contract: _Proxy.contract, event: "SetGDesc", logs: logs, sub: sub}, nil
}

// WatchSetGDesc is a free log subscription operation binding the contract event 0x402035b2534df57a3c7490784b37ac991fa59096f8e251b49d5364967424b85d.
//
// Solidity: event SetGDesc(uint64 indexed gi, bytes desc)
func (_Proxy *ProxyFilterer) WatchSetGDesc(opts *bind.WatchOpts, sink chan<- *ProxySetGDesc, gi []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "SetGDesc", giRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxySetGDesc)
				if err := _Proxy.contract.UnpackLog(event, "SetGDesc", log); err != nil {
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

// ParseSetGDesc is a log parse operation binding the contract event 0x402035b2534df57a3c7490784b37ac991fa59096f8e251b49d5364967424b85d.
//
// Solidity: event SetGDesc(uint64 indexed gi, bytes desc)
func (_Proxy *ProxyFilterer) ParseSetGDesc(log types.Log) (*ProxySetGDesc, error) {
	event := new(ProxySetGDesc)
	if err := _Proxy.contract.UnpackLog(event, "SetGDesc", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxySetGSIterator is returned from FilterSetGS and is used to iterate over the raw logs and unpacked data for SetGS events raised by the Proxy contract.
type ProxySetGSIterator struct {
	Event *ProxySetGS // Event containing the contract specifics and raw log

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
func (it *ProxySetGSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxySetGS)
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
		it.Event = new(ProxySetGS)
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
func (it *ProxySetGSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxySetGSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxySetGS represents a SetGS event raised by the Proxy contract.
type ProxySetGS struct {
	Gi  uint64
	S   uint8
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetGS is a free log retrieval operation binding the contract event 0x6868b79de8578929833c8106714d184d883afab70f4161a0d447c4f2d899417a.
//
// Solidity: event SetGS(uint64 indexed gi, uint8 _s)
func (_Proxy *ProxyFilterer) FilterSetGS(opts *bind.FilterOpts, gi []uint64) (*ProxySetGSIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "SetGS", giRule)
	if err != nil {
		return nil, err
	}
	return &ProxySetGSIterator{contract: _Proxy.contract, event: "SetGS", logs: logs, sub: sub}, nil
}

// WatchSetGS is a free log subscription operation binding the contract event 0x6868b79de8578929833c8106714d184d883afab70f4161a0d447c4f2d899417a.
//
// Solidity: event SetGS(uint64 indexed gi, uint8 _s)
func (_Proxy *ProxyFilterer) WatchSetGS(opts *bind.WatchOpts, sink chan<- *ProxySetGS, gi []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "SetGS", giRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxySetGS)
				if err := _Proxy.contract.UnpackLog(event, "SetGS", log); err != nil {
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

// ParseSetGS is a log parse operation binding the contract event 0x6868b79de8578929833c8106714d184d883afab70f4161a0d447c4f2d899417a.
//
// Solidity: event SetGS(uint64 indexed gi, uint8 _s)
func (_Proxy *ProxyFilterer) ParseSetGS(log types.Log) (*ProxySetGS, error) {
	event := new(ProxySetGS)
	if err := _Proxy.contract.UnpackLog(event, "SetGS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxySetLevelIterator is returned from FilterSetLevel and is used to iterate over the raw logs and unpacked data for SetLevel events raised by the Proxy contract.
type ProxySetLevelIterator struct {
	Event *ProxySetLevel // Event containing the contract specifics and raw log

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
func (it *ProxySetLevelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxySetLevel)
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
		it.Event = new(ProxySetLevel)
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
func (it *ProxySetLevelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxySetLevelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxySetLevel represents a SetLevel event raised by the Proxy contract.
type ProxySetLevel struct {
	Gi  uint64
	Le  uint8
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetLevel is a free log retrieval operation binding the contract event 0x5aee2b03a7ccab0db1135bf51f379d3b74e615f96cb0e7d43db42767638cf899.
//
// Solidity: event SetLevel(uint64 indexed gi, uint8 le)
func (_Proxy *ProxyFilterer) FilterSetLevel(opts *bind.FilterOpts, gi []uint64) (*ProxySetLevelIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "SetLevel", giRule)
	if err != nil {
		return nil, err
	}
	return &ProxySetLevelIterator{contract: _Proxy.contract, event: "SetLevel", logs: logs, sub: sub}, nil
}

// WatchSetLevel is a free log subscription operation binding the contract event 0x5aee2b03a7ccab0db1135bf51f379d3b74e615f96cb0e7d43db42767638cf899.
//
// Solidity: event SetLevel(uint64 indexed gi, uint8 le)
func (_Proxy *ProxyFilterer) WatchSetLevel(opts *bind.WatchOpts, sink chan<- *ProxySetLevel, gi []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "SetLevel", giRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxySetLevel)
				if err := _Proxy.contract.UnpackLog(event, "SetLevel", log); err != nil {
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

// ParseSetLevel is a log parse operation binding the contract event 0x5aee2b03a7ccab0db1135bf51f379d3b74e615f96cb0e7d43db42767638cf899.
//
// Solidity: event SetLevel(uint64 indexed gi, uint8 le)
func (_Proxy *ProxyFilterer) ParseSetLevel(log types.Log) (*ProxySetLevel, error) {
	event := new(ProxySetLevel)
	if err := _Proxy.contract.UnpackLog(event, "SetLevel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxySetPeriodIterator is returned from FilterSetPeriod and is used to iterate over the raw logs and unpacked data for SetPeriod events raised by the Proxy contract.
type ProxySetPeriodIterator struct {
	Event *ProxySetPeriod // Event containing the contract specifics and raw log

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
func (it *ProxySetPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxySetPeriod)
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
		it.Event = new(ProxySetPeriod)
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
func (it *ProxySetPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxySetPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxySetPeriod represents a SetPeriod event raised by the Proxy contract.
type ProxySetPeriod struct {
	Pe  uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetPeriod is a free log retrieval operation binding the contract event 0x11718e279b88c35d197671951471b347900bac4b36c87d6f5a34ede95b3823d0.
//
// Solidity: event SetPeriod(uint64 pe)
func (_Proxy *ProxyFilterer) FilterSetPeriod(opts *bind.FilterOpts) (*ProxySetPeriodIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "SetPeriod")
	if err != nil {
		return nil, err
	}
	return &ProxySetPeriodIterator{contract: _Proxy.contract, event: "SetPeriod", logs: logs, sub: sub}, nil
}

// WatchSetPeriod is a free log subscription operation binding the contract event 0x11718e279b88c35d197671951471b347900bac4b36c87d6f5a34ede95b3823d0.
//
// Solidity: event SetPeriod(uint64 pe)
func (_Proxy *ProxyFilterer) WatchSetPeriod(opts *bind.WatchOpts, sink chan<- *ProxySetPeriod) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "SetPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxySetPeriod)
				if err := _Proxy.contract.UnpackLog(event, "SetPeriod", log); err != nil {
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

// ParseSetPeriod is a log parse operation binding the contract event 0x11718e279b88c35d197671951471b347900bac4b36c87d6f5a34ede95b3823d0.
//
// Solidity: event SetPeriod(uint64 pe)
func (_Proxy *ProxyFilterer) ParseSetPeriod(log types.Log) (*ProxySetPeriod, error) {
	event := new(ProxySetPeriod)
	if err := _Proxy.contract.UnpackLog(event, "SetPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxySetPlePerBIterator is returned from FilterSetPlePerB and is used to iterate over the raw logs and unpacked data for SetPlePerB events raised by the Proxy contract.
type ProxySetPlePerBIterator struct {
	Event *ProxySetPlePerB // Event containing the contract specifics and raw log

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
func (it *ProxySetPlePerBIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxySetPlePerB)
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
		it.Event = new(ProxySetPlePerB)
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
func (it *ProxySetPlePerBIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxySetPlePerBIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxySetPlePerB represents a SetPlePerB event raised by the Proxy contract.
type ProxySetPlePerB struct {
	Gi     uint64
	KpPerB *big.Int
	PpPerB *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetPlePerB is a free log retrieval operation binding the contract event 0x4f8fb948a05297fa2147f00955cede56207c0eb67e57201c6a3fdc53588f27b8.
//
// Solidity: event SetPlePerB(uint64 indexed gi, uint256 kpPerB, uint256 ppPerB)
func (_Proxy *ProxyFilterer) FilterSetPlePerB(opts *bind.FilterOpts, gi []uint64) (*ProxySetPlePerBIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "SetPlePerB", giRule)
	if err != nil {
		return nil, err
	}
	return &ProxySetPlePerBIterator{contract: _Proxy.contract, event: "SetPlePerB", logs: logs, sub: sub}, nil
}

// WatchSetPlePerB is a free log subscription operation binding the contract event 0x4f8fb948a05297fa2147f00955cede56207c0eb67e57201c6a3fdc53588f27b8.
//
// Solidity: event SetPlePerB(uint64 indexed gi, uint256 kpPerB, uint256 ppPerB)
func (_Proxy *ProxyFilterer) WatchSetPlePerB(opts *bind.WatchOpts, sink chan<- *ProxySetPlePerB, gi []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "SetPlePerB", giRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxySetPlePerB)
				if err := _Proxy.contract.UnpackLog(event, "SetPlePerB", log); err != nil {
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

// ParseSetPlePerB is a log parse operation binding the contract event 0x4f8fb948a05297fa2147f00955cede56207c0eb67e57201c6a3fdc53588f27b8.
//
// Solidity: event SetPlePerB(uint64 indexed gi, uint256 kpPerB, uint256 ppPerB)
func (_Proxy *ProxyFilterer) ParseSetPlePerB(log types.Log) (*ProxySetPlePerB, error) {
	event := new(ProxySetPlePerB)
	if err := _Proxy.contract.UnpackLog(event, "SetPlePerB", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxySetRSIterator is returned from FilterSetRS and is used to iterate over the raw logs and unpacked data for SetRS events raised by the Proxy contract.
type ProxySetRSIterator struct {
	Event *ProxySetRS // Event containing the contract specifics and raw log

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
func (it *ProxySetRSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxySetRS)
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
		it.Event = new(ProxySetRS)
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
func (it *ProxySetRSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxySetRSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxySetRS represents a SetRS event raised by the Proxy contract.
type ProxySetRS struct {
	I   uint64
	S   uint8
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetRS is a free log retrieval operation binding the contract event 0xf101cf457bf0c852aed48bf40478c9eb079d6b5a23f9dc8c00c467712b3d293e.
//
// Solidity: event SetRS(uint64 i, uint8 _s)
func (_Proxy *ProxyFilterer) FilterSetRS(opts *bind.FilterOpts) (*ProxySetRSIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "SetRS")
	if err != nil {
		return nil, err
	}
	return &ProxySetRSIterator{contract: _Proxy.contract, event: "SetRS", logs: logs, sub: sub}, nil
}

// WatchSetRS is a free log subscription operation binding the contract event 0xf101cf457bf0c852aed48bf40478c9eb079d6b5a23f9dc8c00c467712b3d293e.
//
// Solidity: event SetRS(uint64 i, uint8 _s)
func (_Proxy *ProxyFilterer) WatchSetRS(opts *bind.WatchOpts, sink chan<- *ProxySetRS) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "SetRS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxySetRS)
				if err := _Proxy.contract.UnpackLog(event, "SetRS", log); err != nil {
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

// ParseSetRS is a log parse operation binding the contract event 0xf101cf457bf0c852aed48bf40478c9eb079d6b5a23f9dc8c00c467712b3d293e.
//
// Solidity: event SetRS(uint64 i, uint8 _s)
func (_Proxy *ProxyFilterer) ParseSetRS(log types.Log) (*ProxySetRS, error) {
	event := new(ProxySetRS)
	if err := _Proxy.contract.UnpackLog(event, "SetRS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxySetRateIterator is returned from FilterSetRate and is used to iterate over the raw logs and unpacked data for SetRate events raised by the Proxy contract.
type ProxySetRateIterator struct {
	Event *ProxySetRate // Event containing the contract specifics and raw log

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
func (it *ProxySetRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxySetRate)
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
		it.Event = new(ProxySetRate)
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
func (it *ProxySetRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxySetRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxySetRate represents a SetRate event raised by the Proxy contract.
type ProxySetRate struct {
	Gi  uint64
	Mr  uint8
	Kr  uint8
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetRate is a free log retrieval operation binding the contract event 0x1ee4799609f0de30722e40275c40d1dd80883a00e6d3ceda4f5d4328968c9a4a.
//
// Solidity: event SetRate(uint64 indexed gi, uint8 mr, uint8 kr)
func (_Proxy *ProxyFilterer) FilterSetRate(opts *bind.FilterOpts, gi []uint64) (*ProxySetRateIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "SetRate", giRule)
	if err != nil {
		return nil, err
	}
	return &ProxySetRateIterator{contract: _Proxy.contract, event: "SetRate", logs: logs, sub: sub}, nil
}

// WatchSetRate is a free log subscription operation binding the contract event 0x1ee4799609f0de30722e40275c40d1dd80883a00e6d3ceda4f5d4328968c9a4a.
//
// Solidity: event SetRate(uint64 indexed gi, uint8 mr, uint8 kr)
func (_Proxy *ProxyFilterer) WatchSetRate(opts *bind.WatchOpts, sink chan<- *ProxySetRate, gi []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "SetRate", giRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxySetRate)
				if err := _Proxy.contract.UnpackLog(event, "SetRate", log); err != nil {
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

// ParseSetRate is a log parse operation binding the contract event 0x1ee4799609f0de30722e40275c40d1dd80883a00e6d3ceda4f5d4328968c9a4a.
//
// Solidity: event SetRate(uint64 indexed gi, uint8 mr, uint8 kr)
func (_Proxy *ProxyFilterer) ParseSetRate(log types.Log) (*ProxySetRate, error) {
	event := new(ProxySetRate)
	if err := _Proxy.contract.UnpackLog(event, "SetRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxySetkpPIterator is returned from FilterSetkpP and is used to iterate over the raw logs and unpacked data for SetkpP events raised by the Proxy contract.
type ProxySetkpPIterator struct {
	Event *ProxySetkpP // Event containing the contract specifics and raw log

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
func (it *ProxySetkpPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxySetkpP)
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
		it.Event = new(ProxySetkpP)
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
func (it *ProxySetkpPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxySetkpPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxySetkpP represents a SetkpP event raised by the Proxy contract.
type ProxySetkpP struct {
	Gi  uint64
	Kp  *big.Int
	Pp  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetkpP is a free log retrieval operation binding the contract event 0xaa3e90b31b32985da56d5c241157e7caec89ed412da1ba26278f59b165a21585.
//
// Solidity: event SetkpP(uint64 indexed gi, uint256 kp, uint256 pp)
func (_Proxy *ProxyFilterer) FilterSetkpP(opts *bind.FilterOpts, gi []uint64) (*ProxySetkpPIterator, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "SetkpP", giRule)
	if err != nil {
		return nil, err
	}
	return &ProxySetkpPIterator{contract: _Proxy.contract, event: "SetkpP", logs: logs, sub: sub}, nil
}

// WatchSetkpP is a free log subscription operation binding the contract event 0xaa3e90b31b32985da56d5c241157e7caec89ed412da1ba26278f59b165a21585.
//
// Solidity: event SetkpP(uint64 indexed gi, uint256 kp, uint256 pp)
func (_Proxy *ProxyFilterer) WatchSetkpP(opts *bind.WatchOpts, sink chan<- *ProxySetkpP, gi []uint64) (event.Subscription, error) {

	var giRule []interface{}
	for _, giItem := range gi {
		giRule = append(giRule, giItem)
	}

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "SetkpP", giRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxySetkpP)
				if err := _Proxy.contract.UnpackLog(event, "SetkpP", log); err != nil {
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

// ParseSetkpP is a log parse operation binding the contract event 0xaa3e90b31b32985da56d5c241157e7caec89ed412da1ba26278f59b165a21585.
//
// Solidity: event SetkpP(uint64 indexed gi, uint256 kp, uint256 pp)
func (_Proxy *ProxyFilterer) ParseSetkpP(log types.Log) (*ProxySetkpP, error) {
	event := new(ProxySetkpP)
	if err := _Proxy.contract.UnpackLog(event, "SetkpP", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
