package collection

import (
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
		dao: GetCollectionDao(),
	}
}

func (s *Service) Add(c *core.Context, vid uint) error {
	//TODO:验证video是否为生效状态(是否存在、删除、审核通过
	exists, err := s.dao.IsCollect(*c.UID, vid)
	//TODO:不需要让用户知道详情的信息，可以直接返回添加失败
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		panic(errors.New("视频不存在"))
	//	}
	//
	//	return err
	//}
	// 已收藏, 直接返回成功
	if exists {
		return nil
	}

	// TODO:
	//// 校验视频是否被锁定、是否审核通过
	//if video.Locked || video.Status != video_status.AuditPermit {
	//	// 可以直接返回添加收藏失败
	//	//panic(errors.New("视频被锁定或者审核未通过"))
	//	c.Log.Info("视频被锁定或者审核未通过")
	//}

	collection := &models.UserVideoCollection{
		UID: *c.UID,
		VID: vid,
	}

	err = s.dao.Add(collection)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(c *core.Context, vid uint) error {
	err := s.dao.Delete(*c.UID, vid)
	if err != nil {
		return err
	}

	return nil
}

// GetAll TODO:分页
func (s *Service) GetAll(c *core.Context) ([]*UserVideoCollectionRes, error) {
	//collections, err := s.dao.GetAll(*c.UID)

	collections, err := s.dao.GetAll2(*c.UID)
	if err != nil {
		return nil, err
	}

	return collections, err
}
