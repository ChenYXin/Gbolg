package main

import (
	"GBolg/dao"
	"GBolg/models"
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
	//同步表
	errDbAutoMigrate := dao.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Article{},
	)
	if errDbAutoMigrate != nil {
		panic(errDbAutoMigrate)
	}
	logrus_logger.LogRus.Infof("create table success")

	//路由
	routes.SetupRouter()
}
