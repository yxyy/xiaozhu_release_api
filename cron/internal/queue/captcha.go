package queue

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"net/smtp"
	"strings"
	"xiaozhu/internal/config/cache"
	"xiaozhu/internal/model/common"
	"xiaozhu/internal/model/key"
	"xiaozhu/utils"
	"xiaozhu/utils/queue"
)

type CaptchaQueue struct {
	common.TopicExtra
	Message CaptchaRequest
}

type CaptchaRequest struct {
	common.RequestForm
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func NewCaptchaQueue() *queue.Queue {
	return queue.NewQueue(key.CodeQueue, &CaptchaQueue{})
}

func (l *CaptchaQueue) Run(q *queue.Queue, topic string) error {
	if err := json.Unmarshal([]byte(topic), &l); err != nil {
		return fmt.Errorf("序列化失败：%w", err)
	}
	dispatcher := l.NewDispatcher()
	if dispatcher == nil {
		return fmt.Errorf("无效的验证码获取类型")
	}

	body, err := l.NewCodeBody()
	if err != nil {
		return fmt.Errorf("生成验证码错误：%w", err)
	}

	if err = dispatcher.Validate(body); err != nil {
		return fmt.Errorf("校验错误：%w", err)
	}

	return dispatcher.Send(body)
}

func (l *CaptchaQueue) NewDispatcher() Dispatcher {
	if l.Message.Phone != "" {
		return &Phone{Phone: l.Message.Phone}
	}

	if l.Message.Email != "" {
		return &Email{Email: l.Message.Email}
	}

	return nil
}

func (l *CaptchaQueue) NewCodeBody() (string, error) {
	captcha := utils.Random(6)
	device := l.Message.Phone
	if device == "" && l.Message.Email != "" {
		device = l.Message.Email
	}
	err := cache.RedisDB00.Set(context.Background(), key.CodePrefix+device, captcha, key.CodeExpress).Err()
	if err != nil {
		return "", fmt.Errorf("设置验证码缓存失败：%s", err)
	}

	body := "小猪验证码：" + captcha + " ,验证码有效期为10分钟，请勿泄漏给别人"

	return body, nil
}

type Dispatcher interface {
	Validate(body string) error
	Send(string) error
}

type Email struct {
	Email string `json:"email"`
}

func (m *Email) Validate(body string) error {
	if m.Email == "" {
		return errors.New("邮箱缺失")
	}
	if body == "" {
		return errors.New("消息主体缺失")
	}
	return nil
}

func (m *Email) Send(body string) error {
	smtpHost := viper.GetString("email.host")
	smtpPort := viper.GetString("email.port")
	from := viper.GetString("email.from")
	code := viper.GetString("email.auth")
	to := m.Email

	// 认证
	auth := smtp.PlainAuth("", from, code, smtpHost)

	// 邮件头部信息
	header := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: 验证码邮件\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n\r\n"

	// 组合邮件
	message := []byte(header + body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil && !strings.Contains(err.Error(), "short response") {
		return err
	}

	return nil

}

type Phone struct {
	Phone string `json:"phone"`
}

func (p *Phone) Validate(body string) error {
	if p.Phone == "" {
		return errors.New("电话缺失")
	}
	if body == "" {
		return errors.New("消息主体缺失")
	}
	return nil
}

func (p *Phone) Send(body string) error {

	return nil
}
