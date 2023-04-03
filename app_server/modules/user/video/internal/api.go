package internal

import (
	"simple-video-server/app_server/modules/user/video"
	"simple-video-server/core"
	"simple-video-server/models"
)

type Api struct {
	service *Service
}

var _api = &Api{
	service: service,
}

func GetApi() *Api {
	return _api
}

func (api *Api) GetAll(c *core.Context) ([]models.Video, error) {

	var query video.VideoQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	videos, err := api.service.GetAll(c, &query)

	return videos, err
}
