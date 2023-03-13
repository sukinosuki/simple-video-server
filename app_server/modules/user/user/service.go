package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"simple-video-server/app_server/cache"
	"simple-video-server/constants/email_action_type"
	"simple-video-server/constants/reset_password_method"
	"simple-video-server/core"
	"simple-video-server/db"
	"simple-video-server/models"
	"simple-video-server/pkg/util"
)

type service struct {
	dao        *_dao
	cache      *redis.Client
	emailCache *cache.EmailCache
}

var Service = &service{
	dao:        Dao,
	cache:      db.GetRedisDB(),
	emailCache: cache.Email,
}

// Register 注册
func (s *service) Register(c *core.Context, userRegister *UserRegister) (*Profile, error) {
	// TODO: 校验email code
	//key := fmt.Sprintf("email_code:%s:%s", email_action_type.Register.Code, userRegister.Email)
	//
	//result, err := s.cache.Get(context.Background(), key).Result()
	result, err := s.emailCache.Get(userRegister.Email, email_action_type.Register.Code)
	if err != nil {
		// TODO
		return nil, errors.New("找不到验证码")
	}

	if userRegister.Code != result {
		return nil, errors.New("邮箱验证码不正确")
	}

	exists, _, err := s.dao.IsExistsByEmail(userRegister.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		//TODO:
		panic(errors.New("该email已被注册"))
	}

	hashedPassword, err := util.Password.Hashed(userRegister.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Nickname: userRegister.Email[0:10], //TODO 随机昵称或者用户提交
		Email:    userRegister.Email,
		Password: hashedPassword,
		Enabled:  true,
	}

	_, err = s.dao.Add(user)

	if err != nil {
		return nil, err
	}

	profile := &Profile{
		User: ProfileUser{
			ID:        user.ID,
			Nickname:  user.Nickname,
			CreatedAt: user.CreatedAt,
			Enabled:   user.Enabled,
			Email:     user.Email,
		},
	}

	return profile, nil
}

// Login 登录
func (s *service) Login(c *core.Context, userLogin *UserLogin) *Profile {
	exists, user, err := s.dao.IsExistsByEmail(userLogin.Email)

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

	videoCount, err := s.dao.GetUserAllVideoCount(user.ID)
	if err != nil {
		panic(err)
	}

	collectionCount, err := s.dao.GetUserAllCollectionCount(user.ID)
	if err != nil {
		panic(err)
	}

	profile := &Profile{
		User: ProfileUser{
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
		FlowerCount:     0, //TODO
	}

	return profile
}

// GetProfile 用户详情
func (s *service) GetProfile(c *core.Context, uid uint) *Profile {
	user, err := UserCache.GetUser(uid)
	if err != nil {
		c.Log.Info("获取user缓存失败 ", zap.Error(err))
	}

	if user == nil {
		user, err = s.dao.GetByID(uid)

		if err != nil {
			panic(err)
		}
	}

	videoCount, err := s.dao.GetUserAllVideoCount(user.ID)

	if err != nil {
		panic(err)
	}
	collectionCount, err := s.dao.GetUserAllCollectionCount(user.ID)
	if err != nil {
		panic(err)
	}

	profile := &Profile{
		User: ProfileUser{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			Enabled:   user.Enabled,
			Nickname:  user.Nickname,
			Email:     user.Email,
		},
		LikeCount:       0,               //TODO
		DislikeCount:    0,               //TODO
		CollectionCount: collectionCount, //TODO
		VideoCount:      videoCount,      //TODO
		FlowerCount:     0,               //TODO
	}

	return profile

}

// ResetPassword 重置密码
func (s *service) ResetPassword(c *core.Context, form *UserResetPassword) error {

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
func (s *service) ResetPasswordByEmailCode(c *core.Context, form *UserResetPassword) error {
	// 1 校验code
	// 2 加密password
	// 3 更新user的password
	// 4 删除code
	uid := *c.UID
	user, err := s.dao.GetByID(uid)
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

	err = s.dao.Updates(user)

	if err != nil {
		return err
	}

	// 删除cache code
	err = s.emailCache.Delete(user.Email, email_action_type.ResetPassword.Code)
	//_, err = s.cache.Del(context.Background(), key).Result()

	return err
}

func (s *service) Logoff(c *core.Context) error {
	//	删除该用户数据、缓存
	err := s.dao.DeleteById(*c.UID)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("user:%d", *c.UID)

	num, err := s.cache.Del(context.Background(), key).Result()
	c.Log.Info(fmt.Sprintf("删除用户缓存: %d ", num))

	if err != nil {
		return err
	}

	return nil
}
