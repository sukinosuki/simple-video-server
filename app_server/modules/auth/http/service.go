package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"simple-video-server/app_server/cache"
	"simple-video-server/app_server/modules/auth"
	"simple-video-server/app_server/modules/email"
	"simple-video-server/constants/email_action_type"
	"simple-video-server/constants/reset_password_method"
	"simple-video-server/core"
	"simple-video-server/db"
	"simple-video-server/models"
	"simple-video-server/pkg/app_jwt"
	"simple-video-server/pkg/util"
)

type _Service struct {
	authDao    *auth.Dao
	cache      *redis.Client
	emailCache *email.Cache
	db         *gorm.DB
}

var service = &_Service{
	authDao:    auth.GetAuthDao(),
	cache:      db.GetRedisDB(),
	emailCache: email.GetCache(),
	db:         db.GetOrmDB(),
}

//func GetService() *_Service {
//	return service
//}

func handleDeferTxError(tx *gorm.DB, err any) {
	if err != nil {
		tx.Rollback()
		panic(err)
	} else {
		tx.Commit()
	}
}

// 验证注册的email code
func (s *_Service) verifyEmailRegisterCode(email string, code string) (bool, error) {
	result, err := s.emailCache.Get(email, email_action_type.Register.Code)
	if err != nil {
		// TODO:
		return false, errors.New("找不到验证码")
	}

	if code != result {
		// TODO
		return false, errors.New("邮箱验证码不正确")
	}
	return false, nil
}

// Register 注册
func (s *_Service) Register(c *core.Context, userRegister *auth.RegisterForm) *auth.LoginRes {
	tx := s.db.Begin()

	defer func() {
		handleDeferTxError(tx, recover())
	}()

	// 从cache获取对应类型的email的验证码
	_, err := s.verifyEmailRegisterCode(userRegister.Email, userRegister.Code)
	core.PanicIfErr(err)

	// 校验email是否注册过
	exists, _, err := s.authDao.IsExistsByEmail(userRegister.Email)
	core.PanicIfErr(err)

	if exists {
		// TODO:
		panic(errors.New("该email已被注册"))
	}

	//hashed password
	hashedPassword, err := util.Password.Hashed(userRegister.Password)
	core.PanicIfErr(err)

	user := &models.User{
		Nickname: userRegister.Email[0:10], //TODO 随机昵称或者用户提交
		Email:    userRegister.Email,
		Password: hashedPassword,
		Enabled:  true,
	}

	_, err = s.authDao.Add(tx, user)
	core.PanicIfErr(err)

	token, err := app_jwt.AppJwt.Create(user.ID)
	core.PanicIfErr(err)

	loginRes := &auth.LoginRes{
		Token: token,
		Profile: &auth.LoginResProfile{
			User: auth.LoginResProfileUser{
				ID:        user.ID,
				Nickname:  user.Nickname,
				CreatedAt: user.CreatedAt,
				Enabled:   user.Enabled,
				Email:     user.Email,
			},
		},
	}

	return loginRes
}

// Login 登录
func (s *_Service) Login(c *core.Context, userLogin *auth.LoginForm) *auth.LoginRes {
	exists, user, err := s.authDao.IsExistsByEmail(userLogin.Email)

	if err != nil {
		panic(err)
	}

	if !exists {
		//TODO
		panic(errors.New("邮箱或密码错误(邮箱不存在"))
	}

	// 对比密码
	err = util.Password.Compare(user.Password, userLogin.Password)

	if err != nil {
		panic(errors.New("邮箱或密码错误(密码错误"))
	}

	// TODO: token

	err = auth.UserCache.SetUser(user.ID, user)

	if err != nil {
		panic(err)
	}

	videoCount, err := s.authDao.GetOneUserAllVideoCount(user.ID)
	if err != nil {
		panic(err)
	}

	collectionCount, err := s.authDao.GetUserAllCollectionCount(user.ID)
	if err != nil {
		panic(err)
	}

	token, err := app_jwt.AppJwt.Create(user.ID)

	if err != nil {
		panic(err)
	}

	loginRes := &auth.LoginRes{
		Token: token,
		Profile: &auth.LoginResProfile{
			User: auth.LoginResProfileUser{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				Enabled:   user.Enabled,
				Nickname:  user.Nickname,
				Email:     user.Email,
				Gender:    *user.Gender,
			},
			LikeCount:       0, //TODO
			DislikeCount:    0, //TODO
			CollectionCount: collectionCount,
			VideoCount:      videoCount,
			FollowerCount:   0, //TODO
		},
	}
	return loginRes
}

