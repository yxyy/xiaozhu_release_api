package auth

import (
	"errors"
	"xiaozhu/backend/internal/model/system"
)

// Mobile 手机登录
type Mobile struct {
	Phone int `json:"phone" form:"phone" gorm:"phone"`
	Code  int `json:"code" form:"code"`
}

func NewMobile() *Mobile {
	return &Mobile{}
}
func (m *Mobile) verify() error {
	if m.Phone <= 0 {
		return errors.New("账号不能为空")
	}
	if m.Code <= 0 {
		return errors.New("账号不能为空")
	}

	return nil
}

func (m *Mobile) login() (user *system.SysUser, err error) {
	if m.Phone <= 0 {
		return user, errors.New("手机号不能为空")
	}

	return
}
