package main

import "fmt"

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
	KL    string
	KR    string
	Arr   []string
}

func doSum(k string, length int) string {
	res := ""
	for length > 0 {
		res += k
		length--
	}
	return res
}

func (k *Kobj) ArrGet(data string, j int) {
	data += doSum(k.KL, j)
	data += doSum(k.KR, j)
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
				data += "()"
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
	fmt.Println("count:",c)
	fmt.Println("arr:", k.Arr)
}

func main() {
	// test1()
	// 使用两个队列模拟栈
	// test2()
	// 把某数量的括号组合成正确的排列 eg: 3组 ["((()))","(())()","()(())","()()()"]
	test3()
}
