package goBTC

import (
	"goBTC/client"
	"goBTC/db"
	"goBTC/db/inscribe"
	"goBTC/global"
	"goBTC/models"
	"goBTC/utils/http"
	"goBTC/utils/logutils"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func MustLoad(confPath string) {
	global.CONFIG = readConfig(confPath)

	// 初始化节点客户端
	global.Client = newBTCClient()

	// 初始化zap日志库
	global.LOG = logutils.Log("", global.CONFIG.Zap)

	// 初始化查询平台
	// service.InitPlatformMap()

	// 初始化HTTPS数据
	http.InitHttps(global.CONFIG.Https.IsHttps, global.CONFIG.Https.CaCert)

	// 数据库连接
	if global.MysqlFlag {
		inscribeInfo := &inscribe.InscribeInfo{}
		db.Gorm(global.CONFIG.Mysql, global.LOG, inscribeInfo.TableName())
		ordBrc20 := &inscribe.OrdToken{}
		db.Gorm(global.CONFIG.Mysql, global.LOG, ordBrc20.TableName())
	}
}

func Shutdown() {
	// 关闭数据链接
	db.CloseAllGormDBConnections()
}

func newBTCClient() *client.BTCClient {
	// 构建节点客户端
	nodeInfo := client.BTC_QUICKNODE_MAIN
	cli, err := client.NewBTCClient(nodeInfo)
	if err != nil {
		log.Fatalf("NewBTCClient error: %+v, nodeInfo: %+v\n", err, nodeInfo)
	}
	// 初始化脚本map
	client.InitBtcScriptMap()
	return cli
}

func readConfig(confPath string) *models.Server {
	var config *models.Server
	data, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatalf("Read conf file[%s] error: %v", confPath, err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Yaml unmarshal error: %v", err)
	}
	return config
}
