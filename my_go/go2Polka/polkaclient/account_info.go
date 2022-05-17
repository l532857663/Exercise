package polkaclient

import (
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/vedhavyas/go-subkey"
	"github.com/vedhavyas/go-subkey/sr25519"
)

// 使用公钥获取账户信息
func (n *Node) GetAccountInfoByPubKey(pubKey string) (*AccountInfo, error) {
	// 公钥转化
	b, err := types.HexDecodeString(pubKey)
	if err != nil {
		fmt.Printf("Hex decode string error: %+v\n", err)
		return nil, err
	}
	fmt.Printf("wch----- data:%+v\n", b)
	return n.GetAccountInfo(b)
}

// 使用种子获取账户信息
func (n *Node) GetAccountKeyPairFromSecret(seedOrPhrase string) (*signature.KeyringPair, *AccountInfo, error) {
	fromAddr, err := signature.KeyringPairFromSecret(seedOrPhrase, uint8(0))
	if err != nil {
		fmt.Printf("Keyring pair from secret: %+v\n", err)
		return nil, nil, err
	}
	// 查询账户信息
	accountInfo, err := n.GetAccountInfo(fromAddr.PublicKey)
	if err != nil {
		fmt.Printf("Get account info by pubKey error: %+v\n", err)
		return nil, nil, err
	}

	return &fromAddr, accountInfo, nil
}

// 使用地址结构获取账户地址
func (n *Node) GetAccountInfoByMultiAddress(multiAddr types.MultiAddress) (*AccountInfo, error) {
	// 地址公钥解析
	pubKey := multiAddr.AsID[:]
	// NOTE: 可能会有后续处理
	accountInfo, err := n.GetAccountInfo(pubKey)
	if err != nil {
		fmt.Printf("Get multiAddr account info error: %+v\n", err)
		return nil, err
	}
	return accountInfo, nil
}

func (n *Node) GetAccountInfo(pubKey []byte) (*AccountInfo, error) {
	api := n.Api
	// 创建交易所需的公共参数 -------------------
	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		fmt.Printf("Get metadata error: %+v\n", err)
		return nil, err
	}
	key, err := types.CreateStorageKey(meta, "System", "Account", pubKey)
	if err != nil {
		fmt.Printf("Create storage key error: %+v\n", err)
		return nil, err
	}
	// 查询最新哈希信息
	hash, err := api.RPC.Chain.GetBlockHashLatest()
	if err != nil {
		fmt.Printf("Get block hash latest error: %+v", err)
		return nil, err
	}
	// 查询最新的账户信息
	raw, err := n.Api.RPC.State.GetStorageRaw(key, hash)
	if err != nil {
		fmt.Printf("Get storage latest error: %+v\n", err)
		return nil, err
	}
	// NOTE: 处理回传结果，匹配依赖库的结构体
	var newRaw []byte
	newRaw = append(newRaw, (*raw)[0:12]...)
	newRaw = append(newRaw, (*raw)[16:]...)

	// 账户信息
	accountOnline := &types.AccountInfo{}
	err = types.DecodeFromBytes(newRaw, accountOnline)
	if err != nil {
		fmt.Printf("Decode from bytes for accountOnline error: %+v\n", err)
		return nil, err
	}

	addr, err := subkey.SS58Address(pubKey, 42)
	if err != nil {
		fmt.Printf("Get ss58 address error: %+v\n", addr)
		return nil, err
	}

	accountInfo := &AccountInfo{
		Address:    addr,
		PubKey:     pubKey,
		Nonce:      uint64(accountOnline.Nonce),
		OnlineInfo: accountOnline,
	}
	return accountInfo, nil
}

func GenerateAccount() (*AccountInfo, error) {
	scheme := sr25519.Scheme{}
	keyring, err := scheme.Generate()
	if err != nil {
		fmt.Printf("Genarate account error: %+v\n", err)
		return nil, err
	}
	fmt.Printf("wch------- keyting: %+v\n", keyring)
	pub := keyring.Public()
	seed := keyring.Seed()
	accountId := keyring.AccountID()
	addr, err := keyring.SS58Address(42)
	if err != nil {
		fmt.Printf("Show address error: %+v\n", err)
		return nil, err
	}
	fmt.Printf("wch-------- pub [%+v]\n seed [%+v]\n accountId [%+v]\n addr [%+v]\n", pub, seed, accountId, addr)
	fmt.Printf("wch-------- pub [%#x]\n seed [%#x]\n accountId [%#x]\n addr [%#x]\n", pub, seed, accountId, addr)

	accountInfo := &AccountInfo{}
	return accountInfo, nil
}
