package main

import (
	"fmt"
	"strings"

	"github.com/gcash/bchd/chaincfg"
	"github.com/gcash/bchd/chaincfg/chainhash"
	"github.com/gcash/bchd/rpcclient"
)

type BchClient struct {
	Client *rpcclient.Client
	Params *chaincfg.Params
}

type Node struct {
	Ip       string
	Port     int
	User     string
	Password string
	Net      string
}

func NewBchClient(conf *Node) (*BchClient, error) {
	var (
		url     string
		isHttps bool
	)
	if conf.Port != 0 {
		url = fmt.Sprintf("%s:%d", conf.Ip, conf.Port)
	} else {
		url = conf.Ip
	}

	// 某些https节点配置需要做一些特殊处理
	if strings.HasPrefix(url, "https://") {
		isHttps = true
		url = strings.TrimPrefix(url, "https://")
	}

	connCfg := &rpcclient.ConnConfig{
		Host:         url,
		User:         conf.User,
		Pass:         conf.Password,
		HTTPPostMode: true,     // Bitcoin core only supports HTTP POST mode
		DisableTLS:   !isHttps, // Bitcoin core does not provide TLS by default
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return nil, err
	}

	bchClient := &BchClient{
		Client: client,
	}
	upperNet := strings.ToUpper(conf.Net)
	switch upperNet {
	case "mainnet":
		bchClient.Params = &chaincfg.MainNetParams
	case "testnet3":
		bchClient.Params = &chaincfg.TestNet3Params
	case "regtest":
		bchClient.Params = &chaincfg.RegressionNetParams
	default:
		bchClient.Params = &chaincfg.Params{}
	}

	return bchClient, nil
}

func main() {
	fmt.Println("vim-go")
	conf := &Node{
		Ip:       "https://bch.getblock.io/mainnet/?api_key=73578b1d-56ce-47b6-9231-aacec95479fa",
		Port:     0,
		User:     "btc",
		Password: "btc2021",
		Net:      "mainnet",
	}

	srv, err := NewBchClient(conf)
	if err != nil {
		fmt.Printf("NewBchClient error: %+v\n", err)
		return
	}

	hash := "000000000000000000a541d38b848c6ba3dece05b68222f560a4ed4f1c13a052"
	h, err := chainhash.NewHashFromStr(hash)
	if err != nil {
		fmt.Println(err)
		return
	}
	msgBlock, err := srv.Client.GetBlock(h)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", msgBlock)
}
