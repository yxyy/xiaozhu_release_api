package log

import (
	"context"
	"gorm.io/gorm/clause"
	"xiaozhu/utils"
)

// Device  游戏设备表
type Device struct {
	Id          int    `json:"id" gorm:"id"`
	PromoteCode string `json:"promote_code" gorm:"promote_code"` // 推广码
	Adid        string `json:"adid" gorm:"adid"`                 // 广告id
	AppId       int    `json:"app_id" gorm:"app_id"`             // 游戏项目ID
	GameId      int    `json:"game_id" gorm:"game_id"`           // 游戏ID
	AppChannel  int    `json:"app_channel" gorm:"app_channel"`   // 应用渠道：1(官方)、2(安卓)、3(苹果)
	AreaCode    string `json:"area_code" gorm:"area_code"`       // 地区
	Area        string `json:"area" gorm:"area"`                 // 地区
	Os          string `json:"os" gorm:"os"`                     // 操作系统：android、ios
	Cause       string `json:"cause" gorm:"cause"`               // 广告归因依据
	DeviceId    string `json:"device_id" gorm:"device_id"`       // 设备号
	Ip          string `json:"ip" gorm:"ip"`                     // IP
	ChannelId   int    `json:"channel_id" gorm:"channel_id"`     // 媒体渠道
	// CampaignId  int64  `json:"campaign_id" gorm:"campaign_id"`   // 广告计划
	// AdgroupId   int64  `json:"adgroup_id" gorm:"adgroup_id"`     // 广告组
	// CreativeId  string `json:"creative_id" gorm:"creative_id"`   // 创意
	CreatedAt int64  `json:"created_at" gorm:"created_at"` // 创建时间
	Ts        int64  `json:"ts" gorm:"ts"`                 // 更新时间时间
	Days      int    `json:"days" gorm:"days"`             // 日期
	RequestId string `json:"request_id"`
}

// TableName 表名称
func (*Device) TableName() string {
	return "log_device"
}

func NewDevice() *Device {
	return &Device{}
}

func (d *Device) Create(ctx context.Context) error {
	return utils.MysqlLogDb.Model(&d).WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&d).Error
}
