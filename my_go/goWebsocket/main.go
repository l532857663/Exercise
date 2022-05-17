package main

import (
	"encoding/json"
	"fmt"
	"goWebSocket/client"

	"github.com/gorilla/websocket"
)

type JsonRpcReq struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int64         `json:"id"`
}

func ChangeMessage(method string, param interface{}) []byte {
	var params []interface{}
	if param != nil {
		params = append(params, param)
	}
	req := JsonRpcReq{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		Id:      66,
	}
	msg, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return msg
}

func main() {
	fmt.Println("vim-go")
	host := "westend-rpc.polkadot.io"
	// host := "127.0.0.1:9944"
	path := ""
	scheme := "wss"
	wsCli, err := client.NewClient(host, path, scheme)
	if err != nil {
		fmt.Printf("New client error: %+v\n", err)
		return
	}
	wsCli.TheSocketConnInfo()
	// rpc_methods 查询请求方法
	// chain_getBlock  Get header and body of a relay chain block
	// "0x9895c428b646cb380b938c63d11288de5af1ce424489cb54b417bc61b9544ba9", // #10319162
	// state_subscribeStorage
	// params := []string{
	// 	"0x26aa394eea5630e07c48ae0c9558cef7b99d880ec681799c0cf30e8886371da96f612df0f4954496fbdfd118bce7b806aa906db4057dc591b516da53f1af0fcaeea158e96236b7938e1cb830794de554",
	// }
	params := []string{"0x26aa394eea5630e07c48ae0c9558cef7b99d880ec681799c0cf30e8886371da96f612df0f4954496fbdfd118bce7b806aa906db4057dc591b516da53f1af0fcaeea158e96236b7938e1cb830794de554"}
	msg := ChangeMessage("state_subscribeStorage", &params)
	fmt.Printf("wch------ params: %s\n", string(msg))
	// res := wsCli.SendMessage(msg)
	conn := wsCli.Conn
	// send message
	err = conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		// handle error
		fmt.Printf("Write message error: %+v\n", err)
		return
	}
	for {
		// receive message
		_, message, err := conn.ReadMessage()
		if err != nil {
			// handle error
			fmt.Printf("Read message error: %+v\n", err)
			return
		}
		fmt.Printf("wch------ res: %+v\n", string(message))
	}

	wsCli.Close()
	return
}
