package system

import (
	"context"
	"xiaozhu/internal/config/mysql"
	"xiaozhu/internal/model/common"
)

type SysMenus struct {
	Id              int    `json:"id" gorm:"id"`
	MenuType        int8   `json:"menu_type" gorm:"menu_type"` // 0代表菜单、1代表iframe、2代表外链、3代表按钮
	ParentId        int    `json:"parent_id" gorm:"parent_id"`
	Title           string `json:"title" gorm:"title"`                       // 显示名称
	Path            string `json:"path" gorm:"path"`                         // 访问的URL
	Redirect        string `json:"redirect" gorm:"redirect"`                 // 重定向路径
	Name            string `json:"name" gorm:"name"`                         // 组件名字
	Component       string `json:"component" gorm:"component"`               // 组件路径
	Rank            int    `json:"rank" gorm:"rank"`                         //  菜单排序，值越高排的越后（只针对顶级路由）
	Status          int8   `json:"status" gorm:"status"`                     // 菜单状态（0正常 1停用）
	Icon            string `json:"icon" gorm:"icon"`                         // icon
	KeepAlive       int8   `json:"keep_alive" gorm:"keep_alive"`             // 缓存页面（是否缓存该路由页面，开启后会保存该页面的整体状态，刷新后会清空状态）
	ShowLink        int8   `json:"show_link" gorm:"show_link"`               // 菜单（是否显示该菜单）
	HiddenTag       int8   `json:"hidden_tag" gorm:"hidden_tag"`             // 标签页（当前菜单名称或自定义信息禁止添加到标签页）
	FixedTag        int8   `json:"fixed_tag" gorm:"fixed_tag"`               // 固定标签页（当前菜单名称是否固定显示在标签页且不可关闭）
	ExtraIcon       string `json:"extra_icon" gorm:"extra_icon"`             // 右侧图标
	EnterTransition string `json:"enter_transition" gorm:"enter_transition"` // 进场动画（页面加载动画）
	LeaveTransition string `json:"leave_transition" gorm:"leave_transition"` // 离场动画（页面加载动画）
	ActivePath      string `json:"active_path" gorm:"active_path"`           // 菜单激活（将某个菜单激活，主要用于通过query或params传参的路由，当它们通过配置showLink: false后不在菜单中显示，就不会有任何菜单高亮，而通过设置activePath指定激活菜单即可获得高亮，activePath为指定激活菜单的path）
	Auths           string `json:"auths" gorm:"auths"`                       // 权限标识（按钮级别权限设置）
	FrameSrc        string `json:"frame_src" gorm:"frame_src"`               // 链接地址（需要内嵌的iframe链接地址）
	FrameLoading    int8   `json:"frame_loading" gorm:"frame_loading"`       // 加载动画（内嵌的iframe页面是否开启首次加载动画）
	ShowParent      int8   `json:"show_parent" gorm:"show_parent"`           // 是否显示父级菜单
	CreatedAt       int64  `json:"created_at" gorm:"created_at"`             // 创建时间
	UpdatedAt       int64  `json:"updated_at" gorm:"updated_at"`             // 更新时间
	DeletedAt       int64  `json:"deleted_at" gorm:"deleted_at"`
}

func NewSysMenus() *SysMenus {
	return &SysMenus{}
}

func (m SysMenus) Create(ctx context.Context) error {

	return mysql.PlatformDB.WithContext(ctx).Model(&m).Create(&m).Error
}

func (m SysMenus) Update(ctx context.Context) error {

	return mysql.PlatformDB.WithContext(ctx).Model(&m).Save(&m).Error
}

type MenuListRequest struct {
	Ids      []int32       `json:"Ids,omitempty"`
	Type     string        `json:"GameClass,omitempty"`
	Title    string        `json:"Title,omitempty"`
	PageInfo common.Params `json:"page_info,omitempty"`
}

type MenuListResponse struct {
	List  []*SysMenus `json:"list,omitempty"`
	Total int64       `json:"total,omitempty"`
}

func (m *SysMenus) List(cxt context.Context, in *MenuListRequest) (MenuListResponse, error) {

	tx := mysql.PlatformDB.WithContext(cxt).Model(&m).Where("deleted_at", 0)

	if in.Type == "tree" {
		tx = tx.Where("show_link = ?", 1)
	}

	if in.Title != "" {
		tx = tx.Where("title like ?", "%"+in.Title+"%")
	}

	if in.Ids != nil {
		tx = tx.Where("id in ?", in.Ids)
	}

	var resp MenuListResponse
	err := tx.Order("sys_menus.rank asc").Find(&resp.List).Error

	return resp, err
}
