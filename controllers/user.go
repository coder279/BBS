package controllers

import (
	"BBS/logic"
	"BBS/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

func SignUpHandler(c *gin.Context){
	//1. 参数校验
	var p models.ParamSignUp
	if err := c.ShouldBind(&p);err != nil {
		zap.L().Error("Signup with invalid param",zap.Error(err))
		errs,ok := err.(validator.ValidationErrors)
		if !ok{
			c.JSON(http.StatusOK,gin.H{
				"msg":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"msg":removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	fmt.Println(p)
	//2. 业务处理
	err := logic.Signup(&p)
	if(err != nil ){
		c.JSON(http.StatusOK,gin.H{
			"error":1,
			"msg":err.Error(),
		})
		return
	}
	//3. 返回响应
	c.JSON(http.StatusOK,gin.H{
		"msg":"success",
	})
}

func LoginHandler(c *gin.Context){
	//1.参数校验
	var p models.ParamLogin
	if err := c.ShouldBind(&p);err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
	}
	//数据比对
	err := logic.Login(&p)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"error":1,
			"msg":err.Error(),
		})
		return
	}
	//3.响应结果
	c.JSON(http.StatusOK,gin.H{
		"error":0,
		"msg":"登陆成功",
	})
	return
}