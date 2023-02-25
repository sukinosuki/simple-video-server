package user

import (
	"fmt"
	"simple-video-server/models"
	"simple-video-server/pkg/redis_util"
)

type userCache struct {
}

var UserCache = &userCache{}

func (c *userCache) getUserCacheKey(uid uint) string {
	return fmt.Sprintf("user::%d", uid)
}

func (c *userCache) GetUser(uid uint) (*models.User, error) {
	key := c.getUserCacheKey(uid)

	user, err := redis_util.Get[models.User](key)

	fmt.Println("user ", &user)
	return user, err
}

func (c *userCache) SetUser(uid uint, user *models.User) error {

	err := redis_util.Set(c.getUserCacheKey(uid), user)

	return err
}
