package inscribe

const (
	// 表名
	InscribeInfoName = "inscribe_info"
	OrdTokenName     = "ord_tokens"
)

type OrdAction string

const (
	// 铭文信息常量
	StateIsFalse = "0"
	StateIsTrue  = "1"
	// 铭文操作类型
	ActionForDeploy   OrdAction = "deploy"
	ActionForMint     OrdAction = "mint"
	ActionForTransfer OrdAction = "transfer"
	ActionForSend     OrdAction = "send"
	ActionForReceive  OrdAction = "receive"
	// 铭文类型
	InscribeTypeNFT   = "NFT"
	InscribeTypeBRC20 = "BRC20"
)
