package main

import (
	"fmt"
	"goBTC"
	"goBTC/client"
	"goBTC/global"
	"goBTC/ord"
	"goBTC/utils"
	"time"

	"go.uber.org/zap"
)

var (
	srv *client.BTCClient
	log *zap.Logger
)

func main() {
	fmt.Println("vim-go")
	global.MysqlFlag = true
	goBTC.MustLoad("./config.yml")
	srv = global.Client
	log = global.LOG
	go GetBlockInfo(767430)
	utils.SignalHandler("scan", goBTC.Shutdown)
}

func GetBlockInfo(startHeight int64) {
	fmt.Println("[GetBlockInfo] Start")
	newHigh, err := srv.GetBlockCount()
	if err != nil {
		fmt.Printf("GetBlockCount error: %+v\n", err)
		return
	}
	for i := startHeight; i < newHigh; i++ {
		startTime := time.Now().Unix()
		blockInfo, err := srv.GetBlockInfoByHeight(i)
		if err != nil {
			log.Error("GetBlockInfoByHash", zap.Error(err))
			i--
			continue
		}
		endTime := time.Now().Unix()
		log.Info("Get block info", zap.Any("block height", i), zap.Any("have tx", len(blockInfo.Transactions)), zap.Any("time", endTime-startTime))
		log.Debug("Get block", zap.Any("header", blockInfo.Header))
		sum := 0
		for j, tx := range blockInfo.Transactions[900:] {
			witnessStr := client.GetTxWitness(tx)
			if witnessStr == "" {
				continue
			}
			res := client.GetScriptString(witnessStr)
			if res != nil {
				txHash := tx.TxHash().String()
				fmt.Printf("[%d] txHash: %s, [ord] : %v\n", j, txHash, res.ContentType)
				txInfo, err := srv.GetRawTransactionByHash(txHash)
				if err != nil {
					log.Error("GetRawTransactionByHash", zap.Error(err))
					continue
				}
				err = ord.SaveInscribeInfoByTxInfo(i, res, txInfo)
				if err != nil {
					log.Error("SaveInscribeInfoByTxInfo", zap.Error(err))
					continue
				}
				if res.Brc20 != nil {
					err := ord.SaveInscribeBrc20ByTxInfo(i, res, txInfo)
					if err != nil {
						log.Error("SaveInscribeBrc20ByTxInfo", zap.Error(err))
						continue
					}
				}
				sum++
			}
		}
		log.Info("the block get inscribe:", zap.Any("sum", sum))
	}
}
