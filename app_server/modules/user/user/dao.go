package user

import (
	"gorm.io/gorm"
	"simple-video-server/db"
	"simple-video-server/models"
)

type _dao struct {
	db *gorm.DB
}

var Dao = &_dao{
	db: db.GetOrmDB(),
}

func (dao *_dao) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := dao.db.Model(&models.User{}).Where("email = ?", email).First(user).Error

	return user, err
}

func (dao *_dao) Create(user *models.User) (uint, error) {
	result := dao.db.Model(&models.User{}).Create(user)

	return user.ID, result.Error
}

func (dao *_dao) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := dao.db.Model(&models.User{}).Where("email = ? ", email).First(&user).Error

	return &user, err
}

func (dao *_dao) FindByEmailAndPassword(email string, password string) (*models.User, error) {
	var user models.User
	err := dao.db.Model(&models.User{}).Where("email = ? and password = ?", email, password).First(&user).Error

	return &user, err
}

func (dao *_dao) FindById(uid uint) (*models.User, error) {
	var user models.User
	err := dao.db.Model(&models.User{}).Where("id = ?", uid).First(&user).Error

	return &user, err
}
