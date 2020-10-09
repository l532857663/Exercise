package main

import (
	"fmt"
	"goStruct/model"
	"time"
)

var (
	s *CronServer
)

type CronServer struct {
	cronSrv *model.Service
}

func main() {
	fmt.Println("vim-go")
	appName := "cron"
	c := &model.Conf{
		Name:   appName,
		Attr:   "test",
		RunEnv: "RUN_ENV_JOIN_DEBUG",
	}

	if s == nil {
		s = &CronServer{
			cronSrv: model.New(c, "cronTest"),
		}
	}

	s.cronSrv.TestDoit(appName)
	time.Sleep(5 * time.Second)
	s.cronSrv.TestDoit(appName)
}
