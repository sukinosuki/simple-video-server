package http

import (
	"errors"
	"simple-video-server/app_server/modules/user/collection"
	"simple-video-server/constants/video_status"
	"simple-video-server/core"
	"simple-video-server/models"
	"simple-video-server/pkg/app_err"
)

type Service struct {
	dao *collection.Dao
}

var _service = &Service{
	//dao: GetCollectionDao(),
	dao: collection.GetDao(),
}

func GetService() *Service {
	return _service
}

func (s *Service) Add(c *core.Context, vid uint) error {
	handlerName := "Add"
	businessErr := app_err.New(nil, handlerName, "")

	uid := *c.AuthUID
	// TODO: 视频是否存在
	exists, video, err := s.dao.IsVideoExists(vid)
	if err != nil {
		return businessErr.NewErr(err, "获取video是否存在失败")
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

func (s *Service) GetAll(c *core.Context, uid uint, query *collection.CollectionQuery) ([]*collection.UserVideoCollectionRes, error) {

	collections, err := s.dao.GetAll(uid, query)
	if err != nil {
		return nil, err
	}

	return collections, err
}
