package logic

import (
	"BBS/dao/mysql"
	"BBS/models"
	"BBS/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error){
	p.ID = int64(snowflake.GenID())
	return mysql.CreatePost(p)
}

func GetPostById(pid int64)(data *models.ApiPostDetail,err error){
	post,err := mysql.GetPostById(pid)
	if  err != nil {
		zap.L().Error("mysql GetPostById failed",zap.Error(err))
		return
	}
	user,err := mysql.GetUserById(post.AuthorID)
	if  err != nil {
		zap.L().Error("mysql GetUserById failed",zap.Error(err))
		return
	}
	community,err := mysql.GetCommunityDetailById(post.CommunityID)
	if  err != nil {
		zap.L().Error("mysql GetCommunityById failed",zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorNmae:user.Username,
		Post:post,
		CommunityDetail:community,
	}

	return
}

func GetPostList(offset int64,limit int64)(data []*models.ApiPostDetail,err error){
	posts,err := mysql.GetPostList(offset,limit)
	if err != nil {
		return nil,err
	}
	data = make([]*models.ApiPostDetail,0,len(posts))
	for _,post := range posts{
		user,err := mysql.GetUserById(post.AuthorID)
		if  err != nil {
			zap.L().Error("mysql GetUserById failed",zap.Error(err))
			continue
		}
		community,err := mysql.GetCommunityDetailById(post.CommunityID)
		if  err != nil {
			zap.L().Error("mysql GetCommunityById failed",zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorNmae:	user.Username,
			Post:	post,
			CommunityDetail:	community,
		}
		data = append(data,postDetail)
	}
	return

}

