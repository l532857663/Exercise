package main

import "fmt"

var AMap map[string][]*A
var BMap map[string][]*A
var CMap map[string]*A

type A struct {
	Name string
	Age  int
}

func initMap() {
	a := make(map[string][]*A)
	tmp := &A{
		Name: "aaa",
		Age:  4,
	}
	tmp1 := &A{
		Name: "bbb",
		Age:  13,
	}
	a["a"] = append(a["a"], tmp)
	a["a"] = append(a["a"], tmp1)
	AMap = a
	CMap = make(map[string]*A)
	CMap["b"] = tmp
}

func showTime(data map[string][]*A, name string) {
	fmt.Printf("name: %+v\n", name)
	for _, v := range data["a"] {
		fmt.Printf("wch---- v: %+v\n", v)
	}
}

func pointUsed() {
	showTime(AMap, "test1")
	showTime(BMap, "test1-1")
	BMap = make(map[string][]*A)
	for i, v := range AMap["a"] {
		v1 := *v
		v1.Name = fmt.Sprintf("c%d", i)
		v1.Age += i
		BMap["a"] = append(BMap["a"], &v1)
	}
	showTime(AMap, "test2")
	showTime(BMap, "test2-1")
}

func showTime1(data map[string]*A, name string) {
	fmt.Printf("wch==== name: %+v\n CMap: %+v\n", name, CMap["b"])
}

func pointUsed1() {
	showTime1(CMap, "test1")
	c := CMap["b"]
	c.Name = "c1"
	showTime1(CMap, "test2")
	d := *(CMap["b"])
	d.Name = "c3"
	showTime1(CMap, "test3")
}

func main() {
	initMap()
	pointUsed()
	pointUsed1()
}
