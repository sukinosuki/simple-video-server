package redis_util

import (
	"context"
	"encoding/json"
	"simple-video-server/pkg/global"
)

var ctx = context.Background()

func Set[T any](key string, v T) error {

	bytes, err := json.Marshal(v)

	if err != nil {
		return err
	}

	err = global.RDB.Set(ctx, key, bytes, 0).Err()

	return err
}

func Get[T any](key string) (*T, error) {
	result, err := global.RDB.Get(ctx, key).Result()

	if err != nil {
		return nil, err
	}

	var v T

	err = json.Unmarshal([]byte(result), &v)

	return &v, err
}
