package cache

import "fmt"

type FollowCache struct{}

var Follow = &FollowCache{}

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
