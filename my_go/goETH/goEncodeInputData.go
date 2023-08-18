package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type Struct0 struct {
	Target   common.Address
	CallData []byte
}

func main() {
	calls := []Struct0{
		Struct0{
			Target:   common.HexToAddress("0x1A72D8B4E0738c8f3871E48D91782CD7F899e48D"),
			CallData: []byte("0x4d2301cc00000000000000000000000039447c3040124057147512c3d1477dac339fcf8c"),
		},
	}
	fmt.Printf("calls: %+v\n", calls)
	params := "aggregate"
	// contractAbi, argsNew, err := GetAbiAndArgs(abiStr, params, calls)
	// if err != nil {
	// 	fmt.Printf("GetAbiAndArgs error: %+v\n", err)
	// 	return
	// }
	// Erc20交易
	fmt.Printf("wch--- Erc20Abi: %+v\n", Erc20Abi.Methods[params])
	input, err := Erc20Abi.Pack(params, calls)
	if nil != err {
		fmt.Printf("contractAbi.Pack error: %+v\n", err)
		return
	}
	fmt.Printf("get input: %+v\n", input)
	fmt.Printf("get input str: %+v\n", hex.Dump(input))
}

// 根据合约方法处理参数类型
func GetAbiAndArgs(abiContent, params string, args []interface{}) (abi.ABI, []interface{}, error) {
	// Abi转化
	contractAbi, err := StringToAbi(abiContent)
	if err != nil {
		fmt.Printf("StringToAbi error: %+v\n", err)
		return contractAbi, nil, err
	}
	var argsNew []interface{}
	abiParam := contractAbi.Methods[params].Inputs
	for i, v := range abiParam {
		arg := ChangeArgType(args[i], v.Type.T)
		_, ok := arg.(error)
		if arg == nil || ok {
			continue
		}
		argsNew = append(argsNew, arg)
	}
	// 检查参数数量是否匹配
	if len(argsNew) != len(abiParam) {
		err := fmt.Errorf("The args len not enough")
		return contractAbi, nil, err
	}
	return contractAbi, argsNew, nil
}

func ChangeArgType(arg interface{}, argType byte) interface{} {
	argStr := arg.(string)
	switch argType {
	case abi.AddressTy:
		addr := EthAddressChange(argStr)
		if addr.String() != argStr {
			return nil
		}
		return addr
	case abi.UintTy:
		val, ok := big.NewInt(0).SetString(argStr, 10)
		if !ok {
			return nil
		}
		return val
	}
	return nil
}

func EthAddressChange(addr string) common.Address {
	return common.HexToAddress(addr)
}

func StringToAbi(abiContent string) (abi.ABI, error) {
	return abi.JSON(strings.NewReader(abiContent))
}

var (
	abiStr      = `[{"constant":true,"inputs":[],"name":"getCurrentBlockTimestamp","outputs":[{"name":"timestamp","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"components":[{"name":"target","type":"address"},{"name":"callData","type":"bytes"}],"name":"calls","type":"tuple[]"}],"name":"aggregate","outputs":[{"name":"blockNumber","type":"uint256"},{"name":"returnData","type":"bytes[]"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getLastBlockHash","outputs":[{"name":"blockHash","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"addr","type":"address"}],"name":"getEthBalance","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getCurrentBlockDifficulty","outputs":[{"name":"difficulty","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getCurrentBlockGasLimit","outputs":[{"name":"gaslimit","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getCurrentBlockCoinbase","outputs":[{"name":"coinbase","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"blockNumber","type":"uint256"}],"name":"getBlockHash","outputs":[{"name":"blockHash","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"}]`
	Erc20Abi, _ = abi.JSON(strings.NewReader(abiStr))
)
