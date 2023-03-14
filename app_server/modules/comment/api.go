package comment

import "simple-video-server/core"

type _api struct {
	service *_service
}

var Api = &_api{
	service: Service,
}

func (api *_api) Add(c *core.Context) (uint, error) {

	var commentAdd CommentAdd
	err := c.ShouldBind(&commentAdd)
	if err != nil {
		panic(err)
	}

	id, err := api.service.Add(c, &commentAdd)

	return id, err
}

func (api *_api) Delete(c *core.Context) (bool, error) {
	err := api.service.Delete(c)

	return true, err
}

func (api *_api) Get(c *core.Context) ([]*CommentResSimple, error) {
	var query CommentQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	comments, err := api.service.Get(c, &query)

	return comments, err
}
