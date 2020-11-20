package routes

import (
	"BBS/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"BBS/logger"
	"BBS/controllers"
)

func Setup()*gin.Engine{
	r := gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	v1.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"ok")
	})
	v1.POST("/signup" , controllers.SignUpHandler)
	v1.POST("/login" , controllers.LoginHandler)
	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("/community",controllers.CommunityHandler)
		v1.GET("/community/:id",controllers.CommunityDetailHandler)

		v1.POST("/post",controllers.CreatePostHandler)
		v1.GET("/post/:id",controllers.GetPostDetailHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"404",
		})
	})
	return r
}