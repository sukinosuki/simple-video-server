package video

import (
	"simple-video-server/core"
	"simple-video-server/models"
)

type _service struct {
	dao *_dao
}

var service = &_service{
	dao: dao,
}

func (s *_service) GetAll(c *core.Context, query *VideoQuery) ([]models.Video, error) {

	videos, err := s.dao.GetAll(*c.UID, query)

	return videos, err
}
