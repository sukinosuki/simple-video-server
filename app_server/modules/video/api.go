package video

import (
	"simple-video-server/core"
)

type _api struct {
	service *service
}

var Api = &_api{
	service: Service,
}

func (api *_api) Add(c *core.Context) (bool, error) {
	var videoAdd VideoAdd

	err := c.ShouldBind(&videoAdd)
	if err != nil {
		panic(err)
	}

	_, err = Service.Add(*c.UID, &videoAdd)

	if err != nil {
		panic(err)
	}

	return true, nil
}

// GetAll get video list
func (api *_api) GetAll(c *core.Context) ([]VideoSimple, error) {

	var query VideoQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	videoSimples, err := api.service.GetAll(c, &query)

	if err != nil {
		panic(err)
	}

	return videoSimples, nil
}

func (api *_api) GetById(c *core.Context) (*VideoRes, error) {

	id := c.GetId()

	videoRes, err := api.service.Get(c, id)

	if err != nil {
		panic(err)
	}

	return videoRes, nil
}

func (api *_api) Update(c *core.Context) (bool, error) {
	var videoUpdate VideoUpdate
	err := c.ShouldBind(&videoUpdate)
	if err != nil {
		panic(err)
	}

	err = api.service.Update(c, c.GetId(), &videoUpdate)
	if err != nil {
		panic(err)
	}

	return true, err
}

func (api *_api) Delete(c *core.Context) (bool, error) {
	err := api.service.Delete(c, c.GetId())
	if err != nil {
		panic(err)
	}

	return true, err
}
