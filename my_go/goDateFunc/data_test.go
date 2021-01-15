package goDateFunc

import (
	"fmt"
	"testing"
	"time"
)

func Test_GetFirstDateOfWeek(t *testing.T) {
	var d DateProcessor
	// 获取某时间的周一8点时间戳
	now := time.Now()
	monday := d.GetFirstDateOfWeek(now)
	addTime := time.Duration(8 * time.Hour)
	fmt.Println("monday: ", monday.Add(addTime).Unix())
}
