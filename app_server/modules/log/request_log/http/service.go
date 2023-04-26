package http

import (
	"simple-video-server/app_server/modules/log/request_log"
	"simple-video-server/core"
	"simple-video-server/models"
)

type Service struct {
	dao *request_log.Dao
}

var _service = &Service{
	dao: request_log.GetDao(),
}

func GetService() *Service {
	return _service
}

func (s *Service) GetAll(c *core.Context, query *request_log.Query) []models.RequestLog {
	handlerName := "GetAll"

	logs, err := s.dao.GetAll(query)

	c.PanicIfErr(err, handlerName, "获取请求log失败")

	return logs
}
