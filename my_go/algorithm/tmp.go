package main

import "fmt"

var A = [][]string{{"张三", "语文", "80"},
	{"张三", "数学", "81"},
	{"张三", "英语", "82"},
	{"李四", "语文", "83"},
	{"李四", "数学", "84"},
	{"李四", "英语", "78"},
	{"王五", "语文", "90"},
	{"王五", "数学", "91"},
	{"王五", "英语", "92"}}

type B struct {
	Name  string
	Score int
}

func main() {
	b := make(map[string][]map[string]string)
	/*
		for _, a := range A {
			c := make(map[string]string)
			c[a[1]] = a[2]
			b[a[0]] = append(b[a[0]], c)
		}

		for name, class := range b {
			fmt.Printf("%s:\t", name)
			// score map[string]string
			for _, scoreMap := range class {
				for _, score := range scoreMap {
					fmt.Printf("%s\t", score)
				}
			}
			fmt.Printf("\n")
		}
	*/
	a := new(map[string]int)
	(*a)["key1"] = 11

	// fmt.Println(b)
	fmt.Println(b["ceshi"])
	fmt.Println(a)
}
