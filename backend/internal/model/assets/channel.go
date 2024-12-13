package assets

import (
	"context"
	"xiaozhu/internal/config/mysql"
	"xiaozhu/internal/model/common"
)

type Channel struct {
	common.Model
	Name        string `json:"name" form:"name"`
	Code        string `json:"code" form:"code"`
	AuthUrl     string `json:"auth_url"`
	Params      string `json:"params"`
	RedirectUri string `json:"redirect_uri"`
	Doc         string `json:"doc"`
	Remark      string `json:"remark"`
}

func (c *Channel) Create(ctx context.Context) error {
	return mysql.PlatformDB.Model(&c).WithContext(ctx).Create(&c).Error
}

func (c *Channel) Update(ctx context.Context) error {
	return mysql.PlatformDB.Model(&c).WithContext(ctx).Updates(&c).Error
}

func (c *Channel) List(ctx context.Context, params *common.Params) (list []*Channel, total int64, err error) {
	tx := mysql.PlatformDB.Model(&c).WithContext(ctx)
	if c.Id > 0 {
		tx = tx.Where("id", c.Id)
	}
	if c.Name != "" {
		tx = tx.Where("name like ?", c.Name+"%")
	}
	if c.Code != "" {
		tx = tx.Where("code like ?", c.Code+"%")
	}
	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err = tx.Offset(params.Offset).Limit(params.Limit).Find(&list).Error

	return
}

func (c *Channel) GetAll(ctx context.Context) (list []*common.IdName, err error) {
	err = mysql.PlatformDB.Model(&c).WithContext(ctx).Find(&list).Error
	return
}
