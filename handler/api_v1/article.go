package api_v1

import (
	"GBolg/conf/errmsg"
	"GBolg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddArticle(c *gin.Context) {
	var data models.Article
	_ = c.ShouldBindJSON(&data)
	code := models.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArticleList 文章列表
func GetArticleList(c *gin.Context) {
	data, total, code := models.GetArticleList(QueryKeyWordCheck(c), QueryPageSizeCheck(c), QueryPageNumCheck(c))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetArticleCategoryList(c *gin.Context) {
	cid := ParamCidCheck(c)

	data, total, code := models.GetArticleCategoryList(cid, QueryPageSizeCheck(c), QueryPageNumCheck(c))
	//code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArticleInfo 查询单个文章
func GetArticleInfo(c *gin.Context) {
	aeticleID := ParamIdCheck(c)
	data, code := models.GetArticleById(aeticleID)
	//code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func UpdateArticle(c *gin.Context) {
	aeticleID := ParamIdCheck(c)

	var data models.Article
	_ = c.ShouldBindJSON(&data)
	code := models.UpdateArticle(aeticleID, &data)
	if code == errmsg.ERROR_CATEGORY_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func DeleteArticle(c *gin.Context) {
	categoryID := ParamIdCheck(c)
	code := models.DeleteArticle(categoryID)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
