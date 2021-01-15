package main

import (
	"encoding/json"
	"fmt"
)

type AccountUtxo struct {
	Txid          string `json:"txid"`
	Value         string `json:"value"`
	Confirmations uint64 `json:"confirmations"`
	Height        uint64 `json:"height"`
	Vout          uint   `json:"vout"`
}

func main() {
	data := []AccountUtxo{}
	err := doChange(&data)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("ceshi: %+v\n", data)
}

func doChange(result interface{}) error {
	a := []byte(`[{"txid":"f5489bd7a89d22c8928deb76c7b15d0df0fc18ad84053e880c054a25c021f7b9","vout":0,"value":"95000","height":1901384,"confirmations":859},{"txid":"2cd60b22d74efcaae087794baabdebcfa86333bac8a644d90a0fa01b38c06ed1","vout":1,"value":"1990955","height":1901330,"confirmations":913}]`)
	err := json.Unmarshal(a, result)
	if err != nil {
		return err
	}
	return nil
}
