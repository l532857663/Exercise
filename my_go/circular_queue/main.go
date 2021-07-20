package main

import "fmt"

type CircularQueue struct {
	front    int64           // 头坐标
	rear     int64           // 尾坐标
	index    int64           // 指针坐标
	stepSize int64           // 步长
	data     [][]interface{} // 数据
}

func NewCircularQueue(total int64, stepSize int64) *CircularQueue {
	// 初始化结构
	data := make([]interface{}, total, total)
	c := CircularQueue{
		front:    0,
		rear:     total,
		index:    0,
		stepSize: stepSize,
		data:     data,
	}

	return &c
}

func main() {
	fmt.Println("vim-go")
}
