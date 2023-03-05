package user

import (
	"gorm.io/gorm"
	"simple-video-server/db"
	"simple-video-server/models"
)

type _userDao struct {
	db *gorm.DB
}

var Dao *_userDao

func GetUserDao() *_userDao {
	if Dao == nil {
		Dao = &_userDao{
			db: db.GetOrmDB(),
		}

		return Dao
	}

	return Dao
}

func (dao *_userDao) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := dao.db.Model(&models.User{}).Where("email = ?", email).First(user).Error

	return user, err
}

func (dao *_userDao) Create(user *models.User) (uint, error) {
	result := dao.db.Model(&models.User{}).Create(user)

	return user.ID, result.Error
}

func (dao *_userDao) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := dao.db.Model(&models.User{}).Where("email = ? ", email).First(&user).Error

	return &user, err
}

func (dao *_userDao) FindByEmailAndPassword(email string, password string) (*models.User, error) {
	var user models.User
	err := dao.db.Model(&models.User{}).Where("email = ? and password = ?", email, password).First(&user).Error

	return &user, err
}

func (dao *_userDao) FindById(uid uint) (*models.User, error) {
	var user models.User
	err := dao.db.Model(&models.User{}).Where("id = ?", uid).First(&user).Error

	return &user, err
}
