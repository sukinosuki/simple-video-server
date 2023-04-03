package internal

import (
	"simple-video-server/app_server/modules/comment"
	"simple-video-server/core"
)

type Api struct {
	service *Service
}

var _api = &Api{
	service: _Service,
}

func GetApi() *Api {
	return _api
}

func (api *Api) Add(c *core.Context) (*comment.CommentResSimple, error) {

	var commentAdd comment.CommentAdd
	err := c.ShouldBind(&commentAdd)
	if err != nil {
		panic(err)
	}

	comment, err := api.service.Add(c, &commentAdd, commentAdd.MediaID, commentAdd.MediaType)

	return comment, err
}

func (api *Api) Delete(c *core.Context) (bool, error) {
	var form comment.CommentDelete
	err := c.ShouldBind(&form)
	if err != nil {
		panic(err)
	}

	err = api.service.Delete(c, form.MediaID, form.MediaType)

	return true, err
}

func (api *Api) GetAll(c *core.Context) ([]*comment.CommentResSimple, error) {
	var query comment.CommentQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	comments, err := api.service.GetAll(c, &query)

	return comments, err
}

func (api *Api) Get(c *core.Context) ([]comment.CommentResSimple, error) {

	var query comment.CommentQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	comments, err := api.service.Get(c, &query)

	return comments, err
}
