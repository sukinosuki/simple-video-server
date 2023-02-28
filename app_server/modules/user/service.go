package user

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"simple-video-server/app_server/modules/video"
	"simple-video-server/constants/video_status"
	"simple-video-server/core"
	"simple-video-server/models"
	"simple-video-server/util"
)

type service struct {
	dao *userDao
}

var Service = &service{
	dao: UserDao,
}

// Register 注册
func (s *service) Register(c *core.Context, userRegister *UserRegister) *models.User {
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
func (s *service) Login(c *core.Context, userLogin *UserLogin) *models.User {
	c.Log.Info("service login start ", zap.Any("data ", *userLogin))

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

func (s *service) AddCollection(c *core.Context, vid uint) error {
	//TODO:验证video是否为生效状态(是否存在、删除、审核通过
	video, err := video.VideoDao.GetById(vid)
	//TODO:不需要让用户知道详情的信息，可以直接返回添加失败
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errors.New("视频不存在"))
		}

		return err
	}

	// 校验视频是否被锁定、是否审核通过
	if video.Locked || video.Status != video_status.AuditPermit {
		// 可以直接返回添加收藏失败
		//panic(errors.New("视频被锁定或者审核未通过"))
		c.Log.Info("视频被锁定或者审核未通过")
	}

	collection := &models.VideoCollection{
		UID: *c.UID,
		VID: vid,
	}

	err = s.dao.AddCollection(collection)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteCollection(c *core.Context, vid uint) error {
	err := s.dao.DeleteCollection(*c.UID, vid)
	if err != nil {
		return err
	}

	return nil
}

// GetAllCollection TODO:分页
func (s *service) GetAllCollection(c *core.Context) ([]models.VideoCollection, error) {
	collections, err := s.dao.GetAllCollection(*c.UID)

	return collections, err
}
