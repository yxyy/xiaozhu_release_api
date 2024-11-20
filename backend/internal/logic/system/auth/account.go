package auth

import (
	"errors"
	"gorm.io/gorm"
	"xiaozhu/backend/internal/model/system"
	"xiaozhu/backend/utils"
)

// Account 账号登录
type Account struct {
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

func (a *Account) login() (user *system.SysUser, err error) {
	if err = utils.MysqlDb.Model(&user).Where("account", a.Account).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("账号不存在")
		}
		return user, err
	}

	a.Password = utils.Md5SaltAndPassword(user.Salt, a.Password)
	if a.Password != user.Password {
		return user, errors.New("密码错误")
	}

	return
}
