package request_log

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"simple-video-server/common"
	"simple-video-server/db"
	"simple-video-server/models"
	"simple-video-server/pkg/app_ctx"
	"time"
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

func (d *Dao) GetAll(query *Query) ([]models.RequestLog, error) {
	var logs []models.RequestLog

	if query.OrderField == "" {
		query.OrderField = "create_time"
	}

	if query.Order == "" {
		query.Order = common.Desc.Code
	}

	err := d.clickDB.Select("uid", "api_url", "api_method", "trace_id", "res_code", "res_status", "ip", "create_time").
		//Find(&logs). // Find放Offset、Limit、Order前的话Offset、Limit、Order会失效
		Offset(query.GetSafeOffset()).
		Limit(query.GetSafeSize()).
		Order(query.GetOrder()).
		Find(&logs). // Find应放在Offset、Limit、Order后
		Error

	return logs, err
}

func (d *Dao) Add(c *gin.Context, payload, resData string) {
	// 获取响应status code
	statusCode := c.Writer.Status()

	var _uid uint = 0
	uid, ok := app_ctx.GetUid(c)
	if ok {
		_uid = *uid
	}
	traceId, _ := app_ctx.GetTraceId(c)
	var res common.AppResponse[any]

	_ = json.Unmarshal([]byte(resData), &res)
	requestLog := &models.RequestLog{
		UID:        uint64(_uid),
		TraceID:    *traceId,
		IP:         c.ClientIP(),
		CreateTime: time.Now(),
		ApiUrl:     c.Request.URL.Path,
		ApiMethod:  c.Request.Method,
		ApiParams:  c.Request.URL.RawQuery,
		ApiPayload: payload,
		ResData:    resData,
		ResCode:    int64(res.Code),
		ResErrMsg:  res.ErrMsg,
		ResMsg:     res.Msg,
		ResStatus:  int64(statusCode),
		UserAgent:  c.Request.UserAgent(),
	}

	err := d.clickDB.Create(requestLog).Error

	if err != nil {
		fmt.Println("记录request log失败 ", err)
	}
}
