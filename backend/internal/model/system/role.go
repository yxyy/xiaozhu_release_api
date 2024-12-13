package system

import (
	"context"
	"errors"
	"xiaozhu/internal/config/mysql"
	"xiaozhu/internal/model/common"
)

type SysRole struct {
	Id               int    `json:"id" gorm:"id"`                                   // 角色ID
	Name             string `json:"name" gorm:"name"`                               // 角色名
	Status           *int   `json:"status" gorm:"status"`                           // 状态：1 正常 0 禁用
	Code             string `json:"code" gorm:"code"`                               // 角色权限字符串
	Sort             int    `json:"sort" gorm:"sort"`                               // 显示顺序
	DataScope        string `json:"data_scope" gorm:"data_scope"`                   // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	DataScopeMenuIds string `json:"data_scope_menu_ids" gorm:"data_scope_menu_ids"` // 菜单权限
	Type             int8   `json:"type" gorm:"type"`                               // 角色类型
	Remark           string `json:"remark" gorm:"remark"`                           // 备注
	Creator          int    `json:"creator" gorm:"creator"`                         // 创建者
	CreatedAt        int64  `json:"created_at" gorm:"created_at"`                   // 创建时间
	Updater          uint64 `json:"updater" gorm:"updater"`                         // 更新者
	UpdatedAt        int64  `json:"updated_at" gorm:"updated_at"`                   // 更新时间
	DeletedAt        int64  `json:"deleted_at" gorm:"deleted_at"`                   // 是否删除
}

func NewSysRole() *SysRole {
	return &SysRole{}
}

type RoleListRequest struct {
	Name   string         `json:"name"`
	Code   string         `json:"code"`
	Status *int           `json:"status"`
	Params *common.Params `json:"params"`
	Ids    []int32        `json:"ids,omitempty"`
}

type RoleListResponse struct {
	List  []*SysRole `json:"list"`
	Total int64      `json:"total"`
}

func (l *SysRole) Create() error {

	if err := mysql.PlatformDB.Model(&l).Where("name", l.Name).Take(&SysRole{}).Error; err == nil {
		return errors.New("该名称已存在")
	}

	if err := mysql.PlatformDB.Model(&l).Where("flag", l.Code).Take(&SysRole{}).Error; err == nil {
		return errors.New("该标识已存在")
	}

	if err := mysql.PlatformDB.Model(&l).Create(&l).Error; err != nil {
		return err
	}

	return nil

}

func (l *SysRole) Update() error {

	return mysql.PlatformDB.Model(&l).Updates(&l).Error
}

// List 表格列表
func (l *SysRole) List(ctx context.Context, in *RoleListRequest) (resp []*SysRole, total int64, err error) {

	tx := mysql.PlatformDB.Model(&l).WithContext(ctx)
	if in.Name != "" {
		tx = tx.Where("name like ?", "%"+in.Name+"%")
	}

	if in.Code != "" {
		tx = tx.Where("code", in.Code)
	}

	if in.Status != nil {
		tx = tx.Where("status", in.Status)
	}

	if in.Ids != nil {
		tx = tx.Where("id in ? ", in.Ids)
	}

	err = tx.Count(&total).Error
	if err != nil || total == 0 {
		return
	}

	if in.Params != nil {
		if in.Params.Page != 0 && in.Params.Limit != 0 {
			tx = tx.Offset((int(in.Params.Page) - 1) * in.Params.Limit).Limit(in.Params.Limit)
		}
	}

	err = tx.Find(&resp).Error

	return
}

// GetAll 获取全部
func (l SysRole) GetAll() (groups []*SysRole, err error) {

	tx := mysql.PlatformDB.Model(&l)
	if l.Id > 0 {
		tx = tx.Where("id", l.Id)
	}
	// if l.PermissionId > 0 {
	// 	tx = tx.Where("permission_id", l.PermissionId)
	// }
	if err = tx.Find(&groups).Error; err != nil {
		return nil, err
	}

	return
}

func (l *SysRole) Get() (err error) {

	tx := mysql.PlatformDB.Model(&l)
	if l.Id > 0 {
		tx = tx.Where("id", l.Id)
	}

	return tx.Find(&l).Error
}
