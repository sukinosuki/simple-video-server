package user

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

	profile := Service.Register(c, &userRegister)

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

// 用户信息
func (api *_api) Profile(c *core.Context) (*Profile, error) {
	log := c.Log.With(zap.String("caller", "user_api/profile"))

	log.Info("获取用户信息")

	profile := Service.GetProfile(c, *c.UID)

	return profile, nil
}
