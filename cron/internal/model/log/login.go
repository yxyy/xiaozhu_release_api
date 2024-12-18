package log

import (
	"context"
	"errors"
	"xiaozhu/internal/config/mysql"
)

// Login  登录日志
type Login struct {
	Id          int    `json:"id" gorm:"id"`
	PromoteCode string `json:"promote_code" gorm:"promote_code"` // 推广码
	AppId       int    `json:"app_id" gorm:"app_id"`             // 游戏项目ID
	GameId      int    `json:"game_id" gorm:"game_id"`           // 游戏ID
	ChannelId   int64  `json:"channel_id" gorm:"channel_id"`     // 媒体渠道
	Os          string `json:"os" gorm:"os"`                     // 操作系统：android、ios
	// Cause       string `json:"cause" gorm:"cause"`               // 广告归因依据
	UserId    int    `json:"user_id"`
	Account   string `json:"account"`
	LoginWay  int    `json:"login_way"`
	DeviceId  string `json:"device_id" gorm:"device_id"`   // 设备号
	Ip        string `json:"ip" gorm:"ip"`                 // IP
	AreaCode  string `json:"area_code" gorm:"area_code"`   // 地区码
	Area      string `json:"area" gorm:"area"`             // 地区
	Ts        int64  `json:"ts" gorm:"ts"`                 // 事件时间，即激活时间
	CreatedAt int64  `json:"created_at" gorm:"created_at"` // 创建时间
	Days      int    `json:"days" gorm:"days"`             // 日期
	RequestId string `json:"request_id"`
}

// TableName 表名称
func (*Login) TableName() string {
	return "log_login"
}

func NewLogin() *Login {
	return &Login{}
}

func (d *Login) Create(ctx context.Context) error {
	return mysql.LogDb.Model(&d).WithContext(ctx).Create(&d).Error
}

func (d *Login) BatchCreate(ctx context.Context, data []*Login) error {
	if data == nil {
		return errors.New("数据不能为空")
	}
	return mysql.LogDb.Model(&d).WithContext(ctx).Create(data).Error
}
