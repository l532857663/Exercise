package main

import (
	"fmt"
)

type ceshi func(string) string

func exec(name string, vs ...ceshi) string {
	var ch = make(chan string)
	defer close(ch)

	fn := func(i int) {
		ch <- vs[i](name)
	}

	for i, _ := range vs {
		go fn(i)
	}

	select {
	case c := <-ch:
		fmt.Println("ch: ", c)
	}

	return <-ch
}

func doing() {
	vs := []string{"1", "2", "3", "4"}
	var ch = make(chan string)
	defer close(ch)

	fn := func(i int) {
		ch <- vs[i]
	}

	/*
		for i, _ := range vs {
			go fn(i)
		}
	*/
	go fn(1)
	go fn(2)

	select {
	case c := <-ch:
		fmt.Println("ch: ", c)
	}

	return
}

func main() {
	fmt.Println("vim-go")

	doing()

	/*
		res := exec("111", func(n string) string {
			return n + "func1"
		}, func(n string) string {
			return n + "func2"
		}, func(n string) string {
			return n + "func3"
		}, func(n string) string {
			return n + "func4"
		})
	*/
}
