package assets

import (
	"context"
	"xiaozhu/internal/config/mysql"
	"xiaozhu/internal/model/common"
)

type Principal struct {
	common.Model
	Name   string `json:"name" form:"name"`
	Code   string `json:"code" form:"code"`
	Remark string `json:"remark"`
}

func (p *Principal) TableName() string {
	return "market_principals"
}

func (p *Principal) Create(ctx context.Context) error {
	return mysql.PlatformDB.Model(&p).WithContext(ctx).Create(&p).Error
}

func (p *Principal) Update(ctx context.Context) error {
	return mysql.PlatformDB.Model(&p).WithContext(ctx).Updates(&p).Error
}

func (p *Principal) List(ctx context.Context, params *common.Params) (list []*Principal, total int64, err error) {
	tx := mysql.PlatformDB.Model(&p).WithContext(ctx)
	if p.Id > 0 {
		tx = tx.Where("id", p.Id)
	}
	if p.Name != "" {
		tx = tx.Where("name like ?", p.Name+"%")
	}
	if p.Code != "" {
		tx = tx.Where("code like ?", p.Code+"%")
	}
	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err = tx.Offset(params.Offset).Limit(params.Limit).Find(&list).Error

	return
}

func (p *Principal) GetAll(ctx context.Context) (list []*common.IdName, err error) {
	err = mysql.PlatformDB.Model(&p).WithContext(ctx).Find(&list).Error
	return
}
