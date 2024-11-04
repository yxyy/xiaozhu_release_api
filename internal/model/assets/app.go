package assets

import (
	"context"
	"xiaozhu/internal/model/common"
	"xiaozhu/utils"
)

type App struct {
	common.Model
	AppName   string `json:"app_name" form:"app_name"`
	CompanyId int    `json:"company_id"  form:"company_id"`
	Alias     string `json:"alias"  form:"alias"`
	Remark    string `json:"remark"  form:"remark"`
	GameClass int    `json:"game_class"  form:"game_class"`
	Status    int    `json:"status"  form:"status"`
}

func (a *App) Create(ctx context.Context) error {
	return utils.MysqlDb.WithContext(ctx).Model(&a).Create(&a).Error
}

func (a *App) Update(ctx context.Context) error {
	return utils.MysqlDb.Model(&a).WithContext(ctx).Where("id", a.Id).Updates(&a).Error
}

func (a *App) List(ctx context.Context, in *common.Params) (resp []*App, total int64, err error) {
	tx := utils.MysqlDb.Model(&a).WithContext(ctx)
	if a.Id > 0 {
		tx = tx.Where("id", a.Id)
	}
	if a.CompanyId > 0 {
		tx = tx.Where("company_id", a.CompanyId)
	}
	if a.AppName != "" {
		tx = tx.Where("name like ?", a.AppName+"%")
	}

	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err = tx.Offset(in.Offset).Limit(in.Limit).Find(&resp).Error

	return

}

func (a *App) GetAll(ctx context.Context) (app []*App, err error) {

	err = utils.MysqlDb.Model(&a).WithContext(ctx).Find(&app).Error
	return
}
