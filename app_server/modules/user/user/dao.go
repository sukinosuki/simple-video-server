package user

import (
	"errors"
	"gorm.io/gorm"
	"simple-video-server/db"
	"simple-video-server/models"
)

type _dao struct {
	db    *gorm.DB
	model *models.User
}

var Dao = &_dao{
	db:    db.GetOrmDB(),
	model: &models.User{},
}

// IsExistsByEmail 邮箱是否已存在
func (dao *_dao) IsExistsByEmail(email string) (bool, *models.User, error) {
	user := &models.User{}
	tx := dao.db.Model(dao.model)

	err := tx.Where("email = ?", email).First(user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, nil
		}
		return false, nil, err
	}

	return true, user, nil
}

func (dao *_dao) GetUserAllVideoCount(uid uint) (int64, error) {

	tx := dao.db.Model(&models.Video{})

	var count int64
	err := tx.Where("uid = ?", uid).Count(&count).Error

	return count, err
}

func (dao *_dao) GetUserAllCollectionCount(uid uint) (int64, error) {
	tx := dao.db.Model(&models.UserVideoCollection{})
	var count int64
	err := tx.Where("uid = ?", uid).Count(&count).Error

	return count, err
}

// GetByEmail get by email
func (dao *_dao) GetByEmail(email string) (*models.User, error) {
	user := models.User{}
	tx := dao.db.Model(dao.model)

	err := tx.Where("email = ?", email).First(user).Error

	return &user, err
}

// GetByID get by id
func (dao *_dao) GetByID(id uint) (*models.User, error) {
	user := models.User{}

	tx := dao.db.Model(dao.model)
	err := tx.First(&user, id).Error

	return &user, err
}

// Add 新增
func (dao *_dao) Add(user *models.User) (uint, error) {
	err := dao.db.Model(dao.model).Create(user).Error

	return user.ID, err
}

//func (dao *_dao) FindByEmail(email string) (*models.User, error) {
//	var user models.User
//	err := dao.db.Model(&models.User{}).Where("email = ? ", email).First(&user).Error
//
//	return &user, err
//}

func (dao *_dao) FindByEmailAndPassword(email string, password string) (*models.User, error) {
	var user models.User
	err := dao.db.Model(&models.User{}).Where("email = ? and password = ?", email, password).First(&user).Error

	return &user, err
}

// Updates 更新user非0值字段
func (dao *_dao) Updates(user *models.User) error {
	err := dao.db.Model(dao.model).Where("id = ?", user.ID).Updates(user).Error

	return err
}

func (dao *_dao) DeleteById(id uint) error {

	err := dao.db.Model(dao.model).Where("id = ?", id).Limit(1).Delete(&models.User{}).Error

	return err
}
