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
	_ = abi.ConvertType
)

// IIdentityManagerRootInfo is an auto generated low-level Go binding around an user-defined struct.
type IIdentityManagerRootInfo struct {
	ReplacedBy *big.Int
	ReplacedAt *big.Int
	IsLatest   bool
	IsExpired  bool
}

// IdentityManagerMetaData contains all meta data concerning the IdentityManager contract.
var IdentityManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"postRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"replacedAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"latestRoot\",\"type\":\"uint256\"}],\"name\":\"SignedRootTransited\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"P\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_EXPIRATION_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceStateContract_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"}],\"name\":\"__IdentityManager_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"}],\"name\":\"__Signers_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"checkSignatureAndIncrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"root_\",\"type\":\"uint256\"}],\"name\":\"getRootInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"replacedBy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLatest\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isExpired\",\"type\":\"bool\"}],\"internalType\":\"structIIdentityManager.RootInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"getSigComponents\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"root_\",\"type\":\"uint256\"}],\"name\":\"isExpiredRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"isLatestRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"root_\",\"type\":\"uint256\"}],\"name\":\"rootExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prevRoot_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"postRoot_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAt_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"}],\"name\":\"signedTransitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceStateContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"validateChangeAddressSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IdentityManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityManagerMetaData.ABI instead.
var IdentityManagerABI = IdentityManagerMetaData.ABI

// IdentityManager is an auto generated Go binding around an Ethereum contract.
type IdentityManager struct {
	IdentityManagerCaller     // Read-only binding to the contract
	IdentityManagerTransactor // Write-only binding to the contract
	IdentityManagerFilterer   // Log filterer for contract events
}

// IdentityManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityManagerSession struct {
	Contract     *IdentityManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityManagerCallerSession struct {
	Contract *IdentityManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IdentityManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityManagerTransactorSession struct {
	Contract     *IdentityManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IdentityManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityManagerRaw struct {
	Contract *IdentityManager // Generic contract binding to access the raw methods on
}

// IdentityManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityManagerCallerRaw struct {
	Contract *IdentityManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityManagerTransactorRaw struct {
	Contract *IdentityManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityManager creates a new instance of IdentityManager, bound to a specific deployed contract.
func NewIdentityManager(address common.Address, backend bind.ContractBackend) (*IdentityManager, error) {
	contract, err := bindIdentityManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityManager{IdentityManagerCaller: IdentityManagerCaller{contract: contract}, IdentityManagerTransactor: IdentityManagerTransactor{contract: contract}, IdentityManagerFilterer: IdentityManagerFilterer{contract: contract}}, nil
}

// NewIdentityManagerCaller creates a new read-only instance of IdentityManager, bound to a specific deployed contract.
func NewIdentityManagerCaller(address common.Address, caller bind.ContractCaller) (*IdentityManagerCaller, error) {
	contract, err := bindIdentityManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityManagerCaller{contract: contract}, nil
}

// NewIdentityManagerTransactor creates a new write-only instance of IdentityManager, bound to a specific deployed contract.
func NewIdentityManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityManagerTransactor, error) {
	contract, err := bindIdentityManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityManagerTransactor{contract: contract}, nil
}

// NewIdentityManagerFilterer creates a new log filterer instance of IdentityManager, bound to a specific deployed contract.
func NewIdentityManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityManagerFilterer, error) {
	contract, err := bindIdentityManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityManagerFilterer{contract: contract}, nil
}

// bindIdentityManager binds a generic wrapper to an already deployed contract.
func bindIdentityManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdentityManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityManager *IdentityManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityManager.Contract.IdentityManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityManager *IdentityManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityManager.Contract.IdentityManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityManager *IdentityManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityManager.Contract.IdentityManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityManager *IdentityManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityManager *IdentityManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityManager *IdentityManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityManager.Contract.contract.Transact(opts, method, params...)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_IdentityManager *IdentityManagerCaller) P(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "P")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_IdentityManager *IdentityManagerSession) P() (*big.Int, error) {
	return _IdentityManager.Contract.P(&_IdentityManager.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_IdentityManager *IdentityManagerCallerSession) P() (*big.Int, error) {
	return _IdentityManager.Contract.P(&_IdentityManager.CallOpts)
}

// ROOTEXPIRATIONTIME is a free data retrieval call binding the contract method 0x783cb137.
//
// Solidity: function ROOT_EXPIRATION_TIME() view returns(uint256)
func (_IdentityManager *IdentityManagerCaller) ROOTEXPIRATIONTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "ROOT_EXPIRATION_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ROOTEXPIRATIONTIME is a free data retrieval call binding the contract method 0x783cb137.
//
// Solidity: function ROOT_EXPIRATION_TIME() view returns(uint256)
func (_IdentityManager *IdentityManagerSession) ROOTEXPIRATIONTIME() (*big.Int, error) {
	return _IdentityManager.Contract.ROOTEXPIRATIONTIME(&_IdentityManager.CallOpts)
}

// ROOTEXPIRATIONTIME is a free data retrieval call binding the contract method 0x783cb137.
//
// Solidity: function ROOT_EXPIRATION_TIME() view returns(uint256)
func (_IdentityManager *IdentityManagerCallerSession) ROOTEXPIRATIONTIME() (*big.Int, error) {
	return _IdentityManager.Contract.ROOTEXPIRATIONTIME(&_IdentityManager.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_IdentityManager *IdentityManagerCaller) ChainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "chainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_IdentityManager *IdentityManagerSession) ChainName() (string, error) {
	return _IdentityManager.Contract.ChainName(&_IdentityManager.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_IdentityManager *IdentityManagerCallerSession) ChainName() (string, error) {
	return _IdentityManager.Contract.ChainName(&_IdentityManager.CallOpts)
}

// GetLatestRoot is a free data retrieval call binding the contract method 0x5445b007.
//
// Solidity: function getLatestRoot() view returns(uint256, uint256)
func (_IdentityManager *IdentityManagerCaller) GetLatestRoot(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "getLatestRoot")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetLatestRoot is a free data retrieval call binding the contract method 0x5445b007.
//
// Solidity: function getLatestRoot() view returns(uint256, uint256)
func (_IdentityManager *IdentityManagerSession) GetLatestRoot() (*big.Int, *big.Int, error) {
	return _IdentityManager.Contract.GetLatestRoot(&_IdentityManager.CallOpts)
}

// GetLatestRoot is a free data retrieval call binding the contract method 0x5445b007.
//
// Solidity: function getLatestRoot() view returns(uint256, uint256)
func (_IdentityManager *IdentityManagerCallerSession) GetLatestRoot() (*big.Int, *big.Int, error) {
	return _IdentityManager.Contract.GetLatestRoot(&_IdentityManager.CallOpts)
}

// GetRootInfo is a free data retrieval call binding the contract method 0xa69db725.
//
// Solidity: function getRootInfo(uint256 root_) view returns((uint256,uint256,bool,bool))
func (_IdentityManager *IdentityManagerCaller) GetRootInfo(opts *bind.CallOpts, root_ *big.Int) (IIdentityManagerRootInfo, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "getRootInfo", root_)

	if err != nil {
		return *new(IIdentityManagerRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IIdentityManagerRootInfo)).(*IIdentityManagerRootInfo)

	return out0, err

}

// GetRootInfo is a free data retrieval call binding the contract method 0xa69db725.
//
// Solidity: function getRootInfo(uint256 root_) view returns((uint256,uint256,bool,bool))
func (_IdentityManager *IdentityManagerSession) GetRootInfo(root_ *big.Int) (IIdentityManagerRootInfo, error) {
	return _IdentityManager.Contract.GetRootInfo(&_IdentityManager.CallOpts, root_)
}

// GetRootInfo is a free data retrieval call binding the contract method 0xa69db725.
//
// Solidity: function getRootInfo(uint256 root_) view returns((uint256,uint256,bool,bool))
func (_IdentityManager *IdentityManagerCallerSession) GetRootInfo(root_ *big.Int) (IIdentityManagerRootInfo, error) {
	return _IdentityManager.Contract.GetRootInfo(&_IdentityManager.CallOpts, root_)
}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_IdentityManager *IdentityManagerCaller) GetSigComponents(opts *bind.CallOpts, methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "getSigComponents", methodId_, contractAddress_)

	outstruct := new(struct {
		ChainName string
		Nonce     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Nonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_IdentityManager *IdentityManagerSession) GetSigComponents(methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	return _IdentityManager.Contract.GetSigComponents(&_IdentityManager.CallOpts, methodId_, contractAddress_)
}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_IdentityManager *IdentityManagerCallerSession) GetSigComponents(methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	return _IdentityManager.Contract.GetSigComponents(&_IdentityManager.CallOpts, methodId_, contractAddress_)
}

