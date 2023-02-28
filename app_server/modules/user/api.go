package user

import (
	"go.uber.org/zap"
	"simple-video-server/common"
	"simple-video-server/core"
	"simple-video-server/models"
	"simple-video-server/pkg/app_jwt"
	"strconv"
)

type _api struct {
	common.BaseApi
	service *service
}

var Api = &_api{
	common.BaseApi{
		Module: "user",
		//Log:  common.BaseApiLog{},
	},
	Service,
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

	err := c.ShouldBindJSON(&userRegister)
	if err != nil {
		//log.Error("用户注册请求参数错误 ", zap.String("msg", err.Error()))
		// TODO: 返回明确的参数错误
		panic(err)
	}

	user := Service.Register(c, &userRegister)

	token, err := app_jwt.AppJwt.Create(user.ID)
	if err != nil {
		panic(err)
	}

	loginRes := &LoginRes{
		user,
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
	log := c.Log.With(zap.String("caller", "user_api/login"))

	log.Info("登录开始")

	var userLogin UserLogin
	err := c.ShouldBind(&userLogin)

	if err != nil {
		//TODO:返回明确的参数错误信息
		panic(err)
	}

	user := Service.Login(c, &userLogin)

	token, err := app_jwt.AppJwt.Create(user.ID)

	if err != nil {
		panic(err)
	}

	loginRes := &LoginRes{
		user,
		token,
	}

	return loginRes, err
}

// 用户信息
func (api *_api) Profile(c *core.Context) (*models.User, error) {
	log := c.Log.With(zap.String("caller", "user_api/profile"))

	log.Info("获取用户信息")

	user := Service.GetProfile(c, *c.UID)

	return user, nil
}

// AddCollection 新增收藏
func (api *_api) AddCollection(c *core.Context) (bool, error) {

	var data AddCollection
	err := c.ShouldBind(&data)
	if err != nil {
		panic(err)
	}

	err = api.service.AddCollection(c, data.VID)
	if err != nil {
		panic(err)
	}

	return true, nil
}

func (api *_api) DeleteCollection(c *core.Context) (bool, error) {

	idStr := c.Param("id")
	vid, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	//var data AddCollection
	//err := c.ShouldBind(&data)
	//if err != nil {
	//	panic(err)
	//}

	err = api.service.DeleteCollection(c, uint(vid))
	if err != nil {
		panic(err)
	}

	return true, nil
}

func (api *_api) GetAllCollection(c *core.Context) ([]models.VideoCollection, error) {

	//	TODO:分页
	collections, err := api.service.GetAllCollection(c)
	if err != nil {
		panic(err)
	}

	return collections, nil
}
