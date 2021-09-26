package main

import (
	"fmt"
)

func doStuff(value []string) {
	fmt.Printf("value=%v\n", value)

	value2 := value[:]
	value2 = append(value2, "b")
	fmt.Printf("value=%v, value2=%v\n", cap(value), cap(value2))
	fmt.Printf("&value=%p, &value2=%p\n", &value[0], &value2[0])
	fmt.Printf("&value2=%p\n", &value2[1])

	value2[0] = "z"
	fmt.Printf("value=%v, value2=%v\n", value, value2)
}

func sliceTest() {
	a := [...]int{1, 2, 3}
	b := a[:]
	func(b []int) {
		c := b
		c = append(c, 6)
		c[0] = 5
		return
	}(b)
	fmt.Printf("a %+v, b %+v\n", a, b)
	fmt.Printf("&a %+v, &b %+v\n", &a, &b)
	c := make([]string, 1, 3)
	c[0] = "1"
	for i := 0; i < 5; i++ {
		a = append(a, "w")
		fmt.Printf("a %+v, %v\n", a, cap(a))
	}
	fmt.Printf("c1 %+v, %v\n", c, cap(c))
	doStuff(c)
	fmt.Printf("c = %v\n", c)
	fmt.Printf("c = %p\n", &c[0])
	d := []string{"a"}
	doStuff(d)
	fmt.Printf("d = %v, d=%p\n", d, &d[0])
}

func main() {
	sliceTest()
	return
}
