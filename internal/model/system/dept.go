package system

import (
	"context"
	"time"
	"xiaozhu/internal/model/common"
	"xiaozhu/utils"
)

// SysDept 部门表
type SysDept struct {
	Id           uint64    `json:"id" gorm:"id"`                       // 部门id
	Name         string    `json:"name" gorm:"name"`                   // 部门名称
	ParentId     int64     `json:"parentId" gorm:"parent_id"`          // 父部门id
	Sort         int       `json:"sort" gorm:"sort"`                   // 显示顺序
	LeaderUserId int64     `json:"leaderUserId" gorm:"leader_user_id"` // 负责人
	Status       int8      `json:"status" gorm:"status"`               // 部门状态（0正常 1停用）
	Creator      string    `json:"creator" gorm:"creator"`             // 创建者
	Updater      string    `json:"updater" gorm:"updater"`             // 更新者
	CreatedAT    int64     `json:"created_at" gorm:"created_at"`       // 创建时间
	UpdatedAT    int64     `json:"updated_at" gorm:"updated_at"`       // 更新时间
	DeletedAt    time.Time `json:"deleted_at" gorm:"deleted_at"`       // 是否删除
}

func NewSysDept() *SysDept {
	return &SysDept{}
}

type SysDeptListRequest struct {
	common.Params
	Name string `json:"name"`
	Sort int    `json:"sort"`
}

type SysDeptListResponse struct {
	List  []*SysDept `json:"list"`
	Total int64      `json:"total"`
}

func (r *SysDept) List(ctx context.Context, in SysDeptListRequest) (resp *SysDeptListResponse, err error) {

	tx := utils.MysqlDb.Model(&r).WithContext(ctx)
	if in.Name != "" {
		tx = tx.Where("name like ?", "%"+in.Name+"%")
	}
	resp = new(SysDeptListResponse)
	err = tx.Count(&resp.Total).Error
	if err != nil {
		return
	}

	if err = tx.Find(&resp.List).Error; err != nil {
		return nil, err
	}

	return
}

func (r *SysDept) Create(ctx context.Context) error {
	return utils.MysqlDb.Model(&r).WithContext(ctx).Create(&r).Error
}

func (r *SysDept) Update(ctx context.Context) error {
	return utils.MysqlDb.Model(&r).WithContext(ctx).Save(&r).Error
}
