package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"simple-video-server/app_server/modules/auth"
	"simple-video-server/app_server/modules/email"
	"simple-video-server/app_server/modules/follow"
	"simple-video-server/constants/email_action_type"
	"simple-video-server/constants/reset_password_method"
	"simple-video-server/core"
	"simple-video-server/db"
	"simple-video-server/models"
	"simple-video-server/pkg/app_jwt"
	"simple-video-server/pkg/business_code"
	"simple-video-server/pkg/util"
)

type _Service struct {
	authDao     *auth.Dao
	userCache   *auth.UserCache
	followCache *follow.Cache
	emailCache  *email.Cache
	db          *gorm.DB
	redisClient *redis.Client
}

var service = &_Service{
	authDao:     auth.GetAuthDao(),
	userCache:   auth.GetUserCache(),
	emailCache:  email.GetCache(),
	followCache: follow.GetCache(),
	db:          db.GetOrmDB(),
	redisClient: db.GetRedisClient(),
}

// 验证注册的email code
func (s *_Service) verifyEmailRegisterCode(email string, code string) (bool, error) {
	_code, err := s.emailCache.Get(email, email_action_type.Register.Code)
	if err != nil {
		// TODO:
		return false, errors.New("找不到验证码")
	}

	if code != _code {
		// TODO
		return false, errors.New("邮箱验证码不正确")
	}

	return false, nil
}

// Register 注册
func (s *_Service) Register(c *core.Context, userRegister *auth.RegisterForm) *auth.LoginRes {
	handlerName := "Register"

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

	// 从cache获取对应类型的email的验证码
	_, err := s.verifyEmailRegisterCode(userRegister.Email, userRegister.Code)
	c.PanicIfErr(err, handlerName, "校验email失败")

	// 校验email是否注册过
	exists, _, err := s.authDao.IsExistsByEmail(userRegister.Email)
	core.PanicIfErr(err)

	if exists {
		err = errors.New("该email已被注册")
		c.Panic(err, handlerName, "该email已被注册")
	}

	//hashed password
	hashedPassword, err := util.Password.Hashed(userRegister.Password)
	c.PanicIfErr(err, handlerName, "加载密码失败")

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
	handlerName := "Login"
	c.Log.Info("登录开始", zap.String("handler", "login"))

	exists, user, err := s.authDao.IsExistsByEmail(userLogin.Email)
	c.PanicIfErr(err, handlerName, "邮箱获取用户失败")

	if !exists {
		c.Panic(business_code.ErrUsernameOrPassword, handlerName, "email不存在")
		return nil
	}

	// 对比密码
	err = util.Password.Compare(user.Password, userLogin.Password)
	if err != nil {
		c.Panic(business_code.ErrUsernameOrPassword, handlerName, "对比密码失败")
		return nil
	}

	err = s.userCache.SetUser(user.ID, user, 0)
	c.PanicIfErr(err, handlerName, "设置user cache失败")

	videoCount, err := s.authDao.GetOneUserAllVideoCount(user.ID)
	c.PanicIfErr(err, handlerName, "获取用户视频数失败")

	collectionCount, err := s.authDao.GetUserAllCollectionCount(user.ID)
	c.PanicIfErr(err, handlerName, "获取用户收藏数失败")

	token, err := app_jwt.AppJwt.Create(user.ID)

	if err != nil {
		c.PanicIfErr(err, handlerName, "生成token失败")
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
			FollowerCount:   0, //TODO
			CollectionCount: collectionCount,
			VideoCount:      videoCount,
		},
	}

	return loginRes
}

