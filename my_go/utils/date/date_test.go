package date

import (
	"fmt"
	"testing"
	"time"
)

func Test_GetTimeByDate(t *testing.T) {
	var d DateProcessor
	var res time.Time
	var err error
	res, err = d.GetTimeByDate("2021-07-28 10:00:00")
	if err != nil {
		fmt.Println("get time by string error: ", err)
		return
	}
	fmt.Println("res1: ", res)
	res, err = d.GetTimeByDate(res.Unix())
	if err != nil {
		fmt.Println("get time by int64 error: ", err)
		return
	}
	fmt.Println("res2: ", res)
}

func Test_GetFirstDateOfWeek(t *testing.T) {
	var d DateProcessor
	// 获取某时间的周一8点时间戳
	now := time.Now()
	monday := d.GetFirstDateOfWeek(now)
	addTime := time.Duration(8 * time.Hour)
	fmt.Println("monday: ", monday.Add(addTime).Unix())
}
