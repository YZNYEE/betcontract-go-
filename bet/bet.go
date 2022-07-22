// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bet

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

// BetMetaData contains all meta data concerning the Bet contract.
var BetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"}],\"name\":\"changeowner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"getmoney\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getplayers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getplayersnum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getstate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inject\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expect\",\"type\":\"uint256\"}],\"name\":\"joinbet\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nouce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600260006101000a81548160ff0219169083151502179055506000600581905550611085806100766000396000f3fe6080604052600436106100a75760003560e01c806393b614771161006457806393b614771461018b578063942057de146101c8578063be9a655514610205578063ddfd70321461021c578063e8c27ba014610247578063fcfff16f14610263576100a7565b8063296d9ac2146100ac5780634d9b3d5d146100d7578063662e4ee41461010257806389b2a3461461012b5780638da5cb5b1461015657806390d4bcc014610181575b600080fd5b3480156100b857600080fd5b506100c161027a565b6040516100ce9190610a76565b60405180910390f35b3480156100e357600080fd5b506100ec610291565b6040516100f99190610aaa565b60405180910390f35b34801561010e57600080fd5b5061012960048036038101906101249190610b28565b610299565b005b34801561013757600080fd5b5061014061036a565b60405161014d9190610aaa565b60405180910390f35b34801561016257600080fd5b5061016b610370565b6040516101789190610b64565b60405180910390f35b610189610394565b005b34801561019757600080fd5b506101b260048036038101906101ad9190610b28565b610424565b6040516101bf9190610aaa565b60405180910390f35b3480156101d457600080fd5b506101ef60048036038101906101ea9190610bab565b61046d565b6040516101fc9190610b64565b60405180910390f35b34801561021157600080fd5b5061021a6104c6565b005b34801561022857600080fd5b506102316105a9565b60405161023e9190610aaa565b60405180910390f35b610261600480360381019061025c9190610bab565b6105b6565b005b34801561026f57600080fd5b506102786107d0565b005b6000600260009054906101000a900460ff16905090565b600047905090565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610327576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031e90610c5b565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60055481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610422576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041990610cc7565b60405180910390fd5b565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600480549050821061048057600080fd5b6004828154811061049457610493610ce7565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610554576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054b90610d62565b60405180910390fd5b60001515600260009054906101000a900460ff1615151461057457600080fd5b6005600081548092919061058790610db1565b91905055506001600260006101000a81548160ff021916908315150217905550565b6000600480549050905090565b600033905060011515600260009054906101000a900460ff16151514610611576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060890610e45565b60405180910390fd5b60018214806106205750600282145b61065f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065690610eb1565b60405180910390fd5b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054146106e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d890610f43565b60405180910390fd5b81600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555034600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506004339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461085e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085590610fd5565b60405180910390fd5b60006001905060005b6004805490508110156109e95760006004828154811061088a57610889610ce7565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205403610990578073ffffffffffffffffffffffffffffffffffffffff166108fc6002600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546109639190610ff5565b9081150290604051600060405180830381858888f1935050505015801561098e573d6000803e3d6000fd5b505b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505080806109e190610db1565b915050610867565b5060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc610a2c610291565b9081150290604051600060405180830381858888f19350505050158015610a57573d6000803e3d6000fd5b5050565b60008115159050919050565b610a7081610a5b565b82525050565b6000602082019050610a8b6000830184610a67565b92915050565b6000819050919050565b610aa481610a91565b82525050565b6000602082019050610abf6000830184610a9b565b92915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610af582610aca565b9050919050565b610b0581610aea565b8114610b1057600080fd5b50565b600081359050610b2281610afc565b92915050565b600060208284031215610b3e57610b3d610ac5565b5b6000610b4c84828501610b13565b91505092915050565b610b5e81610aea565b82525050565b6000602082019050610b796000830184610b55565b92915050565b610b8881610a91565b8114610b9357600080fd5b50565b600081359050610ba581610b7f565b92915050565b600060208284031215610bc157610bc0610ac5565b5b6000610bcf84828501610b96565b91505092915050565b600082825260208201905092915050565b7f546865206f776e6572206d7573742062652070726f766964656420696e206f7260008201527f64657220666f722074686520617574686f7269747920746f206368616e676500602082015250565b6000610c45603f83610bd8565b9150610c5082610be9565b604082019050919050565b60006020820190508181036000830152610c7481610c38565b9050919050565b7f4f6e6c79206f776e6572732063616e20696e6a656374206361706974616c0000600082015250565b6000610cb1601e83610bd8565b9150610cbc82610c7b565b602082019050919050565b60006020820190508181036000830152610ce081610ca4565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4f6e6c79206f776e6572732063616e2073746172742061206265740000000000600082015250565b6000610d4c601b83610bd8565b9150610d5782610d16565b602082019050919050565b60006020820190508181036000830152610d7b81610d3f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610dbc82610a91565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610dee57610ded610d82565b5b600182019050919050565b7f5468657265206d75737420626520612062657420696e2070726f677265737300600082015250565b6000610e2f601f83610bd8565b9150610e3a82610df9565b602082019050919050565b60006020820190508181036000830152610e5e81610e22565b9050919050565b7f657870656374206d7573742062652031206f7220320000000000000000000000600082015250565b6000610e9b601583610bd8565b9150610ea682610e65565b602082019050919050565b60006020820190508181036000830152610eca81610e8e565b9050919050565b7f73656e646572206d757374206e6f742062652020696e766f6c76656420696e2060008201527f63757272656e7420626574000000000000000000000000000000000000000000602082015250565b6000610f2d602b83610bd8565b9150610f3882610ed1565b604082019050919050565b60006020820190508181036000830152610f5c81610f20565b9050919050565b7f4f6e6c7920706f73736573736f72732077696c6c20707265736372696265207260008201527f6573756c74730000000000000000000000000000000000000000000000000000602082015250565b6000610fbf602683610bd8565b9150610fca82610f63565b604082019050919050565b60006020820190508181036000830152610fee81610fb2565b9050919050565b600061100082610a91565b915061100b83610a91565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561104457611043610d82565b5b82820290509291505056fea264697066735822122081647d9b0718fcfd34beec2f701213cf4be814604438a2e91682b7342789d47164736f6c634300080f0033",
}

