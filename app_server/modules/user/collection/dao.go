package collection

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
	"simple-video-server/db"
	"simple-video-server/models"
)

type _dao struct {
	db  *gorm.DB
	sdb *sqlx.DB
}

var Dao *_dao

const (
	_queryVideoCollectionSql = `SELECT
		vc.vid,
		vc.created_at,
		v.title,
		v.cover
	FROM
		video_collection vc
		LEFT JOIN video v ON vc.vid = v.id
	WHERE
		vc.uid = ?
	ORDER BY
		vc.created_at DESC
		LIMIT ?,?`
)

func GetCollectionDao() *_dao {
	if Dao == nil {
		Dao = &_dao{
			db:  db.GetOrmDB(),
			sdb: db.GetSqlxDB(),
		}

		return Dao
	}

	return Dao
}

// Add AddCollection 新增收藏
func (dao *_dao) Add(collection *models.UserVideoCollection) error {
	return dao.db.Model(&models.UserVideoCollection{}).Create(collection).Error
}

// Delete 删除收藏
func (dao *_dao) Delete(uid uint, vid uint) error {
	// delete操作记得加上where条件
	err := dao.db.Model(&models.UserVideoCollection{}).Where("uid = ? AND vid = ?", uid, vid).Limit(1).Delete(&models.UserVideoCollection{}).Error

	return err
}

// GetAll TODO:分页
func (dao *_dao) GetAll(uid uint) ([]models.UserVideoCollection, error) {
	var collection []models.UserVideoCollection
	err := dao.db.Model(&models.UserVideoCollection{}).Find(&collection).Error

	return collection, err
}

func (dao *_dao) GetAll2(uid uint) ([]*UserVideoCollectionRes, error) {

	var userVideoCollection []*UserVideoCollectionRes

	//_queryVideoCollectionSql := "SELECT vc.uid,vc.vid,vc.created_at AS collection_time, v.title, v.cover FROM video_collection vc LEFT JOIN video v ON vc.vid = v.id WHERE vc.uid = ? LIMIT ?,?"

	// sql的select字段,一定要在struct里面声明, 不声明会报 "sqlx missing destination name collection_time in *collection.UserVideoCollectionRes" 异常
	// 一: 使用select查询
	err := dao.sdb.Select(&userVideoCollection, _queryVideoCollectionSql, uid, 0, 10)
	if err != nil {
		panic(err)
	}

	////二: 使用Queryx查询
	//rows, err := dao.sdb.Queryx(_queryVideoCollectionSql, uid, 0, 10)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for rows.Next() {
	//	var entity UserVideoCollectionRes
	//
	//	err = rows.StructScan(&entity)
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	userVideoCollection = append(userVideoCollection, &entity)
	//}

	return userVideoCollection, err
}

// IsCollect 用户是否已收藏
func (dao *_dao) IsCollect(uid, vid uint) (bool, error) {
	var userVideoCollection *models.UserVideoCollection

	err := dao.db.Model(&models.UserVideoCollection{}).Where("uid = ? AND vid = ?", uid, vid).First(&userVideoCollection).Error

	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		return false, err
	//	}
	//	return false, err
	//}

	return err == nil, err
}
