package polkaclient

import (
	"fmt"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/config"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type Node struct {
	Api *gsrpc.SubstrateAPI
}

func NewNode(nodeUrl string) (*Node, error) {
	if nodeUrl == "" {
		nodeUrl = config.Default().RPCURL
	}
	api, err := gsrpc.NewSubstrateAPI(nodeUrl)
	if err != nil {
		fmt.Printf("New substrate api error: %+v\n", err)
		return nil, err
	}
	node := &Node{
		Api: api,
	}
	return node, nil
}

// 查询连接节点信息
func (n *Node) GetConnectedInfo() error {
	api := n.Api
	chain, err := api.RPC.System.Chain()
	if err != nil {
		fmt.Printf("Connected node chain error: %+v\n", err)
		return err
	}
	nodeName, err := api.RPC.System.Name()
	if err != nil {
		fmt.Printf("Get node name error: %+v\n", err)
		return err
	}
	nodeVersion, err := api.RPC.System.Version()
	if err != nil {
		fmt.Printf("Get node version error: %+v\n", err)
		return err
	}

	fmt.Printf("You are connected to chain [%v] using [%v] version [%v]\n", chain, nodeName, nodeVersion)
	return nil
}

// 获取最新块Hash
func (n *Node) GetBlockHashLatest() (string, error) {
	api := n.Api
	hash, err := api.RPC.Chain.GetBlockHashLatest()
	if err != nil {
		return "", err
	}
	return hash.Hex(), nil
}

// 通过块高获取块数据
func (n *Node) GetBlockByBlocknum(num uint64) error {
	api := n.Api
	hash, err := api.RPC.Chain.GetBlockHash(num)
	if err != nil {
		fmt.Printf("Get block hash by number error: %+v\n", err)
		return err
	}
	block, err := api.RPC.Chain.GetBlock(hash)
	if err != nil {
		fmt.Printf("Get block by hash error: %+v\n", err)
		return err
	}
	n.ShowBlockInfo(hash.Hex(), block)
	return nil
}

// 通过HASH获取块数据
func (n *Node) GetBlockByHash(hash string) error {
	api := n.Api
	theHash, err := types.NewHashFromHexString(hash)
	if err != nil {
		fmt.Printf("The hash is Invalid: %+v\n", err)
		return err
	}
	block, err := api.RPC.Chain.GetBlock(theHash)
	if err != nil {
		fmt.Printf("Get block by hash error: %+v\n", err)
		return err
	}
	n.ShowBlockInfo(theHash.Hex(), block)
	return nil
}

// 查询地址余额
func (n *Node) GetBalanceFromPubKey(pubKey string) error {
	accountInfo, err := n.GetAccountInfoByPubKey(pubKey)
	if err != nil {
		fmt.Printf("Get account info by pubKey error: %+v\n", err)
		return err
	}

	previous := accountInfo.OnlineInfo.Data.Free
	fmt.Printf("%v has a balance of %v\n", pubKey, previous)
	return nil
}

// 查询地址nonce值
func (n *Node) GetNonceFromPubKey(pubKey string) error {
	accountInfo, err := n.GetAccountInfoByPubKey(pubKey)
	if err != nil {
		fmt.Printf("Get account info by pubKey error: %+v\n", err)
		return err
	}

	nonce := accountInfo.Nonce
	fmt.Printf("%v has a balance of %v\n", pubKey, nonce)
	return nil
}
