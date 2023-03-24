package video

import (
	"simple-video-server/core"
)

type _api struct {
	service *service
	//commentService *comment.Service
}

var Api = &_api{
	service: Service,
	//commentService: comment.GetService(),
}

// Add 随机返回n条
// SELECT * FROM table_name ORDER BY RAND() LIMIT N;

func (api *_api) Add(c *core.Context) (bool, error) {
	var videoAdd VideoAdd

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
func (api *_api) GetAll(c *core.Context) ([]VideoSimple, error) {

	var query VideoQuery
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

func (api *_api) GetAuthAll(c *core.Context) ([]VideoSimple, error) {
	var query VideoQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	videoSimples, err := api.service.GetAll(c, c.AuthUID, &query)

	return videoSimples, err
}

func (api *_api) Feed(c *core.Context) ([]VideoSimple, error) {
	var query VideoQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}
	query.random = true

	videoSimples, err := api.service.GetAll(c, nil, &query)

	if err != nil {
		panic(err)
	}

	return videoSimples, nil
}

func (api *_api) GetById(c *core.Context) (*VideoRes, error) {

	id := c.GetParamId()

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

	err = api.service.Update(c, c.GetParamId(), &videoUpdate)
	if err != nil {
		panic(err)
	}

	return true, err
}

func (api *_api) Delete(c *core.Context) (bool, error) {
	err := api.service.Delete(c, c.GetParamId())
	if err != nil {
		panic(err)
	}

	return true, err
}

//func (api *_api) GetComment(context *core.Context) (T, error) {
//	var form comment.CommentAdd
//	c.shouldbind
//}
