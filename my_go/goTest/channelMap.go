package main

import (
	"fmt"
	"sync"
	"time"
)

type TmpMap interface {
	Read(key string) string
	Write(key, value string)
}

type Obj struct {
	Data map[string]chan string
}

func (this *Obj) Read(key string) string {
	var value string
	var flag bool
	fmt.Printf("wch------- read key: %v\n", key)
	for !flag {
		fmt.Println("wch------ pong")
		// 是否有该Key
		ch, ok := this.Data[key]
		if !ok {
			time.Sleep(time.Second * 1)
			continue
		}
		// 取值
		select {
		case v, ok := <-ch:
			fmt.Printf("wch---- v: %v, f: %v\n", v, ok)
			flag = true
			value = v
			ch <- v
		default:
		}
		time.Sleep(time.Second * 1)
	}
	return value
}

func (this *Obj) Write(key, value string) {
	ch := make(chan string, 1)
	ch <- value
	this.Data[key] = ch
	return
}

func main() {
	a := Obj{
		Data: make(map[string]chan string),
	}
	key := "haha"
	wg := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(a Obj, key string, i int, wg *sync.WaitGroup) {
			fmt.Printf("wch----- i: %v\n", i)
			value := a.Read(key)
			fmt.Printf("wch----- v: %v\n", value)
			wg.Done()
		}(a, key, i, &wg)
	}
	time.Sleep(5 * time.Second)
	fmt.Printf("wch------ test a %+v\n", a)
	a.Write(key, "ha")
	wg.Wait()
	fmt.Printf("wch------ a %+v\n", a)
	return
}
