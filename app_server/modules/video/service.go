package video

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"simple-video-server/app_server/cache"
	"simple-video-server/constants/like_type"
	"simple-video-server/constants/video_status"
	"simple-video-server/core"
	"simple-video-server/models"
	"simple-video-server/pkg/business_code"
)

type service struct {
	dao   *_dao
	cache *cache.VideoCache
}

var Service = &service{
	dao:   Dao,
	cache: cache.Video,
}

func (s *service) Add2(uid uint, add VideoAdd, url string, cover string) error {

	video := &models.Video{
		Uid:    uid,
		Title:  add.Title,
		Cover:  cover,
		Url:    url,
		Locked: false,
		Status: video_status.Auditing,
	}

	err := s.dao.Add(video)

	return err
	//return db.MysqlDB.Model(&Video{}).Add(video).Error
}

func (s *service) GetAll(c *core.Context, query *VideoQuery) ([]VideoSimple, error) {
	all, err := s.dao.GetAll(query)

	return all, err
}

// Get get video
func (s *service) Get(c *core.Context, vid uint) (*VideoRes, error) {
	// TODO: 校验video是否上架、审核通过、锁定、删除
	video, err := s.dao.GetById(vid)
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
		likeType, err := cache.Like.GetLikeTypeByUserAndVideo(*c.UID, vid)
		if err != nil {
			return nil, err
		}
		isLike = like_type.Like.Is(likeType)
		isDislike = like_type.Dislike.Is(likeType)
		//isCollect, _ = s.dao.IsCollect(*c.UID, vid)
		isCollect, _ = s.dao.IsCollect(*c.UID, vid)
	}

	videoRes := &VideoRes{
		Video:           video,
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
func (s *service) Add(uid uint, add *VideoAdd) (*models.Video, error) {

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
func (s *service) Update(c *core.Context, vid uint, videoUpdate *VideoUpdate) error {
	//vid := c.GetId()

	err := s.dao.Update(*c.UID, vid, videoUpdate)

	if err != nil {
		panic(err)
	}

	return err
}

// Delete delete video
func (s *service) Delete(c *core.Context, vid uint) error {
	err := s.dao.Delete(*c.UID, vid)

	return err
}
