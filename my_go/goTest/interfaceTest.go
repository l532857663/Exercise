package main

import (
	"encoding/json"
	"fmt"

	btcW "github.com/btcsuite/btcd/wire"
	dogeW "github.com/qhxcWallet/doged/wire"
)

type A struct {
	Tx1 *btcW.MsgTx
	Tx2 *dogeW.MsgTx
}

func GetObj(src []byte) (interface{}, error) {
	var tx *btcW.MsgTx
	if err := json.Unmarshal(src, &tx); err != nil {
		fmt.Println("wch----- err:", err)
		return nil, err
	}
	return tx, nil
}

func ChangeContent(a interface{}) {
	tx, ok := a.(*btcW.MsgTx)
	if !ok {
		fmt.Println("wch---- conver err")
		return
	}
	fmt.Printf("wch----: %+v\n", tx.TxIn[0])
	txIn := tx.TxIn[0]
	sign := "asdasd"
	txIn.SignatureScript = []byte(sign)
	fmt.Printf("wch---- new: %+v\n", tx.TxIn[0])

}

func Test() {
	rawTx := []byte(`{"Version":1,"TxIn":[{"PreviousOutPoint":{"Hash":[61,195,33,5,64,177,20,221,173,36,181,96,148,135,237,135,203,192,199,78,231,194,34,12,5,93,229,86,146,128,115,241],"Index":0}}]}`)
	tx, _ := GetObj(rawTx)
	ChangeContent(tx)
	fmt.Printf("wch----- tx new: %+v\n", tx)
}

func main() {
	Test()
}
