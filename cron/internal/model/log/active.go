package log

import (
	"context"
	"xiaozhu/internal/config/mysql"
)

// Active  激活日志
type Active struct {
	Id          int    `json:"id" gorm:"id"`
	PromoteCode string `json:"promote_code" gorm:"promote_code"` // 推广码
	Adid        string `json:"adid" gorm:"adid"`                 // 广告id
	AppId       int    `json:"app_id" gorm:"app_id"`             // 游戏项目ID
	GameId      int    `json:"game_id" gorm:"game_id"`           // 游戏ID
	AppChannel  int64  `json:"app_channel" gorm:"app_channel"`   // 应用渠道：1(官方)、2(安卓)、3(苹果)
	Os          string `json:"os" gorm:"os"`                     // 操作系统：android、ios
	Cause       string `json:"cause" gorm:"cause"`               // 广告归因依据
	DeviceId    string `json:"device_id" gorm:"device_id"`       // 设备号
	Ip          string `json:"ip" gorm:"ip"`                     // IP
	AreaCode    string `json:"area_code" gorm:"area_code"`       // 地区码
	Area        string `json:"area" gorm:"area"`                 // 地区
	Ts          int64  `json:"ts" gorm:"ts"`                     // 事件时间，即激活时间
	CreatedAt   int64  `json:"created_at" gorm:"created_at"`     // 创建时间
	Days        int    `json:"days" gorm:"days"`                 // 日期
	RequestId   string `json:"request_id"`
}

// TableName 表名称
func (*Active) TableName() string {
	return "log_active"
}

func NewActive() *Active {
	return &Active{}
}

func (d *Active) Create(ctx context.Context) error {
	return mysql.LogDb.Model(&d).WithContext(ctx).Create(&d).Error
}
