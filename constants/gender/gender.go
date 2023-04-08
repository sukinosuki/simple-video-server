package gender

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"simple-video-server/common"
	"simple-video-server/pkg/validation"
	"strconv"
	"time"
)

type Gender common.CodeValue[int, string]

var KeepSecret = &Gender{
	Code:        0,
	ValueString: "secret",
}

var Male = &Gender{
	Code:        1,
	ValueString: "male",
}

var Female = &Gender{
	Code:        2,
	ValueString: "female",
}

var m = make(map[int]*Gender)

func init() {
	m[KeepSecret.Code] = KeepSecret
	m[Male.Code] = Male
	m[Female.Code] = Female
}

func (cv *Gender) MarshalJSON() ([]byte, error) {

	bytes, err := json.Marshal(cv.Code)

	return bytes, err
}

func (cv *Gender) UnmarshalJSON(data []byte) error {

	var k int
	err := json.Unmarshal(data, &k)

	if err != nil {
		return err
	}
	value, ok := m[k]
	// TODO: 提交参数、redis的json都会走自定义的UnmarshalJSON方法(如果redis里的值没有得到对应的数据, 也会走该错误
	if !ok {
		//return errors.New(fmt.Sprintf("不合适的gender value: %d", k))
		return validation.NewValidateError(fmt.Sprintf("不合适的gender value: %d", k))
	}

	*cv = *value

	return nil
}

// 写入 mysql 时调用
func (cv *Gender) Value() (driver.Value, error) {
	//// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	//if t.String() == "0001-01-01 00:00:00" {
	//	return nil, nil
	//}
	//return []byte(time.Time(t).Format(TimeFormat)), nil
	// 如果是空值, 解析成nil还是一个默认的cv
	if cv == nil {
		return nil, nil
		//bytes, err := json.Marshal(KeepSecret)
		//if err != nil {
		//	return nil, err
		//}
		//
		//return bytes, nil
	}

	bytes, _ := json.Marshal(cv.Code)

	return bytes, nil
}

// Scan gorm scan
func (cv *Gender) Scan(v any) error {

	v2 := v.(int64)
	value, _ := m[int(v2)]

	*cv = *value

	return nil
}

// 用于 fmt.Println 和后续验证场景
func (cv *Gender) String() string {
	return strconv.Itoa(cv.Code)
}

type User struct {
	Gender Gender `json:"gender"`
	Name   string `json:"name"`
	//Birthday time.Time `json:"birthday"`//不为指针类型时, 无论是否有定义Birthday值, 序列化、反序列化会走 Time类型自定义的MarshalJSON、UnmarshalJSON方法

	//为指针类型时, 有定义Birthday值时,序列化、反序列化会走 Time类型自定义的MarshalJSON、UnmarshalJSON方法.
	//Birthday值为nil时不会走序列化
	Birthday *time.Time `json:"birthday"`
}

func TestGender() {
	//now := time.Now()
	user := User{
		Gender: *Male,
		Name:   "hanami",
		//Birthday: &now,
	}
	bytes, err := json.Marshal(&user)
	if err != nil {
		panic(err)
	}

	var user2 User
	err = json.Unmarshal(bytes, &user2)
	if err != nil {
		panic(err)
	}

	fmt.Println("user2 ", user2)

}
