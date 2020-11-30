package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	data := []string{
		// 原文 私钥 地址 签名数据
		// 导入账户
		//`{"entity_id":"1625","address":"0x44d6e1d986985495f0a1367eec0194a6886cf575","chain_id":"60","public_key":"04cfc82d6bc742dbe2bb4111c2056767963438723e3b31d66b0c62d842f7adc0c410fd570e6421f38451b396efc80c4b73aacf50e32764d65526e38de682c52a9a","description":"123564"}`,
		// 修改参数
		// `{"entity_id":"1625","address":"0x44d6e1d986985495f0a1367eec0194a6886cf575","chain_id":"60","name":"我的账户"}`,
		// "0xec54417b4b1a32ce2c272d35ae5e9a7ad38eefd6974df69d0cf9f407951d06fc",// f575
		// "0x44d6e1d986985495f0a1367eec0194a6886cf575",
		// 导入账户
		`{"entity_id":"1625","address":"0x04dc1cef86255e07a4041f532625b2ecb1203d48","chain_id":"60","public_key":"04362d3225f4a703eaec04a7be7399e709257f42724dcac0a2e9e0c6444c10d1b4628444c6526b69d86033ed9bd503c025222fd30dbddd1950e5831a5b2230cac8","name":"我的账户1"}`,
		// 确认操作
		// `{"entity_id":"1625","address":"0x04dc1cef86255e07a4041f532625b2ecb1203d48","chain_id":"60","id":"62"}`,
		// 修改参数
		// `{"entity_id":"1625","address":"0x04dc1cef86255e07a4041f532625b2ecb1203d48","chain_id":"60","name":"工资卡"}`,
		// `{"entity_id":"1625","address":"0x04dc1cef86255e07a4041f532625b2ecb1203d48","chain_id":"da","is_default":"1"}`,
		"0x00fc666a057b82048b8ea75838851eecad97ae7b53fddd9a1d6ae6b497167b0e", // 3d48
		"0x04dc1cef86255e07a4041f532625b2ecb1203d48",
		// 签名数据
		"0xd95c9bbba0e88d6aa30cb6a0d161fd12b35adb23208c49d2621bc16ee03790d1686c73adc17d72fcabb14d77399349c3020805ae617000ecd54f77bd7272395a01",
	}
	// 处理数据
	dataByte := []byte(data[0])
	msg := getSha256Hash(dataByte)
	// 私钥格式处理,不要私钥的0x
	priKeyStr := data[1]
	if len(priKeyStr) == 66 {
		priKeyStr = priKeyStr[2:]
	}
	priKey, err := crypto.HexToECDSA(priKeyStr)
	if err != nil {
		fmt.Println("err", err)
	}
	// 获取公钥
	pubKey := crypto.FromECDSAPub(&priKey.PublicKey)
	fmt.Println("wch-------------pubKey:", hex.EncodeToString(pubKey))
	// 获取地址
	addr := data[2]
	sigStr := data[3]
	// 签名
	sig, err := crypto.Sign(msg, priKey)
	sigStr1 := "0x" + hex.EncodeToString(sig)
	// fmt.Printf("encode type str: '%s'\n", sigStr1)
	fmt.Println("wch-------------data:", data[0])
	fmt.Println("wch-------------msg:", hex.EncodeToString(msg))
	fmt.Println("wch-------------address", addr)
	fmt.Println("wch-------------sign: signature", sigStr1)
	// 验签
	res := verifySig(
		addr,
		sigStr,
		msg,
	)
	fmt.Println("sign verify:", res)
}

func verifySig(from, sigHex string, msg []byte) bool {
	// 地址格式处理
	fromAddr := common.HexToAddress(from)
	// 签名数据格式处理
	var sig []byte
	var err error
	if len(sigHex) == 132 {
		sig = hexutil.MustDecode(sigHex)
	} else {
		sig, err = base64.StdEncoding.DecodeString(sigHex)
		if err != nil {
			fmt.Println("err:", err)
		}
	}
	// 处理以太坊的V值
	if sig[64] >= 27 {
		sig[64] -= 27
	}
	// 公钥，原文，签名数据 验签
	pubKey, err := crypto.SigToPub(msg, sig)
	if err != nil {
		fmt.Println("err:", err)
		return false
	}
	recoverAddr := crypto.PubkeyToAddress(*pubKey)
	fmt.Println("wch-------------address:", recoverAddr.String())
	return fromAddr == recoverAddr
}

func getSha256Hash(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}
