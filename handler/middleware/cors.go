package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			//可访问域名
			c.Header("Access-Control-Allow-Origin", origin)
			//可访问方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
			//支持的头信息字段
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Content-Length, Token")

			//可被返回前端得到的字段
			c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
			//预检有效期
			c.Header("Access-Control-Max-Age", "172800")
			//是否可发送Cookie
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json") //设置返回格式是json
		}
		if method == "OPTIONS" {
			//c.JSON(200, Controller.R(200, nil, "Options Request"))
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
