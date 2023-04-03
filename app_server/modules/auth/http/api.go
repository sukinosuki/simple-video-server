package http

import (
	"simple-video-server/app_server/modules/auth"
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

// RegisterHandler
// @Tag.name 用户管理
// @Summary 注册summary
// @Description 注册description
// @Param data body RegisterDTO true "登录参数"
// @Router /_Api/v1/user/register [post]
// @Success 200 {object} AuthLoginRes "成功响应"
func (api *_Api) RegisterHandler(c *core.Context) (*auth.LoginRes, error) {
	form := core.MustBindForm[auth.RegisterForm](c)

	loginRes := api.service.Register(c, form)

	return loginRes, nil
}

// Login
// @Summary 登录
// @Tag.name 用户管理
// @Param data body LoginForm true "登录参数"
// @Router /_Api/v1/user/login [post]
// @Success 200 {object} AuthLoginRes
func (api *_Api) Login(c *core.Context) (*auth.LoginRes, error) {
	form := core.MustBindForm[auth.LoginForm](c)
	loginRes := api.service.Login(c, form)

	return loginRes, nil
}

// AuthProfile Profile 用户信息
func (api *_Api) AuthProfile(c *core.Context) (*auth.LoginResProfile, error) {

	profile := api.service.GetProfile(c, *c.AuthUID)

	return profile, nil
}

// UpdateProfile 更新profile
func (api *_Api) UpdateProfile(c *core.Context) (bool, error) {

	form := core.MustBindForm[auth.UpdateForm](c)

	err := api.service.UpdateProfile(c, form)

	return true, err
}

// ResetPassword 重置密码
// TODO: 多种重置密码的方式(邮箱验证码、短信验证码
func (api *_Api) ResetPassword(c *core.Context) (bool, error) {
	form := core.MustBindForm[auth.ResetPasswordForm](c)

	err := api.service.ResetPassword(c, form)

	return true, err
}

// Logoff 注销
func (api *_Api) Logoff(c *core.Context) (bool, error) {

	err := api.service.Logoff(c)

	return true, err
}
