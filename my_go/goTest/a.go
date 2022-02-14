package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func test1() {
	a := "11773422559966"
	b := make(map[string]string)
	for _, c := range a {
		a1 := string(c)
		if _, ok := b[a1]; !ok {
			b[a1] = "OK"
			continue
		}
		delete(b, a1)
	}
}

type Queue struct {
	Length int
	A      []string
	B      []string
}

func (q *Queue) Add(char string) {
	if len(q.A) == 0 {
		q.B = append(q.B, char)
		q.Length = len(q.B)
		return
	}

	q.A = append(q.A, char)
	q.Length = len(q.A)
	return
}

func (q *Queue) Pop() string {
	var res string
	index := q.Length - 1
	if len(q.A) == 0 {
		q.A = make([]string, index)
		copy(q.A, q.B[:index])
		res = q.B[index]
		q.B = make([]string, q.Length, q.Length)
		q.Length = len(q.A)
		return res
	}

	q.B = make([]string, index)
	copy(q.B, q.A[:index])
	res = q.A[index]
	q.A = make([]string, q.Length, q.Length)
	q.Length = len(q.B)
	return res
}

func test2() {
	q := &Queue{}
	q.Add("a")
	fmt.Printf("q: %+v\n", q)
	q.Add("b")
	fmt.Printf("q: %+v\n", q)
	q.Add("c")
	fmt.Printf("q: %+v\n", q)
	res := q.Pop()
	fmt.Printf("res: %v\n", res)
	fmt.Printf("q: %+v\n", q)
	q.Add("d")
	fmt.Printf("q: %+v\n", q)
	res1 := q.Pop()
	fmt.Printf("res1: %v\n", res1)
	fmt.Printf("q: %+v\n", q)
}

type Kobj struct {
	KL  string
	KR  string
	Arr []string
}

func doSum(k string, length int) string {
	res := ""
	for length > 0 {
		res += k
		length--
	}
	return res
}

func (k *Kobj) ArrGet(data string, l int) {
	data += doSum(k.KL, l)
	data += doSum(k.KR, l)
	k.Arr = append(k.Arr, data)
}

func (k *Kobj) GetData(c int, src string) {
	// 数量配置
	for i := c; i > 0; i-- {
		data := src + ""
		data += doSum(k.KL, i)
		data += doSum(k.KR, i)
		j := c - i
		// 第一层，j==0
		if j == 0 {
			k.ArrGet(data, j)
			continue
		}
		// 从第二层开始
		if j >= 1 {
			k.GetData(j, data)
			j--
			if j > 0 {
				data += k.KL + k.KR
			}
		}
	}
}

func test3() {
	// 数据
	k := &Kobj{
		KL: "(",
		KR: ")",
	}
	c := 5
	k.GetData(c, "")
	fmt.Println("count:", c)
	fmt.Println("arr:", k.Arr)
}

type LRUCache struct {
	Head, Tail *Node
	KeyMap     map[int]*Node
	Length     int
}

type Node struct {
	Prve, Next *Node
	Key, Value int
}

func newLRUCache(length int) *LRUCache {
	this := &LRUCache{
		Head:   &Node{},
		Tail:   &Node{},
		KeyMap: make(map[int]*Node),
		Length: length,
	}
	return this
}

func (this *LRUCache) Get(key int) int {
	// 获取缓存中的数据
	v, ok := this.KeyMap[key]
	if !ok {
		// 不存在缓存中，返回空
		return -1
	}
	// 存在缓存中，去掉缓存，挪至第一位
	this.Remove(v)
	// 插入最前边
	this.Insert(v)
	return v.Value
}

func (this *LRUCache) Put(key, value int) {
	// 判断是否存在缓存中
	v, ok := this.KeyMap[key]
	if ok {
		// 存在更新缓存数据
		this.Remove(v)
		v.Value = value
		this.Insert(v)
		return
	}
	// 不存在，判断缓存长度
	if this.Length == len(this.KeyMap) {
		// 缓存填满，删除最后一个节点
		this.Remove(this.Tail.Prve)
	}
	// 把数据插入第一个节点
	this.Insert(&Node{
		Next:  this.Head,
		Key:   key,
		Value: value,
	})
	return
}

func (this *LRUCache) Remove(node *Node) {
	// 删除字典中的key
	delete(this.KeyMap, node.Key)
	// 移除链表对应的数据
	node.Prve.Next = node.Next
	node.Next.Prve = node.Prve
	return
}

func (this *LRUCache) Insert(node *Node) {
	// 添加字典信息
	this.KeyMap[node.Key] = node
	// 把节点加入链表头部
	if this.Head.Next != nil {
		next := this.Head.Next
		next.Prve = node
		node.Next = next
	}
	this.Head.Next = node
	node.Prve = this.Head
	if this.Tail.Prve == nil {
		this.Tail.Prve = node
		node.Next = this.Tail
	}
	return
}

