package comment

import "simple-video-server/core"

type _api struct {
	service *Service
}

var Api = &_api{
	service: _Service,
}

func (api *_api) Add(c *core.Context) (uint, error) {

	var commentAdd CommentAdd
	err := c.ShouldBind(&commentAdd)
	if err != nil {
		panic(err)
	}

	id, err := api.service.Add(c, &commentAdd, commentAdd.MediaID, commentAdd.MediaType)

	return id, err
}

func (api *_api) Delete(c *core.Context) (bool, error) {
	var form CommentDelete
	err := c.ShouldBind(&form)
	if err != nil {
		panic(err)
	}

	err = api.service.Delete(c, form.MediaID, form.MediaType)

	return true, err
}

func (api *_api) GetAll(c *core.Context) ([]*CommentResSimple, error) {
	var query CommentQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	comments, err := api.service.GetAll(c, &query)

	return comments, err
}

func (api *_api) Get(c *core.Context) ([]CommentResSimple, error) {

	var query CommentQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	comments, err := api.service.Get(c, &query)

	return comments, err
}
