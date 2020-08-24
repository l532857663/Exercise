package main

import "fmt"

type ceshi func(string) string

func exec(name string, vs ...ceshi) string {
	// fmt.Println("name", name)
	var ch = make(chan string)
	defer close(ch)

	/*
				fn := func(i int) {
					fmt.Println(i)
					ch <- vs[i](name)
				}

			for i, _ := range vs {
				go fn(i)
			}
		return <-ch
	*/
	for i := 0; i < 4; i++ {
		go func() {
			fmt.Println(i)
			// ch <- vs[i](name)
		}()
	}

	return vs[3](name)
}

func main() {
	fmt.Println("vim-go")
	res := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})

	fmt.Println("res", res)
}
