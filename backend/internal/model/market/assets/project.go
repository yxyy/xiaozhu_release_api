package assets

import (
	"context"
	"xiaozhu/internal/config/mysql"
	"xiaozhu/internal/model/common"
)

type ProxyProject struct {
	common.Model
	Name              string  `json:"name" form:"name"`
	ShortName         string  `json:"short_name" form:"short_name"`
	ChannelId         int     `json:"channel_id"`
	PrincipalId       int     `json:"principal_id"`
	ProxyCompanyId    int     `json:"proxy_company_id"`
	Rebate            float32 `json:"rebate"`
	ContractStartTime int64   `json:"contract_start_time"`
	ContractEndTime   int64   `json:"contract_end_time"`
	ContractStatus    int8    `json:"contract_status"`
	Remark            string  `json:"remark"`
}

type ProxyProjectList struct {
	ProxyProject
	ChannelName      string `json:"channel_name"`
	PrincipalName    string `json:"principal_name"`
	ProxyCompanyName string `json:"proxy_company_name"`
	OptUserName      string `json:"opt_user_name"`
}

func (p *ProxyProject) TableName() string {
	return "market_proxy_project"
}

func (p *ProxyProject) Create(ctx context.Context) error {
	return mysql.PlatformDB.Model(&p).WithContext(ctx).Create(&p).Error
}

func (p *ProxyProject) Update(ctx context.Context) error {
	return mysql.PlatformDB.Model(&p).WithContext(ctx).Updates(&p).Error
}

func (p *ProxyProject) List(ctx context.Context, params *common.Params) (list []*ProxyProjectList, total int64, err error) {
	tx := mysql.PlatformDB.Model(&p).WithContext(ctx).
		Select("market_proxy_project.*",
			"channels.name as channel_name",
			"market_proxy_company.name as proxy_company_name",
			"market_principals.name as principal_name",
		).Joins("left join channels on channels.id = market_proxy_project.channel_id").
		Joins("left join market_proxy_company on market_proxy_company.id = market_proxy_project.proxy_company_id").
		Joins("left join market_principals on market_principals.id = market_proxy_project.principal_id")
	if p.Id > 0 {
		tx = tx.Where("market_proxy_project.id", p.Id)
	}
	if p.Name != "" {
		tx = tx.Where("market_proxy_project.name like ?", p.Name+"%")
	}
	if p.ChannelId > 0 {
		tx = tx.Where("market_proxy_project.channel_id", p.ChannelId)
	}
	if p.PrincipalId > 0 {
		tx = tx.Where("market_proxy_project.principal_id", p.PrincipalId)
	}
	if p.ProxyCompanyId > 0 {
		tx = tx.Where("market_proxy_project.proxy_company_id", p.ProxyCompanyId)
	}

	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err = tx.Offset(params.Offset).Limit(params.Limit).Find(&list).Error

	return
}

func (p *ProxyProject) GetAll(ctx context.Context) (list []*common.IdName, err error) {
	err = mysql.PlatformDB.Model(&p).WithContext(ctx).Find(&list).Error
	return
}
