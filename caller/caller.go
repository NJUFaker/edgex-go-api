package caller

import (
	"github.com/edgex-go-api/config"
	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
)

func InitClient() {
	initRedisClient()
}

func initRedisClient() {
	redisOpt := &redis.Options{
		Addr:     config.RedisSetting.Address,
		Password: config.RedisSetting.Password,
		DB:       config.RedisSetting.DB,
	}
	RedisClient = redis.NewClient(redisOpt)
}
