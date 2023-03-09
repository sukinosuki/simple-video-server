package collection

import (
	"simple-video-server/core"
)

type _api struct {
	service *Service
}

var Api = &_api{
	service: GetCollectionService(),
}

// Add 新增收藏
func (api *_api) Add(c *core.Context) (bool, error) {

	var data AddCollection
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

func (api *_api) Delete(c *core.Context) (bool, error) {

	id := c.GetId()

	err := api.service.Delete(c, id)
	if err != nil {
		panic(err)
	}

	return true, nil
}

func (api *_api) GetAll(c *core.Context) ([]*UserVideoCollectionRes, error) {
	var query CollectionQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	collections, err := api.service.GetAll(c, &query)

	if err != nil {
		panic(err)
	}

	return collections, nil
}
