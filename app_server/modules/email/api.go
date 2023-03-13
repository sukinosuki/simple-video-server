package email

import "simple-video-server/core"

type _api struct {
	service *_service
}

var Api = &_api{
	service: Service,
}

func (api *_api) SendEmail(c *core.Context) (string, error) {

	var data SendEmail
	err := c.ShouldBind(&data)
	if err != nil {
		panic(err)
	}

	code, err := api.service.Send(c, &data)

	return code, err
}
