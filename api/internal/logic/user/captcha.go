package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"net/smtp"
	"strings"
	"xiaozhu/internal/logic/common"
	"xiaozhu/internal/model/key"
	"xiaozhu/utils"
)

type CaptchaLogic struct {
	ctx context.Context
	CaptchaRequest
}

func NewCaptchaLogic(ctx context.Context) *CaptchaLogic {
	return &CaptchaLogic{ctx: ctx}
}

type CaptchaRequest struct {
	common.RequestForm
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func (l *CaptchaLogic) Captcha() error {

	// todo 迁移到队列上
	var email = Email{Email: l.CaptchaRequest.Email}

	return email.Send()
}

type Email struct {
	Email string `json:"email"`
}

func (m *Email) Send() error {

	if m.Email == "" {
		return errors.New("邮箱缺失")
	}
	smtpHost := viper.GetString("email.host")
	smtpPort := viper.GetString("email.port")
	from := viper.GetString("email.from")
	code := viper.GetString("email.auth")
	to := m.Email
	captcha := utils.Random(6)

	// 认证
	auth := smtp.PlainAuth("", from, code, smtpHost)

	// 邮件头部信息
	header := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: 验证码邮件\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n\r\n"

	// 邮件正文
	body := "小猪验证码：" + captcha + " ,验证码有效期为10分钟，请勿泄漏给别人"

	// 组合邮件
	message := []byte(header + body)

	err := utils.RedisClient.Set(context.Background(), key.CodePrefix+m.Email, captcha, key.CodeExpress).Err()
	if err != nil {
		return fmt.Errorf("设置验证码缓存失败：%s", err)
	}

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil && !strings.Contains(err.Error(), "short response") {
		return err
	}

	return nil

}
