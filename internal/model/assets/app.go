package assets

import (
	"xiaozhu/internal/model/common"
	"xiaozhu/utils"
)

type App struct {
	common.Model
	Name      string `json:"name" form:"name"`
	CompanyId int    `json:"company_id"`
	Alias     string `json:"alias"`
	Remark    string `json:"remark"`
	Type      int    `json:"type"`
}

func (a App) Create() error {
	return utils.MysqlDb.Model(&a).Create(&a).Error
}

func (a App) Update() error {
	return utils.MysqlDb.Model(&a).Where("id", a.Id).Updates(&a).Error
}

func (a App) List(params common.Params) (companys []*App, total int64, err error) {
	tx := utils.MysqlDb.Model(&a)
	if a.Id > 0 {
		tx = tx.Where("id", a.Id)
	}
	if a.CompanyId > 0 {
		tx = tx.Where("company_id", a.CompanyId)
	}
	if a.Name != "" {
		tx = tx.Where("name like ?", a.Name+"%")
	}

	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err = tx.Offset(params.Offset).Limit(params.Limit).Find(&companys).Error

	return

}

func (a App) GetAll() (app []*App, err error) {

	err = utils.MysqlDb.Model(&a).Find(&app).Error
	return
}
