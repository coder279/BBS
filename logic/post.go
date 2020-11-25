package logic

import (
	"BBS/dao/mysql"
	"BBS/dao/redis"
	"BBS/models"
	"BBS/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error){
	p.ID = int64(snowflake.GenID())
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID,p.CommunityID)
	return
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

func GetPostList2(list *models.ParamPostList)(data []*models.ApiPostDetail,err error){
	ids,err := redis.GetPostIDsInOrder(list)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostList2 return 0 data")
		return
	}
	//根据ids去数据库查询帖子操作
	posts,err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	//提前查询好每个帖子的投票数
	voteData,err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	for idx,post := range posts{
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
			VoteNum: voteData[idx],
			Post:	post,
			CommunityDetail:	community,
		}
		data = append(data,postDetail)
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

func GetCommunityPostList(p *models.ParamPostList)(data []*models.ApiPostDetail,err error){
	ids,err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostList2 return 0 data")
		return
	}
	//根据ids去数据库查询帖子操作
	posts,err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	//提前查询好每个帖子的投票数
	voteData,err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	for idx,post := range posts{
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
			VoteNum: voteData[idx],
			Post:	post,
			CommunityDetail:	community,
		}
		data = append(data,postDetail)
	}
	return
}

func GetPostListNew(p *models.ParamPostList)(data []*models.ApiPostDetail,err error){
	if p.CommunityID == 0 {
		data,err = GetPostList2(p)

	}else{
		data,err = GetCommunityPostList(p)
	}
	if err != nil {
		zap.L().Error("logic GetPostListNew failed",zap.Error(err))
		return nil,err
	}
	return
}

