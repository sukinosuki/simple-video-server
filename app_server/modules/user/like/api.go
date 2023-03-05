package like

import (
	"simple-video-server/core"
)

type _api struct {
	service *Service
}

var Api = &_api{
	service: GetLikeService(),
}

// Add 点赞
func (api *_api) Add(c *core.Context) (bool, error) {

	var videoLike VideoLike

	err := c.ShouldBind(&videoLike)
	if err != nil {
		panic(err)
	}

	err = api.service.Add(c, &videoLike)
	if err != nil {
		panic(err)
	}

	return true, err
}

// Delete 取消点赞
func (api *_api) Delete(c *core.Context) (bool, error) {
	var videoLike VideoLike
	err := c.ShouldBind(&videoLike)

	if err != nil {
		panic(err)
	}

	err = api.service.Delete(c, &videoLike)

	if err != nil {
		panic(err)
	}

	return true, err
}
