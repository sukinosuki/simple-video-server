package like_type

import (
	"simple-video-server/common"
)

//const (
//	Like = iota + 1
//	Dislike
//)

type LikeType struct {
	common.CodeValue[int, string]
}

var likeTypeMap = make(map[int]*LikeType)

var Like = &LikeType{
	common.CodeValue[int, string]{
		Code:        1,
		ValueString: "like",
	},
}

var Dislike = &LikeType{
	common.CodeValue[int, string]{
		Code:        2,
		ValueString: "dislike",
	},
}

func init() {
	likeTypeMap[Like.Code] = Like

	likeTypeMap[Dislike.Code] = Dislike

	//for k, v := range likeTypeMap {
	//	fmt.Println("k ", k)
	//	fmt.Println("v ", *v)
	//	fmt.Println("value ", v.ValueString)
	//}
}

func GetByCode(code int) (*LikeType, bool) {
	for k, v := range likeTypeMap {
		if code == k {
			return v, true
		}
	}

	return nil, false
}

//var Like = common.CodeValue[string]{Code: 1, ValueString: "like"}
//
//var Dislike = common.CodeValue[string]{Code: 1, ValueString: "like"}

//func (c *LikeType) IsLikeType(code int) int {
//	vo := reflect.ValueOf(c)
//	typeVo := vo.Type()
//
//	for i := 0; i < vo.NumField(); i++ {
//		if typeVo.OrderField(i).Name == id {
//			return vo.OrderField(i).Interface().(int)
//		}
//	}
//	return 0
//}

//var LikeTypeMap = make(map[int]*common.CodeValue[string])
//
//var (
//	//LikeCode = common.CodeValue[string]{Code: 1, ValueString: "like"}
//	LikeCode = common.GenerateNewCodeValue(LikeTypeMap, 1, "like")
//
//	DislikeCode = common.GenerateNewCodeValue(LikeTypeMap, 2, "dislike")
//	//DislikeCode = common.CodeValue[string]{Code: 2, ValueString: "dislike"}
//)
