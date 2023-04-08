package reset_password_method

import "simple-video-server/common"

type ResetPasswordMethod = common.CodeValue[string, string]

var Email = &ResetPasswordMethod{
	Code:        "email",
	ValueString: "email",
}

var Sms = &ResetPasswordMethod{
	Code:        "sms",
	ValueString: "sms",
}
