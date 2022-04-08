package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip32"
)

func main() {
	transferETH()
	return
}

func transferETH() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chainId: ", chainId)

	// priKey, err := generateKeyFromTgSeed()
	priKey, err := crypto.HexToECDSA("b5a0d2d2f0e031e0f6600aa98b50380ae39eebf87fc13f293de956597ba76df8")
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := crypto.PubkeyToAddress(priKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// nonce := uint64(0)

	// value := big.NewInt(11000000000000000) // in wei (0.011 eth)
	value, ok := new(big.Int).SetString("10000000000000", 10) // in wei (0.00001 eth)
	if !ok {
		fmt.Println("转换转账金额失败")
		return
	}
	gasLimit := uint64(21000) // in unit
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// FIXME: only for test
	gasPrice = big.NewInt(10000000000)

	// 设置交易信息
	toAddress := common.HexToAddress("0x77A50402d4d62A1b65F14cE79e8dA0de9337d982") // work metamask MBP
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    value,
	})

	fmt.Println("*************************************************************************************************************")
	fmt.Println("gasLimit: ", gasLimit)
	fmt.Println("gasPrice: ", gasPrice)
	fmt.Println("fromAddress: ", fromAddress)
	fmt.Println("toAddress: ", toAddress)
	fmt.Println("*************************************************************************************************************")

	// 设置signer
	var signer types.Signer
	signer = types.LatestSignerForChainID(chainId)

	// FIXME: only for test
	type eip155Tx struct {
		Nonce    uint64          // nonce of sender account
		GasPrice *big.Int        // wei per gas
		Gas      uint64          // gas limit
		To       *common.Address `rlp:"nil"` // nil means contract creation
		Value    *big.Int        // wei amount
		Data     []byte          // contract invocation input data
		ChainId  *big.Int
		R        uint
		S        uint
	}
	eip155tx := &eip155Tx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    value,
		ChainId:  chainId,
	}
	rawData, _ := json.Marshal(eip155tx)
	fmt.Println("*************************************************************************************************************")
	fmt.Println("rawData: ", rawData)
	fmt.Println("hex.EncodeToString(rawData): ", hex.EncodeToString(rawData))
	fmt.Println("* base64(rawData): ", base64.StdEncoding.EncodeToString(rawData))
	fmt.Println("txHash: ", tx.Hash().String())
	fmt.Println("signer.Hash: ", signer.Hash(tx).String())
	fmt.Println("txHash Bytes(): ", tx.Hash().Bytes())
	fmt.Println("signer.Hash Bytes(): ", signer.Hash(tx).Bytes())
	fmt.Println("signer.Hash string: ", signer.Hash(tx).String())
	fmt.Println("signer.Hash hex2string: ", hex.EncodeToString(signer.Hash(tx).Bytes()))
	// base64 signer.Hash Bytes() 用于HSM签名交易
	fmt.Println("* base64 signer.Hash Bytes(): ", base64.StdEncoding.EncodeToString(signer.Hash(tx).Bytes()))
	fmt.Println("*************************************************************************************************************")

	///////////////////////////////// ///////////////////////////////// ///////////////////////////////// /////////////////////////////////
	///////////////////////////////// ///////////////////////////////// ///////////////////////////////// /////////////////////////////////
	///////////////////////////////// ///////////////////////////////// ///////////////////////////////// /////////////////////////////////
	// 私钥直接签名交易
	pkSignedTx, err := types.SignTx(tx, signer, priKey)
	if err != nil {
		fmt.Println("prikey sign tx error: ", err)
		return
	}
	fmt.Println("*************************************************************************************************************")
	rawPkSignedTx, _ := json.Marshal(pkSignedTx)
	fmt.Println("rawPkSignedTx: ", rawPkSignedTx)
	fmt.Println("hex.EncodeToString(rawpkSignedTx): ", hex.EncodeToString(rawPkSignedTx))
	fmt.Println("base64(rawpkSignedTx): ", base64.StdEncoding.EncodeToString(rawPkSignedTx))

	// v, r, s := pkSignedTx.RawSignatureValues()
	// fmt.Println("*************************************************************************************************************")
	// fmt.Println("pkSignature: ", v.Bytes(), r.Bytes(), s.Bytes())
	// fmt.Println("r: ", hex.EncodeToString(r.Bytes()))
	// fmt.Println("s: ", hex.EncodeToString(s.Bytes()))
	// fmt.Println("v: ", hex.EncodeToString(v.Bytes()))
	// fmt.Println("*************************************************************************************************************")
	// 私钥直接签名交易 end
	///////////////////////////////// ///////////////////////////////// ///////////////////////////////// /////////////////////////////////
	///////////////////////////////// ///////////////////////////////// ///////////////////////////////// /////////////////////////////////
	///////////////////////////////// ///////////////////////////////// ///////////////////////////////// /////////////////////////////////

	// 解析sender 方便调试信息, 这个对不上的话 交易是不会成功的
	// sender, err := types.Sender(signer, signedTx)
	sender, err := types.Sender(signer, pkSignedTx)
	if err != nil {
		fmt.Println("types.Sender error: ", err)
		return
	}
	fmt.Println("*************************************************************************************************************")
	fmt.Println("real sender address: ", sender)
	fmt.Println("*************************************************************************************************************")

	// 输出签名交易内容
	// spew.Dump(signedTx)
	spew.Dump(pkSignedTx)

	// 发送交易
	// err = client.SendTransaction(context.Background(), signedTx)
	err = client.SendTransaction(context.Background(), pkSignedTx)
	if err != nil {
		fmt.Println("Error SendTransaction: ", err)
		return
	}
}

