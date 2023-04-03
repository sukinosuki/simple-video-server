package http

import (
	"simple-video-server/app_server/modules/email"
	"simple-video-server/core"
)

type _Api struct {
	service *_Service
}

var _api = &_Api{
	service: service,
}

//func GetApi() *_Api {
//	return _api
//}

// SendEmail 发送验证码
// 发送验证有的操作是需要登录(重置密码), 有的操作不需要登录(注册)
func (api *_Api) SendEmail(c *core.Context) (string, error) {

	//var data email.SendEmail
	//err := c.ShouldBind(&data)
	//if err != nil {
	//	panic(err)
	//}

	form := core.MustBindForm[email.SendEmail](c)
	code, err := api.service.Send(c, form)

	return code, err
}
