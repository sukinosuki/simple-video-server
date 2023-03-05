package like

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"simple-video-server/app_server/cache"
	"simple-video-server/constants/like_type"
	"simple-video-server/core"
)

type Service struct {
	dao   *Dao
	cache *cache.LikeCache
}

var _service *Service

func (s *Service) panicIfErr(c *core.Context, err error, msg string) {
	if err == nil {
		return
	}

	logger := c.Log.With(zap.String("info", msg))
	logger.Error("like service err", zap.Error(err))

	panic(err)
}

func GetLikeService() *Service {
	if _service != nil {
		return _service
	}

	return &Service{
		dao:   GetLikeDao(),
		cache: cache.Like,
	}
}

// Add 用户点赞过的video
// video总点赞数
// TODO: 是否需要记录某个video被哪些用户点赞(踩)过
func (s *Service) Add(c *core.Context, like *VideoLike) error {
	uid := *c.UID
	vid := like.VID

	existLikeType, _ := s.cache.GetLikeTypeByUserAndVideo(uid, vid)

	// 提交的操作likeType与记录的likeType一致，不需要继续处理直接return
	if existLikeType == like.LikeType {

		err := errors.New(fmt.Sprintf("操作记录与存在记录相同，不需要处理, 当前likeType: %d, vid: %d, like_type: %d,  ", existLikeType, like.VID, like.LikeType))
		// 参数传错
		panic(err)
	}

	// 用户like记录+1
	err := s.cache.AddUserLike(uid, vid, like.LikeType)
	s.panicIfErr(c, err, "用户like记录+1")

	switch true {
	//提交操作为点赞
	//case like_type.Like.Code:
	case like_type.Like.Is(like.LikeType):

		// 加1 video点赞count
		err = cache.Video.IncreaseLikeCount(uid, vid, like.LikeType)
		s.panicIfErr(c, err, "加1 video点赞count")

		// 过去的记录为点踩, 点踩数-1
		//if existLikeType == like_type.Dislike.Code {
		if like_type.Dislike.Is(existLikeType) {
			// 减1 video点踩count
			err = cache.Video.DecreaseDislikeCount(uid, vid)
			s.panicIfErr(c, err, "减1 video点踩count")
		}

		// 提交操作为 点踩
	case like_type.Dislike.Is(like.LikeType):

		// 加1 video点踩count
		err = cache.Video.IncreaseDislikeCount(uid, vid, like.LikeType)
		s.panicIfErr(c, err, "加1 video点踩count")

		// 过去的记录为点赞, 点赞数-1
		//if existLikeType == like_type.Like.Code {
		if like_type.Like.Is(existLikeType) {
			// 减1 video点赞count
			err = cache.Video.DecreaseLikeCount(uid, vid)
			s.panicIfErr(c, err, "减1 video点赞count")
		}
	}

	return nil
}

// Delete 删除点赞、点踩
func (s *Service) Delete(c *core.Context, like *VideoLike) error {

	uid := *c.UID
	vid := like.VID
	existLikeType, _ := s.cache.GetLikeTypeByUserAndVideo(uid, vid)

	// 之前没有点赞记录(却调用了删除点赞记录的接口),
	if existLikeType == 0 {
		return nil
	}

	err := s.cache.DeleteUserLike(uid, vid)
	s.panicIfErr(c, err, fmt.Sprintf("删除用户like记录， uid: %d, vid: %d", uid, vid))

	switch {
	case like_type.Like.Is(existLikeType):
		err = cache.Video.DecreaseLikeCount(uid, vid)
		s.panicIfErr(c, err, fmt.Sprintf("减1 video like count , uid: %d, vid: %d", uid, vid))

	case like_type.Dislike.Is(existLikeType):
		err = cache.Video.DecreaseDislikeCount(uid, vid)
		s.panicIfErr(c, err, fmt.Sprintf("减1 video dislike count, uid: %d, vid:%d", uid, vid))
	}

	return nil
}
