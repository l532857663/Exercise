package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func main() {
	var loopArr = []int{2, 10, 1000, 100000}
	for _, loop := range loopArr {
		s1 := "hello"
		s2 := "world"
		var start time.Time

		// +连接
		var s string
		start = time.Now()
		for i := 0; i < loop; i++ {
			s += s1 + s2
		}
		fmt.Println("+连接方式：", time.Since(start))
		fmt.Println("wch----- s: ", len(s))

		// buffer
		start = time.Now()
		var buf bytes.Buffer
		for i := 0; i < loop; i++ {
			buf.WriteString(s1)
			buf.WriteString(s2)
		}
		ss := buf.String()
		fmt.Println("buffer连接方式：", time.Since(start))
		fmt.Println("wch----- s: ", len(ss))

		// join
		a := []string{}
		for i := 0; i < loop; i++ {
			a = append(a, s1, s2)
		}
		start = time.Now()
		sss := strings.Join(a, "")
		fmt.Println("join连接方式：", time.Since(start))
		fmt.Println("wch----- s: ", len(sss))
		fmt.Println("")
		fmt.Println("")
	}
	return
}
