package email_action_type

import "simple-video-server/common"

type EmailActionType = common.CodeValue[string, string]

var Register = &EmailActionType{
	Code:  "register",
	Value: "register",
}

var ResetPassword = &EmailActionType{
	Code:  "reset_password",
	Value: "reset_password",
}

var emailActionTypeMap = make(map[string]*EmailActionType)

func init() {
	emailActionTypeMap[Register.Code] = Register

	emailActionTypeMap[ResetPassword.Code] = ResetPassword
}
