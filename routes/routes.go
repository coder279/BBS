package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"BBS/logger"
)

func Setup()*gin.Engine{
	r := gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"ok")
	})
	return r
}