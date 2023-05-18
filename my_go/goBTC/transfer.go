package goBTC

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
)

func (c *BTCClient) GetTransactionByHash(hash string) (*btcutil.Tx, error) {
	h, err := chainhash.NewHashFromStr(hash)
	if err != nil {
		return nil, err
	}
	return c.Client.GetRawTransaction(h)
}

func (c *BTCClient) GetRawTransactionByHash(hash string) (*btcjson.TxRawResult, error) {
	h, err := chainhash.NewHashFromStr(hash)
	if err != nil {
		return nil, err
	}
	return c.Client.GetRawTransactionVerbose(h)
}

func (c *BTCClient) SignTx() (string, error) {
	hexPrivateKey := "15oDYaEjTohXrTsnzMZCdZsULqL7cEqJxykXm8zTgtN9CM5D7Hdc"
	wif, err := btcutil.DecodeWIF(hexPrivateKey)
	if err != nil {
		return "", fmt.Errorf("SignTx DecodeWIF fatal, " + err.Error())
	}
	if !wif.IsForNet(c.Params) {
		return "", fmt.Errorf("SignTx IsForNet fatal")
	}
	pubK := wif.PrivKey.PubKey()
	p1, _ := btcutil.NewAddressTaproot(schnorr.SerializePubKey(pubK), c.Params)
	data, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootKeyNoScript(pubK)), c.Params)
	if err != nil {
		return "", err
	}
	fmt.Printf("wch----- pubk: %+v, data: %+v\n", p1, data)
	return "", nil
}
