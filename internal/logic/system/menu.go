package system

import (
	"context"
	"encoding/json"
	"errors"
	"time"
	"xiaozhu/internal/model/system"
	"xiaozhu/utils"
)

type MenuLogic struct {
	ctx context.Context
	MenuList
}

func NewMenuLogic(ctx context.Context) *MenuLogic {
	return &MenuLogic{
		ctx: ctx,
	}
}

type MenuList struct {
	Id              int    `json:"id"`
	ParentId        int    `json:"parentId,optional"`
	MenuType        int8   `json:"menuType"`  // 1：目录 2：菜单 3；按钮
	Title           string `json:"title"`     // 中文名
	Name            string `json:"name"`      // 组件名字
	Path            string `json:"path"`      // 访问的URL
	Component       string `json:"component"` // 组件路径
	Rank            int    `json:"rank"`      // 排序
	Redirect        string `json:"redirect"`
	Status          int8   `json:"status,optional"` // 菜单状态（0正常 1停用）
	Icon            string `json:"icon"`            // icon 暂用于商户后台
	ExtraIcon       string `json:"extraIcon"`
	EnterTransition string `json:"enterTransition"`
	LeaveTransition string `json:"leaveTransition"`
	ActivePath      string `json:"activePath"`
	Auths           string `json:"auths"`
	FrameSrc        string `json:"frameSrc"`
	FrameLoading    bool   `json:"frameLoading"`
	KeepAlive       bool   `json:"keepAlive"`
	HiddenTag       bool   `json:"hiddenTag"`
	FixedTag        bool   `json:"fixedTag"`
	ShowLink        bool   `json:"showLink"` // 可视：是(1)、否(0)
	ShowParent      bool   `json:"showParent"`
}

type MenuListResponse struct {
	List  []*MenuList `json:"list"`
	Total int64       `json:"total"`
}

func (l *MenuLogic) Create() error {
	if l.Name == "" {
		return errors.New("路由名称不能为空")
	}

	if l.Path == "" {
		return errors.New("路径不能为空")
	}

	if l.Title == "" {
		return errors.New("菜单名称不能为空")
	}

	menus := l.NewSysMenus()
	menus.CreatedAt = time.Now().Unix()

	return menus.Create(l.ctx)

}

func (l *MenuLogic) Update() error {

	if l.Id < 1 {
		return errors.New("id不能为空")
	}
	menus := l.NewSysMenus()
	menus.UpdatedAt = time.Now().Unix()

	return menus.Update(l.ctx)

}

func (l *MenuLogic) NewSysMenus() system.SysMenus {
	menus := system.SysMenus{
		Id:        l.Id,
		MenuType:  l.MenuType,
		ParentId:  l.ParentId,
		Title:     l.Title,
		Path:      l.Path,
		Redirect:  l.Redirect,
		Name:      l.Name,
		Component: l.Component,
		Rank:      l.Rank,
		Status:    l.Status,
		Icon:      l.Icon,
		// KeepAlive:       l.KeepAlive == true,
		// ShowLink:        l.ShowLink,
		// HiddenTag:       l.HiddenTag,
		// FixedTag:        l.FixedTag,
		ExtraIcon:       l.ExtraIcon,
		EnterTransition: l.EnterTransition,
		LeaveTransition: l.LeaveTransition,
		ActivePath:      l.ActivePath,
		Auths:           l.Auths,
		FrameSrc:        l.FrameSrc,
		FrameLoading:    l.Status,
		// ShowParent:      l.ParentId,
		// CreateTime:      l.Status,
		// DeletedAt:       l.Status,
	}
	if l.KeepAlive {
		menus.KeepAlive = 1
	}
	if l.ShowLink {
		menus.ShowLink = 1
	}
	if l.HiddenTag {
		menus.HiddenTag = 1
	}
	if l.FixedTag {
		menus.FixedTag = 1
	}
	if l.ShowParent {
		menus.ShowParent = 1
	}

	return menus
}