func pubkey2Address() {
	// Get the public key
	// Base64 Standard Decoding
	// pubKeyBase64 := "MIIBMzCB7AYHKoZIzj0CATCB4AIBATAsBgcqhkjOPQEBAiEA/////////////////////////////////////v///C8wRAQgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHBEEEeb5mfvncu6xVoGKVzocLBwKb/NstzijZWfKBWxb4F5hIOtp3JqPEZV2k+/wOEQio/Re0SKaFVBmcR9CP+xDUuAIhAP////////////////////66rtzmr0igO7/SXozQNkFBAgEBA0IABM2EPsPppO7Ab5OY98WObiVOQ9Yowl9m+Y7ljC10u3FUzwT8YSKNIe5njZaT0VHC2J/MuJh87kY7IzJ7hqCxguY="
	// pubKeyBase64 := "BM2EPsPppO7Ab5OY98WObiVOQ9Yowl9m+Y7ljC10u3FUzwT8YSKNIe5njZaT0VHC2J/MuJh87kY7IzJ7hqCxguY="
	// pubKeyBase64 := "MIIBMzCB7AYHKoZIzj0CATCB4AIBATAsBgcqhkjOPQEBAiEA/////////////////////////////////////v///C8wRAQgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHBEEEeb5mfvncu6xVoGKVzocLBwKb/NstzijZWfKBWxb4F5hIOtp3JqPEZV2k+/wOEQio/Re0SKaFVBmcR9CP+xDUuAIhAP////////////////////66rtzmr0igO7/SXozQNkFBAgEBA0IABFyI0LFESiFXehUQOYu4qM8EqWnkfooNBlpBEvbV9VzL1Vl98AMJeXGunK8uPwxGBldVGwUFEXFVkOwGYOlMfu8="
	pubKeyBase64 := "MEUCIQCQ8XdhV0gIvz6OBQa5S0df0AZHpLgzmfJbx/8rd65GjgIgCZb50LANvUNyRjgY6q+33gGpomA/AGdWlKyaAdesowY="
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKeyBase64)
	if err != nil {
		fmt.Println("Error base64 decoding string: ", err)
		return
	}

	fmt.Printf("pubkey[%v]\n", hex.EncodeToString(pubKeyBytes))
	fmt.Println("***************************************************************************")

	pubKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		fmt.Printf("Unmarshal pubkey error: %+v ", err)
		return
	}

	// Get the address
	address := crypto.PubkeyToAddress(*pubKey)
	fmt.Printf("address[%d][%v]\n", len(address), address)
	fmt.Println("***************************************************************************")

	// 测试密钥格式
	// testKey, err := crypto.HexToECDSA("3e2312ef6bc29c4a2707a6d622b33b17f4945a65b4eadaf2f74a264dfa75aace")
	// if err != nil {
	// 	fmt.Println("PK HexToECDSA error: ", err)
	// 	return
	// }
	// fmt.Println("test pubkey base64: ", base64.StdEncoding.EncodeToString(crypto.FromECDSAPub(&testKey.PublicKey)))
	// fmt.Println("test pubkey: ", hex.EncodeToString(crypto.FromECDSAPub(&testKey.PublicKey)))
	// fmt.Println("test address: ", crypto.PubkeyToAddress(testKey.PublicKey))
}

