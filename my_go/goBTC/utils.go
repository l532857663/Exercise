package goBTC

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/wire"
)

func StuffMapBySlice(signMap map[string]byte, valueMap map[byte]string, signSlice []string, valSlice []byte) {
	if len(signSlice) != len(valSlice) {
		fmt.Printf("StuffMapBySlice lenSign: %+v, lenValue: %+v\n", len(signSlice), len(valSlice))
		return
	}
	for i, sign := range signSlice {
		val := valSlice[i]
		signMap[sign] = val
		valueMap[val] = sign
	}
	return
}

// 解析原始交易数据
func GetTxWitnessByTxHex(txHex string) string {
	// 将原始十六进制数据解析为交易结构体
	txBytes, err := hex.DecodeString(txHex)
	if err != nil {
		return ""
	}
	var tx wire.MsgTx
	tx.Deserialize(bytes.NewReader(txBytes))
	return GetTxWitness(&tx)
}

func GetTxWitness(tx *wire.MsgTx) string {
	// 遍历交易的输入，查找包含WITNESS_V1_TAPROOT数据的输入
	for _, input := range tx.TxIn {
		// 判断输入是否包含WITNESS_V1_TAPROOT数据
		if len(input.Witness) > 1 {
			fmt.Printf("wch---- input len(%v): %x\n", len(input.SignatureScript), input.SignatureScript)
			for _, data := range input.Witness {
				fmt.Printf("wch---- data: %x\n", data)
			}
			return fmt.Sprintf("%x", input.Witness[1])
		}
	}
	return ""
}
