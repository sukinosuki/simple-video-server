package internal

import (
	"simple-video-server/app_server/modules/user/collection"
	"simple-video-server/core"
)

type Api struct {
	service *Service
}

var _api = &Api{
	service: GetService(),
}

func GetApi() *Api {
	return _api
}

// Add 新增收藏
func (api *Api) Add(c *core.Context) (bool, error) {

	var data collection.AddCollection
	err := c.ShouldBind(&data)
	if err != nil {
		panic(err)
	}

	err = api.service.Add(c, data.VID)
	if err != nil {
		panic(err)
	}

	return true, nil
}

func (api *Api) Delete(c *core.Context) (bool, error) {

	id := c.GetParamId()

	err := api.service.Delete(c, id)
	if err != nil {
		panic(err)
	}

	return true, nil
}

func (api *Api) GetAll(c *core.Context) ([]*collection.UserVideoCollectionRes, error) {
	var query collection.CollectionQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	uid := c.GetParamUID()

	collections, err := api.service.GetAll(c, uid, &query)

	if err != nil {
		panic(err)
	}

	return collections, nil
}
