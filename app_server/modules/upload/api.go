package upload

import "simple-video-server/core"

type _api struct {
	service *Service
}

var Api = &_api{
	service: service,
}

func (api *_api) Upload(c *core.Context) (string, error) {
	var form UploadData
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
