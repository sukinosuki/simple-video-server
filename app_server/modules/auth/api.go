package auth

import (
	"go.uber.org/zap"
	"simple-video-server/core"
	"simple-video-server/pkg/app_jwt"
)

type _api struct {
	service *service
}

var Api = &_api{
	service: Service,
}

// RegisterHandler
// @Tag.name 用户管理
// @Summary 注册summary
// @Description 注册description
// @Param data body UserRegister true "登录参数"
// @Router /_api/v1/user/register [post]
// @Success 200 {object} LoginRes "成功响应"
func (api *_api) RegisterHandler(c *core.Context) (*LoginRes, error) {
	log := c.Log.With(zap.String("caller", "user_api/registerHandler"))

	log.Info("用户注册开始")

	var userRegister UserRegister

	err := c.ShouldBind(&userRegister)
	if err != nil {
		//log.Error("用户注册请求参数错误 ", zap.String("msg", err.Error()))
		// TODO: 返回明确的参数错误
		panic(err)
	}

	profile, err := Service.Register(c, &userRegister)

	if err != nil {
		panic(err)
	}

	token, err := app_jwt.AppJwt.Create(profile.User.ID)
	if err != nil {
		panic(err)
	}

	loginRes := &LoginRes{
		profile,
		token,
	}

	return loginRes, nil
}

// Login
// @Summary 登录
// @Tag.name 用户管理
// @Param data body UserLogin true "登录参数"
// @Router /_api/v1/user/login [post]
// @Success 200 {object} LoginRes
func (api *_api) Login(c *core.Context) (*LoginRes, error) {

	var userLogin UserLogin
	err := c.ShouldBind(&userLogin)

	if err != nil {
		panic(err)
	}

	profile := Service.Login(c, &userLogin)

	c.Log.Info("设置token开始")
	token, err := app_jwt.AppJwt.Create(profile.User.ID)
	c.Log.Info("设置token结束")

	if err != nil {
		panic(err)
	}

	loginRes := &LoginRes{
		profile,
		token,
	}

	return loginRes, err
}

// Profile 用户信息
func (api *_api) Profile(c *core.Context) (*Profile, error) {

	id := c.GetParamId()

	profile := Service.GetProfile(c, id)

	return profile, nil
}

// AuthProfile Profile 用户信息
func (api *_api) AuthProfile(c *core.Context) (*Profile, error) {

	profile := Service.GetProfile(c, *c.AuthUID)

	return profile, nil
}

func (api *_api) UpdateProfile(c *core.Context) (bool, error) {
	var form ProfileUpdate
	err := c.ShouldBind(&form)
	if err != nil {
		panic(err)
	}

	err = api.service.UpdateProfile(c, &form)

	return true, err
}

// ResetPassword 重置密码
// TODO: 多种重置密码的方式(邮箱验证码、短信验证码
func (api *_api) ResetPassword(c *core.Context) (bool, error) {
	var form UserResetPassword
	err := c.ShouldBind(&form)
	if err != nil {
		panic(err)
	}

	err = api.service.ResetPassword(c, &form)

	return true, err
}

func (api *_api) Logoff(c *core.Context) (bool, error) {

	err := api.service.Logoff(c)

	return true, err
}
