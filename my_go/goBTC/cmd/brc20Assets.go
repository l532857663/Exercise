package main

import (
	"fmt"
	"goBTC"
	"goBTC/client"
	"goBTC/global"
	"goBTC/service"
	"goBTC/utils"

	"go.uber.org/zap"
)

var (
	srv *client.BTCClient
	log *zap.Logger
)

func main() {
	fmt.Println("vim-go")
	goBTC.MustLoad("./config.yml")
	srv = global.Client
	log = global.LOG
	CheckBrc20Assets()
	if global.MysqlFlag {
		utils.SignalHandler("brc20Assets", goBTC.Shutdown)
	}
}

func CheckBrc20Assets() {
	req := service.GetTransferReq{
		Symbol: "btc",
		Height: "767753",
	}
	blockTxInfo := service.GetTransferInfoForBlock(req)
	if blockTxInfo.TotalPage != "0" {
		for i, txInfo := range blockTxInfo.TransferList {
			fmt.Printf("wch----- txInfo: %+v\n", txInfo)
			if i == 5 {
				break
			}
		}
	}
}
