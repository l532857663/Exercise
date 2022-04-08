package main

import "fmt"

type A struct {
	Name string `json:"name"`
	B    struct {
		Address string `json:"address"`
	} `json:"b"`
}

func main() {
	fmt.Printf("Start\n")
	a := A{
		Name: "haha",
		B: struct {
			Address string `json:"address"`
		}{Address: "aaa"},
	}
	fmt.Printf("a: %+v\n", a)
	fmt.Printf("End\n")
}
