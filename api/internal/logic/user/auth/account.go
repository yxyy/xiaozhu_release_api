package auth

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"xiaozhu/api/internal/model/user"
	"xiaozhu/api/utils"
)

// Account 账号登录
type Account struct {
	ctx      context.Context
	Account  string `json:"account" form:"account" gorm:"account"`
	Password string `json:"password" form:"password" gorm:"password"`
}

func NewAccount() *Account {
	return &Account{}
}

func (a *Account) verify() error {
	if a.Account == "" {
		return errors.New("账号不能为空")
	}
	if a.Password == "" {
		return errors.New("密码不能为空")
	}

	return nil
}

func (a *Account) login() (memberInfo *user.MemberInfo, err error) {
	if err = utils.MysqlDb.Model(&memberInfo).Where("account", a.Account).First(&memberInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return memberInfo, errors.New("账号不存在")
		}
		return memberInfo, err
	}

	memberInfo = user.NewMemberInfo()
	memberInfo.Account = a.Account
	err = memberInfo.MemberInfo(a.ctx)
	if err != nil {
		return nil, err
	}

	a.Password = utils.Md5SaltAndPassword(memberInfo.Salt, a.Password)

	if a.Password != memberInfo.Password {
		return memberInfo, errors.New("密码错误")
	}

	return
}
