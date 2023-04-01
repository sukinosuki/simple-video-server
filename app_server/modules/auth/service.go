package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"simple-video-server/app_server/cache"
	"simple-video-server/constants/email_action_type"
	"simple-video-server/constants/gender"
	"simple-video-server/constants/reset_password_method"
	"simple-video-server/core"
	"simple-video-server/db"
	"simple-video-server/models"
	"simple-video-server/pkg/app_jwt"
	"simple-video-server/pkg/util"
)

type service struct {
	authDao    *Dao
	cache      *redis.Client
	emailCache *cache.EmailCache
	db         *gorm.DB
}

var Service = &service{
	authDao:    GetAuthDao(),
	cache:      db.GetRedisDB(),
	emailCache: cache.Email,
	db:         db.GetOrmDB(),
}

func (s *service) verifyEmailRegisterCode(email string, code string) (bool, error) {
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
func (s *service) Register(c *core.Context, userRegister *RegisterForm) *LoginRes {
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

	//r1 -> 从cache获取对应类型的email的验证码
	_, err := s.verifyEmailRegisterCode(userRegister.Email, userRegister.Code)
	if err != nil {
		panic(err)
	}
	//result, err := s.emailCache.Get(userRegister.Email, email_action_type.Register.Code)
	//if err != nil {
	//	// TODO: 不需要返回太详细
	//	panic(errors.New("找不到验证码"))
	//}
	//
	//if userRegister.Code != result {
	//	panic(errors.New("邮箱验证码不正确"))
	//}

	//r2 -> 校验email是否注册过
	exists, _, err := s.authDao.IsExistsByEmail(userRegister.Email)
	if err != nil {
		panic(err)
	}

	if exists {
		//TODO:
		panic(errors.New("该email已被注册"))
	}

	//hashed password
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

	_, err = s.authDao.Add(tx, user)

	if err != nil {
		panic(err)
	}

	token, err := app_jwt.AppJwt.Create(user.ID)
	if err != nil {
		panic(err)
	}

	loginRes := &LoginRes{
		Token: token,
		Profile: &LoginResProfile{
			User: LoginResProfileUser{
				ID:        user.ID,
				Nickname:  user.Nickname,
				CreatedAt: user.CreatedAt,
				Enabled:   user.Enabled,
				Email:     user.Email,
			},
		},
	}

	//panic(errors.New("自定义错误测试transaction"))

	return loginRes
}

// Login 登录
func (s *service) Login(c *core.Context, userLogin *LoginForm) *LoginRes {
	exists, user, err := s.authDao.IsExistsByEmail(userLogin.Email)

	if err != nil {
		panic(err)
	}

	if !exists {
		//TODO
		panic(errors.New("邮箱或密码错误(邮箱不存在"))
	}

	c.Log.Info("对比密码开始")
	// compare password
	err = util.Password.Compare(user.Password, userLogin.Password)
	c.Log.Info("对比密码结束")

	if err != nil {
		panic(errors.New("邮箱或密码错误(密码错误"))
	}

	// TODO: token

	err = UserCache.SetUser(user.ID, user)

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

	loginRes := &LoginRes{
		Token: token,
		Profile: &LoginResProfile{
			User: LoginResProfileUser{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				Enabled:   user.Enabled,
				Nickname:  user.Nickname,
				Email:     user.Email,
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
func (s *service) GetProfile(c *core.Context, uid uint) *LoginResProfile {
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

	profile := &LoginResProfile{
		User: LoginResProfileUser{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			Enabled:   user.Enabled,
			Nickname:  user.Nickname,
			Email:     user.Email,
			Birthday:  user.Birthday,
			Gender:    user.Gender,
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
func (s *service) ResetPassword(c *core.Context, form *ResetPasswordForm) error {

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
func (s *service) ResetPasswordByEmailCode(c *core.Context, form *ResetPasswordForm) error {
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
func (s *service) UpdateProfile(c *core.Context, form *UpdateForm) error {
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

	_gender := gender.GetByCode(form.Gender)

	if _gender == nil {
		//TODO
		panic(errors.New("gender错误"))
	}

	user := models.User{
		ID:       *c.AuthUID,
		Nickname: form.Nickname,
		Avatar:   form.Avatar,
		Gender:   _gender.Code,
		Birthday: form.Birthday,
	}

	err := s.authDao.UpdateProfile(tx, &user)

	return err
}

// Logoff 注销
func (s *service) Logoff(c *core.Context) error {
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
