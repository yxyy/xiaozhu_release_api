package assets

import (
	"context"
	"gorm.io/gorm/clause"
	"xiaozhu/internal/model/common"
	"xiaozhu/utils"
)

type Account struct {
	common.Model
	Name         string `json:"name"`
	ShortName    string `json:"short_name"`
	ChannelId    int    `json:"channel_id" gorm:"-"`
	ProjectId    int    `json:"project_id"`
	Uid          int    `json:"uid"`
	Account      string `json:"account"`
	Password     string `json:"password"`
	Owner        int    `json:"owner"`
	Status       int8   `json:"status"`
	OauthType    int8   `json:"oauth_type"`
	OauthStatus  int8   `json:"oauth_status"`
	OauthSubject string `json:"oauth_subject"`
	Lot          string `json:"lot"`
	Remark       string `json:"remark"`
}

type AdAccountList struct {
	Account
	ChannelName string `json:"channel_name"`
	ProjectName string `json:"project_name"`
	OwnerName   string `json:"owner_name"`
	OptUserName string `json:"opt_user_name"`
}

func (p *Account) TableName() string {
	return "market_proxy_account"
}

func (p *Account) Create(ctx context.Context) error {
	return utils.MysqlDb.Model(&p).WithContext(ctx).Create(&p).Error
}

func (p *Account) BatchCreate(ctx context.Context, data []*Account) error {
	return utils.MysqlDb.
		Model(&p).
		WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "uid"}, {Name: "deleted_at"}},
			DoNothing: false,
			// DoUpdates:    clause.Assignments(map[string]interface{}{"owner": gorm.Expr("VALUES(owner)"), "lot": gorm.Expr("VALUES(lot)")}),
			UpdateAll: true,
		}).
		Create(&data).
		Error
}

func (p *Account) Update(ctx context.Context) error {
	return utils.MysqlDb.Model(&p).WithContext(ctx).Save(&p).Error
}

func (p *Account) List(ctx context.Context, params *common.Params) (list []*AdAccountList, total int64, err error) {
	tx := utils.MysqlDb.Model(&p).WithContext(ctx).
		Select("market_proxy_account.*",
			"channels.name as channel_name",
			"market_proxy_project.name as project_name",
		).Joins("left join market_proxy_project on market_proxy_project.id = market_proxy_account.project_id").
		Joins("left join channels on channels.id = market_proxy_project.channel_id")
	if p.Id > 0 {
		tx = tx.Where("market_proxy_project.id", p.Id)
	}
	if p.Name != "" {
		tx = tx.Where("market_proxy_project.name like ?", p.Name+"%")
	}
	if p.ChannelId > 0 {
		tx = tx.Where("market_proxy_project.channel_id", p.ChannelId)
	}
	if p.ProjectId > 0 {
		tx = tx.Where("market_proxy_account.project_id", p.ProjectId)
	}

	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err = tx.Offset(params.Offset).Limit(params.Limit).Find(&list).Error

	return
}

func (p *Account) GetAll(ctx context.Context) (list []*common.IdName, err error) {
	err = utils.MysqlDb.Model(&p).WithContext(ctx).Find(&list).Error
	return
}
