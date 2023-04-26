package models

import "time"

type OperationLog struct {
	CreateTime time.Time `json:"time"`    //记录时间
	UID        uint64    `json:"uid"`     //操作人
	Module     string    `json:"module"`  //模块名
	Handler    string    `json:"handler"` //处理方法名
	Level      string    `json:"level"`
	TraceID    string    `json:"trace_id"` //trace id
	Err        string    `json:"err"`
	Msg        string    `json:"msg"`
}

func (r *OperationLog) TableName() string {
	return "operation"
}
