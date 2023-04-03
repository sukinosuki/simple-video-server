package internal

import (
	"simple-video-server/app_server/modules/user/video"
	"simple-video-server/core"
	"simple-video-server/models"
)

type Service struct {
	dao *video.Dao
}

var service = &Service{
	dao: video.GetDao(),
}

func (s *Service) GetAll(c *core.Context, query *video.VideoQuery) ([]models.Video, error) {

	videos, err := s.dao.GetAll(*c.AuthUID, query)

	return videos, err
}
