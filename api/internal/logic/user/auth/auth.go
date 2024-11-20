package auth

import (
	"xiaozhu/api/internal/model/user"
	"xiaozhu/api/utils"
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

func RoleMenu() (resp []user.RoleMenu, err error) {

	var list []user.SysMenus
	menu := user.NewSysMenus()
	if err = utils.MysqlDb.Model(&menu).Find(&list).Error; err != nil {
		return nil, err
	}

	for _, v := range list {
		resp = append(resp, user.RoleMenu{
			Id:       v.Id,
			ParentId: v.ParentId,
			MenuType: v.MenuType,
			Title:    v.Title,
		})
	}

	return resp, nil

}
