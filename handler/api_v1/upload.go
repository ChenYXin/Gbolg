package api_v1

import (
	"GBolg/conf/errmsg"
	"GBolg/models"
	"GBolg/utils/logrus_logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		logrus_logger.LogRus.Errorf("file form err: %v", err)
		panic(errmsg.ERROR)
	}
	fileSize := fileHeader.Size
	url, code := models.UploadToQiniu(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
