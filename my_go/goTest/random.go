package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

var one = new(big.Int).SetInt64(1)

func main() {
	// 获取随机数
	c := crypto.S256()
	params := c.Params()
	b := make([]byte, params.BitSize/8+8)
	fmt.Printf("wch----- b: %+v\n", b)
	fmt.Printf("wch----- r: %+v\n", rand.Reader)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return
	}

	fmt.Printf("wch----- b: %+v\n", b)
	k := new(big.Int).SetBytes(b)
	n := new(big.Int).Sub(params.N, one)
	fmt.Printf("wch------ k: %+v, n: %+v\n", k, n)

	/*
		str := "0x" + "ab12AB"
		data, _ := strconv.ParseInt(str, 0, 64)
		fmt.Printf("wch---- data: %d\n", data)
		// 种子随机
		s1 := rand.NewSource(data)
		r1 := rand.New(s1)

		s2 := rand.NewSource(data)
		r2 := rand.New(s2)

		fmt.Printf("wch------ r1: %v, %v, %v\n", r1.Intn(100), r1.Intn(100), r1.Intn(100))
		fmt.Printf("wch------ r2: %v, %v, %v\n", r2.Intn(100), r2.Intn(100), r2.Intn(100))
	*/
}