// IsExpiredRoot is a free data retrieval call binding the contract method 0xc62fd5a0.
//
// Solidity: function isExpiredRoot(uint256 root_) view returns(bool)
func (_IdentityManager *IdentityManagerCaller) IsExpiredRoot(opts *bind.CallOpts, root_ *big.Int) (bool, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "isExpiredRoot", root_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExpiredRoot is a free data retrieval call binding the contract method 0xc62fd5a0.
//
// Solidity: function isExpiredRoot(uint256 root_) view returns(bool)
func (_IdentityManager *IdentityManagerSession) IsExpiredRoot(root_ *big.Int) (bool, error) {
	return _IdentityManager.Contract.IsExpiredRoot(&_IdentityManager.CallOpts, root_)
}

// IsExpiredRoot is a free data retrieval call binding the contract method 0xc62fd5a0.
//
// Solidity: function isExpiredRoot(uint256 root_) view returns(bool)
func (_IdentityManager *IdentityManagerCallerSession) IsExpiredRoot(root_ *big.Int) (bool, error) {
	return _IdentityManager.Contract.IsExpiredRoot(&_IdentityManager.CallOpts, root_)
}

// IsLatestRoot is a free data retrieval call binding the contract method 0x0cebec86.
//
// Solidity: function isLatestRoot(uint256 root) view returns(bool)
func (_IdentityManager *IdentityManagerCaller) IsLatestRoot(opts *bind.CallOpts, root *big.Int) (bool, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "isLatestRoot", root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLatestRoot is a free data retrieval call binding the contract method 0x0cebec86.
//
// Solidity: function isLatestRoot(uint256 root) view returns(bool)
func (_IdentityManager *IdentityManagerSession) IsLatestRoot(root *big.Int) (bool, error) {
	return _IdentityManager.Contract.IsLatestRoot(&_IdentityManager.CallOpts, root)
}

// IsLatestRoot is a free data retrieval call binding the contract method 0x0cebec86.
//
// Solidity: function isLatestRoot(uint256 root) view returns(bool)
func (_IdentityManager *IdentityManagerCallerSession) IsLatestRoot(root *big.Int) (bool, error) {
	return _IdentityManager.Contract.IsLatestRoot(&_IdentityManager.CallOpts, root)
}

// Nonces is a free data retrieval call binding the contract method 0xed3218a2.
//
// Solidity: function nonces(address , uint8 ) view returns(uint256)
func (_IdentityManager *IdentityManagerCaller) Nonces(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) (*big.Int, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "nonces", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0xed3218a2.
//
// Solidity: function nonces(address , uint8 ) view returns(uint256)
func (_IdentityManager *IdentityManagerSession) Nonces(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _IdentityManager.Contract.Nonces(&_IdentityManager.CallOpts, arg0, arg1)
}

// Nonces is a free data retrieval call binding the contract method 0xed3218a2.
//
// Solidity: function nonces(address , uint8 ) view returns(uint256)
func (_IdentityManager *IdentityManagerCallerSession) Nonces(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _IdentityManager.Contract.Nonces(&_IdentityManager.CallOpts, arg0, arg1)
}

// RootExists is a free data retrieval call binding the contract method 0x76a5cef3.
//
// Solidity: function rootExists(uint256 root_) view returns(bool)
func (_IdentityManager *IdentityManagerCaller) RootExists(opts *bind.CallOpts, root_ *big.Int) (bool, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "rootExists", root_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RootExists is a free data retrieval call binding the contract method 0x76a5cef3.
//
// Solidity: function rootExists(uint256 root_) view returns(bool)
func (_IdentityManager *IdentityManagerSession) RootExists(root_ *big.Int) (bool, error) {
	return _IdentityManager.Contract.RootExists(&_IdentityManager.CallOpts, root_)
}

// RootExists is a free data retrieval call binding the contract method 0x76a5cef3.
//
// Solidity: function rootExists(uint256 root_) view returns(bool)
func (_IdentityManager *IdentityManagerCallerSession) RootExists(root_ *big.Int) (bool, error) {
	return _IdentityManager.Contract.RootExists(&_IdentityManager.CallOpts, root_)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_IdentityManager *IdentityManagerCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_IdentityManager *IdentityManagerSession) Signer() (common.Address, error) {
	return _IdentityManager.Contract.Signer(&_IdentityManager.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_IdentityManager *IdentityManagerCallerSession) Signer() (common.Address, error) {
	return _IdentityManager.Contract.Signer(&_IdentityManager.CallOpts)
}

// SourceStateContract is a free data retrieval call binding the contract method 0xfc228319.
//
// Solidity: function sourceStateContract() view returns(address)
func (_IdentityManager *IdentityManagerCaller) SourceStateContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "sourceStateContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SourceStateContract is a free data retrieval call binding the contract method 0xfc228319.
//
// Solidity: function sourceStateContract() view returns(address)
func (_IdentityManager *IdentityManagerSession) SourceStateContract() (common.Address, error) {
	return _IdentityManager.Contract.SourceStateContract(&_IdentityManager.CallOpts)
}

// SourceStateContract is a free data retrieval call binding the contract method 0xfc228319.
//
// Solidity: function sourceStateContract() view returns(address)
func (_IdentityManager *IdentityManagerCallerSession) SourceStateContract() (common.Address, error) {
	return _IdentityManager.Contract.SourceStateContract(&_IdentityManager.CallOpts)
}

// IdentityManagerInit is a paid mutator transaction binding the contract method 0x56d21a54.
//
// Solidity: function __IdentityManager_init(address signer_, address sourceStateContract_, string chainName_) returns()
func (_IdentityManager *IdentityManagerTransactor) IdentityManagerInit(opts *bind.TransactOpts, signer_ common.Address, sourceStateContract_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "__IdentityManager_init", signer_, sourceStateContract_, chainName_)
}

// IdentityManagerInit is a paid mutator transaction binding the contract method 0x56d21a54.
//
// Solidity: function __IdentityManager_init(address signer_, address sourceStateContract_, string chainName_) returns()
func (_IdentityManager *IdentityManagerSession) IdentityManagerInit(signer_ common.Address, sourceStateContract_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _IdentityManager.Contract.IdentityManagerInit(&_IdentityManager.TransactOpts, signer_, sourceStateContract_, chainName_)
}

// IdentityManagerInit is a paid mutator transaction binding the contract method 0x56d21a54.
//
// Solidity: function __IdentityManager_init(address signer_, address sourceStateContract_, string chainName_) returns()
func (_IdentityManager *IdentityManagerTransactorSession) IdentityManagerInit(signer_ common.Address, sourceStateContract_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _IdentityManager.Contract.IdentityManagerInit(&_IdentityManager.TransactOpts, signer_, sourceStateContract_, chainName_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x509ace95.
//
// Solidity: function __Signers_init(address signer_, string chainName_) returns()
func (_IdentityManager *IdentityManagerTransactor) SignersInit(opts *bind.TransactOpts, signer_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "__Signers_init", signer_, chainName_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x509ace95.
//
// Solidity: function __Signers_init(address signer_, string chainName_) returns()
func (_IdentityManager *IdentityManagerSession) SignersInit(signer_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _IdentityManager.Contract.SignersInit(&_IdentityManager.TransactOpts, signer_, chainName_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x509ace95.
//
// Solidity: function __Signers_init(address signer_, string chainName_) returns()
func (_IdentityManager *IdentityManagerTransactorSession) SignersInit(signer_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _IdentityManager.Contract.SignersInit(&_IdentityManager.TransactOpts, signer_, chainName_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_IdentityManager *IdentityManagerTransactor) CheckSignatureAndIncrementNonce(opts *bind.TransactOpts, methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "checkSignatureAndIncrementNonce", methodId_, contractAddress_, signHash_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_IdentityManager *IdentityManagerSession) CheckSignatureAndIncrementNonce(methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _IdentityManager.Contract.CheckSignatureAndIncrementNonce(&_IdentityManager.TransactOpts, methodId_, contractAddress_, signHash_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_IdentityManager *IdentityManagerTransactorSession) CheckSignatureAndIncrementNonce(methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _IdentityManager.Contract.CheckSignatureAndIncrementNonce(&_IdentityManager.TransactOpts, methodId_, contractAddress_, signHash_, signature_)
}

// SignedTransitRoot is a paid mutator transaction binding the contract method 0xd8cb4948.
//
// Solidity: function signedTransitRoot(uint256 prevRoot_, uint256 postRoot_, uint256 replacedAt_, bytes proof_) returns()
func (_IdentityManager *IdentityManagerTransactor) SignedTransitRoot(opts *bind.TransactOpts, prevRoot_ *big.Int, postRoot_ *big.Int, replacedAt_ *big.Int, proof_ []byte) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "signedTransitRoot", prevRoot_, postRoot_, replacedAt_, proof_)
}

// SignedTransitRoot is a paid mutator transaction binding the contract method 0xd8cb4948.
//
// Solidity: function signedTransitRoot(uint256 prevRoot_, uint256 postRoot_, uint256 replacedAt_, bytes proof_) returns()
func (_IdentityManager *IdentityManagerSession) SignedTransitRoot(prevRoot_ *big.Int, postRoot_ *big.Int, replacedAt_ *big.Int, proof_ []byte) (*types.Transaction, error) {
	return _IdentityManager.Contract.SignedTransitRoot(&_IdentityManager.TransactOpts, prevRoot_, postRoot_, replacedAt_, proof_)
}

// SignedTransitRoot is a paid mutator transaction binding the contract method 0xd8cb4948.
//
// Solidity: function signedTransitRoot(uint256 prevRoot_, uint256 postRoot_, uint256 replacedAt_, bytes proof_) returns()
func (_IdentityManager *IdentityManagerTransactorSession) SignedTransitRoot(prevRoot_ *big.Int, postRoot_ *big.Int, replacedAt_ *big.Int, proof_ []byte) (*types.Transaction, error) {
	return _IdentityManager.Contract.SignedTransitRoot(&_IdentityManager.TransactOpts, prevRoot_, postRoot_, replacedAt_, proof_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_IdentityManager *IdentityManagerTransactor) ValidateChangeAddressSignature(opts *bind.TransactOpts, methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "validateChangeAddressSignature", methodId_, contractAddress_, newAddress_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_IdentityManager *IdentityManagerSession) ValidateChangeAddressSignature(methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _IdentityManager.Contract.ValidateChangeAddressSignature(&_IdentityManager.TransactOpts, methodId_, contractAddress_, newAddress_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_IdentityManager *IdentityManagerTransactorSession) ValidateChangeAddressSignature(methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _IdentityManager.Contract.ValidateChangeAddressSignature(&_IdentityManager.TransactOpts, methodId_, contractAddress_, newAddress_, signature_)
}

// IdentityManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IdentityManager contract.
type IdentityManagerInitializedIterator struct {
	Event *IdentityManagerInitialized // Event containing the contract specifics and raw log

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
func (it *IdentityManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityManagerInitialized)
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
		it.Event = new(IdentityManagerInitialized)
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
func (it *IdentityManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityManagerInitialized represents a Initialized event raised by the IdentityManager contract.
type IdentityManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_IdentityManager *IdentityManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*IdentityManagerInitializedIterator, error) {

	logs, sub, err := _IdentityManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IdentityManagerInitializedIterator{contract: _IdentityManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_IdentityManager *IdentityManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IdentityManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _IdentityManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityManagerInitialized)
				if err := _IdentityManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_IdentityManager *IdentityManagerFilterer) ParseInitialized(log types.Log) (*IdentityManagerInitialized, error) {
	event := new(IdentityManagerInitialized)
	if err := _IdentityManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityManagerSignedRootTransitedIterator is returned from FilterSignedRootTransited and is used to iterate over the raw logs and unpacked data for SignedRootTransited events raised by the IdentityManager contract.
type IdentityManagerSignedRootTransitedIterator struct {
	Event *IdentityManagerSignedRootTransited // Event containing the contract specifics and raw log

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
func (it *IdentityManagerSignedRootTransitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityManagerSignedRootTransited)
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
		it.Event = new(IdentityManagerSignedRootTransited)
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
func (it *IdentityManagerSignedRootTransitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityManagerSignedRootTransitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityManagerSignedRootTransited represents a SignedRootTransited event raised by the IdentityManager contract.
type IdentityManagerSignedRootTransited struct {
	PrevRoot   *big.Int
	PostRoot   *big.Int
	ReplacedAt *big.Int
	LatestRoot *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSignedRootTransited is a free log retrieval operation binding the contract event 0xc6c41b70b30e03e3cb385fb636d8b6388dc80e73605e86eace19b9e2e867126a.
//
// Solidity: event SignedRootTransited(uint256 prevRoot, uint256 postRoot, uint256 replacedAt, uint256 latestRoot)
func (_IdentityManager *IdentityManagerFilterer) FilterSignedRootTransited(opts *bind.FilterOpts) (*IdentityManagerSignedRootTransitedIterator, error) {

	logs, sub, err := _IdentityManager.contract.FilterLogs(opts, "SignedRootTransited")
	if err != nil {
		return nil, err
	}
	return &IdentityManagerSignedRootTransitedIterator{contract: _IdentityManager.contract, event: "SignedRootTransited", logs: logs, sub: sub}, nil
}

// WatchSignedRootTransited is a free log subscription operation binding the contract event 0xc6c41b70b30e03e3cb385fb636d8b6388dc80e73605e86eace19b9e2e867126a.
//
// Solidity: event SignedRootTransited(uint256 prevRoot, uint256 postRoot, uint256 replacedAt, uint256 latestRoot)
func (_IdentityManager *IdentityManagerFilterer) WatchSignedRootTransited(opts *bind.WatchOpts, sink chan<- *IdentityManagerSignedRootTransited) (event.Subscription, error) {

	logs, sub, err := _IdentityManager.contract.WatchLogs(opts, "SignedRootTransited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityManagerSignedRootTransited)
				if err := _IdentityManager.contract.UnpackLog(event, "SignedRootTransited", log); err != nil {
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

// ParseSignedRootTransited is a log parse operation binding the contract event 0xc6c41b70b30e03e3cb385fb636d8b6388dc80e73605e86eace19b9e2e867126a.
//
// Solidity: event SignedRootTransited(uint256 prevRoot, uint256 postRoot, uint256 replacedAt, uint256 latestRoot)
func (_IdentityManager *IdentityManagerFilterer) ParseSignedRootTransited(log types.Log) (*IdentityManagerSignedRootTransited, error) {
	event := new(IdentityManagerSignedRootTransited)
	if err := _IdentityManager.contract.UnpackLog(event, "SignedRootTransited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
