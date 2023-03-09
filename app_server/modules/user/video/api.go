package video

import (
	"simple-video-server/core"
	"simple-video-server/models"
)

type _api struct {
	service *_service
}

var Api = &_api{
	service: service,
}

func (api *_api) GetAll(c *core.Context) ([]models.Video, error) {

	var query VideoQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	videos, err := api.service.GetAll(c, &query)

	return videos, err
}
