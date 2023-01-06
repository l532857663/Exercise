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
		// `{"entity_id":"1625","address":"0x04dc1cef86255e07a4041f532625b2ecb1203d48","chain_id":"60","public_key":"04362d3225f4a703eaec04a7be7399e709257f42724dcac0a2e9e0c6444c10d1b4628444c6526b69d86033ed9bd503c025222fd30dbddd1950e5831a5b2230cac8","name":"我的账户1"}`,
		"fdsafdsafdsa",
		"0x550f14df38b6907dd55156d7cc15632bc5fd7adca4e6ba7e4b2993d50d01ee1f", // 8d42
		"0x03A26A82c474A8a4743f196269aDcd820a098d42",
		// // 签名数据
		// "0x1c68af34541086f129dab79c37fd69af107b382e21299f6eb3745e87886beac70dbacb35dfd8a389521b504c7c09b30b9f689e3d524c953550d396230bd1d57601",
		// "0x9718306f984a08fc3C1555BF7dbEc1e008398bD1",
		// "0xb0f0bc015eda088c2cf32d02543de40064fd161f1d80ec9607b3711db0de6af8482cee80cf9699a5d0b24243d263ba613d35e3ffa660a45f1753d3a84589038d1c",
		"0x2e15943760c05e3070ea09541e4f794a0e65eda21eebdd9749c10706d128cdab01602b4886dd28496e52245843cd58633956d6d8748d2c30f0ee433596c63dfc01",
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
	fmt.Printf("wch------- sig len: %+v\n", len(sigHex))
	if len(sigHex) == 132 {
		sig = hexutil.MustDecode(sigHex)
	} else {
		sig, err = base64.StdEncoding.DecodeString(sigHex)
		if err != nil {
			fmt.Println("err:", err)
		}
	}
	// // 处理以太坊的V值
	// if sig[64] >= 27 {
	// 	sig[64] -= 27
	// }
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
