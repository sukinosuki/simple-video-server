package video

import "simple-video-server/common"

type VideoQuery struct {
	//Page int `json:"page"`
	//Size int `json:"size"`
	*common.Pager
}
