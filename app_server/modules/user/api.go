package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"simple-video-server/common"
	"simple-video-server/models"
	"simple-video-server/pkg/app_ctx"
	"simple-video-server/pkg/app_jwt"
	"simple-video-server/pkg/log"
	"time"
)

type _api struct {
	common.BaseApi
}

var Api = &_api{
	common.BaseApi{
		Module: "user",
		//Log:  common.BaseApiLog{},
	},
}

// RegisterHandler
// @Tag.name 用户管理
// @Summary 注册summary
// @Description 注册description
// @Param data body UserRegister true "登录参数"
// @Router /_api/v1/user/register [post]
// @Success 200 {object} LoginRes "成功响应"
func (api *_api) RegisterHandler(c *gin.Context) (*LoginRes, error) {
	//api.Build(c, "register")
	log := log.GetCtx(c.Request.Context())

	log.Info("用户注册开始")

	traceId, _ := app_ctx.GetTraceId(c)
	log.Info("register1  ", zap.String("trace id ", traceId))

	c.Set("name", 233)
	time.Sleep(3 * time.Second)

	traceId, _ = app_ctx.GetTraceId(c)
	log.Info("register2 ", zap.String("trace id ", traceId))

	var userRegister UserRegister

	err := c.ShouldBindJSON(&userRegister)
	if err != nil {
		panic(err)
	}

	user := Service.Register(&userRegister)

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
func (api *_api) Login(c *gin.Context) (*LoginRes, error) {

	log := log.GetCtx(c.Request.Context())

	traceId, _ := app_ctx.GetTraceId(c)
	fmt.Println("login traceid1 ", traceId)
	log.Info("login 1 ", zap.String("trace id ", traceId))
	// TODO: 不同请求下的ctx.set是否会冲突
	//time.Sleep(5 * time.Second)

	traceId, _ = app_ctx.GetTraceId(c)
	log.Info("login 2", zap.String("trace id ", traceId))

	name, exists := c.Get("name")
	fmt.Println("name ", name)
	fmt.Println("exists ", exists)

	var userLogin UserLogin

	err := c.ShouldBind(&userLogin)

	if err != nil {
		panic(err)
	}

	user := Service.Login(c.Request.Context(), &userLogin)

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

func (api *_api) Profile(c *gin.Context) (*models.User, error) {
	log := log.GetCtx(c.Request.Context())

	uid, _ := app_ctx.GetUid(c)

	log.Info("获取用户信息 ", zap.Uint("uid", uid))

	user := Service.GetProfile(uid)

	return user, nil
}
