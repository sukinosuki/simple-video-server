package collection

import (
	"errors"
	"simple-video-server/constants/video_status"
	"simple-video-server/core"
	"simple-video-server/models"
)

type Service struct {
	dao *_dao
}

var _service *Service

func GetCollectionService() *Service {
	if _service != nil {
		return _service
	}

	return &Service{
		//dao: GetCollectionDao(),
		dao: Dao,
	}
}

func (s *Service) Add(c *core.Context, vid uint) error {
	uid := *c.AuthUID
	// TODO: 视频是否存在
	exists, video, err := s.dao.IsVideoExists(vid)
	if err != nil {
		return err
	}

	if !exists {
		// TODO: 可以直接返回添加收藏失败不返回详情的添加失败信息
		return errors.New("记录不存在")
	}

	// 校验视频是否被锁定、是否审核通过
	if video.Locked || video.Status != video_status.AuditPermit {
		// TODO: 可以直接返回添加收藏失败不返回详情的添加失败信息
		return errors.New("视频被锁定或者审核未通过")
	}

	exists, err = s.dao.IsCollect(uid, vid)
	// 已收藏
	if exists {
		// TODO: 可以直接返回添加收藏失败不返回详情的添加失败信息
		return errors.New("重复收藏")
	}

	collection := &models.UserVideoCollection{
		UID: uid,
		VID: vid,
	}

	err = s.dao.Add(collection)

	return err
}

func (s *Service) Delete(c *core.Context, vid uint) error {
	err := s.dao.Delete(*c.AuthUID, vid)

	return err
}

func (s *Service) GetAll(c *core.Context, uid uint, query *CollectionQuery) ([]*UserVideoCollectionRes, error) {

	collections, err := s.dao.GetAll(uid, query)
	if err != nil {
		return nil, err
	}

	return collections, err
}
