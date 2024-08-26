package api_v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	pageSize = 0
	pageNum  = 10
)

func QueryPageSizeCheck(c *gin.Context) int {
	index, ok := c.GetQuery("pageSize")
	if !ok {
		return pageSize
	}
	num, err := strconv.Atoi(index)
	if err != nil {
		return pageSize
	}
	return num
}

func QueryPageNumCheck(c *gin.Context) int {
	size, ok := c.GetQuery("pageNum")
	if !ok {
		return pageNum
	}
	num, err := strconv.Atoi(size)
	if err != nil {
		return pageNum
	}
	return num
}

func ParamIdCheck(c *gin.Context) uint {
	id := c.Param("id")
	intNum, _ := strconv.Atoi(id)
	ID := uint(intNum)
	//ID, err := strconv.ParseUint(id, 10, 64)
	//ID, err := strconv.Atoi(id)
	//if err != nil {
	//	return -1
	//}
	return ID
}

func ParamCidCheck(c *gin.Context) int {
	cid := c.Param("cid")
	CID, err := strconv.Atoi(cid)
	if err != nil {
		return -1
	}
	return CID
}
