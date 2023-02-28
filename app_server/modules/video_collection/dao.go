package video_collection

type videoCollectionDao struct {
}

var VideoCollectionDao = &videoCollectionDao{}

//func (d *videoCollectionDao) GetById(id uint)  {
//
//	var video
//	global.MysqlDB.Model(&models.VideoCollection{}).Where("id = ?", id)
//}
