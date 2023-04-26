package http

import (
	"simple-video-server/app_server/modules/log/operation_log"
	"simple-video-server/core"
	"simple-video-server/models"
)

type Api struct {
	service *Service
}

var _api = &Api{
	service: _service,
}

func (api *Api) GetAll(c *core.Context) ([]models.OperationLog, error) {
	query := core.MustBindForm[operation_log.Query](c)

	logs := api.service.GetAll(c, query)

	return logs, nil
}
