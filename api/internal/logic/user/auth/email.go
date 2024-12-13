package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"time"
	"xiaozhu/internal/config"
	"xiaozhu/internal/logic/common"
	"xiaozhu/internal/model/key"
	"xiaozhu/internal/model/user"
	"xiaozhu/utils"
)

// Email  邮箱登录
type Email struct {
	ctx    context.Context
	Email  string `json:"email" form:"code"`
	EmCode string `json:"em_code" form:"em_code"`
}

func NewEmail(ctx context.Context) *Email {
	return &Email{ctx: ctx}
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
	result, err := config.RedisDB00.Get(m.ctx, keys).Result()
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

func (m *Email) register(req common.RequestForm) (memberInfo *user.MemberInfo, err error) {

	memberInfo = user.NewMemberInfo()
	memberInfo.Account = m.Email
	err = memberInfo.MemberInfo(m.ctx)
	if err == nil && memberInfo.Id > 0 {
		return memberInfo, errors.New("账号已经存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return memberInfo, err
	}

	memberInfo.Account = m.Email
	memberInfo.Salt = utils.Random(8)
	memberInfo.Password = utils.Md5SaltAndPassword(memberInfo.Salt, utils.Random(12))
	memberInfo.GameId = req.GameId
	memberInfo.RegIp = req.Ip
	memberInfo.RegTime = time.Now().Unix()
	memberInfo.DeviceId = req.DeviceId
	memberInfo.Email = m.Email

	err = memberInfo.Create(m.ctx)
	if err != nil {
		return nil, err
	}

	//
	memberInfo.Salt = ""
	memberInfo.Password = ""

	return

}
