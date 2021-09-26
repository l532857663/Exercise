package main

import "fmt"

func test(a [1]int) {
	c := a
	c[0] = 3
	fmt.Println("1-1", a, c)
}

func test1(a []int) {
	c := a
	c = append(c, 1)
	fmt.Println("1-2", a, c)
}

func main() {
	a := new([1]int)
	test(*a)
	fmt.Println("2", a)

	b := new([]int)
	test1(*b)
	fmt.Println("3", b)
	var c []int
	test1(c)
	fmt.Println("4", c)
	d := []int{}
	test1(d)
	fmt.Println("5", d)
}
