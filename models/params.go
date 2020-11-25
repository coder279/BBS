package models

//定义请求参数的结构体
const (
	OrderTime = "time"
	OrderScore = "score"
)

type ParamSignUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostID string `json:"post_id" binding:"required"` //帖子id
	Direction int8 `json:"direction,string" binding:"required,oneof=1 0 -1"` //赞成票(1) 反对票(-1)
}
//获取帖子列表参数
type ParamPostList struct {
	CommunityID int64 `json:"community_id" form:"community_id"`
	Page int64 `json:"page" form:"page"`
	Size int64 `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}

//根据社区获取帖子
type ParamCommunityPostList struct {
	*ParamPostList
}