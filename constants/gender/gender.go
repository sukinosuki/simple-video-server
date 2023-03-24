package gender

import "simple-video-server/common"

type Gender = common.CodeValue[int, string]

var KeepSecret = &Gender{
	Code:  0,
	Value: "secret",
}

var Male = &Gender{
	Code:  1,
	Value: "male",
}

var Female = &Gender{
	Code:  2,
	Value: "female",
}

var m = make(map[int]*Gender)

func init() {
	m[KeepSecret.Code] = KeepSecret
	m[Male.Code] = Male
	m[Female.Code] = Female
}

func GetByCode(code int) *Gender {

	gender, ok := m[code]
	if ok {
		return gender
	}

	return nil

}
