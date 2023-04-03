package http

import (
	"errors"
	"fmt"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"net/smtp"
	emailModule "simple-video-server/app_server/modules/email"
	"simple-video-server/config"
	"simple-video-server/constants/email_action_type"
	"simple-video-server/core"
	"simple-video-server/pkg/business_code"
	"simple-video-server/pkg/util"
)

type _Service struct {
	//db         *gorm.DB
	dao        *emailModule.Dao
	emailCache *emailModule.Cache
}

var service = &_Service{
	//db:         db.GetOrmDB(),
	dao:        emailModule.GetDao(),
	emailCache: emailModule.GetCache(),
}

func (s *_Service) Send(c *core.Context, data *emailModule.SendEmail) (string, error) {

	switch {
	// 注册操作
	case email_action_type.Register.Is(data.ActionType):
		if data.Email == "" {
			panic(business_code.RequestErr)
		}

		code, err := s.handleRegister(c, data)

		return code, err

	// 重置密码操作
	case email_action_type.ResetPassword.Is(data.ActionType):
		// 重置密码需要登录
		if !c.Authorized {
			return "", business_code.Unauthorized
		}
		code, err := s.handleResetPassword(c, data)

		return code, err

	default:
		// TODO
		panic(errors.New("不支持的action_type"))
	}

	return "", nil
}

func _generateEmailCode() string {
	code := util.RandomNumberString(5)

	return code
}

// 重置密码
func (s *_Service) handleResetPassword(c *core.Context, data *emailModule.SendEmail) (string, error) {
	// 获取到邮箱
	//uid := *c.UID
	//var user models.User

	//err := s.db.Model(&models.User{}).First(&user, uid).Error
	//if err != nil {
	//	return "", err
	//}

	user := c.Auth
	code := _generateEmailCode()
	// cache记录reset password操作的code
	err := s.emailCache.Set(user.Email, email_action_type.ResetPassword.Code, code)

	if err != nil {
		return "", err
	}

	// 异步发送邮件
	go func() {
		// 发送邮件
		text := fmt.Sprintf("您的重置密码验证码为 %s", code)
		subject := "重置密码"
		err = _sendEmail([]string{user.Email}, text, subject)

		if err != nil {
			c.Log.Error("发送邮件失败", zap.String("action_type", data.ActionType), zap.String("email", user.Email))
		}
	}()

	return code, err
}

// 处理注册
func (s *_Service) handleRegister(c *core.Context, data *emailModule.SendEmail) (string, error) {
	//	TODO: 校验邮箱是否已注册
	//var count int64
	//s.db.Model(&models.User{}).Where("email = ?", data.Email).Limit(1).Count(&count)
	//
	//if count != 0 {
	//	return "", business_code.RegisterEmailExists
	//}

	exists, err := s.dao.ExistsByEmail(data.Email)
	if err != nil {
		return "", err
	}
	if exists {
		return "", business_code.RegisterEmailExists
	}

	// 生成5位数字字符串的验证码
	code := _generateEmailCode()
	////TODO: 统一定义cache key
	//key := fmt.Sprintf("email_code:%s:%s", data.ActionType, data.Email)
	//// redis保存该邮箱的注册码 key规则为: "email_code:[action_type]:[email]", 有效期为30分钟
	//_, err := s.cache.Set(context.Background(), key, code, 30*time.Minute).Result() //TODO: 有效时间配置化
	err = s.emailCache.Set(data.Email, email_action_type.Register.Code, code)

	if err != nil {
		return "", err
	}

	// code := util.RandomNumberString(5)
	go func() {
		text := fmt.Sprintf("您的注册验证码为 %s", code)
		subject := "用户注册"
		err = _sendEmail([]string{data.Email}, text, subject)

		if err != nil {
			c.Log.Error("发送邮件失败")
		}
	}()

	return code, err
}

func _sendEmail(targetEmails []string, text, subject string) error {

	newEmail := email.NewEmail()

	newEmail.From = config.Email.Email // from必须为授权配置的email
	newEmail.To = targetEmails         // 目标email
	newEmail.Subject = subject
	newEmail.Text = []byte(text)

	addr := "smtp.qq.com:25"
	host := "smtp.qq.com"

	err := newEmail.Send(addr, smtp.PlainAuth("", config.Email.Email, config.Email.AuthCode, host))

	return err
}
