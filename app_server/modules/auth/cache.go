package auth

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"simple-video-server/db"
	"simple-video-server/models"
	"simple-video-server/pkg/redis_util"
	"time"
)

type UserCache struct {
	client *redis.Client
}

var _userCache = &UserCache{
	client: db.GetRedisClient(),
}

func GetUserCache() *UserCache {
	return _userCache
}

func _generateUserCacheKeyByUID(uid uint) string {
	key := fmt.Sprintf("user:%d:info", uid)

	return key
}

func (c *UserCache) GetUser(uid uint) (*models.User, error) {
	key := _generateUserCacheKeyByUID(uid)

	user, err := redis_util.Get[models.User](key)

	return user, err
}

func (c *UserCache) SetUser(uid uint, user *models.User, expiration time.Duration) error {
	key := _generateUserCacheKeyByUID(uid)

	err := redis_util.Set(key, user, expiration)

	return err
}

func (c *UserCache) Delete(tx redis.Pipeliner, uid uint) error {
	key := _generateUserCacheKeyByUID(uid)

	err := tx.Del(context.Background(), key).Err()

	return err
}
