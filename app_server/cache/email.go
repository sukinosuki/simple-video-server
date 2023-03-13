package cache

import (
	"context"
	"fmt"
	"simple-video-server/pkg/global"
	"time"
)

type EmailCache struct {
}

var Email = &EmailCache{}

func getKey(email string, actionType string) string {

	key := fmt.Sprintf("email_code:%s:%s", actionType, email)

	return key
}

func (c *EmailCache) Set(email string, actionType string, value string) error {
	//key := fmt.Sprintf("email_code:%s:%s", actionType, email)
	key := getKey(email, actionType)

	duration := 30 * time.Minute

	_, err := global.RDB.Set(context.Background(), key, value, duration).Result() //TODO: 有效时间配置化

	return err
}

func (c *EmailCache) Get(email string, actionType string) (string, error) {
	key := getKey(email, actionType)

	result, err := global.RDB.Get(context.Background(), key).Result()

	return result, err
}

func (c *EmailCache) Delete(email string, actionType string) error {
	key := getKey(email, actionType)

	num, err := global.RDB.Del(context.Background(), key).Result()
	fmt.Println("删除key ", num)

	return err
}
