package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"simple-video-server/app_server/modules/follow"
	"simple-video-server/core"
	"simple-video-server/db"
	"simple-video-server/models"
	"simple-video-server/pkg/arr"
	"sort"
	"strconv"
)

type _service struct {
	cache *redis.Client
	db    *gorm.DB
}

var Service = &_service{
	cache: db.GetRedisDB(),
	db:    db.GetOrmDB(),
}

const scoreKey = "fans_count_statistic"

func getFollowingKey(uid uint) string {
	key := fmt.Sprintf("user:%d:following", uid)

	return key
}

func getFollowerKey(uid uint) string {
	key := fmt.Sprintf("user:%d:follower", uid)

	return key
}

// TODO:抽离cache层

// Follow 添加一个关注
func (s *_service) Follow(c *core.Context, targetUid uint) error {

	// TODO: 是否已关注
	followingKey := getFollowingKey(*c.AuthUID)
	followerKey := getFollowerKey(targetUid)

	exists, err := s.cache.SIsMember(context.Background(), followingKey, targetUid).Result()

	if err != nil {
		panic(err)
	}

	if exists {
		//TODO
		panic(errors.New("操作不允许, 已关注过该用户"))
	}

	// uid的关注set加上targetUid
	_, err = s.cache.SAdd(context.Background(), followingKey, targetUid).Result()
	// TODO: 事务
	if err != nil {
		return err
	}

	// targetUid的粉丝set加上uid
	_, err = s.cache.SAdd(context.Background(), followerKey, *c.AuthUID).Result()
	if err != nil {
		return err
	}

	// targetUid 用户的粉丝数自增1
	result2, err := s.cache.ZIncrBy(context.Background(), scoreKey, 1, strconv.Itoa(int(targetUid))).Result()
	if err != nil {
		return err
	}

	fmt.Println("result2 ", result2)
	return nil
}

// Unfollow 取消一个关注
func (s *_service) Unfollow(c *core.Context, targetUid uint) error {
	followingKey := getFollowingKey(*c.AuthUID)
	followerKey := getFollowerKey(targetUid)

	// 用户是否已经关注目标用户
	exists, err := s.cache.SIsMember(context.Background(), followerKey, *c.AuthUID).Result()

	if err != nil {
		return err
	}

	if !exists {
		//TODO
		panic(errors.New("你还不是该用户的粉丝"))
	}

	//用户的关注列表删除目标用户
	result, err := s.cache.SRem(context.Background(), followingKey, targetUid).Result()
	if err != nil {
		return err
	}

	fmt.Println("result ", result)

	//目标用户的关注列表删除用户
	result, err = s.cache.SRem(context.Background(), followerKey, *c.AuthUID).Result()
	if err != nil {
		return err
	}

	fmt.Println("result ", result)

	//score的目标用户粉丝数-1
	result2, err := s.cache.ZIncrBy(context.Background(), scoreKey, -1, strconv.Itoa(int(targetUid))).Result()
	if err != nil {
		return err
	}

	fmt.Println("result2 ", result2)
	return nil
}

// FollowScores 获取 follow 排名
func (s *_service) FollowScores(c *core.Context, query *follow.UserFollowRankQuery) (users []*follow.UserFollowerRankResSimple, err error) {

	var start int64 = 0
	var end int64 = 10

	// 校验range
	if query.Range != nil {
		// 如果只传了一个值, 该值默认做为start
		if len(query.Range) == 1 {
			end = query.Range[0] + 10

			// 传了两个值或以上
		} else {
			start = query.Range[0]
			end = query.Range[1]
		}
	}

	// 限制不允许查询超过start 100的排名
	if end-start > 100 {
		end = start + 100
	}

	// redis 从高到低获取排名
	result, err := s.cache.ZRevRangeWithScores(context.Background(), scoreKey, start, end).Result()
	if err != nil {
		return nil, err
	}

	// 排名列表长度为0, 直接返回
	if len(result) == 0 {
		return nil, nil
	}

	var ids []uint
	var scoreMap = make(map[uint]float64)

	for _, v := range result {
		idStr, ok := v.Member.(string)
		if !ok {
			panic(errors.New(fmt.Sprintf("id不合法, id = %s", v.Member)))
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
		uintId := uint(id)

		ids = append(ids, uintId)
		scoreMap[uintId] = v.Score
	}

	//按score排名前n条得到的user id查询user
	//var users []*UserFollowerRankResSimple
	err = s.db.Model(&models.User{}).Where("id in ?", ids).Find(&users).Error

	//每个user赋值对应的score
	arr.ForEach(users, func(item *follow.UserFollowerRankResSimple, index int) {
		item.Score = scoreMap[item.ID]
	})

	// 按score排序
	sort.Slice(users, func(i, j int) bool {
		return users[i].Score > users[j].Score
	})

	return users, err
}

// GetUserFollower 获取某个用户的粉丝列表
func (s *_service) GetUserFollower(c *core.Context, userId uint, query *follow.UserFollowerQuery) ([]*follow.UserFollowerResSimple, error) {

	targetUserFollowerKey := getFollowerKey(userId)
	selfFollowerKey := getFollowerKey(*c.AuthUID)

	var ids []string
	var err error
	if query.IsInter {
		ids, err = s.cache.SInter(context.Background(), targetUserFollowerKey, selfFollowerKey).Result()
	} else {

		ids, err = s.cache.SMembers(context.Background(), targetUserFollowerKey).Result()
	}

	if err != nil {
		return nil, err
	}

	var users []*follow.UserFollowerResSimple
	err = s.db.Model(&models.User{}).Where("id in ?", ids).Find(&users).Error

	return users, err
}

// GetUserFollowing 获取某个用户的关注列表
func (s *_service) GetUserFollowing(c *core.Context, userId uint, query *follow.UserFollowerQuery) ([]*follow.UserFollowerResSimple, error) {
	targetUserFollowingKey := getFollowingKey(userId)
	selfFollowingKey := getFollowingKey(*c.AuthUID)

	var ids []string
	var err error
	if query.IsInter {
		ids, err = s.cache.SInter(context.Background(), targetUserFollowingKey, selfFollowingKey).Result()
	} else {

		ids, err = s.cache.SMembers(context.Background(), targetUserFollowingKey).Result()
	}

	if err != nil {
		return nil, err
	}

	var users []*follow.UserFollowerResSimple
	err = s.db.Model(&models.User{}).Where("id in ?", ids).Find(&users).Error

	return users, err
}

// IsFollower 自己是否是关注了某个用户
//func (s *_service) IsFollowingOneUser(c *core.Context, targetUID uint) {
//	//exists, err := s.cache.SIsMember(context.Background(), followingKey, targetUid).Result()
//
//}
