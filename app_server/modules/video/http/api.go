package http

import (
	"simple-video-server/app_server/modules/video"
	"simple-video-server/core"
)

type Api struct {
	service *service
	//commentService *comment.Service
}

var _api = &Api{
	service: Service,
	//commentService: comment.GetService(),
}

func GetApi() *Api {
	return _api
}

// Add 随机返回n条
// SELECT * FROM table_name ORDER BY RAND() LIMIT N;

func (api *Api) Add(c *core.Context) (bool, error) {
	var videoAdd video.VideoAdd

	err := c.ShouldBind(&videoAdd)
	if err != nil {
		panic(err)
	}

	_, err = Service.Add(*c.AuthUID, &videoAdd)

	if err != nil {
		panic(err)
	}

	return true, nil
}

// GetAll get video list
func (api *Api) GetAll(c *core.Context) ([]video.VideoSimple, error) {

	var query video.VideoQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}
	uid := c.GetParamUID()
	videoSimples, err := api.service.GetAll(c, &uid, &query)

	if err != nil {
		panic(err)
	}

	return videoSimples, nil
}

func (api *Api) GetAuthAll(c *core.Context) ([]video.VideoSimple, error) {
	var query video.VideoQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	videoSimples, err := api.service.GetAll(c, c.AuthUID, &query)

	return videoSimples, err
}

func (api *Api) Feed(c *core.Context) ([]video.VideoSimple, error) {
	var query video.VideoQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}
	query.Random = true

	videoSimples, err := api.service.GetAll(c, nil, &query)

	if err != nil {
		panic(err)
	}

	return videoSimples, nil
}

func (api *Api) GetById(c *core.Context) (*video.VideoRes, error) {

	id := c.GetParamId()

	videoRes, err := api.service.Get(c, id)

	if err != nil {
		panic(err)
	}

	return videoRes, nil
}

func (api *Api) Update(c *core.Context) (bool, error) {
	var videoUpdate video.VideoUpdate
	err := c.ShouldBind(&videoUpdate)
	if err != nil {
		panic(err)
	}

	err = api.service.Update(c, c.GetParamId(), &videoUpdate)
	if err != nil {
		panic(err)
	}

	return true, err
}

func (api *Api) Delete(c *core.Context) (bool, error) {
	err := api.service.Delete(c, c.GetParamId())
	if err != nil {
		panic(err)
	}

	return true, err
}

//func (api *Api) GetComment(context *core.Context) (T, error) {
//	var form comment.CommentAdd
//	c.shouldbind
//}
