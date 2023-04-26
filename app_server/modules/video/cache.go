package video

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"simple-video-server/db"
	"simple-video-server/pkg/util"
	"strconv"
)

type Cache struct {
	client *redis.Client
}

var _video = &Cache{
	client: db.GetRedisClient(),
}

func GetCache() *Cache {
	return _video
}

// vid生成视频点赞key
func _generateVideoLikeCountKeyByVid(vid uint) string {
	videoLikeCountKey := fmt.Sprintf("video:%d:like_count", vid)

	return videoLikeCountKey
}

// vid生成视频点踩key
func _generateVideoDislikeCountKeyByVid(vid uint) string {
	videoDislikeCountKey := fmt.Sprintf("video:%d:dislike_count", vid)

	return videoDislikeCountKey
}

//func parseJson[T any](data string) (*T, error) {
//
//
//}

// GetVideoLikeCount video id获取视频点赞数
func (vc *Cache) GetVideoLikeCount(vid uint) (int, error) {
	ctx := context.Background()
	key := _generateVideoLikeCountKeyByVid(vid)

	result, err := vc.client.Get(ctx, key).Result()

	if err != nil {
		return 0, err
	}

	// TODO: 使用泛型获取对应格式
	count, err := strconv.Atoi(result)

	count2, err := util.ParseJson[int](result)
	if err != nil {
		return 0, err
	}

	fmt.Println("count2 ", count2)

	return count, err
}

// GetVideoDislikeCount video id获取视频点踩数
func (vc *Cache) GetVideoDislikeCount(vid uint) (int, error) {
	ctx := context.Background()
	key := _generateVideoDislikeCountKeyByVid(vid)

	result, err := vc.client.Get(ctx, key).Result()

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(result)
}

// IncreaseVideoLikeCount 加1视频点赞数
func (vc *Cache) IncreaseVideoLikeCount(tx redis.Pipeliner, vid uint) error {
	var ctx = context.Background()

	key := _generateVideoLikeCountKeyByVid(vid)

	err := tx.Incr(ctx, key).Err()

	return err
}

// DecreaseVideoLikeCount 减1视频点赞数
func (vc *Cache) DecreaseVideoLikeCount(tx redis.Pipeliner, vid uint) error {
	var ctx = context.Background()
	key := _generateVideoLikeCountKeyByVid(vid)

	err := tx.Decr(ctx, key).Err()

	return err
}

// IncreaseDislikeCount 加1视频点踩数
func (vc *Cache) IncreaseDislikeCount(tx redis.Pipeliner, uid, vid uint) error {
	var ctx = context.Background()

	key := _generateVideoDislikeCountKeyByVid(vid)

	err := tx.Incr(ctx, key).Err()

	return err
}

// DecreaseDislikeCount 减1视频点踩数
func (vc *Cache) DecreaseDislikeCount(tx redis.Pipeliner, uid, vid uint) error {
	var ctx = context.Background()
	key := _generateVideoDislikeCountKeyByVid(vid)

	err := tx.Decr(ctx, key).Err()

	return err
}
