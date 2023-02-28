package business_code

//var (
//	RecodeNotFound = New(404)
//)

var (
	OK             = add(0)
	RequestErr     = add(-400)
	Unauthorized   = add(-401)
	AccessDenied   = add(-403)
	NothingFound   = add(-404)
	ServerErr      = add(-500)
	RecodeNotFound = New(404)
)

var (
	ErrUsernameOrPassword = add(10001)
	UsernameAlreadyExist  = add(10002)
	TokenExpired          = add(10020)
	TokenMalformed        = add(10021)
	TokenNotValidYet      = add(10022)
	TokenInvalid          = add(10023)
	ClaimsCoverErr        = add(10024)
	UploadError           = add(10050)
	TagNameExist          = add(10060)

	EmptyUploadFile = add(40010)
)

var ECodeMap = map[int]string{
	RecodeNotFound.Code():        "数据不存在",
	RequestErr.Code():            "参数错误",
	ServerErr.Code():             "服务器错误",
	NothingFound.Code():          "查询数据不存在",
	ErrUsernameOrPassword.Code(): "用户名或密码错误",
	UsernameAlreadyExist.Code():  "用户名已存在",

	Unauthorized.Code():     "请先登录",
	AccessDenied.Code():     "暂无权限",
	TokenExpired.Code():     "登录过期(token)",
	TokenInvalid.Code():     "Token 无效",
	TokenMalformed.Code():   "Token 格式错误",
	TokenNotValidYet.Code(): "Token 无效",
	ClaimsCoverErr.Code():   "Token 转换失败",

	UploadError.Code():     "上传失败",
	TagNameExist.Code():    "tag已存在",
	EmptyUploadFile.Code(): "上传文件为空",
}
