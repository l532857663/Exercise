package date

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

// @Description 获取时间数据
// @Author Wangch
// @Version 1.0
// @Update Wangch 2020-09-14 init
func (d *DateProcessor) GetTimeByDate(date interface{}) (time.Time, error) {
	var res time.Time
	var err error
	// 判断传入的时间参数是什么类型
	dateType := reflect.TypeOf(date).Name()
	switch dateType {
	case "string":
		res, err = time.ParseInLocation(GO_TIME_FORMAT, date.(string), time.Local)
	case "int64":
		res = time.Unix(date.(int64), 0)
	}
	return res, err
}

// @Description 判断处理时间分页参数
// @Author Wangch
// @Version 1.0
// @Update Wangch 2020-09-14 init
func (d *DateProcessor) CheckTimePageParameter(req *TimePageReq, flag string) (*TimePageFilter, error) {
	// 判断请求处理
	var err error
	var reqFilter = &TimePageFilter{}
	if flag != REQ_FLAG_ALL && flag != REQ_FLAG_TIME && flag != REQ_FLAG_PAGE {
		// flag参数不正确
		return reqFilter, errors.New("check time page param flag error!")
	}
	// 时间参数处理
	if flag == REQ_FLAG_ALL || flag == REQ_FLAG_TIME {
		// 开始时间为空，或者开始时间大于结束时间，则开始时间置为今天0点
		if req.StartTime <= 0 || req.StartTime >= req.EndTime {
			reqFilter.StartTime, err = time.Parse(GO_DATE_FORMAT, time.Now().Format(GO_DATE_FORMAT))
			if err != nil {
				return reqFilter, err
			}
		} else {
			reqFilter.StartTime, err = time.Parse(GO_TIME_FORMAT, time.Unix(req.StartTime, 0).Format(GO_TIME_FORMAT))
			if err != nil {
				return reqFilter, err
			}
		}
		// 结束时间为空，或者开始时间大于结束时间，则结束时间置为今天0点
		if req.EndTime <= 0 || req.StartTime >= req.EndTime {
			reqFilter.EndTime, err = time.Parse(GO_TIME_FORMAT, time.Now().Format(GO_TIME_FORMAT))
			if err != nil {
				return reqFilter, err
			}
		} else {
			reqFilter.EndTime, err = time.Parse(GO_TIME_FORMAT, time.Unix(req.EndTime, 0).Format(GO_TIME_FORMAT))
			if err != nil {
				return reqFilter, err
			}
		}
	}
	// 分页参数处理
	if flag == REQ_FLAG_ALL || flag == REQ_FLAG_PAGE {
		// 如果分页数量是空，每页数量默认是10
		if req.PageSize == "0" || req.PageSize == "" {
			reqFilter.PageSize = 10
		} else {
			reqFilter.PageSize, err = strconv.ParseInt(req.PageSize, 10, 64)
			if err != nil {
				return reqFilter, err
			}
		}
		// 如果分页页码是空，默认是第一页
		if req.PageNum == "0" || req.PageNum == "" {
			reqFilter.PageNum = 0
		} else {
			reqFilter.PageNum, err = strconv.ParseInt(req.PageNum, 10, 64)
			if err != nil {
				return reqFilter, err
			}
			// 偏移量处理
			reqFilter.PageNum = (reqFilter.PageNum - 1) * reqFilter.PageSize
		}
	}
	return reqFilter, err
}

// @Description 获取时间戳周的周一日期
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-01-15 init
func (d *DateProcessor) GetFirstDateOfWeek(now time.Time) time.Time {
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStartDate
}
