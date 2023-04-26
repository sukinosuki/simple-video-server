package operation_log

import (
	"gorm.io/gorm"
	"simple-video-server/common"
	"simple-video-server/db"
	"simple-video-server/models"
)

type Dao struct {
	clickDB *gorm.DB
}

var _dao = &Dao{
	clickDB: db.GetClickhouseDB(),
}

func GetDao() *Dao {
	return _dao
}

func (dao *Dao) Add(log *models.OperationLog) error {

	err := dao.clickDB.Create(log).Error

	if err != nil {
		//	TODO: 记录记录操作日志失败
	}

	return err
}

func (dao *Dao) GetAll(query *Query) ([]models.OperationLog, error) {
	var logs []models.OperationLog
	if query.OrderField == "" {
		query.OrderField = "create_time"
	}

	if query.Order == "" {
		query.Order = common.Desc.Code
	}

	err := dao.clickDB.
		//Find(&logs). // Find放Offset、Limit、Order前的话Offset、Limit、Order会失效
		Offset(query.GetSafeOffset()).
		Limit(query.GetSafeSize()).
		Order(query.GetOrder()).
		Find(&logs). // Find应放在Offset、Limit、Order后
		Error

	return logs, err
}
