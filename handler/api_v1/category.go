package api_v1

import (
	"GBolg/conf/errmsg"
	"GBolg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//查询分类是否存在

//查询单个分类下的文章

// 添加分类
func AddCategory(c *gin.Context) {
	var data models.Category
	_ = c.ShouldBindJSON(&data)
	code := models.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		models.CreateCategory(&data)
	}
	if code == errmsg.ERROR_CATEGORY_USED {
		code = errmsg.ERROR_CATEGORY_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCategoryList 查询分类列表
func GetCategoryList(c *gin.Context) {
	data := models.GetCategoryList(QueryPageSizeCheck(c), QueryPageNumCheck(c))
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// UpdateCategory 编辑分类
func UpdateCategory(c *gin.Context) {
	categoryID := ParamIdCheck(c)

	var data models.Category
	_ = c.ShouldBindJSON(&data)
	code := models.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		code = models.UpdateCategory(categoryID, &data)
	}
	if code == errmsg.ERROR_CATEGORY_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	categoryID := ParamIdCheck(c)
	code := models.DeleteCategory(categoryID)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
