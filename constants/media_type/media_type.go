package media_type

import "simple-video-server/common"

type MediaType = common.CodeValue[int, string]

var Video = &MediaType{
	Code:        1,
	ValueString: "video",
}

var Post = &MediaType{
	Code:        2,
	ValueString: "post",
}

var MediaTypeMaps = make(map[int]*MediaType)

func init() {

	MediaTypeMaps[Video.Code] = Video

	MediaTypeMaps[Post.Code] = Post
}
