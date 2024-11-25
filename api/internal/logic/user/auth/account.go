package auth

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
	"xiaozhu/api/internal/logic/common"
	"xiaozhu/api/internal/model/user"
	"xiaozhu/api/utils"
)

// Account 账号登录
type Account struct {
	ctx      context.Context
	Account  string `json:"account" form:"account" gorm:"account"`
	Password string `json:"password" form:"password" gorm:"password"`
}

func NewAccount(ctx context.Context) *Account {
	return &Account{ctx: ctx}
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
	// if err = utils.MysqlDb.Model(&memberInfo).Where("account", a.Account).First(&memberInfo).Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return memberInfo, errors.New("账号不存在")
	// 	}
	// 	return memberInfo, err
	// }

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

func (a *Account) register(req common.RequestForm) (memberInfo *user.MemberInfo, err error) {

	memberInfo = user.NewMemberInfo()
	memberInfo.Account = a.Account
	err = memberInfo.MemberInfo(a.ctx)
	if err == nil && memberInfo.Id > 0 {
		return memberInfo, errors.New("账号已经存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return memberInfo, err
	}

	// memberInfo = user.NewMemberInfo()
	// memberInfo.Account = a.Account
	memberInfo.Salt = utils.Random(8)
	memberInfo.Password = utils.Md5SaltAndPassword(memberInfo.Salt, a.Password)
	memberInfo.GameId = req.GameId
	memberInfo.RegIp = req.Ip
	memberInfo.RegTime = time.Now().Unix()
	memberInfo.DeviceId = req.DeviceId

	err = memberInfo.Create(a.ctx)
	if err != nil {
		return nil, err
	}

	//
	memberInfo.Salt = ""
	memberInfo.Password = ""

	return
}
