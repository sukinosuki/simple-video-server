package email

import (
	"context"
	"fmt"
	"simple-video-server/pkg/global"
	"simple-video-server/pkg/redis_util"
	"time"
)

type Cache struct {
}

var cache = &Cache{}

func GetCache() *Cache {
	return cache
}

func _generateKeyByEmailAndActionType(email string, actionType string) string {

	key := fmt.Sprintf("email_code:%s:%s", actionType, email)

	return key
}

func (c *Cache) Set(email string, actionType string, value string) error {
	key := _generateKeyByEmailAndActionType(email, actionType)

	//TODO: 有效时间配置化
	duration := 30 * time.Minute

	err := redis_util.Set(key, value, duration)

	return err
}

func (c *Cache) Get(email string, actionType string) (string, error) {
	key := _generateKeyByEmailAndActionType(email, actionType)

	result, err := redis_util.Get[string](key)
	if err != nil {
		return "", err
	}

	return *result, err
}

func (c *Cache) Delete(email string, actionType string) (int64, error) {
	key := _generateKeyByEmailAndActionType(email, actionType)

	return global.RDB.Del(context.Background(), key).Result()
}
