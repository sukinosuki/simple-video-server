package like

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
	"simple-video-server/db"
)

type Dao struct {
	db  *gorm.DB
	sdb *sqlx.DB
}

var _dao *Dao

func GetLikeDao() *Dao {
	if _dao == nil {
		_dao = &Dao{
			db:  db.GetOrmDB(),
			sdb: db.GetSqlxDB(),
		}

		return _dao
	}

	return _dao
}

//func (dao *Dao) Add(userVideoLike *models.UserVideoLike) error {
//	err := dao.db.Model(&models.UserVideoLike{}).Create(userVideoLike).Error
//	return err
//}
//
//func (dao *Dao) Delete(uid, vid uint) error {
//	err := dao.db.Model(&models.UserVideoLike{}).Where("uid = ? AND vid = ?", uid, vid).Limit(1).Delete(&models.UserVideoLike{}).Error
//	//err := dao.db.Unfollow(like).Error
//
//	return err
//}

//func (dao *Dao)GetAll(uid uint)  {
//
//	var sql= `
//`
//}
