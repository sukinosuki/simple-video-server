package auth

import (
	"simple-video-server/core"
)

type _api struct {
	service *service
}

var Api = &_api{
	service: Service,
}

func mustBindForm[T any](c *core.Context) *T {
	var t T

	err := c.ShouldBind(&t)
	if err != nil {
		panic(err)
	}

	return &t
}

// RegisterHandler
// @Tag.name 用户管理
// @Summary 注册summary
// @Description 注册description
// @Param data body RegisterDTO true "登录参数"
// @Router /_api/v1/user/register [post]
// @Success 200 {object} AuthLoginRes "成功响应"
func (api *_api) RegisterHandler(c *core.Context) (*LoginRes, error) {
	//var userRegister common_models.AuthRegisterForm
	//
	//err := c.ShouldBind(&userRegister)
	//if err != nil {
	//	panic(err)
	//}
	userRegister := mustBindForm[RegisterForm](c)

	loginRes := Service.Register(c, userRegister)

	//if err != nil {
	//	panic(err)
	//}

	return loginRes, nil
}

// Login
// @Summary 登录
// @Tag.name 用户管理
// @Param data body LoginForm true "登录参数"
// @Router /_api/v1/user/login [post]
// @Success 200 {object} AuthLoginRes
func (api *_api) Login(c *core.Context) (*LoginRes, error) {

	var userLogin LoginForm
	err := c.ShouldBind(&userLogin)

	if err != nil {
		panic(err)
	}

	loginRes := Service.Login(c, &userLogin)

	return loginRes, err
}

// Profile 用户信息
func (api *_api) Profile(c *core.Context) (*LoginResProfile, error) {

	id := c.GetParamId()

	profile := Service.GetProfile(c, id)

	return profile, nil
}

// AuthProfile Profile 用户信息
func (api *_api) AuthProfile(c *core.Context) (*LoginResProfile, error) {

	profile := Service.GetProfile(c, *c.AuthUID)

	return profile, nil
}

// UserProfile AuthProfile Profile 用户信息
func (api *_api) UserProfile(c *core.Context) (*LoginResProfile, error) {

	uid := c.GetParamUID()
	profile := Service.GetProfile(c, uid)

	return profile, nil
}

func (api *_api) UpdateProfile(c *core.Context) (bool, error) {
	//var form UpdateForm
	//err := c.ShouldBindJSON(&form)
	//if err != nil {
	//	panic(err)
	//}

	form := mustBindForm[UpdateForm](c)
	//if form.Birthday != nil {
	//	value := *form.Birthday
	//
	//	birthday, err := time.Parse("2006-01-02", value)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	form.Birthday2 = &birthday
	//}

	err := api.service.UpdateProfile(c, form)

	return true, err
}

// ResetPassword 重置密码
// TODO: 多种重置密码的方式(邮箱验证码、短信验证码
func (api *_api) ResetPassword(c *core.Context) (bool, error) {
	var form ResetPasswordForm
	err := c.ShouldBind(&form)
	if err != nil {
		panic(err)
	}

	err = api.service.ResetPassword(c, &form)

	return true, err
}

// Logoff 注销
func (api *_api) Logoff(c *core.Context) (bool, error) {

	err := api.service.Logoff(c)

	return true, err
}
