package user

import (
	"errors"
	"go.uber.org/zap"
	"simple-video-server/core"
	"simple-video-server/models"
	"simple-video-server/pkg/util"
)

type service struct {
	dao *_dao
}

var Service = &service{
	dao: Dao,
}

// Register 注册
func (s *service) Register(c *core.Context, userRegister *UserRegister) *Profile {
	exists, _, err := s.dao.IsExistsByEmail(userRegister.Email)
	if err != nil {
		panic(err)
	}

	if exists {
		//TODO:
		panic(errors.New("该email已被注册"))
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

	_, err = Dao.Add(user)

	if err != nil {
		panic(err)
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

	return profile
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
