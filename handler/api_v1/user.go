package api_v1

import (
	"GBolg/conf/errmsg"
	"GBolg/handler/middleware"
	"GBolg/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data models.User
	_ = c.ShouldBindJSON(&data)
	code := models.CheckUser(data.UserName)
	if code == errmsg.SUCCESS {
		middleware.GenerateToken(&middleware.UserClaims{
			UserName:       data.UserName,
			StandardClaims: jwt.StandardClaims{},
		})

		models.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func UserLogin(c *gin.Context) {
	var data models.User
	_ = c.ShouldBindJSON(&data)
	user, code := models.CheckLogin(data.UserName, data.Password)
	var token string
	if code == errmsg.SUCCESS {
		currentToken := models.GetToken(user.UserName)
		token = middleware.Refresh(currentToken)
		models.UpdateUser(user.ID, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    user,
		"token":   token,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUserList 查询用户列表
func GetUserList(c *gin.Context) {
	data, code := models.GetUserList(QueryPageSizeCheck(c), QueryPageNumCheck(c))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// UpdateUser 编辑用户
func UpdateUser(c *gin.Context) {
	userID := ParamIdCheck(c)

	var data models.User
	_ = c.ShouldBindJSON(&data)
	code := models.CheckUser(data.UserName)
	if code == errmsg.SUCCESS {
		middleware.GenerateToken(&middleware.UserClaims{
			UserName:       data.UserName,
			StandardClaims: jwt.StandardClaims{},
		})
		code = models.UpdateUser(userID, &data)
	}
	if code == errmsg.ErrorUserNameIsExist {
		panic(code)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	userID := ParamIdCheck(c)
	code := models.DeleteUser(userID)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
