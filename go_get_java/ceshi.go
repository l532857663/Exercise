package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("Start")
	key := []byte("encryptedKey1234encryptedKey1234")
	keyBase64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println("key base64:", keyBase64, len(keyBase64))
	keyString := "6Ie5hEWp6VsaM5NYD0i/Kw=="
	keyBytes, err := base64.StdEncoding.DecodeString(keyString)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println("key bytes:", keyBytes, len(keyBytes), string(keyBytes))
	fmt.Println("End")
}
