package http

import (
	"simple-video-server/app_server/modules/user/like"
	"simple-video-server/core"
)

type Api struct {
	service *Service
}

var _api = &Api{
	service: GetLikeService(),
}

func GetApi() *Api {
	return _api
}

// Add 点赞
func (api *Api) Add(c *core.Context) (bool, error) {

	//var videoLike like.VideoLike
	//
	//err := c.ShouldBind(&videoLike)
	//if err != nil {
	//	panic(err)
	//}

	form := core.MustBindForm[like.VideoLike](c)

	err := api.service.Add(c, *c.AuthUID, c.GetParamId(), form)
	if err != nil {
		panic(err)
	}

	return true, err
}

// Delete 取消点赞
func (api *Api) Delete(c *core.Context) (bool, error) {
	form := core.MustBindForm[like.VideoLike](c)

	err := api.service.Delete(c, *c.AuthUID, c.GetParamId(), form)

	return true, err
}
