package date

import (
	"reflect"
	"time"
)

// 获取时间数据
func GetTimeByDate(date interface{}) (time.Time, error) {
	var res time.Time
	var err error
	// 判断传入的时间参数是什么类型
	dateType := reflect.TypeOf(date).Name()
	switch dateType {
	case "string":
		res, err = time.ParseInLocation(layout, date.(string), time.Local)
	case "int64":
		res = time.Unix(date.(int64), 0)
	}
	return res, err
}
