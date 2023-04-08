package follow

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"simple-video-server/db"
)

type Cache struct {
	cache *redis.Client
}

var _cache = &Cache{
	cache: db.GetRedisDB(),
}

func GetFollowCache() *Cache {
	return _cache
}

// GetUserFollowingKey 获取用户的关注key
func (c *Cache) GetUserFollowingKey(uid uint) string {

	key := fmt.Sprintf("user:%d:following", uid)

	return key
}

// GetUserFollowerKey 获取用户的粉丝key
func (c *Cache) GetUserFollowerKey(uid uint) string {

	key := fmt.Sprintf("user:%d:follower", uid)

	return key
}

func (c *Cache) IsFollowingOneUser(uid, targetUid uint) (bool, error) {
	//followingKey := getFollowingKey(*c.UID)
	//followerKey := getFollowerKey(targetUID)
	followingKey := c.GetUserFollowingKey(uid)

	result, err := c.cache.SIsMember(context.Background(), followingKey, targetUid).Result()

	return result, err
}

func (c *Cache) OneUserFollowersCount(uid uint) (int64, error) {
	followerKey := c.GetUserFollowerKey(uid)
	result, err := c.cache.SCard(context.Background(), followerKey).Result()

	return result, err
}

func (c *Cache) OneUserFollowingCount(uid uint) (int64, error) {
	followingKey := c.GetUserFollowingKey(uid)
	return c.cache.SCard(context.Background(), followingKey).Result()
}
