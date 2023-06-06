package service

// *************************  请求结果 ************************
type PageReq struct {
	PageIndex string `json:"pageIndex"   binding:"required"` // 页码
	PageSize  string `json:"pageSize"    binding:"required"` // 每页长度
}

// 查询余额信息
type GetAddressInfoReq struct {
	PageReq
	Symbol               string `json:"symbol"       binding:"required"` // 主链
	Address              string `json:"address"      binding:"required"` // 地址
	ProtocolType         string `json:"protocolType"`                    // 合约协议类型
	TokenContractAddress string `json:"tokenContractAddress"`            // 合约地址
}

type GetUnspentReq struct {
	Symbol  string `json:"symbol"       binding:"required"` // 主链
	Address string `json:"address"      binding:"required"` // 地址
}

// 查询区块交易
type GetTransferReq struct {
	PageReq
	Symbol       string `json:"symbol"       binding:"required"` // 主链
	Height       string `json:"height"`                          // 区块高度
	ProtocolType string `json:"protocolType"`                    // 合约协议类型
}
