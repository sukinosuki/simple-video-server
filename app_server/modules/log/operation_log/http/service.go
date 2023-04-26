package http

import (
	"simple-video-server/app_server/modules/log/operation_log"
	"simple-video-server/core"
	"simple-video-server/models"
)

type Service struct {
	dao *operation_log.Dao
}

var _service = &Service{
	dao: operation_log.GetDao(),
}

func GetService() *Service {
	return _service
}

func (s *Service) Add(log *models.OperationLog) error {

	err := s.dao.Add(log)

	return err
}

func (s *Service) GetAll(c *core.Context, query *operation_log.Query) []models.OperationLog {
	handlerName := "GetAll"
	logs, err := s.dao.GetAll(query)

	c.PanicIfErr(err, handlerName, "获取操作日志失败")

	return logs
}
