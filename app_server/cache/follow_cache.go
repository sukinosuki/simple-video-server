package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"simple-video-server/db"
)

type FollowCache struct {
	cache *redis.Client
}

var Follow = &FollowCache{
	cache: db.GetRedisDB(),
}

// GetUserFollowingKey 获取用户的关注key
func (c *FollowCache) GetUserFollowingKey(uid uint) string {

	key := fmt.Sprintf("user:%d:following", uid)

	return key
}

// GetUserFollowerKey 获取用户的粉丝key
func (c *FollowCache) GetUserFollowerKey(uid uint) string {

	key := fmt.Sprintf("user:%d:follower", uid)

	return key
}

//func (c *FollowCache) GetUserFollowers(uid uint)  {
//
//	key:=c.GetUserFollowerKey(uid)
//
//	global.RDB.
//}

func (c *FollowCache) IsFollowingOneUser(uid, targetUid uint) (bool, error) {
	//followingKey := getFollowingKey(*c.UID)
	//followerKey := getFollowerKey(targetUID)
	followingKey := c.GetUserFollowingKey(uid)

	result, err := c.cache.SIsMember(context.Background(), followingKey, targetUid).Result()

	return result, err
}

func (c *FollowCache) OneUserFollowersCount(uid uint) (int64, error) {
	followerKey := c.GetUserFollowerKey(uid)
	result, err := c.cache.SCard(context.Background(), followerKey).Result()

	return result, err
}

func (c *FollowCache) OneUserFollowingCount(uid uint) (int64, error) {
	followingKey := c.GetUserFollowingKey(uid)
	return c.cache.SCard(context.Background(), followingKey).Result()
}
