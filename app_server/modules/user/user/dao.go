package user

import (
	"gorm.io/gorm"
	"simple-video-server/common"
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

func GetDao() *Dao {
	return _dao
}

func (d *Dao) GetRanks(query *UserQuery) ([]*RankUsers, error) {

	tx := d.db.Model(d.model)

	var users []*RankUsers
	err := tx.
		Select("user.id user_id, user.nickname user_nickname, user.avatar user_avatar, count(1) video_count").
		Joins("LEFT JOIN video ON user.id = video.uid").
		Group("user.id").
		Offset(query.GetSafeOffset()).
		Limit(query.GetSafeSize()).
		Find(&users).Error

	return users, err
}

func (d *Dao) GetByIdIn(ids []uint, pager *common.Pager) ([]models.User, error) {
	var users []models.User

	err := d.db.Model(&models.User{}).Where("id in ?", ids).
		Offset(pager.GetSafeOffset()).
		Limit(pager.GetSafeSize()).
		Find(&users).Error

	return users, err
}
