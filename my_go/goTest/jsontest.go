package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := map[string]interface{}{
		"a": nil,
		"b": "aaa",
	}

	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Printf("b: %+v\n", b)
	fmt.Printf("b: %+v\n", string(b))
}
