package auth

import (
	"xiaozhu/internal/config/mysql"
	"xiaozhu/internal/model/system"
)

type Loginer interface {
	verify() error
	login() (*system.SysUser, error)
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
func Login(l Loginer) (user *system.SysUser, err error) {
	if err = l.verify(); err != nil {
		return nil, err
	}

	return l.login()
}

func RoleMenu() (resp []system.RoleMenu, err error) {

	var list []system.SysMenus
	menu := system.NewSysMenus()
	if err = mysql.PlatformDB.Model(&menu).Find(&list).Error; err != nil {
		return nil, err
	}

	for _, v := range list {
		resp = append(resp, system.RoleMenu{
			Id:       v.Id,
			ParentId: v.ParentId,
			MenuType: v.MenuType,
			Title:    v.Title,
		})
	}

	return resp, nil

}
