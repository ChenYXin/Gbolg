package middleware

import (
	"GBolg/conf/errmsg"
	"GBolg/utils/logrus_logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// 用户信息类，作为生成token的参数
type UserClaims struct {
	UserName string `json:"userName"`
	//jwt-go提供的标准claim
	jwt.StandardClaims
}

var (
	//自定义的token秘钥
	secret = []byte("16qwxc413Axcee8JOD9036f487PKwi")
	//该路由下不校验token
	//noVerify = []interface{}{"/login", "/ping"}
	//token有效时间（纳秒）
	effectTime = 24 * time.Hour * 365
)

// GenerateToken 生成token
func GenerateToken(claims *UserClaims) string {
	//设置token有效期，也可不设置有效期，采用redis的方式
	//   1)将token存储在redis中，设置过期时间，token如没过期，则自动刷新redis过期时间，
	//   2)通过这种方式，可以很方便的为token续期，而且也可以实现长时间不登录的话，强制登录
	//本例只是简单采用 设置token有效期的方式，只是提供了刷新token的方法，并没有做续期处理的逻辑
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		logrus_logger.LogRus.Errorf("jwt signing failed: %v", err)
		code := errmsg.ErrorTokenSigningFail
		panic(code)
	}
	return sign
}

// JwtVerify 验证token
func JwtVerify(c *gin.Context) {
	var code int
	//过滤是否验证token
	tokenString := c.GetHeader("token")
	if tokenString == "" {
		logrus_logger.LogRus.Errorf("jwt token is empty")
		code = errmsg.ErrorTokenNotExist
		panic(code)
	}
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		logrus_logger.LogRus.Errorf("jwt parse failed: %v", err)
		code = errmsg.ErrorTokenParseFail
		panic(code)
	}
	_, ok := token.Claims.(*UserClaims)
	if !ok {
		logrus_logger.LogRus.Errorf("jwt parse failed: %v", err)
		code = errmsg.ErrorTokenParseFail
		panic(code)
	}
	c.Next()
}
