package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// data := []byte(`{"entity_id":"1121","address":"0x1018649f26e9455fedc885ae372e2dbbf3a62892","chain_id":"60","name":"哦哦哦"}`)
	data := []byte(`{"address":"0x04dc1cef86255e07a4041f532625b2ecb1203d48","chain_id":"60","entity_id":"1624","is_default":"1"}`)
	msg := GetSha256Hash(data)
	res := verifySig(
		"0x04dc1cef86255e07a4041f532625b2ecb1203d48",
		"0x451bc24d7c86bce3797b04760ccc0647237c64e83290205b280bd39f2f6285006eb791b1978ccc12e919c29552a23c8898f26e89b2eb4e0420cba850bdf53d6b1b",
		// "0x98eB720D8AAD7fF5374F59657C45196Cb9cE8983",
		// "1844760377298905130805959797594268268221065089234216538189440842217493831405595422202112113281727268467818414345556854904111585729505377862744732123385349888",
		msg,
	)
	fmt.Println(res)
	/*
		//生成私钥
		key, err := crypto.GenerateKey()
		if err != nil {
			t.Fatalf("failed GenerateKey with %s.", err)
		}
		//不含0x的私钥
		fmt.Println("private key no 0x \n", hex.EncodeToString(crypto.FromECDSA(key)))
	*/
	key := "6a865a48a8fd9f422232b6b38978646fd57ab7b75c3fa785619fbbc56011e9e3"
	// 私钥 to 公钥
	Pri, err := crypto.HexToECDSA(key)
	if err != nil {
		fmt.Println("err", err)
	}
	pub := Pri.PublicKey
	// 公钥 to 地址
	address1 := crypto.PubkeyToAddress(Pri.PublicKey)
	fmt.Printf("pri: %+v\n pub: %+v\n address: %+v\n", Pri, pub, address1.String())
	// 签名
	sig, err := crypto.Sign(msg, Pri)
	sigStr := hex.EncodeToString(sig)
	// sigStr := new(big.Int).SetBytes(sig).String()

	//推出公钥字节
	recoveredPub, err := crypto.Ecrecover(msg, sig)
	//字节转 公钥
	pubKey, _ := crypto.UnmarshalPubkey(recoveredPub)
	fmt.Printf("sig: %+v\n pubKey:%+v\n", sigStr, pubKey)
	return
}

func verifySig(from, sigHex string, msg []byte) bool {
	fromAddr := common.HexToAddress(from)
	fmt.Println("sig:", sigHex)

	sig := hexutil.MustDecode(sigHex)
	//sigInt, _ := new(big.Int).SetString(sigHex, 0)
	//sig := sigInt.Bytes()

	fmt.Println("msg:", msg)

	pubKey, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	fmt.Println("address:", recoveredAddr)

	return fromAddr == recoveredAddr
}

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func GetSha256Hash(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}
