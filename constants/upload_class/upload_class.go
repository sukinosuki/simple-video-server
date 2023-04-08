package upload_class

import "simple-video-server/common"

// UploadClass 上传分类(用来区分上传到oss的目录): 个人视频、番剧、音乐、专栏...
type UploadClass struct {
	common.CodeValue[int, string]
}

var uploadClassMap = make(map[int]*UploadClass)

var OneVideo = &UploadClass{
	common.CodeValue[int, string]{
		Code:        1,
		ValueString: "moment-video",
	},
}

var VideoCover = &UploadClass{
	common.CodeValue[int, string]{
		Code:        2,
		ValueString: "moment-video-cover",
	},
}

var UserAvatar = &UploadClass{
	common.CodeValue[int, string]{
		Code:        3,
		ValueString: "user-avatar",
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
	uploadClassMap[VideoCover.Code] = VideoCover
}
