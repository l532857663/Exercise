package main

import (
	"fmt"
	"testing"
)

func Test_GenNetworkInfo(t *testing.T) {
	dogeClient, err := NewDogeClient()
	if err != nil {
		fmt.Println(err.Error())
	}

	result, err := dogeClient.Client.GetNetworkInfo()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)

	height, err := dogeClient.Client.GetBlockCount()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(height)

	hash, err := dogeClient.Client.GetBlockHash(height)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(hash)
}
