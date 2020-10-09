package model

import (
	"fmt"
	"time"
)

var (
	s *Service
)

func New(c *Conf, appletName string) *Service {
	if s == nil {
		s = &Service{
			conf: c,
		}

		if c.RunEnv == "RUN_ENV_JOIN_DEBUG" {
			fmt.Println(c)
		} else {
			// 加载数据库配置信息
			s.LoadDbConfig()
		}
	}

	return s
}

// LoadDbConfig 加载数据库配置信息
func (s *Service) LoadDbConfig() {
	fmt.Println("ceshi conf start")
	s.Ceshi = &TestConf{
		testMap: map[string]string{
			"haha": "ha1",
		},
		Data: "123456",
	}
	s.Ceshi.Ch = make(chan string)
	fmt.Println("ceshi conf end")
}

func (s *Service) TestDoit(appName string) {
	fmt.Println("test app name:", appName)
	fmt.Printf("do it show conf: %+v\n", s.conf)
	fmt.Printf("do it show ceshi: %v\n", s.Ceshi)
}

func (s *Service) TestChannel() {
	fmt.Println("chan start")
	time.Sleep(5 * time.Second)
	s.Ceshi.Ch <- "you"
	fmt.Println("chan end")
	return
}
