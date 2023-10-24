package main

import "fmt"

type StructInterface interface {
	Name()
	Say()
}

type StructBase struct {
}

func (s *StructBase) Name(f interface{}) {
	fmt.Println("Name base")
	fmt.Println(s)
	f()
}

func (s *StructBase) Say() {
	fmt.Println("aaaaa")
}

type DD struct {
	StructBase
}

func (s *DD) Say() {
	fmt.Println("bbbbb")
}

func main() {
	var base StructInterface
	base = new(DD)
	fmt.Println(base)
	base.Name(base.Say())
}
