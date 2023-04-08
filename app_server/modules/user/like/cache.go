package like

import (
	"context"
	"fmt"
	"simple-video-server/pkg/global"
	"strconv"
)

type Cache struct {
}

var _likeCache = &Cache{}

func GetCache() *Cache {
	return _likeCache
}

// uid获取key: user:1:like_video
func (ca *Cache) getUserLikeVideoKey(uid uint) string {
	likeKey := fmt.Sprintf("user:%d:like_video", uid)

	return likeKey
}

// uid获取key: user:1:like_video
func getUserLikeVideoKey(uid uint) string {
	likeKey := fmt.Sprintf("user:%d:like_video", uid)

	return likeKey
}

// AddUserLike 增加用户点赞(点踩)
func (ca *Cache) AddUserLike(uid, vid uint, likeType int) error {
	var ctx = context.Background()

	// field为video id
	field := strconv.Itoa(int(vid))

	likeKey := getUserLikeVideoKey(uid)

	err := global.RDB.HSet(ctx, likeKey, field, likeType).Err()

	return err
}

// GetLikeTypeByUserAndVideo 获取当前某个用户对某个视频的like type
// 返回0: 用户对视频没有过点赞|点踩记录
// 返回1: 用户对视频like操作为点赞
// 返回2: 用户对视频like操作为点踩
// @return
func (ca *Cache) GetLikeTypeByUserAndVideo(uid, vid uint) (int, error) {
	var ctx = context.Background()

	likeKey := getUserLikeVideoKey(uid)

	field := strconv.Itoa(int(vid))

	// result是字符串
	result, err := global.RDB.HGet(ctx, likeKey, field).Result()

	if err == nil {
		likeType, err2 := strconv.Atoi(result)

		return likeType, err2
	}

	return 0, err
}

// DeleteUserLike 删除用户点赞
func (ca *Cache) DeleteUserLike(uid, vid uint) error {
	var ctx = context.Background()
	likeKey := getUserLikeVideoKey(uid)
	field := strconv.Itoa(int(vid))

	err := global.RDB.HDel(ctx, likeKey, field).Err()

	return err
}
