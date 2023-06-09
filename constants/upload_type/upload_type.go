package upload_type

import "simple-video-server/common"

// UploadType 上传类型: 视频、图片...
type UploadType struct {
	common.CodeValue[int, string]
}

var uploadCateMap = make(map[int]*UploadType)

var Video = &UploadType{
	common.CodeValue[int, string]{
		Code:        1,
		ValueString: "video",
	},
}

var Picture = &UploadType{
	common.CodeValue[int, string]{
		Code:        2,
		ValueString: "picture",
	},
}

func NewByCode(code int) *UploadType {
	value, ok := uploadCateMap[code]

	if ok {
		return value
	}

	return nil
}

func init() {
	uploadCateMap[Video.Code] = Video

	uploadCateMap[Picture.Code] = Picture
}
