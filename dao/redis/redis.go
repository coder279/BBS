package redis

import (
	"BBS/settings"
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var (
	rdb *redis.Client
	Nil = redis.Nil
	)

const (
	KeyPrefix = "bbs:"
	KeyPostTimeZset = "post:time"
	KeyPostScoreZset = "post:score"
	KeyPostVotedZsetPrefix = "post:voted:"
	KeyCommunitySetPF = "community:"
	)

// 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	_, err = rdb.Ping().Result()
	return
}

func Close(){
	_ = rdb.Close()
}

func getRedisKey(key string)string{
	return KeyPrefix + key
}