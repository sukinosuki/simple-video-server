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
