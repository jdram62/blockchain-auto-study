// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uni

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

// UniswapV2ERC20MetaData contains all meta data concerning the UniswapV2ERC20 contract.
var UniswapV2ERC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[]},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Burn\",\"inputs\":[{\"type\":\"address\",\"name\":\"sender\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount0\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount1\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Mint\",\"inputs\":[{\"type\":\"address\",\"name\":\"sender\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount0\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount1\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Swap\",\"inputs\":[{\"type\":\"address\",\"name\":\"sender\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount0In\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount1In\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount0Out\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount1Out\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Sync\",\"inputs\":[{\"type\":\"uint112\",\"name\":\"reserve0\",\"internalType\":\"uint112\",\"indexed\":false},{\"type\":\"uint112\",\"name\":\"reserve1\",\"internalType\":\"uint112\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"MINIMUM_LIQUIDITY\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"PERMIT_TYPEHASH\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"allowance\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"approve\",\"inputs\":[{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"balanceOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"amount0\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amount1\",\"internalType\":\"uint256\"}],\"name\":\"burn\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"decimals\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"factory\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint112\",\"name\":\"_reserve0\",\"internalType\":\"uint112\"},{\"type\":\"uint112\",\"name\":\"_reserve1\",\"internalType\":\"uint112\"},{\"type\":\"uint32\",\"name\":\"_blockTimestampLast\",\"internalType\":\"uint32\"}],\"name\":\"getReserves\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialize\",\"inputs\":[{\"type\":\"address\",\"name\":\"_token0\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_token1\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"kLast\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"liquidity\",\"internalType\":\"uint256\"}],\"name\":\"mint\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"name\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"nonces\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"permit\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"v\",\"internalType\":\"uint8\"},{\"type\":\"bytes32\",\"name\":\"r\",\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"s\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"price0CumulativeLast\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"price1CumulativeLast\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"skim\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"swap\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount0Out\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amount1Out\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"symbol\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"sync\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"token0\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"token1\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"totalSupply\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"transferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"}]}]",
}

// UniswapV2ERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV2ERC20MetaData.ABI instead.
var UniswapV2ERC20ABI = UniswapV2ERC20MetaData.ABI

// UniswapV2ERC20 is an auto generated Go binding around an Ethereum contract.
type UniswapV2ERC20 struct {
	UniswapV2ERC20Caller     // Read-only binding to the contract
	UniswapV2ERC20Transactor // Write-only binding to the contract
	UniswapV2ERC20Filterer   // Log filterer for contract events
}

// UniswapV2ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV2ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV2ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV2ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV2ERC20Session struct {
	Contract     *UniswapV2ERC20   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV2ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV2ERC20CallerSession struct {
	Contract *UniswapV2ERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UniswapV2ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV2ERC20TransactorSession struct {
	Contract     *UniswapV2ERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UniswapV2ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV2ERC20Raw struct {
	Contract *UniswapV2ERC20 // Generic contract binding to access the raw methods on
}

// UniswapV2ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV2ERC20CallerRaw struct {
	Contract *UniswapV2ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// UniswapV2ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV2ERC20TransactorRaw struct {
	Contract *UniswapV2ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV2ERC20 creates a new instance of UniswapV2ERC20, bound to a specific deployed contract.
func NewUniswapV2ERC20(address common.Address, backend bind.ContractBackend) (*UniswapV2ERC20, error) {
	contract, err := bindUniswapV2ERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV2ERC20{UniswapV2ERC20Caller: UniswapV2ERC20Caller{contract: contract}, UniswapV2ERC20Transactor: UniswapV2ERC20Transactor{contract: contract}, UniswapV2ERC20Filterer: UniswapV2ERC20Filterer{contract: contract}}, nil
}

// NewUniswapV2ERC20Caller creates a new read-only instance of UniswapV2ERC20, bound to a specific deployed contract.
func NewUniswapV2ERC20Caller(address common.Address, caller bind.ContractCaller) (*UniswapV2ERC20Caller, error) {
	contract, err := bindUniswapV2ERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2ERC20Caller{contract: contract}, nil
}

// NewUniswapV2ERC20Transactor creates a new write-only instance of UniswapV2ERC20, bound to a specific deployed contract.
func NewUniswapV2ERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV2ERC20Transactor, error) {
	contract, err := bindUniswapV2ERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2ERC20Transactor{contract: contract}, nil
}

// NewUniswapV2ERC20Filterer creates a new log filterer instance of UniswapV2ERC20, bound to a specific deployed contract.
func NewUniswapV2ERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV2ERC20Filterer, error) {
	contract, err := bindUniswapV2ERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV2ERC20Filterer{contract: contract}, nil
}

