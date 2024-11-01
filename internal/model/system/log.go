package system

import (
	"golang.org/x/net/context"
	"xiaozhu/internal/model/common"
	"xiaozhu/utils"
)

type SysUserLog struct {
	Id        int    `json:"id" form:"id" gorm:"primarykey"`
	LogType   int    `json:"log_type" gorm:"log_type"`
	UserId    int    `json:"user_id,omitempty" form:"user_id"`
	Account   string `json:"account" form:"account"`
	Ip        string `json:"ip"`
	Path      string `json:"path"`
	Request   string `json:"request" form:"request"`
	Response  string `json:"response" form:"response"`
	Status    int    `json:"status" form:"status"`
	CreatedAt int64  `json:"created_at" form:"created_at" `
	RequestId string `json:"request_id"`
}

func (l *SysUserLog) Create() error {
	return utils.MysqlDb.Model(&l).Create(&l).Error
}

func (l *SysUserLog) List(ctx context.Context, in *common.Params) (resp []*SysUserLog, total int64, err error) {
	tx := utils.MysqlDb.Model(&l).WithContext(ctx)
	if l.Path != "" {
		tx = tx.Where("path like ?", "%"+l.Path+"%")
	}
	if l.LogType > 0 {
		tx = tx.Where("log_type", l.LogType)
	}

	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err = tx.Offset(in.Offset).Limit(in.Limit).Find(&resp).Error; err != nil {
		return nil, 0, err
	}

	return

}
