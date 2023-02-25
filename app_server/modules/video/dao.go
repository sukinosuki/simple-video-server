package video

import (
	"simple-video-server/models"
	"simple-video-server/pkg/global"
)

type videoDao struct {
}

var VideoDao = &videoDao{}

func (d *videoDao) Add(video *models.Video) error {

	err := global.MysqlDB.Model(&models.Video{}).Create(video).Error

	return err
}

//func (d *videoDao) Update(video) {
//
//}
