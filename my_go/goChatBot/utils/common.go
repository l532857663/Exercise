package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

func StringToMD5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func Base64ToByte(imageStr string) ([]byte, error) {
	// 解码 Base64 编码的字符串
	imgData, err := base64.StdEncoding.DecodeString(imageStr)
	if err != nil {
		fmt.Println("decode error:", err)
		return nil, err
	}
	return imgData, err
}

func DeleteStringBot(text, str string) string {
	// 检查消息文本中是否包含 str
	if text != "" && strings.Contains(text, str) {
		return strings.Replace(text, str, "", -1)
	}
	return text
}
