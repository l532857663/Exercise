package main

import "fmt"

type Map struct {
	len int
	cap int
}

func main() {
	a := make(map[string]string, 10)
	b := make(map[string]string, 10)
	a["dsa"] = "dsa"

	b = a

	fmt.Printf("a: %v\nb: %v\n", a, b)
	fmt.Printf("a: %p\n", a)
	fmt.Printf("&a: %p\n", &a)
	fmt.Printf("b: %p\n", b)
	fmt.Printf("&b: %p\n", &b)
	b["dsa1"] = "dsa"
	b["dsa2"] = "dsa"
	b["dsa3"] = "dsa"
	b["dsa4"] = "dsa"
	b["dsa5"] = "dsa"
	b["dsa6"] = "dsa"
	b["dsa7"] = "dsa"
	b["dsa8"] = "dsa"
	b["dsa9"] = "dsa"
	b["dsa10"] = "dsa"
	b["dsa11"] = "dsa"
	b["dsa12"] = "dsa"
	b["dsa13"] = "dsa"
	fmt.Printf("a: %v\nb: %v\n", a, b)
	fmt.Printf("a: %p\n", a)
	fmt.Printf("&a: %p\n", &a)
	fmt.Printf("b: %p\n", b)
	fmt.Printf("&b: %p\n", &b)
}