func generateEthereumPrivateKeyByBip39Seed(seed string) (*ecdsa.PrivateKey, common.Address, error) {
	/*
		// 随机生成助记词组
		// entropy, err := bip39.NewEntropy(128)
		// if err != nil {
		//  log.Fatal(err)
		// }
		// mnemonic, _ := bip39.NewMnemonic(entropy)
		// //var mnemonic = "pepper hair process town say voyage exhibit over carry property follow define"
		// fmt.Println("mnemonic:", mnemonic)
		// 这里可以选择传入指定密码或者空字符串，不同密码生成的种子不同
		// seed := bip39.NewSeed(mnemonic, "")
	*/

	// 生成种子 (这里可以选择传入指定密码或者空字符串，不同密码生成的种子不同)
	// seed := bip39.NewSeed("seed", salt)

	seedBytes, err := base64.StdEncoding.DecodeString("gmVe/dZ0sX6Z5zFne2ILrT4B1jjJIqJgZhaRyKH3R98=")
	if err != nil {
		log.Fatal(err)
	}
	wallet, err := hdwallet.NewFromSeed(seedBytes)
	if err != nil {
		log.Fatal(err)
	}

	// 生成地址 (最后一位是同一个助记词的地址id，从0开始，相同助记词可以生产无限个地址)
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	// 解析私钥
	prvKey, err := wallet.PrivateKey(account)
	if err != nil {
		log.Fatal(err)
	}

	// address := account.Address.Hex()
	// privateKey, _ := wallet.PrivateKeyHex(account)
	// publicKey, _ := wallet.PublicKeyHex(account)

	// fmt.Println("address0:", address)      // id为0的钱包地址
	// fmt.Println("privateKey:", privateKey) // 私钥
	// fmt.Println("publicKey:", publicKey)   // 公钥

	// path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1") //生成id为1的钱包地址
	// account, err = wallet.Derive(path, false)
	// if err != nil {
	//  log.Fatal(err)
	// }

	fmt.Println("address1:", account.Address.Hex())

	return prvKey, account.Address, nil
}

// Example address creation for a fictitious company ComputerVoice Inc. where
// each department has their own wallet to manage
func generateEthereumPrivateKeyByBip32Seed(salt string) (*ecdsa.PrivateKey, common.Address, error) {
	// Generate a seed to determine all keys from.
	// This should be persisted, backed up, and secured
	// seed, err := bip32.NewSeed()
	// if err != nil {
	// 	log.Fatalln("Error generating seed:", err)
	// }

	seedBytes, err := base64.StdEncoding.DecodeString("gmVe/dZ0sX6Z5zFne2ILrT4B1jjJIqJgZhaRyKH3R98=")
	if err != nil {
		log.Fatal(err)
	}
	// Create master private key from seed
	// computerVoiceMasterKey, _ := bip32.NewMasterKey(seed)
	masterKey, _ := bip32.NewMasterKey(seedBytes)

	fmt.Println("pubkey: ", masterKey.PublicKey().B58Serialize())
	fmt.Println("pubkey: ", masterKey.PublicKey())
	fmt.Println("pubkey: ", hex.EncodeToString(masterKey.PublicKey().Key))

	// // Map departments to keys
	// // There is a very small chance a given child index is invalid
	// // If so your real program should handle this by skipping the index
	// departmentKeys := map[string]*bip32.Key{}
	// departmentKeys["Sales"], _ = computerVoiceMasterKey.NewChildKey(0)
	// departmentKeys["Marketing"], _ = computerVoiceMasterKey.NewChildKey(1)
	// departmentKeys["Engineering"], _ = computerVoiceMasterKey.NewChildKey(2)
	// departmentKeys["Customer Support"], _ = computerVoiceMasterKey.NewChildKey(3)

	// // Create public keys for record keeping, auditors, payroll, etc
	// departmentAuditKeys := map[string]*bip32.Key{}
	// departmentAuditKeys["Sales"] = departmentKeys["Sales"].PublicKey()
	// departmentAuditKeys["Marketing"] = departmentKeys["Marketing"].PublicKey()
	// departmentAuditKeys["Engineering"] = departmentKeys["Engineering"].PublicKey()
	// departmentAuditKeys["Customer Support"] = departmentKeys["Customer Support"].PublicKey()

	// // Print public keys
	// for department, pubKey := range departmentAuditKeys {
	// 	fmt.Println(department, pubKey)
	// }

	return nil, common.Address{}, nil
}

