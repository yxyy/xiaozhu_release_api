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
		return &Email{To: l.Message.Email}
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
	SmtpHost string
	SmtpPort string
	From     string
	To       string
	Auth     string
}

func (m *Email) Validate(body string) error {
	m.SmtpHost = viper.GetString("email.host")
	m.SmtpPort = viper.GetString("email.port")
	m.From = viper.GetString("email.From")
	m.Auth = viper.GetString("email.Auth")

	if m.To == "" {
		return errors.New("接收邮箱缺失")
	}
	if body == "" {
		return errors.New("消息主体缺失")
	}

	if m.SmtpHost == "" {
		return errors.New("发送服务器地址缺失")
	}
	if m.SmtpPort == "" {
		m.SmtpPort = "587"
	}
	if m.From == "" {
		return errors.New("发送邮箱缺失")
	}
	if m.Auth == "" {
		return errors.New("发送验证缺失")
	}

	return nil
}

func (m *Email) Send(body string) error {
	// 认证
	auth := smtp.PlainAuth("", m.From, m.Auth, m.SmtpHost)

	// 邮件头部信息
	header := "From: " + m.From + "\r\n" +
		"To: " + m.To + "\r\n" +
		"Subject: 验证码邮件\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n\r\n"

	// 组合邮件
	message := []byte(header + body)

	err := smtp.SendMail(m.SmtpHost+":"+m.SmtpPort, auth, m.From, []string{m.To}, message)
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
