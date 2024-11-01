package assets

import (
	"xiaozhu/internal/model/common"
	"xiaozhu/utils"
)

type Channel struct {
	common.Model
	Name string `json:"name" form:"name"`
	Flag string `json:"flag" form:"flag"`
}

func (c Channel) Create() error {
	return utils.MysqlDb.Model(&c).Create(&c).Error
}

func (c Channel) Update() error {
	return utils.MysqlDb.Model(&c).Updates(&c).Error
}

func (c Channel) List(params common.Params) (list []*Channel, total int64, err error) {
	tx := utils.MysqlDb.Model(&c)
	if c.Id > 0 {
		tx = tx.Where("id", c.Id)
	}
	if c.Name != "" {
		tx = tx.Where("name like ?", c.Name+"%")
	}
	if c.Flag != "" {
		tx = tx.Where("flag like ?", c.Flag+"%")
	}
	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err = tx.Offset(params.Offset).Limit(params.Limit).Find(&list).Error

	return
}

func (c Channel) GetAll() (list []*Channel, err error) {
	err = utils.MysqlDb.Model(&c).Find(&list).Error
	return
}
