package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	// 还原交易对象
	encodedTxStr := "0xf889188504a817c800832dc6c09405e56888360ae54acf2a389bab39bd41e3934d2b80a4ee919d50000000000000000000000000000000000000000000000000000000000000007b25a041c4a2eb073e6df89c3f467b3516e9c313590d8d57f7c217fe7e72a7b4a6b8eda05f20a758396a5e681ce1ab4cec749f8560e28c9eb91072ec7a8acc002a11bb1d"
	encodedTx, err := hexutil.Decode(encodedTxStr)
	if err != nil {
		fmt.Println("hexutil.Decode failed: ", err.Error())
		return
	}
	fmt.Printf("wch------ test encoded %+v\n", string(encodedTx))
	// rlp解码
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(encodedTx, tx); err != nil {
		fmt.Println("rlp.DecodeBytes failed: ", err.Error())
		return
	}
	fmt.Printf("wch----- test tx %+v\n", tx)
	// chainId为1的EIP155签名器
	signer := types.NewEIP155Signer(big.NewInt(1))
	// 使用签名器从已签名的交易中还原账户公钥
	from, err := types.Sender(signer, tx)
	if err != nil {
		fmt.Println("types.Sender: ", err.Error())
		return
	}
	fmt.Println("from: ", from.Hex())
	jsonTx, _ := tx.MarshalJSON()
	fmt.Println("tx: ", string(jsonTx))
}
