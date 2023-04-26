package comment

import (
	"errors"
	"gorm.io/gorm"
	"simple-video-server/db"
	"simple-video-server/models"
)

type Dao struct {
	db    *gorm.DB
	model *models.Comment
}

var _dao = &Dao{
	db:    db.GetOrmDB(),
	model: &models.Comment{},
}

func GetDao() *Dao {
	return _dao
}

func (d *Dao) Create(tx *gorm.DB, comment *models.Comment) error {

	//tx := d.db.Model(d.model)

	err := tx.Model(d.model).Create(comment).Error

	return err
}

func (d *Dao) DeleteByIdAndUidAndIdAndType(tx *gorm.DB, uid uint, mediaType int, mediaId uint, id uint) error {
	err := tx.Model(d.model).
		Where("id = ? AND media_id = ? AND uid = ? AND media_type = ?", id, mediaId, uid, mediaType).
		Delete(&models.Comment{}).Error

	return err
}

func (d *Dao) GetMediaComment(mediaId uint) ([]*models.Comment, error) {
	tx := d.db.Model(d.model)
	var comment []*models.Comment
	err := tx.
		Where("media_id = ?", mediaId).
		Find(&comment).Error

	return comment, err
}

func (d *Dao) IsVideoExists(mediaType int, mediaId uint) (bool, *models.Video, error) {
	var video models.Video
	err := d.db.Model(&models.Video{}).Where("id = ?", mediaId).First(&video).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, nil
		}

		return false, nil, err
	}

	return true, &video, err
}

func (d *Dao) GetUserById(uid uint) (*models.User, error) {
	var user models.User
	tx := d.db.Model(&models.User{})
	err := tx.First(&user, uid).Error

	return &user, err
}
