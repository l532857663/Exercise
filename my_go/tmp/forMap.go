package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_s() map[string]*student {
	sMap := make(map[string]*student)

	stus := []student{
		{Name: "li", Age: 26},
		{Name: "zhao", Age: 21},
		{Name: "wang", Age: 22},
	}

	for _, stu := range stus {
		sMap[stu.Name] = &stu
	}
	return sMap
}

func main() {
	a := pase_s()
	for k, v := range a {
		fmt.Printf("k: %s, v: %v\n", k, v)
	}
}
