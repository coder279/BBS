package redis

import (
	"BBS/models"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

func GetPostIDsInOrder(p *models.ParamPostList)([]string ,error){
	key := getRedisKey(KeyPostTimeZset)
	if p.Order == models.OrderScore{
		key = getRedisKey(KeyPostScoreZset)
	}
	return getIDsFormKey(key,p.Page,p.Size)
}
//根据ids查询每篇梯子的投票赞成数
func GetPostVoteData(ids []string)(data []int64 ,err error){
	//data = make([]int64,0,len(ids))
	//for _,id := range ids{
	//	key := getRedisKey(KeyPostVotedZsetPrefix + id)
	//	v := rdb.ZCount(key,"1","1").Val()
	//	data = append(data,v)
	//	rdb.ZCount(key,"1","1").Val()
	//}
	//使用Pipline一次发送多条数据
	keys := make([]string,0,len(ids))
	pipline := rdb.Pipeline()
	for _,id := range ids{
		key := getRedisKey(KeyPostVotedZsetPrefix + id)
		pipline.ZCount(key,"1","1")
		keys = append(keys,key)
	}
	cmders,err := pipline.Exec()
	if err != nil {
		return nil,err
	}
	data = make([]int64,0,len(cmders))
	for _,cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data,v)
	}
	return
}
func getIDsFormKey(key string,page,size int64)([]string,error){
	//2. 确定查询的索引起始点
	start := (page -1) * size
	end := start + size - 1
	return rdb.ZRevRange(key,start,end).Result()
}

func GetCommunityPostIDsInOrder(p *models.ParamPostList)([]string,error){
	orderKey := getRedisKey(KeyPostTimeZset)
	if p.Order == models.OrderScore{
		orderKey = getRedisKey(KeyPostScoreZset)
	}
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	ckey := getRedisKey(KeyCommunitySetPF+strconv.Itoa(int(p.CommunityID)))
	if rdb.Exists(orderKey).Val() < 1 {
		pipline := rdb.TxPipeline()
		pipline.ZInterStore(key,redis.ZStore{
			Aggregate:"MAX",
		},ckey,orderKey)
		pipline.Expire(key,60*time.Second)
		_,err := pipline.Exec()
		if err != nil {
			return nil,err
		}
	}
	return getIDsFormKey(key,p.Page,p.Size)
}