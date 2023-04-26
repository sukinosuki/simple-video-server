package http

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

	form := core.MustBindForm[comment.CommentAdd](c)

	commentSimple := api.service.Add(c, form, form.MediaID, form.MediaType)

	return commentSimple, nil
}

func (api *Api) Delete(c *core.Context) (bool, error) {
	form := core.MustBindForm[comment.CommentDelete](c)

	api.service.Delete(c, form.MediaID, form.MediaType)

	return true, nil
}

func (api *Api) GetAll(c *core.Context) ([]*comment.CommentResSimple, error) {

	query := core.MustBindForm[comment.CommentQuery](c)

	comments := api.service.GetAll(c, query)

	return comments, nil
}

func (api *Api) Get(c *core.Context) ([]comment.CommentResSimple, error) {

	query := core.MustBindForm[comment.CommentQuery](c)

	comments := api.service.Get(c, query)

	return comments, nil
}
