package video

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
	"simple-video-server/db"
	"simple-video-server/models"
	"simple-video-server/pkg/global"
	"time"
)

type VideoDao struct {
	db  *gorm.DB
	sdb *sqlx.DB
}

var Dao = &VideoDao{
	db:  db.GetOrmDB(),
	sdb: db.GetSqlxDB(),
}

func (d *VideoDao) Add(video *models.Video) error {

	err := global.MysqlDB.Model(&models.Video{}).Create(video).Error

	return err
}

func (d *VideoDao) GetById(id uint) (*models.Video, error) {

	var video models.Video
	err := global.MysqlDB.Model(&models.Video{}).Where("id = ?", id).First(&video).Error

	return &video, err
}

func (d *VideoDao) IsCollect(uid, vid uint) (bool, error) {

	var userVideoCollection models.UserVideoCollection
	err := d.db.Model(&models.UserVideoCollection{}).Where("uid = ? AND vid = ?", uid, vid).First(&userVideoCollection).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (d *VideoDao) IsCollectBySqlx(uid, vid uint) (bool, error) {
	sql := `
	SELECT
		count(*) 
	FROM
		user_video_collection 
	WHERE
		uid = ? 
		AND vid = ? 
		LIMIT 1
	`
	var count int64
	err := db.SqlSelectLog(d.sdb.Get, &count, sql, uid, vid)
	if err != nil {
		return false, err
	}

	return count > 0, err
}

func (d *VideoDao) AllCollectionCount(vid uint) (int64, error) {
	var count int64
	err := d.db.Model(&models.UserVideoCollection{}).Where("vid = ? ", vid).Count(&count).Error

	return count, err
}

func (d *VideoDao) Update(uid, vid uint, update *VideoUpdate) error {
	//sql := `UPDATE video set title = ?, cover = ?, updated_at = ? WHERE id = ?`

	//_, err := d.sdb.Exec(sql, update.Title, update.Cover, time.Now(), vid)

	_, err := db.SqlExecLog(d.sdb.Exec, updateVideo, update.Title, update.Cover, time.Now(), vid)
	if err != nil {
		return err
	}

	return nil
}

func (d *VideoDao) Delete(uid, vid uint) error {
	sql := `
	DELETE 
	FROM
		video 
	WHERE
		id = ? 
	AND uid = ? 
	LIMIT 1
`
	result, err := db.SqlExecLog(d.sdb.Exec, sql, vid, uid)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	fmt.Println("删除成功 ", affected)

	return nil
}
