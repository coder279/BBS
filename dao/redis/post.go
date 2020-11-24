package redis

import "BBS/models"

func GetPostIDsInOrder(p *models.ParamPostList)([]string ,error){
	key := getRedisKey(KeyPostTimeZset)
	if p.Order == models.OrderScore{
		key = getRedisKey(KeyPostScoreZset)
	}
	//2. 确定查询的索引起始点
	start := (p.Page -1) * p.Size
	end := start + p.Size - 1
	return rdb.ZRevRange(key,start,end).Result()
}