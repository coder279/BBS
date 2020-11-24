package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

const (
	oneWeenkInSeconds = 7 * 24 * 3600
	scorePervote = 432
)
var (
	ErrVoteTimeExpire = errors.New("超出投票时间")
)

func CreatePost(id int64)(err error){
	rdb.ZAdd(getRedisKey(KeyPostTimeZset),redis.Z{
		Score:float64(time.Now().Unix()),
		Member:id,
	}).Result()
	return
}

func VoteForPost(userID,postID string ,value float64) (err error) {
	pipline := rdb.TxPipeline()
	//取redis帖子发布时间
	postTime := pipline.ZScore(getRedisKey(KeyPostTimeZset),postID).Val()
	print(postTime)
	if float64(time.Now().Unix()) - float64(postTime) < oneWeenkInSeconds {
		return ErrVoteTimeExpire
	}
	ov := pipline.ZScore(getRedisKey(KeyPostVotedZsetPrefix+postID),userID).Val()
	var dir float64
	if value > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - value)
	pipline.ZIncrBy(getRedisKey(KeyPostScoreZset),dir*diff*scorePervote,postID)
	if value == 0 {
		_, _ = pipline.ZRem(getRedisKey(KeyPostVotedZsetPrefix+postID),userID).Result()
	}else{
		_,_ = pipline.ZAdd(getRedisKey(KeyPostVotedZsetPrefix+postID),redis.Z{
			Score:value,
			Member:userID,
		}).Result()
	}
	return nil

}


