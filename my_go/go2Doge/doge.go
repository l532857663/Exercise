package main

import (
	"fmt"

	"github.com/qhxcWallet/doged/btcec/v2"
	"github.com/qhxcWallet/doged/chaincfg"
)

func main() {
	NewAddress(nil)
}

// 地址生成
func NewAddress(netConfig *chaincfg.Params) (*blocktype.Account, error) {
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		return nil, err
	}
	fmt.Printf("wch------ privateKey: %+v\n", privateKey)
	return nil, nil
	// wif, err := btcutil.NewWIF(privateKey, netConfig, true)
	// if err != nil {
	//     return nil, err
	// }
	// addr, err := btcutil.NewAddressPubKey(wif.SerializePubKey(), netConfig)
	// if err != nil {
	//     return nil, err
	// }
	// acct := &blocktype.Account{
	//     PrivateKey: wif.String(),
	//     Address:    addr.EncodeAddress(),
	// }
	// return acct, nil
}
