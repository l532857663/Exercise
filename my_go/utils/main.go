package main

import (
	"fmt"
	"utils/date"
	"utils/sort"
)

func main() {
	fmt.Println("vim-go")
	// 时间格式化
	// test1()

	// 排序
	test2()

	return
}

func test1() {
	var dateObj date.DateProcessor
	res, err := dateObj.GetTimeByDate("2021-07-29 05:00:00")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("res: ", res)
	return
}

func test2() {
	a := []int{9, 4, 3, 5, 6, 7, 1, 0, 8, 2, 4, 24, 13, 10}
	fmt.Printf("wch------- a %+v\n", a)
	// sort.ShellSort(a)
	// sort.BubbleSort(a)
	sort.QuickSort(a)
	fmt.Printf("wch------- b %+v\n", a)
	return
}
