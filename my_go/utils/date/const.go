package date

import "time"

type DateProcessor struct {
}

// 时间过滤器类型：
//  "0":日
//  "1":月
//  "2":季度
//  "3":年
type DateFilterType string

const (
	// DateFilterType
	DATE_FILTER_HOUR  DateFilterType = "0" // "0":小时
	DATE_FILTER_DAY   DateFilterType = "1" // "1":日
	DATE_FILTER_MONTH DateFilterType = "2" // "2":月
	DATE_FILTER_YEAR  DateFilterType = "3" // "3":年

	// Go标准时间格式
	GO_DATE_FORMAT              = "2006-01-02"
	GO_DATE_FORMAT_WITHOUT_BARS = "20060102"
	GO_TIME_FORMAT              = "2006-01-02 15:04:05"

	// 时间分页参数flag
	REQ_FLAG_ALL  = "all"
	REQ_FLAG_TIME = "time"
	REQ_FLAG_PAGE = "page"
)

var (
	// 亚洲时区
	AsiaLocation *time.Location
)
