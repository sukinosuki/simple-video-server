package db

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"simple-video-server/config"
)

var _redisClient *redis.Client

func GetRedisClient() *redis.Client {
	return _redisClient
}

func init() {

	addr := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
	})

	fmt.Println("redis init complete")

	_redisClient = client
}
