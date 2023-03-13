package reset_password_method

import "simple-video-server/common"

type ResetPasswordMethod = common.CodeValue[string, string]

var Email = &ResetPasswordMethod{
	Code:  "email",
	Value: "email",
}

var Sms = &ResetPasswordMethod{
	Code:  "sms",
	Value: "sms",
}