func (l *MenuLogic) List(in *system.MenuListRequest) (resp MenuListResponse, err error) {
	menus := system.NewSysMenus()

	response, err := menus.List(l.ctx, in)
	if err != nil {
		return MenuListResponse{}, err
	}

	for _, v := range response.List {
		resp.List = append(resp.List, &MenuList{
			Id:       v.Id,
			ParentId: v.ParentId,
			// Permission:   v.Permission,
			Path:      v.Path,
			Component: v.Component,
			Status:    v.Status,
			// ResourceType: v.ResourceType,
			// CreateTime:   time.Unix(v.CreateTime, 0).Format("2006-01-02 15:04:05"),
			Name:       v.Name,
			Title:      v.Title,
			Icon:       v.Icon,
			ShowLink:   v.ShowLink == 1,
			KeepAlive:  v.KeepAlive == 1,
			ShowParent: v.ShowParent == 1,
			Rank:       v.Rank,
		})
	}

	return
}

type Meta struct {
	Title      string   `json:"title"` // 中文名
	Icon       string   `json:"icon"`  // icon 暂用于商户后台
	ExtraIcon  string   `json:"extraIcon,omitempty"`
	ShowLink   bool     `json:"showLink,omitempty"` // 可视：是(1)、否(0)
	ShowParent bool     `json:"showParent,omitempty"`
	KeepAlive  bool     `json:"keepAlive,omitempty"`
	Rank       int      `json:"rank,omitempty"` // 排序
	Roles      []string `json:"roles,omitempty"`
}

type MenuTree struct {
	Path      string      `json:"path"`               // 访问的URL
	Name      string      `json:"name,omitempty"`     // 组件名字
	Redirect  string      `json:"redirect,omitempty"` // 访问的URL
	Meta      Meta        `json:"meta"`
	Component string      `json:"component,omitempty"` // 组件路径
	Children  []*MenuTree `json:"children,omitempty"`
}

func (l *MenuLogic) MenuTree() (resp []*MenuTree, err error) {

	role, ok := l.ctx.Value("roleIds").(string)
	if !ok || role == "" {
		return nil, errors.New("无效的角色")
	}

	var roleIds []int32
	err = json.Unmarshal([]byte(role), &roleIds)
	if err != nil {
		return nil, err
	}

	var in system.RoleListRequest
	in.Ids = roleIds
	sysRole := system.NewSysRole()

	roleList, _, err := sysRole.List(l.ctx, &system.RoleListRequest{Ids: roleIds})
	if err != nil {
		return nil, err
	}

	if roleList == nil {
		return
	}

	var menuId []int32
	for _, v := range roleList {
		var Ids []int32
		err = json.Unmarshal([]byte(v.DataScopeMenuIds), &Ids)
		if err != nil {
			return nil, err
		}
		menuId = append(menuId, Ids...)
	}

	menuId = utils.ArrayUnique(menuId)
	request := system.MenuListRequest{
		Type: "tree",
		Ids:  menuId,
	}

	menu := system.NewSysMenus()
	list, err := menu.List(l.ctx, &request)
	if err != nil {
		return nil, err
	}

	resp = tree(0, list.List)

	return
}

func tree(pid int, list []*system.SysMenus) []*MenuTree {
	var data []*MenuTree
	for _, v := range list {
		if v.ParentId == pid {
			meta := Meta{
				Title:    v.Title,
				Icon:     v.Icon,
				ShowLink: v.ShowLink == 1,
				// KeepAlive: v.KeepAlive == 1,
				// Rank:  v.Rank,
				// Roles: []string{"admin"},
			}

			menuTree := &MenuTree{
				Path: v.Path,
				// Component: v.Component,
				// AppName:      v.AppName,
				// Meta:     meta,
				Children: tree(v.Id, list),
			}

			if pid == 0 {
				meta.Rank = v.Rank
			} else {
				menuTree.Name = v.Name
				menuTree.Component = v.Component

				meta.KeepAlive = v.KeepAlive == 1
				meta.ShowParent = v.ShowParent == 1
			}

			menuTree.Meta = meta

			data = append(data, menuTree)
		}
	}

	return data
}

type ListAllResponse struct {
	Id       int    `json:"id"`
	ParentId int    `json:"parentId"`
	MenuType int8   `json:"menuType"`
	Title    string `json:"title"`
}

func (l *MenuLogic) MenuAll() (resp []*ListAllResponse, err error) {
	menu := system.NewSysMenus()
	list, err := menu.List(l.ctx, &system.MenuListRequest{})
	if err != nil {
		return nil, err
	}

	for _, v := range list.List {
		resp = append(resp, &ListAllResponse{
			Id:       v.Id,
			ParentId: v.ParentId,
			MenuType: v.MenuType,
			Title:    v.Title,
		})
	}

	return
}
