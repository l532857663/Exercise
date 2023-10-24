package main

import (
	"fmt"
	"math/big"
)

type A big.Int

func (a A) DoHA() {
	fmt.Println("haha", a)
}

func (a *A) DoHAHA() {
	fmt.Println("haha1", a)
}

func testIt(a A) {
	a.DoHAHA()
}

func main() {
	i := new(big.Int).SetInt64(10)
	a := A(*i)
	a.DoHA()

	testIt(a)
}
