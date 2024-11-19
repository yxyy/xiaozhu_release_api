package assets

import (
	"context"
	"xiaozhu/internal/model/common"
	"xiaozhu/utils"
)

type MarketAppId struct {
	common.Model
	Name      string `json:"name" gorm:"name"`             // 名字
	ChannelId int64  `json:"channel_id" gorm:"channel_id"` // 媒体渠道
	Account   string `json:"account" gorm:"account"`       // 开发者账号
	AppId     string `json:"app_id" gorm:"app_id"`         // app_id
	Secret    string `json:"secret" gorm:"secret"`         // 秘钥
	ButlerId  string `json:"butler_id" gorm:"butler_id"`   // 授权账户管家ID
	State     string `json:"code" gorm:"code"`             // 标识
	Status    int8   `json:"status" gorm:"status"`         // 是否可用 0 否 1是
	Params    string `json:"params" gorm:"params"`         // 授权参数配置
	Remark    string `json:"remark" gorm:"remark"`         // 备注
}

type MarketAppIdList struct {
	MarketAppId
	AuthUrl     string `json:"auth_url"`
	Params      string `json:"params"`
	RedirectUri string `json:"redirect_uri"`
	ChannelCode string `json:"channel_code"`
	ChannelName string `json:"channel_name"`
	OptUserName string `json:"opt_user_name"`
}

func (p *MarketAppId) TableName() string {
	return "market_app_id"
}

func (p *MarketAppId) Create(ctx context.Context) error {
	return utils.MysqlDb.Model(&p).WithContext(ctx).Create(&p).Error
}

func (p *MarketAppId) Update(ctx context.Context) error {
	return utils.MysqlDb.Model(&p).WithContext(ctx).Save(&p).Error
}

func (p *MarketAppId) List(ctx context.Context, params *common.Params) (list []*MarketAppIdList, total int64, err error) {
	tx := utils.MysqlDb.Model(&p).WithContext(ctx).
		Select("market_app_id.*",
			"channels.name as channel_name",
			"channels.code as channel_code",
			"channels.auth_url as auth_url",
			"channels.redirect_uri as redirect_uri",
			"channels.params as params",
		).
		Joins("left join channels on channels.id = market_app_id.channel_id")
	if p.Id > 0 {
		tx = tx.Where("market_app_id.id", p.Id)
	}
	if p.Name != "" {
		tx = tx.Where("market_app_id.name like ?", "%"+p.Name+"%")
	}
	if p.State != "" {
		tx = tx.Where("market_app_id.code like ?", "%"+p.State+"%")
	}
	if p.ChannelId > 0 {
		tx = tx.Where("channel_id", p.ChannelId)
	}

	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err = tx.Offset(params.Offset).Limit(params.Limit).Find(&list).Error

	return
}

func (p *MarketAppId) GetAll(ctx context.Context) (list []*common.IdName, err error) {
	err = utils.MysqlDb.Model(&p).WithContext(ctx).Find(&list).Error
	return
}

func (p *MarketAppId) Get(ctx context.Context) (err error) {
	tx := utils.MysqlDb.Model(&p).WithContext(ctx)
	if p.Id != 0 {
		tx = tx.Where("id", p.Id)
	}
	if p.State != "" {
		tx = tx.Where("state", p.State)
	}

	return tx.First(&p).Error

	// return utils.MysqlDb.Model(&p).WithContext(ctx).First(&p).Error
	// if errors.Is(err,gorm.ErrRecordNotFound) {
	// 	return nil
	// }
	//
	// return err
}
