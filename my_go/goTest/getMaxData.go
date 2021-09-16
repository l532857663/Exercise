package main

import "fmt"

func test1(arr []int) (int, int) {
	// 数组长度小于2，返回
	arrLength := len(arr)
	if arrLength == 0 || arrLength < 2 {
		return -1, -1
	}
	max1 := arr[0]
	max2 := arr[0]
	// 循环判断
	for i := 1; i < arrLength; i++ {
		// 取最第二大值
		if arr[i] > max2 {
			max2 = arr[i]
		}
		// 取最大值
		if max2 >= max1 {
			max1, max2 = max2, max1
		}
	}
	return max1, max2
}

func test2(arr []int, index int, max1, max2 *int) {
	length := len(arr)
	if length == 0 || length < 2 {
		return
	}
	// 最后一个返回当前值
	if index < length-1 {
		test2(arr, index+1, max1, max2)
	}
	// 对比下一个返回大值
	max := arr[index]
	if max > *max2 {
		*max2 = max
	}
	if *max2 >= *max1 {
		*max1, *max2 = *max2, *max1
	}
	return
}

func main() {
	// 定义数组
	a := []int{9, 4, 3, 5, 6, 7, 25, 1, 0, 8, 2, 4, 24, 13, 10}
	var max1, max2 int
	// 循环算法
	max1, max2 = test1(a)
	// 输出最大的两个数据
	fmt.Printf("Max %v, %v\n", max1, max2)

	// 递归算法
	max1, max2 = a[0], a[0]
	test2(a, 0, &max1, &max2)
	// 输出最大的两个数据
	fmt.Printf("Max %v, %v\n", max1, max2)
	return
}
