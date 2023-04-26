package auth

import (
	"errors"
	"gorm.io/gorm"
	"simple-video-server/db"
	"simple-video-server/models"
)

type Dao struct {
	db    *gorm.DB
	model *models.User
}

var _dao = &Dao{
	db:    db.GetOrmDB(),
	model: &models.User{},
}

func GetAuthDao() *Dao {
	return _dao
}

// IsExistsByEmail 邮箱是否已存在
func (dao *Dao) IsExistsByEmail(email string) (bool, *models.User, error) {
	var user models.User

	tx := dao.db.Model(dao.model)

	err := tx.Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, nil
		}

		return false, nil, err
	}

	return true, &user, nil
}

// GetOneUserAllVideoCount 获取某个用户所有视频数
func (dao *Dao) GetOneUserAllVideoCount(uid uint) (int64, error) {

	tx := dao.db.Model(&models.Video{})

	var count int64
	err := tx.Where("uid = ?", uid).Count(&count).Error

	return count, err
}

// GetUserAllCollectionCount 获取用户视频收藏数
func (dao *Dao) GetUserAllCollectionCount(uid uint) (int64, error) {
	var count int64

	tx := dao.db.Model(&models.UserVideoCollection{})

	err := tx.Where("uid = ?", uid).Count(&count).Error

	return count, err
}

// GetOneByEmail get by email
func (dao *Dao) GetOneByEmail(email string) (*models.User, error) {
	var user models.User

	tx := dao.db.Model(dao.model)

	err := tx.Where("email = ?", email).First(user).Error

	return &user, err
}

// GetOneByID get by id
func (dao *Dao) GetOneByID(id uint) (*models.User, error) {
	var user models.User

	tx := dao.db.Model(dao.model)

	err := tx.First(&user, id).Error

	return &user, err
}

// GetOneByEmailAndPassword find by email and password
func (dao *Dao) GetOneByEmailAndPassword(email string, password string) (*models.User, error) {
	var user models.User

	tx := dao.db.Model(dao.model)

	err := tx.Where("email = ? and password = ?", email, password).
		First(&user).
		Error

	return &user, err
}

// Add 新增
func (dao *Dao) Add(tx *gorm.DB, user *models.User) (id uint, err error) {
	err = tx.Model(dao.model).Create(user).Error

	return user.ID, err
}

// UpdateProfile Update 更新user
func (dao *Dao) UpdateProfile(tx *gorm.DB, user *models.User) error {
	// 更新指定字段
	result := tx.
		//Where("id = ?", user.ID).
		Select("avatar", "nickname", "gender", "birthday").
		Updates(user)

	return result.Error
}

// Updates 更新user非0值字段
func (dao *Dao) Updates(tx *gorm.DB, user *models.User) error {

	err := tx.Model(dao.model).Where("id = ?", user.ID).Updates(user).Error

	return err
}

// DeleteById 删除用户
func (dao *Dao) DeleteById(tx *gorm.DB, uid uint) error {
	err := tx.Model(dao.model).
		Where("id = ?", uid).
		Limit(1).
		Delete(&models.User{}).Error

	return err
}
