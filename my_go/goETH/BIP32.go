package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
)

// Example address creation for a fictitious company ComputerVoice Inc. where
// each department has their own wallet to manage
func main() {
	// Generate a seed to determine all keys from.
	// This should be persisted, backed up, and secured
	// seed, err := bip32.NewSeed()
	// if err != nil {
	// 	log.Fatalln("Error generating seed:", err)
	// }
	seed, _ := hex.DecodeString("7b99603952e4944ff478627c66a83dee63554d412cbd592de37268dbaf8cf0ccee1506ee43f3e7b9ca6800f7ea86c4845cda1b98b841400286f363bb4031641c")
	fmt.Printf("wch-------- seed: %d, %s, %x\n", len(seed), base64.StdEncoding.EncodeToString(seed), seed)

	// Create master private key from seed
	computerVoiceMasterKey, _ := bip32.NewMasterKey(seed)
	priB58 := computerVoiceMasterKey.B58Serialize()
	fmt.Printf("wch----- m depth: %+v\n", computerVoiceMasterKey.Depth)
	fmt.Printf("master is prikey: %v, key: %+v\n", computerVoiceMasterKey.IsPrivate, priB58)
	GetEthAddr(priB58)

	// Map departments to keys
	// There is a very small chance a given child index is invalid
	// If so your real program should handle this by skipping the index
	departmentKeys := map[string]*bip32.Key{}
	departmentKeys["Sales"], _ = computerVoiceMasterKey.NewChildKey(0)
	GetEthAddr(departmentKeys["Sales"].B58Serialize())
	departmentKeys["Marketing"], _ = computerVoiceMasterKey.NewChildKey(1)
	departmentKeys["Engineering"], _ = computerVoiceMasterKey.NewChildKey(2)
	departmentKeys["Customer Support"], _ = computerVoiceMasterKey.NewChildKey(3)

	// Create public keys for record keeping, auditors, payroll, etc
	departmentAuditKeys := map[string]*bip32.Key{}
	departmentAuditKeys["Sales"] = departmentKeys["Sales"].PublicKey()
	departmentAuditKeys["Marketing"] = departmentKeys["Marketing"].PublicKey()
	departmentAuditKeys["Engineering"] = departmentKeys["Engineering"].PublicKey()
	departmentAuditKeys["Customer Support"] = departmentKeys["Customer Support"].PublicKey()

	// // Print public keys
	// for department, pubKey := range departmentAuditKeys {
	// 	fmt.Println(department, pubKey)
	// }
	for name, key := range departmentKeys {
		p58 := key.B58Serialize()
		fmt.Printf("name: %+v, depth: %+v\n", name, key.Depth)
		GetEthAddr(p58)
	}
}

func GetEthAddr(priB58 string) {
	key, _ := hdkeychain.NewKeyFromString(priB58)
	// fmt.Printf("priB58: %+v, priKey: %+v\n", priB58, key.String())
	priKey, _ := key.ECPrivKey()
	ethAdd := crypto.PubkeyToAddress(priKey.ToECDSA().PublicKey).Hex()
	fmt.Printf("wch---- eth: %+v\n", ethAdd)
}
