package http

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"simple-video-server/app_server/modules/user/like"
	"simple-video-server/app_server/modules/video"
	"simple-video-server/constants/like_type"
	"simple-video-server/constants/video_status"
	"simple-video-server/core"
	"simple-video-server/models"
	"simple-video-server/pkg/business_code"
)

type service struct {
	dao       *video.Dao
	cache     *video.Cache
	likeCache *like.Cache
}

var Service = &service{
	dao:       video.GetDao(),
	cache:     video.GetCache(),
	likeCache: like.GetCache(),
}

func (s *service) Add2(uid uint, add video.VideoAdd, url string, cover string) error {

	_video := &models.Video{
		Uid:    uid,
		Title:  add.Title,
		Cover:  cover,
		Url:    url,
		Locked: false,
		Status: video_status.Auditing,
	}

	err := s.dao.Add(_video)

	return err
}

func (s *service) GetAll(c *core.Context, uid *uint, query *video.VideoQuery) ([]video.VideoSimple, error) {
	all, err := s.dao.GetAll(uid, query)

	return all, err
}

// Get get video
func (s *service) Get(c *core.Context, vid uint) (*video.VideoRes, error) {
	// TODO: 校验video是否上架、审核通过、锁定、删除
	_video, err := s.dao.GetById(vid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(business_code.RecodeNotFound)
		}

		panic(err)
	}

	var isLike = false    //是否点赞
	var isDislike = false // 是否点踩
	var isCollect = false // 是否收藏
	likeCount, _ := s.cache.GetVideoLikeCount(vid)
	if err != nil {
		panic(err)
	}
	dislikeCount, _ := s.cache.GetVideoDislikeCount(vid)
	collectionCount, err := s.dao.GetVideoCollectionCountById(vid)
	if err != nil {
		panic(err)
	}
	// 已登录
	if c.Authorized {
		//TODO: 捕获错误
		likeType, _ := s.likeCache.GetLikeTypeByUserAndVideo(*c.AuthUID, vid)
		//if err != nil {
		//	return nil, err
		//}
		isLike = like_type.Like.Is(likeType)
		isDislike = like_type.Dislike.Is(likeType)
		//isCollect, _ = s.dao.IsCollect(*c.UID, vid)
		isCollect, _ = s.dao.IsCollect(*c.AuthUID, vid)
	}

	videoRes := &video.VideoRes{
		Video:           _video,
		IsLike:          isLike,
		IsDislike:       isDislike,
		IsCollect:       isCollect,
		LikeCount:       likeCount,
		DislikeCount:    dislikeCount,
		CollectionCount: collectionCount,
	}

	return videoRes, err

}

// Add add video
func (s *service) Add(uid uint, add *video.VideoAdd) (*models.Video, error) {

	snapshot := fmt.Sprintf("%s?x-oss-process=video/snapshot,t_7000,f_jpg,w_800,h_600,m_fast", add.Url)

	video := &models.Video{
		Uid:      uid,
		Title:    add.Title,
		Cover:    add.Cover,
		Url:      add.Url,
		Locked:   false,                    // 默认未锁定
		Status:   video_status.AuditPermit, // 没后台审核默认审核通过
		Snapshot: snapshot,
	}

	err := s.dao.Add(video)

	return video, err
}

// Update update video
func (s *service) Update(c *core.Context, vid uint, videoUpdate *video.VideoUpdate) error {
	//vid := c.GetId()

	err := s.dao.Update(*c.AuthUID, vid, videoUpdate)

	if err != nil {
		panic(err)
	}

	return err
}

// Delete delete video
func (s *service) Delete(c *core.Context, vid uint) error {
	err := s.dao.Delete(*c.AuthUID, vid)

	return err
}
