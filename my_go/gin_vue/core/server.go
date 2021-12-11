package core

import (
	"fmt"
	"gin_vue/global"
	"gin_vue/initialize"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

var shutdownCh chan struct{}

func RunWindowsServer() {
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	signalHandler(appletName)
}

func ServerStart(s server) {
	shutdownCh = make(chan struct{})
	go s.ListenAndServe()
	// global.GVA_LOG.Error(s.ListenAndServe().Error())
}

func ServerClose() {
	// 关闭订单状态检测服务
	initialize.OrdersPaymentClose()
	time.Sleep(2 * time.Second)
	global.GVA_WIAT.Wait()
}

// @Description 系统信号捕获
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-12-11 init
func signalHandler(appletName string) {
	var (
		ch = make(chan os.Signal)
	)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

WAIT_SIGNAL:
	for {
		sig := <-ch
		global.GVA_LOG.Info(fmt.Sprintf("get a signal \"%s\", stop the %s process", sig.String(), appletName))
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP:
			// 安全关闭服务
			go func() {
				ServerClose()
				close(shutdownCh)
				return
			}()

			break WAIT_SIGNAL

		default:
			return
		}
	}

	select {
	case <-shutdownCh:
		fmt.Println("Gracefully shutting down server ...")
	}

	return
}
