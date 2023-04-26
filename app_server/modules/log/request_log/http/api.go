package http

import (
	"simple-video-server/app_server/modules/log/request_log"
	"simple-video-server/core"
	"simple-video-server/models"
)

type Api struct {
	service *Service
}

var _api = &Api{
	service: _service,
}

func GetApi() *Api {
	return _api
}

func (api *Api) GetAll(c *core.Context) ([]models.RequestLog, error) {
	query := core.MustBindForm[request_log.Query](c)

	logs := api.service.GetAll(c, query)

	return logs, nil
}
