package main

import (
	"fmt"
	"go2Polka/polkaclient"
)

var (
	chain      = "polkadot"
	apiNodeUrl = "https://polkadot.subscan.io/api"
	// wssNodeUrl = "wss://rpc.polkadot.io"
	wssNodeUrl        = "wss://westend-rpc.polkadot.io"
	address           = "5FvLuFtdv91A84PG6RmsDbzS7Q23PdBUxLSaWtfBB2BDhe7e"
	blockhash         = "0xbcaa7531d6caa18327eff14a22903aa7cad1fe7963d54a673f82cb9741f558e3"
	txhash            = "0x5edbcd340b68f9700b39aac69801d27c71352803dea80805678cb86e25bb9e56" // 10405237-2
	blocknum   uint64 = 10763536
	blocknum1  uint64 = 10405237
	// blocknum   uint64 = 10735225 // 广播成功 交易失败
	// blocknum   uint64 = 10267021 // 广播成功 多交易示例
	pubKey = "0xaa906db4057dc591b516da53f1af0fcaeea158e96236b7938e1cb830794de554"
	// pubKey = "0x705046c0c12d4ef9b46bcbfe74d481f5576c19c12c288ca11dc9c89ff1448b36"
	// priKey = "farm bonus shoulder liquid chapter depart measure race candy risk discover print" // 5FvLuFtdv91A84PG6RmsDbzS7Q23PdBUxLSaWtfBB2BDhe7e
	priKey = "shine quantum convince tattoo public boost remain lunar prefer wrong orient fame" // 5HDx4Pb4azp5tEwQCTr6xwTKvizrxsWjqhpPhr6GpJ7DDmzQ
)

func main() {
	fmt.Printf("wch-----: Start\n")
	node, err := polkaclient.NewNode(wssNodeUrl)
	if err != nil {
		fmt.Printf("New node error: %+v\n", wssNodeUrl)
		return
	}

	//  查询连接节点信息
	err = node.GetConnectedInfo()
	if err != nil {
		fmt.Printf("Get connected node info error: %+v\n", wssNodeUrl)
		return
	}

	// 查询最新块HASH
	// latestBlockHash, err := node.GetBlockHashLatest()
	// if err != nil {
	// 	fmt.Printf("Get block hash latest: %+v\n", err)
	// 	return
	// }
	// fmt.Printf("Get block hash latest: %+v\n", latestBlockHash)

	// 使用hash 查询块数据
	err = node.GetBlockByHash(blockhash)
	if err != nil {
		fmt.Printf("Get block by block hash [%+v]: %+v\n", blockhash, err)
		return
	}

	// accountInfo, _ := node.GetAccountInfoByPubKey(pubKey)
	// fmt.Printf("wch------- accountInfo: %+v\n", accountInfo)
	// // 使用块高 查询块数据
	// fmt.Printf("wch---- GetBlockByBlocknum- %v-------------------------------------------:\n", blocknum)
	// err = node.GetBlockByBlocknum(blocknum)
	// if err != nil {
	// 	fmt.Printf("Get block by block num [%+v]: %+v\n", blocknum, err)
	// 	return
	// }
	// fmt.Printf("wch---- GetBlockByBlocknum- %v-------------------------------------------:\n\n", blocknum)

	// // 使用公钥签名交易
	// fmt.Printf("wch---- DoTransfer--------------------------------------------:\n")
	// err = node.DoTransfer(priKey, pubKey)
	// if err != nil {
	// 	fmt.Printf("Do transfer by pubKey [%+v]: %+v\n", pubKey, err)
	// }
	// fmt.Printf("wch---- DoTransfer--------------------------------------------:\n\n")

	// node.GetAccountKeyPairFromSecret(priKey)

	// fmt.Printf("wch----- my address: %+v\n", address)

	// // 使用公钥获取账户余额
	// fmt.Println("wch---- GetBalanceFronPubKey--------------------------------------------:\n")
	// err = node.GetBalanceFromPubKey(pubKey)
	// if err != nil {
	// 	fmt.Printf("Get balance from pubKey error: %+v\n", pubKey)
	// 	return
	// }
	// fmt.Println("wch---- GetBalanceFronPubKey--------------------------------------------:\n\n")

	// // 使用公钥获取账户nonce
	// fmt.Println("wch---- GetNonceFromPubKey--------------------------------------------:\n")
	// err = node.GetNonceFromPubKey(pubKey)
	// if err != nil {
	// 	fmt.Printf("Get nonce from pubKey error: %+v\n", pubKey)
	// 	return
	// }
	// fmt.Println("wch---- GetNonceFromPubKey--------------------------------------------:\n\n")

	// 创建账户
	// polkaclient.GenerateAccount()
	// node.GetResultInfo("0x73010000020000000100000000000000ff4650f550000000000000000000000000000000000000000000000000000000354548414a0000000000000000000000354548414a0000000000000000000000")
	// node.GetResultInfo("0x010000000000000001000000000000005bf32530b90200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")

	fmt.Printf("wch-----: End\n")
}
