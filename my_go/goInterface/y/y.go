package y

import (
	"fmt"
)

type ObjA struct {
	ID string
}

func (a *ObjA) Haha() string {
	return fmt.Sprintf("name is %s", a.ID)
}

func GetObjA(name string) *ObjA {
	objA := &ObjA{
		ID: name,
	}

	return objA
}
