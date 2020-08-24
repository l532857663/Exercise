package main

import (
	"fmt"
	"time"
)

var A = []int{5, 3, 7, 6, 4, 1, 0, 2, 9, 10, 8}

func QuickSort(data []int, left, right int) {
	temp := data[left]
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && data[j] >= temp {
			j--
		}
		if j >= p {
			data[p], p = data[j], j
		}
		for i <= p && data[i] <= temp {
			i++
		}
		if i <= p {
			data[p], p = data[i], i
		}
	}
	data[p] = temp
	if p-left > 1 {
		QuickSort(data, left, p-1)
	}
	if right-p > 1 {
		QuickSort(data, p+1, right)
	}
}

func main() {
	fmt.Println("vim-go")
	// 生成一个随机数数组
	// str := common.GetRandomString(10, 0)
	// fmt.Println(strings.Split(str, ""))
	fmt.Println(A)
	fmt.Printf("start: %s\n", time.Now().String())
	QuickSort(A, 0, len(A)-1)
	fmt.Printf("end  : %s\n", time.Now().String())
	fmt.Println(A)
}
