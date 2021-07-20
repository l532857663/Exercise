package main

import (
	"fmt"
	"time"
	"utils/date"
)

func main() {
	fmt.Println("vim-go")
	var res time.Time
	var err error
	res, err = date.GetTimeByDate("2021-07-28 10:00:00")
	if err != nil {
		fmt.Println("get time by string error: ", err)
		return
	}
	fmt.Println("res1: ", res)
	res, err = date.GetTimeByDate(res.Unix())
	if err != nil {
		fmt.Println("get time by int64 error: ", err)
		return
	}
	fmt.Println("res2: ", res)

	return
}
