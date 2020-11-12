package controllers

import (
	"BBS/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context){
	//1. 参数校验

	//2. 业务处理
	logic.Signup()
	//3. 返回响应
	c.JSON(http.StatusOK,"ok")
}