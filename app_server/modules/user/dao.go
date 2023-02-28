package user

import (
	"simple-video-server/models"
	"simple-video-server/pkg/global"
)

type userDao struct {
}

var UserDao = &userDao{}

func (d *userDao) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := global.MysqlDB.Model(&models.User{}).Where("email = ?", email).First(user).Error

	return user, err
}

func (d *userDao) Create(user *models.User) (uint, error) {
	result := global.MysqlDB.Model(&models.User{}).Create(user)

	return user.ID, result.Error
}

func (d *userDao) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := global.MysqlDB.Model(&models.User{}).Where("email = ? ", email).First(&user).Error

	return &user, err
}

func (d *userDao) FindByEmailAndPassword(email string, password string) (*models.User, error) {
	var user models.User
	err := global.MysqlDB.Model(&models.User{}).Where("email = ? and password = ?", email, password).First(&user).Error

	return &user, err
}

func (d *userDao) FindById(uid uint) (*models.User, error) {
	var user models.User
	err := global.MysqlDB.Model(&models.User{}).Where("id = ?", uid).First(&user).Error

	return &user, err
}

// AddCollection 新增收藏
func (d *userDao) AddCollection(collection *models.VideoCollection) error {
	return global.MysqlDB.Model(&models.VideoCollection{}).Create(collection).Error
}

// DeleteCollection 删除收藏
func (d *userDao) DeleteCollection(uid uint, vid uint) error {
	// delete操作记得加上where条件
	err := global.MysqlDB.Model(&models.VideoCollection{}).Where("uid = ? AND vid = ?", uid, vid).Limit(1).Delete(&models.VideoCollection{}).Error

	return err
}

// GetAllCollection TODO:分页
func (d *userDao) GetAllCollection(uid uint) ([]models.VideoCollection, error) {
	var collection []models.VideoCollection
	err := global.MysqlDB.Model(&models.VideoCollection{}).Find(&collection).Error

	return collection, err
}
