package controllers

import (
	"BBS/logic"
	"BBS/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func PostVoteHandler(c *gin.Context){
	//1.参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p);err == nil {
		errs,ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c,CodeInvalidParams)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c,CodeInvalidParams,errData)
		return
	}
	//获取当前用户id
	userID,err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c,CodeNotLogin)
		return
	}
	//具体投票业务逻辑
	if err := logic.VoteForPost(userID,p);err != nil {
		zap.L().Error("logic.VoteForPost",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,nil)
	return
}
