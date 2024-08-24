package main

import (
	"GBolg/dao"
	"GBolg/routes"
	"GBolg/utils/logrus_logger"
	"GBolg/utils/viper_config"
)

func main() {
	//初始化yaml文件
	viper_config.InitViperConfig()
	//初始化日志
	logrus_logger.InitLog()
	//连接数据库
	dao.InitMySql()

	//路由
	routes.SetupRouter()
}
