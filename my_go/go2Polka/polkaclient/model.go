package polkaclient

import "github.com/centrifuge/go-substrate-rpc-client/v4/types"

type AccountInfo struct {
	Address    string             // 地址字符串
	PubKey     []byte             // 公钥信息
	Nonce      uint64             // 最新nonce值
	OnlineInfo *types.AccountInfo // 线上信息
}

type BlockHeader struct {
	Number            int64  // 块高
	BlockHash         string // 哈希
	ParentHash        string // 父哈希
	StateRoot         string // 状态根
	ExtrinsicsRoot    string // 交易根
	TransactionLength int    // 交易数量
}

type BlockTransaction struct {
	Version      byte   // 版本
	Nonce        uint64 // 交易nonce
	TransferType string // 交易类型
	FromAddress  string // 发送地址
	ToAddress    string // 目标地址
	Value        string // 金额
	AsSr25519    string // 签名数据
	AsEd25519    string // 签名数据
}

type CallArgs struct {
	Dest  types.MultiAddress // 目标账户地址
	Value types.UCompact     // 金额
}
