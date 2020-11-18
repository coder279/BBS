package controllers

import (
	"BBS/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// --- 社区相关

func CommunityHandler(c *gin.Context){
	data,err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList()",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	print(data)
	ResponseSuccess(c,data)
	return
}

func CommunityDetailHandler(c *gin.Context){
	communityID := c.Param("id")
	id,err := strconv.ParseInt(communityID,10,64)
	if err != nil {
		ResponseError(c,CodeInvalidParams)
		return
	}
	data,err := logic.GetCommunityDetail(id)
	if err != nil {
		ResponseError(c,CodeInvalidParams)
		return
	}
	ResponseSuccess(c,data)

}