// GetProfile 用户详情
func (s *_Service) GetProfile(c *core.Context, uid uint) *auth.LoginResProfile {
	handlerName := "GetProfile"
	// TODO: 可以从cache获取
	user, err := s.authDao.GetOneByID(uid)
	c.PanicIfErr(err, handlerName, "获取用户失败")

	videoCount, err := s.authDao.GetOneUserAllVideoCount(user.ID)
	c.PanicIfErr(err, handlerName, "获取用户全部视频数失败")

	collectionCount, err := s.authDao.GetUserAllCollectionCount(user.ID)
	c.PanicIfErr(err, handlerName, "获取用户全部收藏数失败")

	followersCount, _ := s.followCache.GetOneUserFollowersCount(uid)
	followingCount, _ := s.followCache.GetOneUserFollowingCount(uid)

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
	//email更新密码
	case reset_password_method.Email.Is(form.Method):
		// todo
		err := s.ResetPasswordByEmail(c, form)

		return err
		// sms更新密码
	case reset_password_method.Sms.Is(form.Code):
	//	TODO
	default:
		panic(errors.New("不支持的method"))
	}

	return nil
}

// ResetPasswordByEmail 通过邮箱重置密码
func (s *_Service) ResetPasswordByEmail(c *core.Context, form *auth.ResetPasswordForm) error {
	handlerName := "ResetPasswordByEmail"

	tx := s.db.Begin()

	defer func() {
		err := recover()

		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
			panic(err)
		}
	}()

	// 1 校验code
	// 2 加密password
	// 3 更新user的password
	// 4 删除code
	uid := *c.AuthUID
	user, err := s.authDao.GetOneByID(uid)
	c.PanicIfErr(err, handlerName, "uid获取用户失败")

	// 比对缓存code和提交的code
	cachedEmailCode, err := s.emailCache.Get(user.Email, email_action_type.ResetPassword.Code)
	c.PanicIfErr(err, handlerName, fmt.Sprintf("缓存获取email code失败, email: %s, code: %s, cached_code: %s", user.Email, form.Code, cachedEmailCode))

	if cachedEmailCode != form.Code {
		c.Panic(errors.New("code不正确"), handlerName, "提交code与缓存code对应失败")
	}

	// 加密密码
	hashedPassword, err := util.Password.Hashed(form.Password)
	c.PanicIfErr(err, handlerName, "加密密码错误")

	// 更新user的密码字段
	user = &models.User{
		ID:       user.ID,
		Password: hashedPassword,
	}

	err = s.authDao.Updates(tx, user)
	c.PanicIfErr(err, handlerName, "更新用户密码失败")

	// 删除cache code
	_, err = s.emailCache.Delete(user.Email, email_action_type.ResetPassword.Code)
	c.PanicIfErr(err, handlerName, "email cache删除code失败")

	// 操作记录
	c.Log.Info("用户重置密码成功", zap.String("handlerName", handlerName))

	return nil
}

// UpdateProfile 更新profile
func (s *_Service) UpdateProfile(c *core.Context, form *auth.UpdateForm) error {
	handlerName := "UpdateProfile"
	tx := s.db.Begin()
	log := c.Log.With(zap.String("handler", handlerName))

	defer func() {
		err := recover()

		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
			panic(err)
		}
	}()

	user := models.User{
		ID:       *c.AuthUID,
		Nickname: form.Nickname,
		Avatar:   form.Avatar,
		Gender:   form.Gender,
		Birthday: form.Birthday,
	}

	err := s.authDao.UpdateProfile(tx, &user)
	c.PanicIfErr(err, handlerName, "更新用户profile失败")

	log.Info("用户更新profile")

	return err
}

// Logoff 注销
func (s *_Service) Logoff(c *core.Context) error {
	handlerName := "Logoff"
	uid := *c.AuthUID
	redisTx := s.redisClient.TxPipeline()

	tx := s.db.Begin()

	defer func() {
		err := recover()

		if err == nil {
			// redis提交事务
			_, _err := redisTx.Exec(context.Background())
			if _err != nil {
				panic(_err)
			}
			// gorm提交事务
			tx.Commit()
		} else {
			tx.Rollback()
			panic(err)
		}
	}()

	//	删除该用户数据、缓存
	err := s.authDao.DeleteById(tx, *c.AuthUID)
	c.PanicIfErr(err, handlerName, "删除用户失败")

	err = s.userCache.Delete(redisTx, uid)
	c.PanicIfErr(err, handlerName, "删除用户缓存失败")

	return nil
}
