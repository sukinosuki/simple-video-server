package user

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"simple-video-server/core"
	"simple-video-server/models"
	"simple-video-server/pkg/util"
)

type service struct {
	dao *_userDao
}

//	var Service = &service{
//		dao: Dao,
//	}
var Service *service

func GetUserService() *service {
	if Service != nil {
		return Service
	}

	return &service{
		dao: GetUserDao(),
	}
}

// Register 注册
func (s *service) Register(c *core.Context, userRegister *UserRegister) *models.User {
	_, err := Dao.GetByEmail(userRegister.Email)

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			panic(err)
		}
	}

	hashedPassword, err := util.Password.Hashed(userRegister.Password)
	if err != nil {
		panic(err)
	}

	user := &models.User{
		Nickname: userRegister.Email[0:10], //TODO 随机昵称或者用户提交
		Email:    userRegister.Email,
		Password: hashedPassword,
		Enabled:  true,
	}

	_, err = Dao.Create(user)

	if err != nil {
		panic(err)
	}

	return user
}

// Login 登录
func (s *service) Login(c *core.Context, userLogin *UserLogin) *models.User {
	c.Log.Info("service login start ", zap.Any("data ", *userLogin))

	user, err := Dao.FindByEmail(userLogin.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errors.New("邮箱或密码错误(邮箱不存在"))
		}

		panic(err)
	}

	// TODO: compare password
	err = util.Password.Compare(user.Password, userLogin.Password)

	if err != nil {
		c.Log.Error("邮箱或密码错误")
		//panic(err)
		panic(errors.New("邮箱或密码错误(密码错误"))
	}

	// TODO: token

	err = UserCache.SetUser(user.ID, user)

	if err != nil {
		panic(err)
	}

	return user
}

// GetProfile 用户详情
func (s *service) GetProfile(c *core.Context, uid uint) *models.User {

	user, err := UserCache.GetUser(uid)

	if err == nil {
		return user
	}

	//TODO: 查询数据库返回
	panic(err)
}
