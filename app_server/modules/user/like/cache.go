package like

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"simple-video-server/db"
	"strconv"
)

type Cache struct {
	client *redis.Client
}

var _likeCache = &Cache{
	client: db.GetRedisClient(),
}

func GetCache() *Cache {
	return _likeCache
}

//// uid获取key: user:1:like_video
//func (ca *Cache) getUserLikeVideoKey(uid uint) string {
//	likeKey := fmt.Sprintf("user:%d:like_video", uid)
//
//	return likeKey
//}

// uid获取key, 格式: user:[:id]:like_video
func getUserLikeVideoKey(uid uint) string {
	likeKey := fmt.Sprintf("user:%d:like_video", uid)

	return likeKey
}

// AddUserLike 增加用户点赞(点踩)
func (ca *Cache) AddUserLike(tx redis.Pipeliner, uid, vid uint, likeType int) error {
	var ctx = context.Background()

	// field为video id
	//field := strconv.Itoa(int(vid))

	key := getUserLikeVideoKey(uid)

	// hash: 根据key和field字段设置，field字段的值
	err := tx.HSet(ctx, key, vid, likeType).Err()

	return err
}

// GetLikeTypeByUserAndVideo 获取当前某个用户对某个视频的like type
// 返回0: 用户对视频没有过点赞|点踩记录
// 返回1: 用户对视频like操作为点赞
// 返回2: 用户对视频like操作为点踩
// @return
func (ca *Cache) GetLikeTypeByUserAndVideo(uid, vid uint) (int, error) {
	var ctx = context.Background()

	key := getUserLikeVideoKey(uid)

	field := strconv.Itoa(int(vid))

	// 根据key和field字段，查询field字段的值
	// field没获取到值时, result是"", err为 error类型的字符串"redis: nil"
	result, err := ca.client.HGet(ctx, key, field).Result()

	if err == nil {
		likeType, err2 := strconv.Atoi(result)

		return likeType, err2
	}

	return 0, err
}

// DeleteUserLike 删除用户点赞
func (ca *Cache) DeleteUserLike(tx redis.Pipeliner, uid, vid uint) error {
	var ctx = context.Background()
	likeKey := getUserLikeVideoKey(uid)
	field := strconv.Itoa(int(vid))

	err := tx.HDel(ctx, likeKey, field).Err()

	return err
}