func test4() {
	// arr := []int{1, 2, 1, 3, 2}
	arr := []int{1, 2, 1, 3, 2}
	s := newLRUCache(2)
	fmt.Printf("wch------ new %+v\n", s)
	for _, a := range arr {
		// GET
		data := s.Get(a)
		if data != -1 {
			fmt.Printf("wch----- get %v, s %+v\n", a, s.Head.Next)
			continue
		}
		// PUT
		s.Put(a, a)
		fmt.Printf("wch----- put %v, s %+v\n", a, s.Head.Next)
	}
	head := s.Head
	for head.Next != nil {
		fmt.Printf("wch------ head %+v\n", head)
		head = head.Next
	}
}

type ANode struct {
	Next  *ANode
	Value int
}

func SetANodeValue(data int, next *ANode) *ANode {
	node := &ANode{
		Next:  next,
		Value: data,
	}
	return node
}

func showLink(node *ANode) {
	fmt.Printf("wch---- %+v\n", node)
	for node.Next != nil {
		fmt.Printf("wch---- %+v\n", node.Next)
		node = node.Next
	}
}

func reverseLink(head *ANode) *ANode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseLink(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}

func ReANode(link *ANode, k int) *ANode {
	temp := link
	for i := 1; i < k && temp != nil; i++ {
		temp = temp.Next
	}
	if temp == nil {
		return link
	}
	link2 := temp.Next
	temp.Next = nil
	// 逆序
	showLink(link2)
	newHead := reverseLink(link)
	// 递归
	newLink := ReANode(link2, k)
	link.Next = newLink
	return newHead
}

func test5() {
	node := SetANodeValue(1, SetANodeValue(2, SetANodeValue(3, SetANodeValue(4, SetANodeValue(5, nil)))))

	node = ReANode(node, 2)
	showLink(node)
	return
}

func test6() {
	a := "()())((()))"
	var res bool
	var sum, max, tmp int
	for _, str := range strings.Split(a, "") {
		if str == "(" {
			sum += 1
			tmp += 1
		} else {
			sum -= 1
			tmp += 1
		}
		if sum < 0 {
			if tmp > max {
				max = tmp
			}
			tmp = 0
		}
	}
	if tmp > max {
		max = tmp
	}
	if sum == 0 {
		res = true
	}
	fmt.Println("wch------ res max\n", res, max)
}

func doit(ch, ch1 chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		ch <- "asdasd"
		if i == 5 {
			close(ch1)
		}
	}
	return
}

func test7() {
	ch := make(chan string, 10)
	ch1 := make(chan string)
	go doit(ch, ch1)
	i := 0
	for {
		select {
		//		case res := <-ch:
		//			fmt.Println("wch------ 1", res)
		case <-ch:
			fmt.Println("wch------ 1-1")
		case <-ch1:
			fmt.Println("wch------ 2", i)
			i += 1
		default:
			fmt.Println("wch----- default")
		}
		if i == 10 {
			break
		}
		time.Sleep(time.Second * 1)
	}
	return
}

func getChan(ch1, ch2, ch3 chan int) {
	fmt.Println("wch---- getChan start")
	go func() {
		for i := 0; i < 100; i++ {
			switch {
			case i%2 == 0:
				ch1 <- i
			case i%3 == 0:
				ch2 <- i
			case i%5 == 0:
				ch3 <- i
			}
		}
		close(ch1)
		close(ch2)
		close(ch3)
	}()
	fmt.Println("wch---- getChan end")
}

func test8() {
	ch1 := make(chan int, 50)
	ch2 := make(chan int, 50)
	ch3 := make(chan int, 50)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go getChan(ch1, ch2, ch3)
	go func(ch1, ch2, ch3 chan int, wg *sync.WaitGroup) {
		for {
			c1, ok1 := <-ch1
			if ok1 {
				fmt.Println("wch---- c1:", c1)
				continue
			}
			c2, ok2 := <-ch2
			if ok2 {
				fmt.Println("wch---- c2:", c2)
				continue
			}
			c3, ok3 := <-ch3
			if ok3 {
				fmt.Println("wch---- c3:", c3)
				continue
			}
			if !ok1 && !ok2 && !ok3 {
				break
			}
		}
		wg.Done()
	}(ch1, ch2, ch3, &wg)
	wg.Wait()
	// time.Sleep(time.Second * 10)
}

func test9() {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6:10]
	fmt.Printf("myslice为 %d, 其长度为: %d\n", myslice, len(myslice))

	fmt.Printf("test------- %+v, %+v\n", myslice, cap(myslice))
	myslice = myslice[:cap(myslice)]
	fmt.Printf("myslice的第四个元素为: %d", myslice[5])
}

func main() {
	// test1()
	// 使用两个队列模拟栈
	// test2()
	// 把某数量的括号组合成正确的排列 eg: 3组 ["((()))","(())()","()(())","()()()"]
	// test3()
	// LRU 缓存淘汰算法 最近最少使用
	// test4()
	// 单链表反转
	// test5()
	// 有效括号字符串
	// test6()
	// select跟switch的区别
	// test7()
	// 输出100内2、3、5的倍数?
	// test8()
	// 数组和切片
	test9()
}