// GetProfile 用户详情
func (s *_Service) GetProfile(c *core.Context, uid uint) *auth.LoginResProfile {
	user, err := s.authDao.GetOneByID(uid)
	videoCount, err := s.authDao.GetOneUserAllVideoCount(user.ID)

	if err != nil {
		panic(err)
	}
	collectionCount, err := s.authDao.GetUserAllCollectionCount(user.ID)
	if err != nil {
		panic(err)
	}
	followersCount, _ := cache.Follow.OneUserFollowersCount(uid)
	followingCount, _ := cache.Follow.OneUserFollowingCount(uid)

	profile := &auth.LoginResProfile{
		User: auth.LoginResProfileUser{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			Enabled:   user.Enabled,
			Nickname:  user.Nickname,
			Email:     user.Email,
			Birthday:  user.Birthday,
			Gender:    *user.Gender,
			Avatar:    user.Avatar,
		},
		LikeCount:       0, // TODO
		DislikeCount:    0, // TODO
		CollectionCount: collectionCount,
		VideoCount:      videoCount,
		FollowerCount:   followersCount,
		FollowingCount:  followingCount,
	}

	return profile

}

// ResetPassword 重置密码
func (s *_Service) ResetPassword(c *core.Context, form *auth.ResetPasswordForm) error {

	switch {
	case reset_password_method.Email.Is(form.Method):
		// todo
		err := s.ResetPasswordByEmailCode(c, form)

		return err
	case reset_password_method.Sms.Is(form.Code):
	//	TODO
	default:
		panic(errors.New("不支持的method"))
	}

	return nil
}

// ResetPasswordByEmailCode 邮箱重置密码
func (s *_Service) ResetPasswordByEmailCode(c *core.Context, form *auth.ResetPasswordForm) error {
	// 1 校验code
	// 2 加密password
	// 3 更新user的password
	// 4 删除code
	uid := *c.AuthUID
	user, err := s.authDao.GetOneByID(uid)
	if err != nil {
		return err
	}

	result, err := s.emailCache.Get(user.Email, email_action_type.ResetPassword.Code)
	// 比对缓存code和提交的code
	//key := fmt.Sprintf("email_code:%s:%s", email_action_type.ResetPassword.Code, user.Email)
	//result, err := s.cache.Get(context.Background(), key).Result()
	if err != nil {
		// TODO: 没有key时会报error|proto.RedisError redis: nil错误
		return errors.New(fmt.Sprintf("code不正确, cache key不存在, email: %s", user.Email))
	}

	if result != form.Code {
		//TODO
		return errors.New("code不正确")
	}

	// 加密密码
	hashedPassword, err := util.Password.Hashed(form.Password)
	if err != nil {
		return err
	}

	// 更新user的密码字段
	user = &models.User{
		ID:       user.ID,
		Password: hashedPassword,
	}

	err = s.authDao.Updates(user)

	if err != nil {
		return err
	}

	// 删除cache code
	err = s.emailCache.Delete(user.Email, email_action_type.ResetPassword.Code)
	//_, err = s.cache.Del(context.Background(), key).Result()

	return err
}

// UpdateProfile 更新profile
func (s *_Service) UpdateProfile(c *core.Context, form *auth.UpdateForm) error {
	tx := s.db.Begin()

	defer func() {
		err := recover()

		if err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	//_gender := gender.GetByCode(form.Gender)
	//
	//if _gender == nil {
	//	//TODO
	//	panic(errors.New("gender错误"))
	//}

	user := models.User{
		ID:       *c.AuthUID,
		Nickname: form.Nickname,
		Avatar:   form.Avatar,
		Gender:   form.Gender,
		Birthday: form.Birthday,
	}

	err := s.authDao.UpdateProfile(tx, &user)

	return err
}

// Logoff 注销
func (s *_Service) Logoff(c *core.Context) error {
	//	删除该用户数据、缓存
	err := s.authDao.DeleteById(*c.AuthUID)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("user:%d", *c.AuthUID)

	num, err := s.cache.Del(context.Background(), key).Result()
	c.Log.Info(fmt.Sprintf("删除用户缓存: %d ", num))

	if err != nil {
		return err
	}

	return nil
}
