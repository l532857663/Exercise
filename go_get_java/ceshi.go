package main

import(
	"fmt"
	"encoding/base64"
)

func main() {
	fmt.Println("Start")
	key := "pMBAusbsM3xHwI9s/5nzapcH3E2fkRI+OfcdF9L1jD0="
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println("key bytes:", keyBytes, len(keyBytes))
	fmt.Println("End")
}
