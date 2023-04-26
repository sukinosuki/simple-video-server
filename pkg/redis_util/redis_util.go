package redis_util

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"simple-video-server/db"
	"time"
)

var client *redis.Client

func init() {
	client = db.GetRedisClient()
}

func Set[T any](key string, v T, expiration time.Duration) error {
	var ctx = context.Background()

	bytes, err := json.Marshal(v)

	if err != nil {
		return err
	}

	err = client.Set(ctx, key, bytes, expiration).Err()

	return err
}

func Get[T any](key string) (*T, error) {
	var ctx = context.Background()
	result, err := client.Get(ctx, key).Result()

	if err != nil {
		return nil, err
	}

	var v T

	err = json.Unmarshal([]byte(result), &v)

	return &v, err
}
