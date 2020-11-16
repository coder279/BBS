package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
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