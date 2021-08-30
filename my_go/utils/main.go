package main

import (
	"fmt"
	"utils/date"
)

func main() {
	fmt.Println("vim-go")

	var dateObj date.DateProcessor
	res, err := dateObj.GetTimeByDate("2021-07-29 05:00:00")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Println("res: ", res)
	return
}
