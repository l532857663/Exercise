package model

import (
	"sync"
)

type Service struct {
	conf  *Conf     // 配置文件
	Ceshi *TestConf // 测试数据
}

type Conf struct {
	Name   string
	Attr   string
	RunEnv string
}

type TestConf struct {
	sync.Mutex
	testMap map[string]string
	Data    DataType
	Ch      chan string
}

type DataType string

type TimePageReq struct {
	StartTime int64 `form:"start_time" json:"start_time"`
	EndTime   int64 `form:"end_time" json:"end_time"`
	PageNo    int   `form:"page_no" json:"page_no"`
	PageSize  int   `form:"page_size" json:"page_size"`
}

type CeshiStruct struct {
	ID   string
	Name string
	TimePageReq
}
