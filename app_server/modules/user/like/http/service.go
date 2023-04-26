package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"simple-video-server/app_server/modules/user/like"
	"simple-video-server/app_server/modules/video"
	"simple-video-server/constants/like_type"
	"simple-video-server/core"
	"simple-video-server/db"
)

type Service struct {
	likeCache   *like.Cache
	videoCache  *video.Cache
	redisClient *redis.Client
}

var _service = &Service{
	likeCache:   like.GetCache(),
	videoCache:  video.GetCache(),
	redisClient: db.GetRedisClient(),
}

func GetLikeService() *Service {
	return _service
}

func (s *Service) panicIfErr(c *core.Context, err error, msg string) {
	if err == nil {
		return
	}

	logger := c.Log.With(zap.String("info", msg))
	logger.Error("err", zap.Error(err))

	panic(err)
}

func (s *Service) _checkIsSameLikeType(uid, vid uint, likeType like_type.LikeType) (bool, int) {
	_prevLikeType, _ := s.likeCache.GetLikeTypeByUserAndVideo(uid, vid)

	// -> 提交的操作likeType与记录的likeType一致，不需要继续处理直接return
	if likeType.Is(_prevLikeType) {

		//panic(err)
		return true, _prevLikeType
	}

	return false, _prevLikeType
}

// Add 用户点赞过的video
// video总点赞数
// TODO: 是否需要记录某个video被哪些用户点赞(踩)过
func (s *Service) Add(c *core.Context, uid uint, vid uint, like *like.VideoLike) error {
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

	isSameLikeType, _prevLikeType := s._checkIsSameLikeType(uid, vid, like.LikeType)
	if isSameLikeType {
		err := errors.New(fmt.Sprintf("操作记录与存在记录相同，不需要处理, 当前likeType: %d, vid: %d, like_type: %d,  ", _prevLikeType, vid, like.LikeType.Code))

		panic(err)
	}

	// 用户like记录+1
	err := s.likeCache.AddUserLike(redisTx, uid, vid, like.LikeType.Code)
	s.panicIfErr(c, err, "用户like记录+1")

	switch true {
	//提交操作为点赞
	case like_type.Like.Is(like.LikeType.Code):

		// 加1 video点赞数
		err = s.videoCache.IncreaseVideoLikeCount(redisTx, vid)
		s.panicIfErr(c, err, "加1 video点赞count")

		// 过去的记录为点踩, 点踩数-1
		if like_type.Dislike.Is(_prevLikeType) {
			// 减1 video点踩数
			err = s.videoCache.DecreaseDislikeCount(redisTx, uid, vid)
			s.panicIfErr(c, err, "减1 video点踩count")
		}

		// 提交操作为 点踩
	case like_type.Dislike.Is(like.LikeType.Code):
		// 加1 video点踩count
		err = s.videoCache.IncreaseDislikeCount(redisTx, uid, vid)
		s.panicIfErr(c, err, "加1 video点踩count")

		// 过去的记录为点赞, 点赞数-1
		if like_type.Like.Is(_prevLikeType) {
			// 减1 video点赞count
			err = s.videoCache.DecreaseVideoLikeCount(redisTx, vid)
			s.panicIfErr(c, err, "减1 video点赞count")
		}
	}

	//panic(errors.New("手动抛出错误"))

	//return errors.New("手动抛出错误")
	return nil
}

// Delete 删除点赞、点踩
func (s *Service) Delete(c *core.Context, uid uint, vid uint, like *like.VideoLike) error {
	tx := s.redisClient.TxPipeline()

	defer func() {
		err := recover()
		if err == nil {
			_, err := tx.Exec(context.Background())
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}()

	//uid := *c.AuthUID
	//vid := like.VID
	// TODO: redis使用事务
	// TODO: 使用枚举
	// uid和vid获取点赞记录
	existLikeType, _ := s.likeCache.GetLikeTypeByUserAndVideo(uid, vid)

	// 之前没有点赞记录(却调用了删除点赞记录的接口),
	if existLikeType == 0 {
		return nil
	}

	// 删除user(uid)点赞的video(vid)记录
	err := s.likeCache.DeleteUserLike(tx, uid, vid)
	s.panicIfErr(c, err, fmt.Sprintf("删除用户like记录， uid: %d, vid: %d", uid, vid))

	switch {
	// 之前是点赞: 点赞数-1
	case like_type.Like.Is(existLikeType):
		err = s.videoCache.DecreaseVideoLikeCount(tx, vid)
		s.panicIfErr(c, err, fmt.Sprintf("减1 video like count , uid: %d, vid: %d", uid, vid))

	// 之前是点踩: 点踩数-1
	case like_type.Dislike.Is(existLikeType):
		err = s.videoCache.DecreaseDislikeCount(tx, uid, vid)
		s.panicIfErr(c, err, fmt.Sprintf("减1 video dislike count, uid: %d, vid:%d", uid, vid))
	}

	return err
}
