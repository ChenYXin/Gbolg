package api_v1

import (
	"GBolg/conf/errmsg"
	"GBolg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data models.User
	_ = c.ShouldBindJSON(&data)
	code := models.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		models.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUserList 查询用户列表
func GetUserList(c *gin.Context) {

	data := models.GetUserList(QueryPageSizeCheck(c), QueryPageNumCheck(c))
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// UpdateUser 编辑用户
func UpdateUser(c *gin.Context) {}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {}
