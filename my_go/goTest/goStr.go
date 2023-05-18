package main

import (
	"encoding/hex"
	"fmt"
)

// 字符串和16进制的转换
func ChangeHexAndString() {
	// 将字符串转换为16进制数据
	str := "Hello, world!"
	hexData := hex.EncodeToString([]byte(str))
	fmt.Println("String to Hex:", hexData)

	// 将16进制数据转换为字符串
	hexStr := "0063036f7264010118746578742f706c61696e3b636861727365743d7574662d38"
	data, _ := hex.DecodeString(hexStr)
	fmt.Println("Hex to String:", string(data))
}

// 十六进制字符串小端序转数字
func HexToInt() {
	hexString := "0x0802"
	bytes, _ := hex.DecodeString(hexString[2:]) // 将十六进制字符串转换为字节数组
	var num uint32
	for i := len(bytes) - 1; i >= 0; i-- {
		num = num<<8 + uint32(bytes[i]) // 将字节数组转换为数字
	}
	fmt.Println(num)
}

func main() {
	// ChangeHexAndString()
	HexToInt()
}
