package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func longRunningCalculation(timeCost int) chan string {
	result := make(chan string)
	go func() {
		time.Sleep(time.Second * (time.Duration(timeCost)))
		fmt.Printf("wch-------- go2\n")
		result <- "Done"
	}()
	fmt.Printf("wch-------- go1\n")
	return result
}

func jobWithTimeoutHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 解析请求数据
	err := r.ParseForm()
	if err != nil {
		log.Println("wch----- form error %v\n", err)
		return
	}
	fmt.Printf("wch-------- ctx req:%+v\n", r.Form)
	if len(r.Form) == 0 {
		log.Println("wch----- form nil\n")
		return
	}
	fmt.Fprintf(w, "Hello golang http![%s]", r.Form["test"][0])

	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		return
	case <-time.After(3 * time.Second):
		fmt.Printf("wch-------- timeout\n")
		return
		/*
			case result := <-longRunningCalculation(5):
				fmt.Printf("wch-------- long\n")
				io.WriteString(w, result)
		*/
	}
	fmt.Printf("wch-------- end\n")
	return
}

func main() {
	http.HandleFunc("/", jobWithTimeoutHandler)
	http.ListenAndServe(":2333", nil)
}
