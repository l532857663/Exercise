package main

import "sync"

func main() {
	num := 100
	flag := *string
	allMap := sync.Map
	for i := 0; i <= num; i++ {
		go Worker(flag, i)
	}
}
