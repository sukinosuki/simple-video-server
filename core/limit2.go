package core

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"simple-video-server/db"
	"time"
)

const (
	// 请求量限制
	Threshold = 10

	// 时间窗口，单位秒
	Window = 60

	// 子窗口大小，单位秒
	SubWindow = 10
)

type RateLimiter struct {
	redisClient *redis.Client
}

// func NewRateLimiter(redisClient *redis.Client) *RateLimiter {
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		redisClient: db.GetRedisClient(),
	}
}

func (r *RateLimiter) SlidingWindowTryAcquire(userID string) bool {
	// 计算当前时间在哪个窗口
	currentTime := time.Now().Unix()
	currentWindow := currentTime / SubWindow * SubWindow

	ctx := context.Background()
	// Redis 增加该用户当前窗口的计数
	_, err := r.redisClient.IncrBy(ctx, fmt.Sprintf("%s:%d", userID, currentWindow), 1).Result()
	if err != nil {
		panic(err)
	}

	// 删除过期的子窗口计数器
	startTime := currentWindow - (Window/SubWindow-1)*SubWindow
	for i := 0; i < Window/SubWindow-1; i++ {
		r.redisClient.Del(ctx, fmt.Sprintf("%s:%d", userID, startTime))
		startTime += SubWindow
	}

	// 统计当前窗口的请求数
	currentWindowCount := 0
	for i := 0; i < Window/SubWindow; i++ {
		val, err := r.redisClient.Get(ctx, fmt.Sprintf("%s:%d", userID, currentWindow-int64(SubWindow*i))).Int64()
		if err == redis.Nil {
			continue
		} else if err != nil {
			panic(err)
		}
		currentWindowCount += int(val)
	}

	// 判断是否超过阀值
	if currentWindowCount > Threshold {
		return false
	}

	return true
}
