package user

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"simple-video-server/models"
	"simple-video-server/pkg/log"
	"simple-video-server/util"
)

type service struct {
}

var Service = &service{}

// Register 注册
func (s *service) Register(userRegister *UserRegister) *models.User {
	_, err := UserDao.GetByEmail(userRegister.Email)

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

	_, err = UserDao.Create(user)

	if err != nil {
		panic(err)
	}

	return user
}

// Login 登录
func (s *service) Login(ctx context.Context, userLogin *UserLogin) *models.User {
	log := log.GetCtx(ctx)
	log.Info("service login start ", zap.Any("data ", *userLogin))

	user, err := UserDao.FindByEmail(userLogin.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errors.New("邮箱或密码错误(邮箱不存在"))
		}

		panic(err)
	}

	// TODO: compare password
	err = util.Password.Compare(user.Password, userLogin.Password)

	if err != nil {
		log.Error("邮箱或密码错误")
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

func (s *service) GetProfile(uid uint) *models.User {

	user, err := UserCache.GetUser(uid)

	if err == nil {
		return user
	}

	//TODO: 查询数据库返回
	panic(err)
}
