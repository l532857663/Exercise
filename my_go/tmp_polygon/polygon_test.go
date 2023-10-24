package polygon

import (
	"common/salt"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sign/config"
	"testing"
	utilsCrypto "utils/utils/crypto"
	"utils/utils/kms"

	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

func TestCreateAccounts(t *testing.T) {
	polygonAddress, err := newEthAddr()
	if err != nil {
		config.LOG.Error("new btc address err", zap.Any("err", err))
		return
	}
	fmt.Println(polygonAddress.Address, polygonAddress.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncryptByKms(t *testing.T) {
	priKey := "4ec73a187419116ee24245854762f3ed189f46efd01e2a8f5d72d954b24e86dd"
	AccessKeyID := "3+HARloBycmWttvZv5HpF7Y872BEatI2GuaVsoB2Suw="
	SecretAccessKey := "st9s0dUHsox6nswb5tl7pI3XhHCXzEPPOyb/67fLMsP3AMh72CkIW6yitQIUHyMb"
	KMS_ACCESSKEY_PWD := "%5&b_ZbB#>NH"
	//AesIv := "g&tL%@WJkIzEg#hx"
	awsConfig := kms.AwsConfig{
		EndPoint: "kms.ap-east-1.amazonaws.com",
		Region:   "ap-east-1",
		KeyId:    "arn:aws:kms:ap-east-1:550012843142:key/5e454ce5-5e34-4e42-be83-48c8edab2718",
	}
	awsConfig.AccessKeyID = utilsCrypto.AesHashDecryptTool(AccessKeyID, KMS_ACCESSKEY_PWD)
	awsConfig.SecretAccessKey = utilsCrypto.AesHashDecryptTool(SecretAccessKey, KMS_ACCESSKEY_PWD)
	kmsClient, err := kms.GetKMSClient(awsConfig)
	if err != nil {
		return
	}

	//dk, err := kmsClient.GenerateDataKey()
	//if err != nil {
	//	return
	//}
	kmsKey, _ := base64.StdEncoding.DecodeString("AQIDAHgfn2tIqQP/X7gHEzgVxHRrNxBnMFFUiABMgjSD0CMpBwGPuGNUE5bDSyVJtQqUg+VwAAAAfjB8BgkqhkiG9w0BBwagbzBtAgEAMGgGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMXTB6XwHAPh9a9qB6AgEQgDv2xehgaGMyWsWOD2hJDO/fJ/rDirXKtpptrxUzkjSvJGDGRNVFMD5jCam6cHa7UlAiW8sAufRbo3fVJg==")
	Plaintext, _ := kmsClient.Decrypt(kmsKey)
	key := string(Plaintext) + salt.GetGenerateSalt()
	encryptPriKey := utilsCrypto.AesHashEncodeTool(priKey, key)

	privateKey := utilsCrypto.AesHashDecryptTool("lSCc9QWv7xygS0zt/RVpQnUrOI5UWzD/lUfzFVkIAxjjma2aMA+KrJAegRgYlZmSEfqIknUgKxcjbN/LAp8L2A==", key)
	fmt.Println("-----", encryptPriKey, privateKey)
	//fmt.Println(base64.StdEncoding.EncodeToString(dk.CiphertextBlob))
}

func TestDecryptPrikey(t *testing.T) {
	targetPriKey := "cVagL87yWdzMQNmHSgPsBzjp7ohvQ2Ybxpx9hzYmqjDfkinetEb8" // todo
	encryptPriKey := "UWkg8/+QSTIwZto4kr1hjrlAOo3ci2bVdOWMu98QHKdcLVJvlrd7cVwxijAhkhqV7pcKQJkTyABjR2pE/zeVVlpXpIDRCxsu/+kt4NzIBJg="
	encryptKmsKey := "AQIDAHgfn2tIqQP/X7gHEzgVxHRrNxBnMFFUiABMgjSD0CMpBwFlxUgfSDSf/4qr25mqvlh2AAAAfjB8BgkqhkiG9w0BBwagbzBtAgEAMGgGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMTKy8b5W2OZ62ABr1AgEQgDuVCOqgtLV19NMsz0qWC1H1kKN9Ef+FA+s0OQ/xYx/CN4mAaqU/4faQy7gnIe9B9s3FouZCsjykvN4JEw=="
	AccessKeyID := "2e6WqC6dKqWJ6VnDZ7U4DNaiZ0AijNrGDnEKC0u8Irw="
	SecretAccessKey := "PxM9dBJaad56VOt3lyrwD/rfo6b9YA7g+xd13zYr6MLXaYyn1MO4PojhpDpps11a"
	KMS_ACCESSKEY_PWD := "abcd"
	awsConfig := kms.AwsConfig{
		EndPoint: "kms.ap-east-1.amazonaws.com",
		Region:   "ap-east-1",
		KeyId:    "arn:aws:kms:ap-east-1:550012843142:key/6adee5ef-a003-4dae-adfb-b3be584eb322",
	}
	awsConfig.AccessKeyID = utilsCrypto.AesHashDecryptTool(AccessKeyID, KMS_ACCESSKEY_PWD)
	awsConfig.SecretAccessKey = utilsCrypto.AesHashDecryptTool(SecretAccessKey, KMS_ACCESSKEY_PWD)
	kmsClient, err := kms.GetKMSClient(awsConfig)
	if err != nil {
		return
	}
	kmsKey, _ := base64.StdEncoding.DecodeString(encryptKmsKey)
	kmsSecret, err := kmsClient.Decrypt(kmsKey)
	if err != nil {
		return
	}
	key := string(kmsSecret) + salt.GetGenerateSalt()
	privateKey := utilsCrypto.AesHashDecryptTool(encryptPriKey, key)
	if err != nil {
		return
	}

	fmt.Printf("targetPriKey: %s, priKey: %s", targetPriKey, privateKey)
}

//func TestSign(t *testing.T) {
//	conf1 := &conf.Node{
//		Ip:       "192.168.10.48",
//		Port:     8545,
//		User:     "",
//		Password: "",
//		Net:      "",
//	}
//	polygon, err := client.NewEthClient(conf1)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	//n,_:=client.GetNonce("0x3ee3a47e69d39c22d2a376ae6ccb718c56f8b10a")
//	//fmt.Println(n)
//	//
//	//b,_:=client.GetBalance("0x3ee3a47e69d39c22d2a376ae6ccb718c56f8b10a","latest")
//	//fmt.Println(b)
//
//	address := "0xb67AED487525AC99C688b6f11Eeb58F84aecb09b"
//	//privateKey := "71e8169cf4de54940e1c027c1907a308ea10ef93c64bb64cf661b9d375d080db"
//	nonce, _ := polygon.GetNonce(address)
//	var i uint64 = 0
//	for i = 0; i < 3; i++ {
//
//		gasePrice := polygon.SuggestGasPrice()
//		fmt.Println(nonce+i, "nonce")
//
//		var input = []byte{}
//		tx, err := polygon.GenTransaction("0x169d9846e56baebafbe5c039c2355765560b89d8", big.NewInt(100000000000000000), nonce+i, gasePrice, client.GasLimit, input)
//
//		params := model.RawDataReq{
//			Symbol: "ETH",
//			Sub_symbol: "ETH",
//			From: address,
//			Raw: tx,
//		}
//		signedTx, err := Sign(params)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println(signedTx)
//	}
//}

func TestLegacyTxConvert(t *testing.T) {
	//to := common.BytesToAddress([]byte("0x0c53a7908d2e5ba58e75291a306d19d8cc3a6650"))
	//var Ltx interface{}
	//Ltx = types.LegacyTx{
	//	Nonce: 16,
	//	GasPrice: big.NewInt(1000000000),
	//	Gas:23000,
	//	To: &to,
	//	Value: big.NewInt(200000000000000000),
	//	Data: []byte(""),
	//	V:nil,
	//	R:nil,
	//	S:nil,
	//}

	tt := fmt.Sprintf("{\"Nonce\":16,\"GasPrice\":1000000000,\"Gas\":23000,\"To\":\"0x0c53a7908d2e5ba58e75291a306d19d8cc3a6650\",\"Value\":200000000000000000,\"Data\":\"\",\"V\":null,\"R\":null,\"S\":null}")
	var tx types.LegacyTx
	if json.Unmarshal([]byte(tt), &tx) == nil {
		fmt.Println("ok", tx)
	} else {
		fmt.Println("err")
	}
}

func TestVerifyAddressSign(t *testing.T) {
	address := "0xbd46f927799041a1f4e637a9f1ea6b28ec6dbd04"
	encryptAddrSignPriKey := "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDV3dJQkFBS0JnUUNlaHdDeU01aUdZckxwS2lqTjBrWk5PbkJnWUxtUnVDOEZtNWVQcm1WSEp0RkdNemgxCmp2dnNPU2ZUS1V6ZWphUFdRNWZpZTZIenk1dW1Qa3hLd0lIZ0YxRDFoUGZZZUxxMWhQM1owdTdTY3JReVJ5Y3UKVEJ1cC9EOHNLZHVnTXdkb3NROHlIU2I0dHd0bDdRaW5lZEVVc1dmVFEzMEFmM1FnekJrTERpVG9kUUlEQVFBQgpBb0dBRGxzdGZmWGNOemRTK25pZDcxMitqaG5mdVdxcEE0QkppZGw0VlVPMjJrV3lxQWZWY2hmN3luMjJsSkhsCjZ6V3FnNm5sWkZaTDZWY2tCbDhYNjZFWllieGNHU1RXY3drdnVjbWs3aUtodm0wK002bW9GVHpEZFB3NXd5N08KMGRvSEJzcXF1RXdMV0xUWCtQUW0zU0pRTFhBbGVBRjlwdW9DL0J3QWRTMEw0VUVDUVFETm00UE1ORW8vTkRpUgpUMnZ4QzBQaThyclNGNWlhMXdGT1B6RjN1ZFNzK2lMaWJXaFQvNlIxQWZzRFA4V0VpcnZBT3gwSWJaYzNTZHJ2CnNoZm1rQmtGQWtFQXhXR0c1enlZU21hb0J0aFNZYjQrVlQwdHltUUxxNkYrZHhudHBRL3p4aktCVFZaa1FqUHYKWEFSc1Bya3JhRmU5dTI4TG1ETi9OVnp6TjRVTXRoZnNzUUpBUHVvMWZFa0w1eWM5b0FsamlGRGdKeFFMWXJwaApzZy9Va0hMNTJoNzlHeWszZjMzbkRMMFBQOWFwVHFjMjg0WFlTY3hNNkFWUTNsUTFNRitZdks2ZldRSkFGbWtSCldSWGZNS3RoTG8zSEpNUGw3ZVdwV2s1cnFNd0lRTnVYeU9MN3lhZ1lXRUNUMTVSdisrR2dyS3AzakR4U0ZZTHgKTFJIRHdycFAveURESmJXbnNRSkFHSkVCS1hhM0NkeUlsVElQZnNhK2ZINEdZR002YmxRUllUR1ZJNmMwWG5TVQp6UkFwc0ZoSWxMZExPOXJES2Y1b2FkVUNyVGs5dUNTVHIzOGV5VzhWc2c9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="
	encryptAddrSignPubKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDZWh3Q3lNNWlHWXJMcEtpak4wa1pOT25CZwpZTG1SdUM4Rm01ZVBybVZISnRGR016aDFqdnZzT1NmVEtVemVqYVBXUTVmaWU2SHp5NXVtUGt4S3dJSGdGMUQxCmhQZlllTHExaFAzWjB1N1NjclF5UnljdVRCdXAvRDhzS2R1Z013ZG9zUTh5SFNiNHR3dGw3UWluZWRFVXNXZlQKUTMwQWYzUWd6QmtMRGlUb2RRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="

	addrSignPriKey, _ := base64.StdEncoding.DecodeString(encryptAddrSignPriKey)
	addrSignPubKey, _ := base64.StdEncoding.DecodeString(encryptAddrSignPubKey)
	addrSign := utilsCrypto.RsaSignWithSha256([]byte(address), addrSignPriKey)
	err := utilsCrypto.RsaVerySignWithSha256([]byte(address), addrSign, addrSignPubKey)

	println(address, addrSign, err == nil)
}
