package upload_class

import "simple-video-server/common"

// UploadClass 上传分类(用来区分上传到oss的目录): 个人视频、番剧、音乐、专栏...
type UploadClass struct {
	common.CodeValue[string]
}

var uploadClassMap = make(map[int]*UploadClass)

var OneVideo = &UploadClass{
	common.CodeValue[string]{
		Code:  1,
		Value: "moment-video",
	},
}

func GetByCode(code int) *UploadClass {
	value, ok := uploadClassMap[code]

	if ok {
		return value
	}

	return nil
}

func init() {
	uploadClassMap[OneVideo.Code] = OneVideo
}
