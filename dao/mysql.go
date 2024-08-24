package dao

import (
	"GBolg/utils/logrus_logger"
	"GBolg/utils/viper_config"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySql() {
	bUserName := viper_config.VC.GetString("blog.userName")
	bPassword := viper_config.VC.GetString("blog.password")
	bHost := viper_config.VC.GetString("blog.host")
	bPort := viper_config.VC.GetInt("blog.port")
	bDbName := viper_config.VC.GetString("blog.dbName")

	//将yaml配置参数拼接成连接数据库的url
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		bUserName, //blog.UserName,
		bPassword, //blog.Password,
		bHost,     //blog.Host,
		bPort,     //blog.Port,
		bDbName,   //blog.DbName,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus_logger.LogRus.Errorf("gorm open failed: %s\n", err)
		panic(err)
	}
	var sqlDB *sql.DB
	sqlDB, err = DB.DB()
	if err != nil {
		logrus_logger.LogRus.Errorf("gorm DB.DB failed: %s\n", err)
		panic(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		logrus_logger.LogRus.Errorf("gorm sqlDB.Ping failed: %s\n", err)
		panic(err)
	}

	logrus_logger.LogRus.Infof("connect to mysql success")
}
