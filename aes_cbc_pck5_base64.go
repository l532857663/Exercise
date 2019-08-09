package cryptoutil

import (
	"encoding/base64"

	"github.com/chentaihan/aesCbc"
)

/**
 * AES加密 ，CBC模式，PCK5填充数据，base64加密
 */

func EncryptColumn(originDataString, encryptKeyString string) (string, error) {
	// string2[]byte
	originData := []byte(originDataString)
	encryptKey := []byte(encryptKeyString)
	// Aes加密
	iv := "12345678901234561234567890123456"
	aesCipher := aesCbc.NewAesCipher(encryptKey, []byte(iv))
	encrData := aesCipher.Encrypt([]byte(originData))

	encodeString := base64.StdEncoding.EncodeToString(encrData)
	return encodeString, nil
}

/*
func DecryptColumn(encryptData, encryptKeyString string) (string, error) {
	// string2[]byte
	encryptKey := []byte(encryptKeyString)
	// Aes解密
	decodeString, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		logger.Errorf("base64解密失败:%s", err.Error())
		return "", err
	}
	originData, err := AesCBCDecrypt(decodeString, encryptKey)
	if err != nil {
		logger.Errorf("aes decrypt error:%s", err.Error())
		return "", err
	}
	return string(originData), nil
}
*/
