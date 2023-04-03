package video

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
	"simple-video-server/db"
	"simple-video-server/models"
	"simple-video-server/pkg/global"
)

type Dao struct {
	db    *gorm.DB
	sdb   *sqlx.DB
	model *models.Video
}

var _dao = &Dao{
	db:    db.GetOrmDB(),
	sdb:   db.GetSqlxDB(),
	model: &models.Video{},
}

func GetDao() *Dao {
	return _dao
}

func (d *Dao) Add(video *models.Video) error {
	tx := d.db.Model(d.model)

	err := tx.Create(video).Error

	return err
}

func (d *Dao) GetById2(id uint) (*models.Video, error) {

	var video models.Video
	tx := d.db.Model(d.model)

	err := tx.Where("id = ?", id).First(&video).Error

	return &video, err
}

// GetById get video(with user) by id
func (d *Dao) GetById(id uint) (*VideoResVideo, error) {

	var video VideoResVideo
	err := global.MysqlDB.Model(d.model).
		//Where("video.id = ?", id).
		Select(
			"video.id",
			"video.title",
			"video.created_at",
			"video.cover",
			"video.url",
			"video.uid",
			"user.id user_id",
			"user.nickname user_nickname",
			"user.avatar user_avatar").
		Joins("left join user on user.id = video.uid").
		Where("video.id = ? ", id).
		First(&video).
		Error

	return &video, err
}

func (d *Dao) IsCollect(uid, vid uint) (bool, error) {

	var userVideoCollection models.UserVideoCollection

	tx := d.db.Model(&models.UserVideoCollection{})

	err := tx.
		Where("uid = ? AND vid = ?", uid, vid).
		First(&userVideoCollection).Error

	if err != nil {
		//记录不存在, 返回false为未收藏
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		//返回其它错误在是否需要在service层处理该错误
		return false, err
	}

	//返回true为已收藏
	return true, nil
}

//	GetVideoCollectionCountById func (d *Dao) IsCollectBySqlx(uid, vid uint) (bool, error) {
//		sql := `
//		SELECT
//			count(*)
//		FROM
//			user_video_collection
//		WHERE
//			uid = ?
//			AND vid = ?
//			LIMIT 1
//		`
//		var count int64
//		err := db.SqlSelectLog(d.sdb.Get, &count, sql, uid, vid)
//		if err != nil {
//			return false, err
//		}
//
//		return count > 0, err
//	}
//

// GetVideoCollectionCountById get video's collection count by id
func (d *Dao) GetVideoCollectionCountById(vid uint) (int64, error) {
	var count int64
	tx := d.db.Model(&models.UserVideoCollection{})

	err := tx.
		Where("vid = ? ", vid).
		Count(&count).Error

	return count, err
}

// GetAll 返回视频列表
func (d *Dao) GetAll(uid *uint, query *VideoQuery) ([]VideoSimple, error) {
	var videos []VideoSimple

	tx := d.db.Model(d.model)

	tx.Select("video.id",
		"video.title",
		"video.locked",
		"video.cover",
		"video.url",
		"video.created_at",
		"user.id user_id",
		"user.nickname user_nickname", "user.avatar user_avatar").
		Joins("left join user on user.id = video.uid")

	if uid != nil {
		tx.Where("video.uid = ?", uid)
	}

	if query.Random {
		tx.
			Order("RAND()").
			Limit(query.GetSafeSize())
	} else {
		tx.Order(query.GetOrder()).
			Offset(query.GetSafeOffset()).
			Limit(query.GetSafeSize())
	}
	err := tx.
		Find(&videos).Error

	return videos, err
}

func (d *Dao) Update(uid, vid uint, update *VideoUpdate) error {

	db := d.db.Model(d.model)

	video := models.Video{
		Title: update.Title,
		Cover: update.Cover,
	}

	// 1.save 会保存所有字段，即使是0值
	//err := db.Where("uid = ? AND id = ?", uid, vid).Save(video).Error

	// 2.Select 和 Struct （可以选中更新零值字段）
	//db.Model(&result).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})

	// 3.Updates 方法支持 struct 和 map[string]interface{} 参数。当使用 struct 更新时，默认情况下，GORM 只会更新非零值的字段
	result := db.Where("uid = ? AND id = ?", uid, vid).Updates(video)

	//传参错误会更新0 rows
	//if result.Error == nil && result.RowsAffected == 0 {
	//	// TODO: 更新无效
	//	panic(errors.New("update video err, the video is not belong to the user"))
	//}

	return result.Error
}

// Delete delete one user's video by id
func (d *Dao) Delete(uid, vid uint) error {

	db := d.db.Model(d.model)

	err := db.
		Where("uid = ? AND id = ?", uid, vid).
		Limit(1).
		Delete(d.model).Error

	return err
}
