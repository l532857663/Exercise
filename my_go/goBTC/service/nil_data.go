package service

import "goBTC/models"

var (
	NilTransferList = []models.TransferInfo{}
	NilTokenList    = []models.TokenInfo{}
	NilTransferResp = &models.GetTransferResp{
		ChainFullName:  "",
		ChainShortName: "",
		TotalPage:      "0",
		TransferList:   NilTransferList,
	}
)
