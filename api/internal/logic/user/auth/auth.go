package auth

import (
	"xiaozhu/api/internal/model/user"
)

type Loginer interface {
	verify() error
	login() (*user.SysUser, error)
}

type Params struct {
	*Account
	*Mobile
	*WeChat
}

func NewParams() *Params {
	return &Params{}
}

// Login 登录控制
func Login(l Loginer) (user *user.SysUser, err error) {
	if err = l.verify(); err != nil {
		return nil, err
	}

	return l.login()
}
