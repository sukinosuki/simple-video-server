package email

import (
	"errors"
	"gorm.io/gorm"
	"simple-video-server/db"
	"simple-video-server/models"
)

type Dao struct {
	db *gorm.DB
}

var dao = &Dao{
	db: db.GetOrmDB(),
}

func GetDao() *Dao {
	return dao
}

func (d *Dao) ExistsByEmail(email string) (bool, error) {

	//var count int64
	var user models.User
	err := d.db.Model(&models.User{}).Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
