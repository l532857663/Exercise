package main

import "fmt"

// 二叉树节点
type Node struct {
	Value       string
	Left, Right *Node
}

// 前序展示
func (this *Node) PrvShow() {
	if this == nil {
		return
	}
	fmt.Println(this.Value)
	this.Left.PrvShow()
	this.Right.PrvShow()
	return
}

// 按层展示
func (this *Node) FloorShow() {
	if this == nil {
		return
	}

	fmt.Println(this.Value)
	floor := []*Node{this}
	for len(floor) > 0 {
		// 新数据缓存下一层数据
		tmpFloor := []*Node{}
		// 遍历该层级
		for i := 0; i < len(floor); i++ {
			if floor[i].Left != nil {
				tmpFloor = append(tmpFloor, floor[i].Left)
				fmt.Println(floor[i].Left)
			}
			if floor[i].Right != nil {
				tmpFloor = append(tmpFloor, floor[i].Right)
				fmt.Println(floor[i].Right)
			}
		}
		floor = tmpFloor
	}
	return
}

func (this *Node) SetValue(value string) {
	this.Value = value
	return
}

func CreateNode(value string) *Node {
	node := &Node{
		Value: value,
	}
	return node
}

func main() {
	// 创建节点
	root := CreateNode("1")
	root.Left = CreateNode("2")
	root.Left.Left = CreateNode("4")
	root.Right = CreateNode("3")

	// root.PrvShow()
	root.FloorShow()
	return
}
