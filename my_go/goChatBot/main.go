package main

import (
	"goChatBot/global"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 初始化服务
	err := global.Init()
	if err != nil {
		return
	}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		global.Stop()
		os.Exit(0)
	}()
	global.CreateBotReceive()
}
