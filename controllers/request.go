package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

var ErrorUserNotLogin = errors.New("用户未登录")
var ContextUserIDKey string
func GetCurrentUser(c *gin.Context)(UserId int64,err error){
	uid,ok := c.Get(ContextUserIDKey)
	if !ok{
		return
	}
	return uid.(int64),nil
}

func getPageInfo(c *gin.Context)(int64,int64){
	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")
	var (
		offset int64
		limit int64
		err error
	)
	offset,err = strconv.ParseInt(offsetStr,10,64)
	if err != nil {
		offset = 0
	}
	limit,err = strconv.ParseInt(limitStr,10,64)
	if err != nil {
		limit = 10
	}
	return offset,limit
}