package http

import (
	"simple-video-server/app_server/modules/user/user"
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

func (api *Api) GetAll(c *core.Context) ([]*user.UserSimple, error) {

	var query user.UserQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	users, err := api.service.GetAll(c, &query)

	return users, err
}
