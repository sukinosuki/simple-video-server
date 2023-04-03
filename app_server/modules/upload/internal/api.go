package internal

import (
	upload2 "simple-video-server/app_server/modules/upload"
	"simple-video-server/core"
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

func (api *Api) Upload(c *core.Context) (string, error) {
	var form upload2.UploadData
	err := c.ShouldBind(&form)
	if err != nil {
		panic(err)
	}

	filename, err := api.service.Upload(c, &form)
	if err != nil {
		panic(err)
	}

	return filename, err
}
