package e

import (
	"fmt"
	"goInterface/model"
)

type ObjA struct {
	Name string
	Age  int
	*Class
}

type Class struct {
	Chinese int
	Math    int
	Englist int
}

func (a *ObjA) Haha() string {
	return fmt.Sprintf("name is %s", a.Name)
}

func (a *ObjA) Ohou() string {
	return fmt.Sprintf("ohou,%s on here, I'm %d!", a.Name, a.Age)
}

func GetObjA(name string) *ObjA {
	objA := &ObjA{
		Name: name,
	}

	return objA
}

func (a *ObjA) Ceshi() model.Service {
	return model.Service{
		Name: a.Name,
	}
}
