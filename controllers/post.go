package controllers

import (
	"BBS/logic"
	"BBS/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context){
	//1.获取参数并校验
	p := new (models.Post)
	if err := c.ShouldBindJSON(p);err != nil {
		ResponseError(c,CodeInvalidParams)
		return
	}
	userID,err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c,CodeNotLogin)
		return
	}
	p.AuthorID = userID
	//2.创建帖子
	if err := logic.CreatePost(p);err != nil {
		zap.L().Error("logic createPost(c) failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}

	//3.返回响应
	ResponseSuccess(c,CodeSuccess)
}
