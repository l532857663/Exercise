package main

import (
	"fmt"
	"goStruct/model"
	"sync"
	"time"
)

var (
	s *GatewayServer
)

type GatewayServer struct {
	gatewaySrv *model.Service
}

func main() {
	fmt.Println("vim-go")
	appName := "gateway"
	c := &model.Conf{
		Name:   appName,
		Attr:   "test",
		RunEnv: "R",
	}

	if s == nil {
		s = &GatewayServer{
			gatewaySrv: model.New(c, "gatewayTest"),
		}
	}

	wg := sync.WaitGroup{}
	s.gatewaySrv.TestDoit(appName)
	wg.Add(1)
	go func() {
		s.gatewaySrv.TestChannel()
		wg.Done()
	}()
	var i int
BothControl:
	for {
		select {
		case <-s.gatewaySrv.Ceshi.Ch:
			wg.Wait()
			break BothControl
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("didi", i)
			break
			i++
		}
	}
}
