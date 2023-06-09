package collection

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
	"simple-video-server/db"
	"simple-video-server/models"
)

type Dao struct {
	db    *gorm.DB
	sdb   *sqlx.DB
	model *models.UserVideoCollection
}

var _dao = &Dao{
	db:    db.GetOrmDB(),
	sdb:   db.GetSqlxDB(),
	model: &models.UserVideoCollection{},
}

func GetDao() *Dao {
	return _dao
}

// Add AddCollection 用户新增收藏
func (d *Dao) Add(collection *models.UserVideoCollection) error {
	return d.db.Model(d.model).Create(collection).Error
}

// Delete 删除用户收藏
func (d *Dao) Delete(uid uint, vid uint) error {
	// delete操作记得加上where条件
	tx := d.db.Model(d.model)
	err := tx.Where("uid = ? AND vid = ?", uid, vid).
		Limit(1).
		Delete(d.model).Error

	return err
}

// GetAll get user's video collection
func (d *Dao) GetAll(uid uint, query *CollectionQuery) ([]*UserVideoCollectionRes, error) {
	var collection []*UserVideoCollectionRes

	tx := d.db.Model(d.model)

	err := tx.Where("user_video_collection.uid = ?", uid).
		Select("video.id", "video.title", "video.cover", "video.created_at", "user.id user_id", "user.nickname user_nickname", "user.avatar user_avatar").
		Joins("left join video on user_video_collection.vid = video.id").
		Joins("left join user on user.id = video.uid").
		Order("created_at desc").
		Offset(query.GetSafeOffset()).
		Limit(query.GetSafeSize()).
		Find(&collection).Error

	return collection, err
}

func (d *Dao) IsVideoExists(vid uint) (bool, *models.Video, error) {
	tx := d.db.Model(&models.Video{})

	var video models.Video

	err := tx.First(&video, vid).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, nil
		}

		return false, nil, err
	}

	return true, &video, nil
}

//func (d *Dao) GetAll2(uid uint) ([]*UserVideoCollectionRes, error) {
//
//	var userVideoCollection []*UserVideoCollectionRes
//
//	//_queryVideoCollectionSql := "SELECT vc.uid,vc.vid,vc.created_at AS collection_time, v.title, v.cover FROM video_collection vc LEFT JOIN video v ON vc.vid = v.id WHERE vc.uid = ? LIMIT ?,?"
//
//	// sql的select字段,一定要在struct里面声明, 不声明会报 "sqlx missing destination name collection_time in *collection.UserVideoCollectionRes" 异常
//	// 一: 使用select查询
//	err := d.sdb.Select(&userVideoCollection, _queryVideoCollectionSql, uid, 0, 10)
//	if err != nil {
//		panic(err)
//	}
//
//	////二: 使用Queryx查询
//	//rows, err := d.sdb.Queryx(_queryVideoCollectionSql, uid, 0, 10)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//
//	//for rows.Next() {
//	//	var entity UserVideoCollectionRes
//	//
//	//	err = rows.StructScan(&entity)
//	//
//	//	if err != nil {
//	//		panic(err)
//	//	}
//	//
//	//	userVideoCollection = append(userVideoCollection, &entity)
//	//}
//
//	return userVideoCollection, err
//}

// IsCollect 用户是否已收藏
func (d *Dao) IsCollect(uid, vid uint) (bool, error) {
	tx := d.db.Model(d.model)

	var collection models.UserVideoCollection
	//err := tx.Where("uid = ? AND vid = ?", uid, vid).First(&userVideoCollection).Error
	err := tx.Where("uid = ? AND vid = ?", uid, vid).First(&collection).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
