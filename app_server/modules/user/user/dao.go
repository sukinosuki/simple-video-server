package user

import (
	"gorm.io/gorm"
	"simple-video-server/db"
	"simple-video-server/models"
)

type Dao struct {
	db    *gorm.DB
	model *models.User
}

var dao = &Dao{
	db:    db.GetOrmDB(),
	model: &models.User{},
}

func (d *Dao) GetAll(query *UserQuery) ([]*UserSimple, error) {

	tx := d.db.Model(d.model)

	var users []*UserSimple
	err := tx.
		Select("user.id, user.nickname, user.avatar, count(1) video_count").
		Joins("LEFT JOIN video ON user.id = video.uid").
		Group("user.id").
		Offset(query.GetSafeOffset()).
		Limit(query.GetSafeSize()).
		Find(&users).Error

	return users, err
}
