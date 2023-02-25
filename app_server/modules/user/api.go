package user

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/models"
	"simple-video-server/pkg/app_ctx"
	"simple-video-server/pkg/app_jwt"
)

type controller struct {
}

var Controller = &controller{}

// RegisterHandler
// @Tag.name 用户管理
// @Summary 注册summary
// @Description 注册description
// @Param data body UserRegister true "登录参数"
// @Router /api/v1/user/register [post]
// @Success 200 {object} LoginRes "成功响应"
func (ctl *controller) RegisterHandler(c *gin.Context) (*LoginRes, error) {
	userRegister := &UserRegister{}

	err := c.ShouldBindJSON(userRegister)
	if err != nil {
		panic(err)
		//if errs, ok := err.(validator.ValidationErrors); ok {
		//	errorsMap := errs.Translate(validation.Trans)
		//
		//	msg := ""
		//
		//	for _, v := range errorsMap {
		//		msg = v
		//		break
		//	}
		//	fmt.Println("msg ", msg)
		//
		//	//return nil, errors.New(msg)
		//	//return nil, business_code.RequestErr
		//	return nil, err_code.RequestErr(msg, err.Error())
		//}
		//
		//return nil, err
	}

	user := Service.Register(userRegister)

	token, err := app_jwt.AppJwt.Create(user.ID)

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
// @Router /api/v1/user/login [post]
// @Success 200 {object} LoginRes
func (ctl *controller) Login(c *gin.Context) (*LoginRes, error) {
	userLogin := &UserLogin{}

	err := c.ShouldBind(userLogin)

	if err != nil {
		panic(err)
	}

	user := Service.Login(userLogin)

	token, err := app_jwt.AppJwt.Create(user.ID)

	if err != nil {
		return nil, err
	}

	loginRes := &LoginRes{
		user,
		token,
	}

	return loginRes, err
}

func (ctl *controller) Info(c *gin.Context) (*models.User, error) {
	uid, _ := app_ctx.GetUid(c)

	user := Service.GetProfile(uid)

	return user, nil
}
