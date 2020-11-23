package logic

import (
	"BBS/dao/redis"
	"BBS/models"
	"strconv"
)

func VoteForPost(userID int64,p *models.ParamVoteData)error{
	return redis.VoteForPost(strconv.Itoa(int(userID)), string(p.PostID),float64(p.Direction))
}
