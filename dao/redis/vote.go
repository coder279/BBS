package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"math"
	"strconv"
	"time"
)

const (
	oneWeenkInSeconds = 7 * 24 * 3600
	scorePervote = 432
)
var (
	ErrVoteTimeExpire = errors.New("超出投票时间")
	ErrVoteRepeat = errors.New("不允许重复投票")
)

func CreatePost(id,communityID int64)(err error){
	pipline := rdb.TxPipeline()
	pipline.ZAdd(getRedisKey(KeyPostTimeZset),redis.Z{
		Score:float64(time.Now().Unix()),
		Member:id,
	}).Result()
	//待完善
	pipline.ZAdd(getRedisKey(KeyPostScoreZset),redis.Z{
		Score:float64(time.Now().Unix()),
		Member:id,
	}).Result()
	_,err = pipline.Exec()
	return
}

func VoteForPost(userID,postID string ,value float64) (err error) {
	pipline := rdb.TxPipeline()
	//取redis帖子发布时间
	postTime := pipline.ZScore(getRedisKey(KeyPostTimeZset),postID).Val()
	zap.L().Info(strconv.FormatFloat(float64(time.Now().Unix()) - float64(postTime),'f',6,64))
	if float64(time.Now().Unix()) - float64(postTime) > oneWeenkInSeconds {
		return ErrVoteTimeExpire
	}
	ov := pipline.ZScore(getRedisKey(KeyPostVotedZsetPrefix+postID),userID).Val()
	var dir float64
	if value == ov{
		return ErrVoteRepeat
	}
	if value > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - value)
	pipline.ZIncrBy(getRedisKey(KeyPostScoreZset),dir*diff*scorePervote,postID)
	if value == 0 {
		pipline.ZRem(getRedisKey(KeyPostVotedZsetPrefix+postID),userID)
	}else{
		pipline.ZAdd(getRedisKey(KeyPostVotedZsetPrefix+postID),redis.Z{
			Score:value,
			Member:userID,
		})
	}
	_,err = pipline.Exec()
	return err

}


