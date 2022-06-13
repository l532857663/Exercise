package main

import (
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
)

type DogeClient struct {
	Client *rpcclient.Client
	Params *chaincfg.Params
}

type Conf struct {
	Ip       string
	Port     int64
	User     string
	Password string
	Net      string
}

func NewDogeClient() (*DogeClient, error) {
	conf := Conf{
		Ip:       "https://doge.getblock.io/mainnet/?api_key=73578b1d-56ce-47b6-9231-aacec95479fa",
		User:     "btc",
		Password: "btc2021",
		Net:      "regtest",
	}
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

	fmt.Println("btc node request url: ", url)

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

	btcClient := &DogeClient{
		Client: client,
	}
	btcClient.Params = &chaincfg.MainNetParams

	return btcClient, nil
}

func (b *DogeClient) GetBlockCount() (int64, error) {
	return b.Client.GetBlockCount()
}

func (b *DogeClient) GetBlockHash(height int64) (string, error) {
	hash, err := b.Client.GetBlockHash(height)
	if err != nil {
		return "", err
	}
	return hash.String(), err
}
