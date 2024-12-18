package log

import (
	"context"
	"errors"
	"xiaozhu/internal/config/mysql"
)

// Register  注册日志
type Register struct {
	Id          int    `json:"id" gorm:"id"`
	PromoteCode string `json:"promote_code" gorm:"promote_code"` // 推广码
	AppId       int    `json:"app_id" gorm:"app_id"`             // 游戏项目ID
	GameId      int    `json:"game_id" gorm:"game_id"`           // 游戏ID
	ChannelId   int    `json:"channel_id" gorm:"channel_id"`     // 媒体渠道：巨量、腾讯
	Os          string `json:"os" gorm:"os"`                     // 操作系统：android、ios
	UserId      int    `json:"user_id"`
	Account     string `json:"account"`
	DeviceId    string `json:"device_id" gorm:"device_id"`   // 设备号
	Ip          string `json:"ip" gorm:"ip"`                 // IP
	AreaCode    string `json:"area_code" gorm:"area_code"`   // 地区码
	Area        string `json:"area" gorm:"area"`             // 地区
	Ts          int64  `json:"ts" gorm:"ts"`                 // 事件时间，即激活时间
	CreatedAt   int64  `json:"created_at" gorm:"created_at"` // 创建时间
	Days        int    `json:"days" gorm:"days"`             // 日期
	RequestId   string `json:"request_id"`
}

// TableName 表名称
func (*Register) TableName() string {
	return "log_register"
}

func NewRegister() *Register {
	return &Register{}
}

func (d *Register) BatchCreate(ctx context.Context, data []*Register) error {
	if data == nil {
		return errors.New("数据不能为空")
	}
	return mysql.LogDb.Model(&d).WithContext(ctx).Create(data).Error
}
