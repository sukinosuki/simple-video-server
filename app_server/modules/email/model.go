package email

type SendEmail struct {
	//TODO: 注册时是需要传email字段的, 重置密码也需要传email字段吗
	//TODO: 注册时需要传email字段, 重置密码不用传email字段
	Email string `json:"email" form:"email" binding:"omitempty,lt=50,email"`

	ActionType string `json:"action_type" form:"action_type" binding:"required,oneof=register reset_password"`
}

//type Query struct {
//	Action string `form:"action" binding:"omitempty,oneof=register reset_password"` // 当不为空时只能是register或者reset_password
//}
