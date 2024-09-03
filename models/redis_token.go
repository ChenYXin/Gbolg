package models

import (
	"GBolg/conf"
	"GBolg/dao"
	"GBolg/utils/logrus_logger"
	"context"
)

/**
什么时候使用Redis?
1. 高并发，低延时。redis比mysql快一到两个数量级。
2. redis可靠性没mysql高，万一redis挂了对业务影响不大，好修复。
3. redis通常存储string型value，此时它相对于mysql的性能优势更明显。
*/

const (
	tokenPrefix = "GBlog:username:"
)

// SetToken 把<username, token>写入redis
func SetToken(username, token string) {
	if err := dao.Redis.Set(context.Background(), tokenPrefix+username, token, conf.EffectTime).Err(); err != nil {
		logrus_logger.LogRus.Errorf("write token pair(%s, %s) to redis failed: %s", username, token, err)
	}

}

// GetToken 根据username获取token
func GetToken(username string) string {
	token, err := dao.Redis.Get(context.Background(), tokenPrefix+username).Result()
	if err != nil {
		logrus_logger.LogRus.Errorf("write token pair(%s, %s) to redis failed: %s", username, token, err)
	}
	return token
}
