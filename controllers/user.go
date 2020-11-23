package controllers

import (
	"BBS/logic"
	"BBS/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context){
	//1. 参数校验
	var p models.ParamSignUp
	if err := c.ShouldBind(&p);err != nil {
		zap.L().Error("Signup with invalid param",zap.Error(err))
		errs,ok := err.(validator.ValidationErrors)
		if !ok{
			ResponseError(c,CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c,CodeInvalidParams,removeTopStruct(errs.Translate(trans)))
		return
	}
	fmt.Println(p)
	//2. 业务处理
	err := logic.Signup(&p)
	if(err != nil ){
		ResponseError(c,CodeInvalidParams)
		return
	}
	//3. 返回响应
	ResponseSuccess(c,CodeSuccess)
}

func LoginHandler(c *gin.Context){
	//1.参数校验
	var p models.ParamLogin
	if err := c.ShouldBind(&p);err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c,CodeInvalidParams,err.Error())
			return
		}
		ResponseErrorWithMsg(c,CodeInvalidParams,removeTopStruct(errs.Translate(trans)))
		return
	}
	//数据比对
	user,err := logic.Login(&p)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,err.Error())
		return
	}
	//3.响应结果
	ResponseSuccess(c,gin.H{
		"user_id":fmt.Sprintf("%d",user.UserID),
		"user_name":user.Username,
		"token":user.Token,
	})
	return
}