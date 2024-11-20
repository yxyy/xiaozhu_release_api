package assets

import (
	"context"
	"xiaozhu/backend/internal/model/common"
	"xiaozhu/backend/utils"
)

type ProxyCompany struct {
	common.Model
	Name      string `json:"name" form:"name"`
	ShortName string `json:"short_name" form:"short_name"`
	Code      string `json:"code" form:"code"`
	Remark    string `json:"remark"`
}

func (p *ProxyCompany) TableName() string {
	return "market_proxy_company"
}

func (p *ProxyCompany) Create(ctx context.Context) error {
	return utils.MysqlDb.Model(&p).WithContext(ctx).Create(&p).Error
}

func (p *ProxyCompany) Update(ctx context.Context) error {
	return utils.MysqlDb.Model(&p).WithContext(ctx).Updates(&p).Error
}

func (p *ProxyCompany) List(ctx context.Context, params *common.Params) (list []*ProxyCompany, total int64, err error) {
	tx := utils.MysqlDb.Model(&p).WithContext(ctx)
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

func (p *ProxyCompany) GetAll(ctx context.Context) (list []*common.IdName, err error) {
	err = utils.MysqlDb.Model(&p).WithContext(ctx).Find(&list).Error
	return
}
