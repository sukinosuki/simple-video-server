package follow

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"simple-video-server/db"
	"strconv"
)

type Cache struct {
	redisClient *redis.Client
}

var _cache = &Cache{
	redisClient: db.GetRedisClient(),
}

func GetCache() *Cache {
	return _cache
}

const scoreKey = "fans_count_statistic"

// _generateUserFollowingKeyByUID 获取用户的关注key
func _generateUserFollowingKeyByUID(uid uint) string {

	key := fmt.Sprintf("user:%d:following", uid)

	return key
}

// _generateUserFollowerKeyByUID 获取用户的粉丝key
func _generateUserFollowerKeyByUID(uid uint) string {

	key := fmt.Sprintf("user:%d:follower", uid)

	return key
}

// IsUserFollowingAnotherUser 用户是否关注了某个用户
func (cache *Cache) IsUserFollowingAnotherUser(uid, targetUid uint) (bool, error) {
	followingKey := _generateUserFollowingKeyByUID(uid)

	// SIsMember 判断元素member是否在集合set中
	result, err := cache.redisClient.SIsMember(context.Background(), followingKey, targetUid).Result()

	return result, err
}

// GetOneUserFollowersCount 获取用户粉丝数
func (cache *Cache) GetOneUserFollowersCount(uid uint) (int64, error) {
	followerKey := _generateUserFollowerKeyByUID(uid)
	// SCard 获取集合set元素个数
	result, err := cache.redisClient.SCard(context.Background(), followerKey).Result()

	return result, err
}

// GetOneUserFollowingCount 获取某个用户关注数
func (cache *Cache) GetOneUserFollowingCount(uid uint) (int64, error) {
	followingKey := _generateUserFollowingKeyByUID(uid)

	return cache.redisClient.SCard(context.Background(), followingKey).Result()
}

// AddFollowing 增加关注
func (cache *Cache) AddFollowing(tx redis.Pipeliner, uid, targetUid uint) (int64, error) {
	key := _generateUserFollowingKeyByUID(uid)

	return tx.SAdd(context.Background(), key, targetUid).Result()
}

// AddFollower 增加粉丝
func (cache *Cache) AddFollower(tx redis.Pipeliner, uid, targetUid uint) (int64, error) {
	key := _generateUserFollowerKeyByUID(targetUid)

	return tx.SAdd(context.Background(), key, uid).Result()
}

// IncreaseFollowerCount 粉丝数+1
func (cache *Cache) IncreaseFollowerCount(tx redis.Pipeliner, userId uint) (float64, error) {
	return tx.ZIncrBy(context.Background(), scoreKey, 1, strconv.Itoa(int(userId))).Result()
}

// DecreaseFollowerCount 粉丝数-1
func (cache *Cache) DecreaseFollowerCount(tx redis.Pipeliner, userId uint) (float64, error) {
	return tx.ZIncrBy(context.Background(), scoreKey, -1, strconv.Itoa(int(userId))).Result()
}

// GetFollowerCountRank 获取粉丝数排名
func (cache *Cache) GetFollowerCountRank(start, end int64) ([]redis.Z, error) {
	return cache.redisClient.ZRevRangeWithScores(context.Background(), scoreKey, start, end).Result()
}