// BetABI is the input ABI used to generate the binding from.
// Deprecated: Use BetMetaData.ABI instead.
var BetABI = BetMetaData.ABI

// BetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BetMetaData.Bin instead.
var BetBin = BetMetaData.Bin

// DeployBet deploys a new Ethereum contract, binding an instance of Bet to it.
func DeployBet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bet, error) {
	parsed, err := BetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bet{BetCaller: BetCaller{contract: contract}, BetTransactor: BetTransactor{contract: contract}, BetFilterer: BetFilterer{contract: contract}}, nil
}

// Bet is an auto generated Go binding around an Ethereum contract.
type Bet struct {
	BetCaller     // Read-only binding to the contract
	BetTransactor // Write-only binding to the contract
	BetFilterer   // Log filterer for contract events
}

// BetCaller is an auto generated read-only Go binding around an Ethereum contract.
type BetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BetSession struct {
	Contract     *Bet              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BetCallerSession struct {
	Contract *BetCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BetTransactorSession struct {
	Contract     *BetTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BetRaw is an auto generated low-level Go binding around an Ethereum contract.
type BetRaw struct {
	Contract *Bet // Generic contract binding to access the raw methods on
}

// BetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BetCallerRaw struct {
	Contract *BetCaller // Generic read-only contract binding to access the raw methods on
}

// BetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BetTransactorRaw struct {
	Contract *BetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBet creates a new instance of Bet, bound to a specific deployed contract.
func NewBet(address common.Address, backend bind.ContractBackend) (*Bet, error) {
	contract, err := bindBet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bet{BetCaller: BetCaller{contract: contract}, BetTransactor: BetTransactor{contract: contract}, BetFilterer: BetFilterer{contract: contract}}, nil
}

// NewBetCaller creates a new read-only instance of Bet, bound to a specific deployed contract.
func NewBetCaller(address common.Address, caller bind.ContractCaller) (*BetCaller, error) {
	contract, err := bindBet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BetCaller{contract: contract}, nil
}

// NewBetTransactor creates a new write-only instance of Bet, bound to a specific deployed contract.
func NewBetTransactor(address common.Address, transactor bind.ContractTransactor) (*BetTransactor, error) {
	contract, err := bindBet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BetTransactor{contract: contract}, nil
}

// NewBetFilterer creates a new log filterer instance of Bet, bound to a specific deployed contract.
func NewBetFilterer(address common.Address, filterer bind.ContractFilterer) (*BetFilterer, error) {
	contract, err := bindBet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BetFilterer{contract: contract}, nil
}

// bindBet binds a generic wrapper to an already deployed contract.
func bindBet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bet *BetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bet.Contract.BetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bet *BetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bet.Contract.BetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bet *BetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bet.Contract.BetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bet *BetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bet *BetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bet *BetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bet.Contract.contract.Transact(opts, method, params...)
}

// Getbalance is a free data retrieval call binding the contract method 0x4d9b3d5d.
//
// Solidity: function getbalance() view returns(uint256)
func (_Bet *BetCaller) Getbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bet.contract.Call(opts, &out, "getbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Getbalance is a free data retrieval call binding the contract method 0x4d9b3d5d.
//
// Solidity: function getbalance() view returns(uint256)
func (_Bet *BetSession) Getbalance() (*big.Int, error) {
	return _Bet.Contract.Getbalance(&_Bet.CallOpts)
}

// Getbalance is a free data retrieval call binding the contract method 0x4d9b3d5d.
//
// Solidity: function getbalance() view returns(uint256)
func (_Bet *BetCallerSession) Getbalance() (*big.Int, error) {
	return _Bet.Contract.Getbalance(&_Bet.CallOpts)
}

// Getmoney is a free data retrieval call binding the contract method 0x93b61477.
//
// Solidity: function getmoney(address player) view returns(uint256)
func (_Bet *BetCaller) Getmoney(opts *bind.CallOpts, player common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bet.contract.Call(opts, &out, "getmoney", player)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Getmoney is a free data retrieval call binding the contract method 0x93b61477.
//
// Solidity: function getmoney(address player) view returns(uint256)
func (_Bet *BetSession) Getmoney(player common.Address) (*big.Int, error) {
	return _Bet.Contract.Getmoney(&_Bet.CallOpts, player)
}

// Getmoney is a free data retrieval call binding the contract method 0x93b61477.
//
// Solidity: function getmoney(address player) view returns(uint256)
func (_Bet *BetCallerSession) Getmoney(player common.Address) (*big.Int, error) {
	return _Bet.Contract.Getmoney(&_Bet.CallOpts, player)
}

// Getplayers is a free data retrieval call binding the contract method 0x942057de.
//
// Solidity: function getplayers(uint256 id) view returns(address)
func (_Bet *BetCaller) Getplayers(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bet.contract.Call(opts, &out, "getplayers", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Getplayers is a free data retrieval call binding the contract method 0x942057de.
//
// Solidity: function getplayers(uint256 id) view returns(address)
func (_Bet *BetSession) Getplayers(id *big.Int) (common.Address, error) {
	return _Bet.Contract.Getplayers(&_Bet.CallOpts, id)
}

// Getplayers is a free data retrieval call binding the contract method 0x942057de.
//
// Solidity: function getplayers(uint256 id) view returns(address)
func (_Bet *BetCallerSession) Getplayers(id *big.Int) (common.Address, error) {
	return _Bet.Contract.Getplayers(&_Bet.CallOpts, id)
}

// Getplayersnum is a free data retrieval call binding the contract method 0xddfd7032.
//
// Solidity: function getplayersnum() view returns(uint256)
func (_Bet *BetCaller) Getplayersnum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bet.contract.Call(opts, &out, "getplayersnum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Getplayersnum is a free data retrieval call binding the contract method 0xddfd7032.
//
// Solidity: function getplayersnum() view returns(uint256)
func (_Bet *BetSession) Getplayersnum() (*big.Int, error) {
	return _Bet.Contract.Getplayersnum(&_Bet.CallOpts)
}

// Getplayersnum is a free data retrieval call binding the contract method 0xddfd7032.
//
// Solidity: function getplayersnum() view returns(uint256)
func (_Bet *BetCallerSession) Getplayersnum() (*big.Int, error) {
	return _Bet.Contract.Getplayersnum(&_Bet.CallOpts)
}

// Getstate is a free data retrieval call binding the contract method 0x296d9ac2.
//
// Solidity: function getstate() view returns(bool)
func (_Bet *BetCaller) Getstate(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bet.contract.Call(opts, &out, "getstate")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Getstate is a free data retrieval call binding the contract method 0x296d9ac2.
//
// Solidity: function getstate() view returns(bool)
func (_Bet *BetSession) Getstate() (bool, error) {
	return _Bet.Contract.Getstate(&_Bet.CallOpts)
}

// Getstate is a free data retrieval call binding the contract method 0x296d9ac2.
//
// Solidity: function getstate() view returns(bool)
func (_Bet *BetCallerSession) Getstate() (bool, error) {
	return _Bet.Contract.Getstate(&_Bet.CallOpts)
}

// Nouce is a free data retrieval call binding the contract method 0x89b2a346.
//
// Solidity: function nouce() view returns(uint256)
func (_Bet *BetCaller) Nouce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bet.contract.Call(opts, &out, "nouce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nouce is a free data retrieval call binding the contract method 0x89b2a346.
//
// Solidity: function nouce() view returns(uint256)
func (_Bet *BetSession) Nouce() (*big.Int, error) {
	return _Bet.Contract.Nouce(&_Bet.CallOpts)
}

// Nouce is a free data retrieval call binding the contract method 0x89b2a346.
//
// Solidity: function nouce() view returns(uint256)
func (_Bet *BetCallerSession) Nouce() (*big.Int, error) {
	return _Bet.Contract.Nouce(&_Bet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bet *BetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bet *BetSession) Owner() (common.Address, error) {
	return _Bet.Contract.Owner(&_Bet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bet *BetCallerSession) Owner() (common.Address, error) {
	return _Bet.Contract.Owner(&_Bet.CallOpts)
}

// Changeowner is a paid mutator transaction binding the contract method 0x662e4ee4.
//
// Solidity: function changeowner(address next) returns()
func (_Bet *BetTransactor) Changeowner(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Bet.contract.Transact(opts, "changeowner", next)
}

// Changeowner is a paid mutator transaction binding the contract method 0x662e4ee4.
//
// Solidity: function changeowner(address next) returns()
func (_Bet *BetSession) Changeowner(next common.Address) (*types.Transaction, error) {
	return _Bet.Contract.Changeowner(&_Bet.TransactOpts, next)
}

// Changeowner is a paid mutator transaction binding the contract method 0x662e4ee4.
//
// Solidity: function changeowner(address next) returns()
func (_Bet *BetTransactorSession) Changeowner(next common.Address) (*types.Transaction, error) {
	return _Bet.Contract.Changeowner(&_Bet.TransactOpts, next)
}

// Inject is a paid mutator transaction binding the contract method 0x90d4bcc0.
//
// Solidity: function inject() payable returns()
func (_Bet *BetTransactor) Inject(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bet.contract.Transact(opts, "inject")
}

// Inject is a paid mutator transaction binding the contract method 0x90d4bcc0.
//
// Solidity: function inject() payable returns()
func (_Bet *BetSession) Inject() (*types.Transaction, error) {
	return _Bet.Contract.Inject(&_Bet.TransactOpts)
}

// Inject is a paid mutator transaction binding the contract method 0x90d4bcc0.
//
// Solidity: function inject() payable returns()
func (_Bet *BetTransactorSession) Inject() (*types.Transaction, error) {
	return _Bet.Contract.Inject(&_Bet.TransactOpts)
}

// Joinbet is a paid mutator transaction binding the contract method 0xe8c27ba0.
//
// Solidity: function joinbet(uint256 _expect) payable returns()
func (_Bet *BetTransactor) Joinbet(opts *bind.TransactOpts, _expect *big.Int) (*types.Transaction, error) {
	return _Bet.contract.Transact(opts, "joinbet", _expect)
}

// Joinbet is a paid mutator transaction binding the contract method 0xe8c27ba0.
//
// Solidity: function joinbet(uint256 _expect) payable returns()
func (_Bet *BetSession) Joinbet(_expect *big.Int) (*types.Transaction, error) {
	return _Bet.Contract.Joinbet(&_Bet.TransactOpts, _expect)
}

// Joinbet is a paid mutator transaction binding the contract method 0xe8c27ba0.
//
// Solidity: function joinbet(uint256 _expect) payable returns()
func (_Bet *BetTransactorSession) Joinbet(_expect *big.Int) (*types.Transaction, error) {
	return _Bet.Contract.Joinbet(&_Bet.TransactOpts, _expect)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Bet *BetTransactor) Open(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bet.contract.Transact(opts, "open")
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Bet *BetSession) Open() (*types.Transaction, error) {
	return _Bet.Contract.Open(&_Bet.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Bet *BetTransactorSession) Open() (*types.Transaction, error) {
	return _Bet.Contract.Open(&_Bet.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Bet *BetTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bet.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Bet *BetSession) Start() (*types.Transaction, error) {
	return _Bet.Contract.Start(&_Bet.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Bet *BetTransactorSession) Start() (*types.Transaction, error) {
	return _Bet.Contract.Start(&_Bet.TransactOpts)
}
