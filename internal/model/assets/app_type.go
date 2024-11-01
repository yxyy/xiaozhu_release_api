package assets

import (
	"xiaozhu/internal/model/common"
	"xiaozhu/utils"
)

type AppType struct {
	common.Model
	Name string `json:"name" form:"name"`
}

func (t AppType) Create() error {
	return utils.MysqlDb.Model(&t).Create(&t).Error
}

func (t AppType) Update() error {
	return utils.MysqlDb.Model(&t).Where("id", t.Id).Updates(&t).Error
}

func (t AppType) List(params common.Params) (companys []*AppType, total int64, err error) {
	tx := utils.MysqlDb.Model(&t)
	if t.Id > 0 {
		tx = tx.Where("id", t.Id)
	}
	if t.Name != "" {
		tx = tx.Where("name like ?", t.Name+"%")
	}

	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err = tx.Offset(params.Offset).Limit(params.Limit).Find(&companys).Error

	return

}

func (t AppType) GetAll() (appType []*AppType, err error) {

	err = utils.MysqlDb.Model(&t).Find(&appType).Error
	return
}
