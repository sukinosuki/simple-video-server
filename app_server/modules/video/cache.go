package video

import (
	"context"
	"fmt"
	"simple-video-server/pkg/global"
	"strconv"
)

type Cache struct {
}

var _video = &Cache{}

func GetCache() *Cache {
	return _video
}

// get like count key
func (vc *Cache) getVideoLikeCountKey(vid uint) string {
	videoLikeCountKey := fmt.Sprintf("video:%d:like_count", vid)
	return videoLikeCountKey
}

// get dislike count key
func (vc *Cache) getVideoDislikeCountKey(vid uint) string {

	videoDislikeCountKey := fmt.Sprintf("video:%d:dislike_count", vid)

	return videoDislikeCountKey
}

// GetVideoLikeCount video id获取视频点赞数
func (vc *Cache) GetVideoLikeCount(vid uint) (int, error) {

	ctx := context.Background()
	likeCountKey := vc.getVideoLikeCountKey(vid)

	result, err := global.RDB.Get(ctx, likeCountKey).Result()

	if err != nil {
		return 0, err
	}

	count, err := strconv.Atoi(result)

	return count, err
}

// GetVideoDislikeCount video id获取视频点踩数
func (vc *Cache) GetVideoDislikeCount(vid uint) (int, error) {
	ctx := context.Background()
	key := vc.getVideoDislikeCountKey(vid)
	result, err := global.RDB.Get(ctx, key).Result()

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(result)
}

// IncreaseLikeCount 加1视频点赞数
func (vc *Cache) IncreaseLikeCount(uid, vid uint, likeType int) error {
	var ctx = context.Background()

	key := vc.getVideoLikeCountKey(vid)

	err := global.RDB.Incr(ctx, key).Err()

	return err
}

// DecreaseLikeCount 减1视频点赞数
func (vc *Cache) DecreaseLikeCount(uid, vid uint) error {
	var ctx = context.Background()
	key := vc.getVideoLikeCountKey(vid)

	err := global.RDB.Decr(ctx, key).Err()

	return err
}

// IncreaseDislikeCount 加1视频点踩数
func (vc *Cache) IncreaseDislikeCount(uid, vid uint, likeType int) error {
	var ctx = context.Background()

	key := vc.getVideoDislikeCountKey(vid)

	err := global.RDB.Incr(ctx, key).Err()

	return err
}

// DecreaseDislikeCount 减1视频点踩数
func (vc *Cache) DecreaseDislikeCount(uid, vid uint) error {
	var ctx = context.Background()
	key := vc.getVideoDislikeCountKey(vid)

	err := global.RDB.Decr(ctx, key).Err()

	return err
}
