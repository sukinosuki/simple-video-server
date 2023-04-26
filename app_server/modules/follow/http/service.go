package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"simple-video-server/app_server/modules/follow"
	"simple-video-server/app_server/modules/user/user"
	"simple-video-server/common"
	"simple-video-server/core"
	"simple-video-server/db"
	"simple-video-server/models"
	"simple-video-server/pkg/arr"
	"sort"
	"strconv"
)

type _service struct {
	redisClient *redis.Client
	db          *gorm.DB
	followCache *follow.Cache
	userDao     *user.Dao
}

var Service = &_service{
	redisClient: db.GetRedisClient(),
	db:          db.GetOrmDB(),
	followCache: follow.GetCache(),
	userDao:     user.GetDao(),
}

func _generateFollowingKeyByUID(uid uint) string {
	key := fmt.Sprintf("user:%d:following", uid)

	return key
}

func _generateFollowerKeyByUID(uid uint) string {
	key := fmt.Sprintf("user:%d:follower", uid)

	return key
}

// TODO:抽离cache层

// Follow 添加一个关注
func (s *_service) Follow(c *core.Context, targetUid uint) error {
	handlerName := "Follow"

	redisTx := s.redisClient.TxPipeline()
	defer func() {
		err := recover()

		if err == nil {
			_, _err := redisTx.Exec(context.Background())

			if _err != nil {
				panic(_err)
			}
		} else {
			panic(err)
		}
	}()

	// TODO: 是否已关注
	//followingKey := _generateFollowingKeyByUID(*c.AuthUID)
	//followerKey := _generateFollowerKeyByUID(targetUid)
	//
	//exists, err := s.redisClient.SIsMember(context.Background(), followingKey, targetUid).Result()
	//c.PanicIfErr(err, handlerName, "判断是否是成员失败")
	isFollowingTargetUser, err := s.followCache.IsUserFollowingAnotherUser(*c.AuthUID, targetUid)
	c.PanicIfErr(err, handlerName, "判断是否是成员失败")

	if isFollowingTargetUser {
		c.Panic(errors.New("操作不允许, 已关注过该用户"), handlerName, "重复的关注操作")
	}

	_, err = s.followCache.AddFollowing(redisTx, *c.AuthUID, targetUid)
	c.PanicIfErr(err, handlerName, "缓存添加关注失败")

	// targetUid的粉丝set加上uid
	_, err = s.followCache.AddFollower(redisTx, *c.AuthUID, targetUid)
	c.PanicIfErr(err, handlerName, "目录用户粉丝set添加粉丝uid失败")

	// targetUid 用户的粉丝数自增1
	_, err = s.followCache.IncreaseFollowerCount(redisTx, targetUid)
	c.PanicIfErr(err, handlerName, "用户的粉丝数自增1失败")

	c.Info("关注用户", handlerName)

	return nil
}

// Unfollow 取消一个关注
func (s *_service) Unfollow(c *core.Context, targetUid uint) error {
	handlerName := "Unfollow"
	redisTx := s.redisClient.TxPipeline()

	defer func() {
		err := recover()
		if err == nil {
			_, _err := redisTx.Exec(context.Background())
			if _err != nil {
				c.Panic(_err, handlerName, "redis提交事务失败")
			}
		} else {
			panic(err)
		}
	}()

	followingKey := _generateFollowingKeyByUID(*c.AuthUID)
	followerKey := _generateFollowerKeyByUID(targetUid)

	// 用户是否已经关注目标用户
	exists, err := s.redisClient.SIsMember(context.Background(), followerKey, *c.AuthUID).Result()

	if err != nil {
		return err
	}

	if !exists {
		//TODO
		panic(errors.New("你还不是该用户的粉丝"))
	}

	//用户的关注列表删除目标用户
	_, err = redisTx.SRem(context.Background(), followingKey, targetUid).Result()
	c.PanicIfErr(err, handlerName, "用户的关注列表删除目标用户失败")

	//目标用户的关注列表删除用户
	_, err = redisTx.SRem(context.Background(), followerKey, *c.AuthUID).Result()
	c.PanicIfErr(err, handlerName, "目标用户的关注列表删除用户失败")

	//score的目标用户粉丝数-1
	_, err = s.followCache.DecreaseFollowerCount(redisTx, targetUid)
	c.PanicIfErr(err, handlerName, "score的目标用户粉丝数-1")

	c.Info("用户取消关注", handlerName)

	return nil
}

// FollowScores 获取 follow 排名
func (s *_service) FollowScores(c *core.Context, query *follow.UserFollowRankQuery) ([]*follow.UserFollowerRankResSimple, error) {
	handlerName := "FollowScores"

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
	//result, err := s.redisClient.ZRevRangeWithScores(context.Background(), scoreKey, start, end).Result()
	result, err := s.followCache.GetFollowerCountRank(start, end)
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
	//err = s.db.Model(&models.User{}).Where("id in ?", ids).Find(&users).Error
	_users, err := s.userDao.GetByIdIn(ids, &common.Pager{Page: 1, Size: len(ids)})
	c.PanicIfErr(err, handlerName, "获取用户失败")

	users := arr.Map(_users, func(item models.User, index int) *follow.UserFollowerRankResSimple {

		return &follow.UserFollowerRankResSimple{
			ID:       item.ID,
			Avatar:   item.Avatar,
			Nickname: item.Nickname,
		}
	})

	c.PanicIfErr(err, handlerName, "获取用户列表失败")

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

	targetUserFollowerKey := _generateFollowerKeyByUID(userId)
	selfFollowerKey := _generateFollowerKeyByUID(*c.AuthUID)

	var ids []string
	var err error
	if query.IsInter {
		ids, err = s.redisClient.SInter(context.Background(), targetUserFollowerKey, selfFollowerKey).Result()
	} else {

		ids, err = s.redisClient.SMembers(context.Background(), targetUserFollowerKey).Result()
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
	targetUserFollowingKey := _generateFollowingKeyByUID(userId)
	selfFollowingKey := _generateFollowingKeyByUID(*c.AuthUID)

	var ids []string
	var err error
	if query.IsInter {
		ids, err = s.redisClient.SInter(context.Background(), targetUserFollowingKey, selfFollowingKey).Result()
	} else {

		ids, err = s.redisClient.SMembers(context.Background(), targetUserFollowingKey).Result()
	}

	if err != nil {
		return nil, err
	}

	var users []*follow.UserFollowerResSimple
	err = s.db.Model(&models.User{}).Where("id in ?", ids).Find(&users).Error

	return users, err
}
