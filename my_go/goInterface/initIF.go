package main

import (
	"fmt"
	"goInterface/e"
	"goInterface/model"
	"goInterface/y"
	"reflect"
)

type (
	ObjA interface {
		Ceshi() model.Service
	}

	LaughType interface {
		ObjA
		Haha() string
	}

	DoitType interface {
		ObjA
		Ohou() string
	}

	ObjAs map[string]interface{}
)

func main() {
	fmt.Println("vim-go")

	oS := GetInf()
	for key, s := range oS {
		l, okl := s.(LaughType)
		if okl {
			fmt.Printf("true l: %+v, type: %v\n", l, reflect.TypeOf(l))
			fmt.Printf("key: %s, s: %+v\n", key, s)
			// 传指针 赋值未初始化的字段
			// 传值 调用指针方法
			srv := l.Ceshi()
			fmt.Printf("srv: %+v, func: %s\n", srv, srv.Ceshi())
			// 传值 赋值未初始化的字段
			srv.Age = 22
			fmt.Printf("srv: %+v\n", srv)
		} else {
			fmt.Printf("false l: %+v\n", l)
		}
		/*
			l, okl := s.(DoitType)
			if okl {
				fmt.Printf("true l: %+v\n", l)
			} else {
				fmt.Printf("false l: %+v\n", l)
			}
		*/
	}
}

func GetInf() ObjAs {
	a := e.GetObjA("zhang")
	a.Age = 25
	a.Class = &e.Class{
		Chinese: 99,
		Math:    98,
	}
	fmt.Printf("a obj: %+v, class: %+v\n", a, a.Class)
	fmt.Println(a.Ohou())

	return ObjAs{
		"a": y.GetObjA("zhang"),
		"b": e.GetObjA("wang"),
	}
}
