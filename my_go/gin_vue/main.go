package main

import (
	"gin_vue/core"
	"gin_vue/global"
	"gin_vue/initialize"
	"sync"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	global.GVA_WIAT = &(sync.WaitGroup{}) // 初始化线程锁
	if global.GVA_DB != nil {
		// NOTE: 配置文件设置后，启动会初始化数据库
		initialize.MysqlTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
		// 循环队列现基于redis服务
		global.GVA_OrderPayment = initialize.OrdersPaymentStart()
		go initialize.OrdersPayment()
	}
	core.RunWindowsServer()
}
