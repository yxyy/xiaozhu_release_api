package auth

import (
	"context"
	"errors"
	"xiaozhu/api/internal/logic/common"
	"xiaozhu/api/internal/model/user"
)

// Mobile 手机登录
type Mobile struct {
	ctx    context.Context
	Phone  int `json:"phone" form:"phone" gorm:"phone"`
	MbCode int `json:"mb_code" form:"mb_code"`
}

func NewMobile(ctx context.Context) *Mobile {
	return &Mobile{ctx: ctx}
}
func (m *Mobile) verify() error {
	if m.Phone <= 0 {
		return errors.New("账号不能为空")
	}
	if m.MbCode <= 0 {
		return errors.New("账号不能为空")
	}

	return nil
}

func (m *Mobile) login() (memberInfo *user.MemberInfo, err error) {
	if m.Phone <= 0 {
		return memberInfo, errors.New("手机号不能为空")
	}

	return
}

func (m *Mobile) register(req common.RequestForm) (memberInfo *user.MemberInfo, err error) {
	if m.Phone <= 0 {
		return memberInfo, errors.New("手机号不能为空")
	}

	return
}
