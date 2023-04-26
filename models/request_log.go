package models

import "time"

type RequestLog struct {
	UID        uint64    `json:"uid"`
	TraceID    string    `json:"trace_id"`
	ApiUrl     string    `json:"api_url"`
	ApiMethod  string    `json:"api_method"`
	ApiParams  string    `json:"api_params"`
	ApiPayload string    `json:"api_payload"`
	ResCode    int64     `json:"res_code"`
	ResData    string    `json:"res_data"`
	ResStatus  int64     `json:"res_status"`
	ResMsg     string    `json:"res_msg"`
	ResErrMsg  string    `json:"res_err_msg"`
	CreateTime time.Time `json:"create_time"`
	IP         string    `json:"ip"`
	UserAgent  string    `json:"user_agent"`
}

func (r *RequestLog) TableName() string {
	return "request"
}
