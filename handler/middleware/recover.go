package middleware

import (
	"GBolg/conf/errmsg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			code, msg := errorToString(r)
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": msg,
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) (int, string) {
	switch v := r.(type) {
	case error:
		return http.StatusInternalServerError, v.Error()
	case int:
		return v, errmsg.GetErrMsg(v)
	default:
		return http.StatusInternalServerError, r.(string)
	}
}
