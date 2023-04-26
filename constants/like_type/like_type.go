package like_type

import (
	"encoding/json"
	"fmt"
	"simple-video-server/common"
	"simple-video-server/pkg/validation"
)

//const (
//	Like = iota + 1
//	Dislike
//)

type LikeType common.CodeValue[int, string]

var likeTypeMap = make(map[int]*LikeType)

var Like = &LikeType{
	Code:        1,
	ValueString: "like",
}

var Dislike = &LikeType{
	Code:        2,
	ValueString: "dislike",
}

func init() {
	likeTypeMap[Like.Code] = Like

	likeTypeMap[Dislike.Code] = Dislike
}

func (cv *LikeType) Is(code int) bool {

	return cv.Code == code
}

func (cv *LikeType) MarshalJSON() ([]byte, error) {

	bytes, err := json.Marshal(cv.Code)

	return bytes, err
}

func (cv *LikeType) UnmarshalJSON(data []byte) error {

	var k int
	err := json.Unmarshal(data, &k)

	if err != nil {
		return err
	}
	value, ok := likeTypeMap[k]
	// TODO: 提交参数、redis的json都会走自定义的UnmarshalJSON方法(如果redis里的值没有得到对应的数据, 也会走该错误
	if !ok {
		//return errors.New(fmt.Sprintf("不合适的gender value: %d", k))
		return validation.NewValidateError(fmt.Sprintf("不合适的gender value: %d", k))
	}

	*cv = *value

	return nil
}

//func GetByCode(code int) (*LikeType, bool) {
//	for k, v := range likeTypeMap {
//		if code == k {
//			return v, true
//		}
//	}
//
//	return nil, false
//}

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