func parseBase64Seed() ([]byte, error) {
	seed, err := base64.StdEncoding.DecodeString("gmVe/dZ0sX6Z5zFne2ILrT4B1jjJIqJgZhaRyKH3R98=")
	if err != nil {
		return nil, err
	}

	fmt.Println("seed: ", hex.EncodeToString(seed))

	return seed, nil
}

// 发现一个好库：https://github.com/ygcool/go-hdwallet
func generateKeyFromTgSeed() (*ecdsa.PrivateKey, error) {
	seed, err := parseBase64Seed()
	if err != nil {
		return nil, err
	}

	// masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.TestNet3Params)
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	// 以太坊格式私钥
	ecPriKey, err := masterKey.ECPrivKey()
	if err != nil {
		return nil, err
	}

	// ecPubKey := ecPriKey.PubKey()

	// fmt.Println("seed base64: ", "gmVe/dZ0sX6Z5zFne2ILrT4B1jjJIqJgZhaRyKH3R98=")
	// fmt.Println("tsb pubkey base64:               ", "BFyI0LFESiFXehUQOYu4qM8EqWnkfooNBlpBEvbV9VzL1Vl98AMJeXGunK8uPwxGBldVGwUFEXFVkOwGYOlMfu8=")
	// fmt.Println("this pubkey uncompressed base64: ", base64.StdEncoding.EncodeToString(ecPubKey.SerializeUncompressed()))

	// ethAddress := crypto.PubkeyToAddress(*ecPubKey.ToECDSA())
	// fmt.Println("eth address: ", ethAddress)

	return ecPriKey.ToECDSA(), nil
}

func parseBase64DerData() ([]byte, error) {
	var data []byte

	// echo "MEUCIQCVNZwhUTqsONhGKp2TApBO++uPdzBlOWzGrH6pxvSFFwIgVmZS7Kd3vdLIL0/JszkWe3rDAtB3PXwAGia2DNdFu64=" | base64 -D | openssl asn1parse -inform DER
	derData, err := base64.StdEncoding.DecodeString("MEQCIF5cjDxP/oQ4mickDjRf8saNZ1oCTlLh5d0MyAjzjRrLAiBU0NaqpFEoXCGEltMJiARxV3dPlFj3/gbyANrJYJTb4Q==")
	if err != nil {
		return nil, err
	}
	fmt.Println("data hex: ", hex.EncodeToString(derData))

	rsSig, err := btcec.ParseDERSignature(derData, btcec.S256())
	if err != nil {
		fmt.Println("Error ParseDERSignature: ", err)
		return nil, err
	}
	fmt.Println("*************************************************************************************************************")
	fmt.Println("rsSig.R: ", hex.EncodeToString(rsSig.R.Bytes()))
	fmt.Println("rsSig.S: ", hex.EncodeToString(rsSig.S.Bytes()))
	fmt.Println("*************************************************************************************************************")

	// rest, err := asn1.Unmarshal(derData, &data)
	// if err != nil {
	// 	fmt.Println("asn1.Unmarshal err: ", err)
	// 	return nil, err
	// } else if len(rest) != 0 {
	// 	return nil, err
	// }
	// fmt.Println("unmarshal data: ", hex.EncodeToString(data))

	return data, nil
}
