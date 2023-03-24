package user

import "simple-video-server/core"

type Api struct {
	service *Service
}

var api = &Api{
	service: service,
}

func (api *Api) GetAll(c *core.Context) ([]*UserSimple, error) {

	var query UserQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	users, err := api.service.GetAll(c, &query)

	return users, err
}
