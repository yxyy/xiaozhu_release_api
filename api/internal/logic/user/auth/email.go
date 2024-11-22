package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"xiaozhu/api/internal/model/key"
	"xiaozhu/api/internal/model/user"
	"xiaozhu/api/utils"
)

// Email  邮箱登录
type Email struct {
	ctx    context.Context
	Email  string `json:"email" form:"code"`
	EmCode string `json:"em_code" form:"em_code"`
}

func NewEmail() *Email {
	return &Email{}
}
func (m *Email) verify() error {
	if m.Email == "" {
		return errors.New("邮箱不能为空")
	}

	fmt.Printf("%#v\n", m)

	if m.EmCode == "" || len(m.EmCode) != 6 {
		return errors.New("验证码无效")
	}

	keys := key.CodePrefix + m.Email
	result, err := utils.RedisClient.Get(m.ctx, keys).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return fmt.Errorf("验证码错误")
		}
		return fmt.Errorf("获取验证码失败：%s", err)
	}

	if m.EmCode != result {
		return errors.New("验证码不正确")
	}

	return nil
}

func (m *Email) login() (memberInfo *user.MemberInfo, err error) {
	memberInfo = user.NewMemberInfo()
	memberInfo.Email = m.Email
	err = memberInfo.MemberInfo(m.ctx)

	return
}
