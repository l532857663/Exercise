package main

import (
	"fmt"
	"runtime"
	"sync"
)

func process(id int, wg *sync.WaitGroup) {
	fmt.Printf("id: %d\n", id)
	// flag := true
	// for flag {
	// 	time.Sleep(10 * time.Second)
	// 	flag = false
	// }
	wg.Done()
	return
}

func main() {
	fmt.Println("vim-go")

	r := runtime.GOMAXPROCS(1)
	fmt.Printf("wch %v\n", r)
	var wg sync.WaitGroup
	a := 5
	wg.Add(a)
	for i := 0; i < a; i++ {
		go process(i, &wg)
	}
	wg.Wait()
	return
}
