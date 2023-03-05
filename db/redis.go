package db

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"simple-video-server/config"
)

var rdb *redis.Client

func GetRedisDB() *redis.Client {
	return rdb
}

func init() {

	db := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
	})

	fmt.Println("redis init complete")

	rdb = db
}
