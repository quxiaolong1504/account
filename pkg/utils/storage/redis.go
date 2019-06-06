package storage

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/quxiaolong/account/pkg/config"
	"github.com/quxiaolong/account/pkg/utils/logger"
)

var Cache *redis.Client

func InitRedis(conf *config.RedisConfig) {
	logger.Logger.Infof("Init  redis client")
	// init cache redis
	Cache = NewClient(conf.Cache)
	pong, err := Cache.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Print(pong)
}

func NewClient(redisConf *config.RedisConf) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       int(redisConf.DB),
	})
}

func ShutdownRedis() {
	Cache.Close()
}