// bindUniswapV2ERC20 binds a generic wrapper to an already deployed contract.
func bindUniswapV2ERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapV2ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2ERC20 *UniswapV2ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2ERC20.Contract.UniswapV2ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2ERC20 *UniswapV2ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.UniswapV2ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2ERC20 *UniswapV2ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.UniswapV2ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2ERC20 *UniswapV2ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2ERC20 *UniswapV2ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2ERC20 *UniswapV2ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _UniswapV2ERC20.Contract.DOMAINSEPARATOR(&_UniswapV2ERC20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _UniswapV2ERC20.Contract.DOMAINSEPARATOR(&_UniswapV2ERC20.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _UniswapV2ERC20.Contract.MINIMUMLIQUIDITY(&_UniswapV2ERC20.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _UniswapV2ERC20.Contract.MINIMUMLIQUIDITY(&_UniswapV2ERC20.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) PERMITTYPEHASH() ([32]byte, error) {
	return _UniswapV2ERC20.Contract.PERMITTYPEHASH(&_UniswapV2ERC20.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _UniswapV2ERC20.Contract.PERMITTYPEHASH(&_UniswapV2ERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _UniswapV2ERC20.Contract.Allowance(&_UniswapV2ERC20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _UniswapV2ERC20.Contract.Allowance(&_UniswapV2ERC20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _UniswapV2ERC20.Contract.BalanceOf(&_UniswapV2ERC20.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _UniswapV2ERC20.Contract.BalanceOf(&_UniswapV2ERC20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Decimals() (uint8, error) {
	return _UniswapV2ERC20.Contract.Decimals(&_UniswapV2ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) Decimals() (uint8, error) {
	return _UniswapV2ERC20.Contract.Decimals(&_UniswapV2ERC20.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Factory() (common.Address, error) {
	return _UniswapV2ERC20.Contract.Factory(&_UniswapV2ERC20.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) Factory() (common.Address, error) {
	return _UniswapV2ERC20.Contract.Factory(&_UniswapV2ERC20.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _UniswapV2ERC20.Contract.GetReserves(&_UniswapV2ERC20.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _UniswapV2ERC20.Contract.GetReserves(&_UniswapV2ERC20.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) KLast() (*big.Int, error) {
	return _UniswapV2ERC20.Contract.KLast(&_UniswapV2ERC20.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) KLast() (*big.Int, error) {
	return _UniswapV2ERC20.Contract.KLast(&_UniswapV2ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Name() (string, error) {
	return _UniswapV2ERC20.Contract.Name(&_UniswapV2ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) Name() (string, error) {
	return _UniswapV2ERC20.Contract.Name(&_UniswapV2ERC20.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _UniswapV2ERC20.Contract.Nonces(&_UniswapV2ERC20.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _UniswapV2ERC20.Contract.Nonces(&_UniswapV2ERC20.CallOpts, arg0)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Price0CumulativeLast() (*big.Int, error) {
	return _UniswapV2ERC20.Contract.Price0CumulativeLast(&_UniswapV2ERC20.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _UniswapV2ERC20.Contract.Price0CumulativeLast(&_UniswapV2ERC20.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Price1CumulativeLast() (*big.Int, error) {
	return _UniswapV2ERC20.Contract.Price1CumulativeLast(&_UniswapV2ERC20.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _UniswapV2ERC20.Contract.Price1CumulativeLast(&_UniswapV2ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Symbol() (string, error) {
	return _UniswapV2ERC20.Contract.Symbol(&_UniswapV2ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) Symbol() (string, error) {
	return _UniswapV2ERC20.Contract.Symbol(&_UniswapV2ERC20.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Token0() (common.Address, error) {
	return _UniswapV2ERC20.Contract.Token0(&_UniswapV2ERC20.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) Token0() (common.Address, error) {
	return _UniswapV2ERC20.Contract.Token0(&_UniswapV2ERC20.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Token1() (common.Address, error) {
	return _UniswapV2ERC20.Contract.Token1(&_UniswapV2ERC20.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) Token1() (common.Address, error) {
	return _UniswapV2ERC20.Contract.Token1(&_UniswapV2ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV2ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) TotalSupply() (*big.Int, error) {
	return _UniswapV2ERC20.Contract.TotalSupply(&_UniswapV2ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UniswapV2ERC20 *UniswapV2ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _UniswapV2ERC20.Contract.TotalSupply(&_UniswapV2ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UniswapV2ERC20 *UniswapV2ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapV2ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Approve(&_UniswapV2ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UniswapV2ERC20 *UniswapV2ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Approve(&_UniswapV2ERC20.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_UniswapV2ERC20 *UniswapV2ERC20Transactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _UniswapV2ERC20.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Burn(to common.Address) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Burn(&_UniswapV2ERC20.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_UniswapV2ERC20 *UniswapV2ERC20TransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Burn(&_UniswapV2ERC20.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_UniswapV2ERC20 *UniswapV2ERC20Transactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _UniswapV2ERC20.contract.Transact(opts, "initialize", _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Initialize(&_UniswapV2ERC20.TransactOpts, _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_UniswapV2ERC20 *UniswapV2ERC20TransactorSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Initialize(&_UniswapV2ERC20.TransactOpts, _token0, _token1)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_UniswapV2ERC20 *UniswapV2ERC20Transactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _UniswapV2ERC20.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Mint(to common.Address) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Mint(&_UniswapV2ERC20.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_UniswapV2ERC20 *UniswapV2ERC20TransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Mint(&_UniswapV2ERC20.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_UniswapV2ERC20 *UniswapV2ERC20Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV2ERC20.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Permit(&_UniswapV2ERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_UniswapV2ERC20 *UniswapV2ERC20TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Permit(&_UniswapV2ERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_UniswapV2ERC20 *UniswapV2ERC20Transactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _UniswapV2ERC20.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Skim(to common.Address) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Skim(&_UniswapV2ERC20.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_UniswapV2ERC20 *UniswapV2ERC20TransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Skim(&_UniswapV2ERC20.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_UniswapV2ERC20 *UniswapV2ERC20Transactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _UniswapV2ERC20.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Swap(&_UniswapV2ERC20.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_UniswapV2ERC20 *UniswapV2ERC20TransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Swap(&_UniswapV2ERC20.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_UniswapV2ERC20 *UniswapV2ERC20Transactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2ERC20.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Sync() (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Sync(&_UniswapV2ERC20.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_UniswapV2ERC20 *UniswapV2ERC20TransactorSession) Sync() (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Sync(&_UniswapV2ERC20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_UniswapV2ERC20 *UniswapV2ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapV2ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Transfer(&_UniswapV2ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_UniswapV2ERC20 *UniswapV2ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.Transfer(&_UniswapV2ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_UniswapV2ERC20 *UniswapV2ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapV2ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_UniswapV2ERC20 *UniswapV2ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.TransferFrom(&_UniswapV2ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_UniswapV2ERC20 *UniswapV2ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapV2ERC20.Contract.TransferFrom(&_UniswapV2ERC20.TransactOpts, from, to, value)
}

// UniswapV2ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the UniswapV2ERC20 contract.
type UniswapV2ERC20ApprovalIterator struct {
	Event *UniswapV2ERC20Approval // Event containing the contract specifics and raw log

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
func (it *UniswapV2ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV2ERC20Approval)
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
		it.Event = new(UniswapV2ERC20Approval)
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
func (it *UniswapV2ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV2ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV2ERC20Approval represents a Approval event raised by the UniswapV2ERC20 contract.
type UniswapV2ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*UniswapV2ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UniswapV2ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV2ERC20ApprovalIterator{contract: _UniswapV2ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UniswapV2ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UniswapV2ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV2ERC20Approval)
				if err := _UniswapV2ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) ParseApproval(log types.Log) (*UniswapV2ERC20Approval, error) {
	event := new(UniswapV2ERC20Approval)
	if err := _UniswapV2ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV2ERC20BurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the UniswapV2ERC20 contract.
type UniswapV2ERC20BurnIterator struct {
	Event *UniswapV2ERC20Burn // Event containing the contract specifics and raw log

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
func (it *UniswapV2ERC20BurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV2ERC20Burn)
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
		it.Event = new(UniswapV2ERC20Burn)
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
func (it *UniswapV2ERC20BurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV2ERC20BurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV2ERC20Burn represents a Burn event raised by the UniswapV2ERC20 contract.
type UniswapV2ERC20Burn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*UniswapV2ERC20BurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UniswapV2ERC20.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV2ERC20BurnIterator{contract: _UniswapV2ERC20.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *UniswapV2ERC20Burn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UniswapV2ERC20.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV2ERC20Burn)
				if err := _UniswapV2ERC20.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) ParseBurn(log types.Log) (*UniswapV2ERC20Burn, error) {
	event := new(UniswapV2ERC20Burn)
	if err := _UniswapV2ERC20.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV2ERC20MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the UniswapV2ERC20 contract.
type UniswapV2ERC20MintIterator struct {
	Event *UniswapV2ERC20Mint // Event containing the contract specifics and raw log

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
func (it *UniswapV2ERC20MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV2ERC20Mint)
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
		it.Event = new(UniswapV2ERC20Mint)
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
func (it *UniswapV2ERC20MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV2ERC20MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV2ERC20Mint represents a Mint event raised by the UniswapV2ERC20 contract.
type UniswapV2ERC20Mint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*UniswapV2ERC20MintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _UniswapV2ERC20.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV2ERC20MintIterator{contract: _UniswapV2ERC20.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *UniswapV2ERC20Mint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _UniswapV2ERC20.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV2ERC20Mint)
				if err := _UniswapV2ERC20.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) ParseMint(log types.Log) (*UniswapV2ERC20Mint, error) {
	event := new(UniswapV2ERC20Mint)
	if err := _UniswapV2ERC20.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV2ERC20SwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the UniswapV2ERC20 contract.
type UniswapV2ERC20SwapIterator struct {
	Event *UniswapV2ERC20Swap // Event containing the contract specifics and raw log

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
func (it *UniswapV2ERC20SwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV2ERC20Swap)
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
		it.Event = new(UniswapV2ERC20Swap)
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
func (it *UniswapV2ERC20SwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV2ERC20SwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV2ERC20Swap represents a Swap event raised by the UniswapV2ERC20 contract.
type UniswapV2ERC20Swap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*UniswapV2ERC20SwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UniswapV2ERC20.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV2ERC20SwapIterator{contract: _UniswapV2ERC20.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *UniswapV2ERC20Swap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UniswapV2ERC20.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV2ERC20Swap)
				if err := _UniswapV2ERC20.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) ParseSwap(log types.Log) (*UniswapV2ERC20Swap, error) {
	event := new(UniswapV2ERC20Swap)
	if err := _UniswapV2ERC20.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV2ERC20SyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the UniswapV2ERC20 contract.
type UniswapV2ERC20SyncIterator struct {
	Event *UniswapV2ERC20Sync // Event containing the contract specifics and raw log

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
func (it *UniswapV2ERC20SyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV2ERC20Sync)
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
		it.Event = new(UniswapV2ERC20Sync)
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
func (it *UniswapV2ERC20SyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV2ERC20SyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV2ERC20Sync represents a Sync event raised by the UniswapV2ERC20 contract.
type UniswapV2ERC20Sync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) FilterSync(opts *bind.FilterOpts) (*UniswapV2ERC20SyncIterator, error) {

	logs, sub, err := _UniswapV2ERC20.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &UniswapV2ERC20SyncIterator{contract: _UniswapV2ERC20.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) WatchSync(opts *bind.WatchOpts, sink chan<- *UniswapV2ERC20Sync) (event.Subscription, error) {

	logs, sub, err := _UniswapV2ERC20.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV2ERC20Sync)
				if err := _UniswapV2ERC20.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) ParseSync(log types.Log) (*UniswapV2ERC20Sync, error) {
	event := new(UniswapV2ERC20Sync)
	if err := _UniswapV2ERC20.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV2ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the UniswapV2ERC20 contract.
type UniswapV2ERC20TransferIterator struct {
	Event *UniswapV2ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *UniswapV2ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV2ERC20Transfer)
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
		it.Event = new(UniswapV2ERC20Transfer)
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
func (it *UniswapV2ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV2ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV2ERC20Transfer represents a Transfer event raised by the UniswapV2ERC20 contract.
type UniswapV2ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UniswapV2ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UniswapV2ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV2ERC20TransferIterator{contract: _UniswapV2ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UniswapV2ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UniswapV2ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV2ERC20Transfer)
				if err := _UniswapV2ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UniswapV2ERC20 *UniswapV2ERC20Filterer) ParseTransfer(log types.Log) (*UniswapV2ERC20Transfer, error) {
	event := new(UniswapV2ERC20Transfer)
	if err := _UniswapV2ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
