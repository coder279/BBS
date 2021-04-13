package controllers

import (
	"BBS/logic"
	"BBS/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	//1.获取参数并校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}
	p.AuthorID = userID
	//2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic createPost(c) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//3.返回响应
	ResponseSuccess(c, CodeSuccess)
}

func GetPostDetailHandler(c *gin.Context) {
	//1.获取帖子的id
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error(" get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
	}
	//2.根据id取得帖子信息
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回数据
	ResponseSuccess(c, data)
	return

}

//获取帖子列表数据
func GetPostListHandler2(c *gin.Context) {
	p := models.ParamPostList{
		Page:  0,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := c.ShouldBind(&p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	data, err := logic.GetPostListNew(&p)
	if err != nil {
		return
	}
	ResponseSuccess(c, data)
	return

}

//获取帖子列表数据
func GetPostListHandler(c *gin.Context) {
	//1.获取数据
	offset, limit := getPageInfo(c)
	data, err := logic.GetPostList(offset, limit)
	if err != nil {
		zap.L().Error("GetPostListHandler failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	return

}
