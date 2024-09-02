package dao

import (
	"GBolg/utils/logrus_logger"
	"GBolg/utils/viper_config"
	"context"
	"github.com/redis/go-redis/v9"
)

var (
	Redis *redis.Client
)

func InitMyRedis() {
	address := viper_config.VC.GetString("redis.addr")
	password := viper_config.VC.GetString("redis.pass")
	db := viper_config.VC.GetInt("redis.db")

	Redis = redis.NewClient(&redis.Options{
		Username: "default",
		Addr:     address,
		Password: password,
		DB:       db,
	})
	var err error
	if err = Redis.Ping(context.Background()).Err(); err != nil {
		logrus_logger.LogRus.Panicf("connect to redis %d failed %v", db, err)
	} else {
		logrus_logger.LogRus.Infof("connect to redis %d", db) //能ping成功才说明连接成功
	}
}
