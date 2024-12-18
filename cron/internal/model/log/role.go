package log

import (
	"context"
	"errors"
	"gorm.io/gorm/clause"
	"xiaozhu/internal/config/mysql"
)

type RoleInfo struct {
	Id          int    `json:"id" gorm:"id"`
	AppId       int    `json:"app_id" gorm:"app_id"`             // 游戏项目ID
	GameId      int    `json:"game_id" gorm:"game_id"`           // 游戏ID
	ChannelId   int    `json:"channel_id" gorm:"channel_id"`     // 媒体渠道：巨量、腾讯
	Os          string `json:"os" gorm:"os"`                     // 操作系统：android、ios
	DeviceId    string `json:"device_id" gorm:"device_id"`       // 设备号
	Ip          string `json:"ip" gorm:"ip"`                     // IP
	AreaCode    string `json:"area_code" gorm:"area_code"`       // 地区码
	Area        string `json:"area" gorm:"area"`                 // 地区
	UserId      int    `json:"user_id" gorm:"user_id"`           // 用户ID
	Account     string `json:"account" gorm:"account"`           // 帐户名
	ZoneId      int    `json:"zone_id" gorm:"zone_id"`           // 区服
	ZoneName    string `json:"zone_name" gorm:"zone_name"`       // 区服名
	RoleId      int    `json:"role_id" gorm:"role_id"`           // 角色名id
	RoleName    string `json:"role_name" gorm:"role_name"`       // 角色名
	RoleLevel   int    `json:"role_level" gorm:"role_level"`     // 角色级别
	PromoteCode string `json:"promote_code" gorm:"promote_code"` // 推广码
	Ts          int64  `json:"ts" gorm:"ts"`                     // 事件时间，即激活时间
	CreatedAt   int64  `json:"created_at" gorm:"created_at"`     // 创建时间
}

// MemberGameRole 角色登录日志
type MemberGameRole struct {
	RoleInfo
	Online      int    `json:"online"`
	LastIp      string `json:"last_ip"`
	LastPayAt   int64  `json:"last_pay_at"`
	LastLoginAt int64  `json:"last_login_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type RoleEvent struct {
	RoleInfo
	Event     int    `json:"event"`            // 事件类型： 0-upgrade、1-create、2-enter、3-online
	Days      int    `json:"days" gorm:"days"` // 日期
	RequestId string `json:"request_id"`
}

// TableName 表名称
func (*RoleEvent) TableName() string {
	return "log_role_event"
}

func (*MemberGameRole) TableName() string {
	return "member_game_role"
}

func NewRoleEvent() *RoleEvent {
	return &RoleEvent{}
}

func (r *RoleEvent) BatchCreate(ctx context.Context, data []*RoleEvent) error {
	if data == nil {
		return errors.New("数据不能为空")
	}
	return mysql.LogDb.Model(&r).WithContext(ctx).Create(data).Error
}

func NewMemberGameRole() *MemberGameRole {
	return &MemberGameRole{}
}

func (m *MemberGameRole) Save(ctx context.Context, data []*MemberGameRole) error {
	return mysql.PlatformDB.Model(&m).WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "game_id"}, {Name: "zone_id"}, {Name: "role_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"online", "last_ip", "last_pay_at", "last_login_at", "updated_at"}),
	}).Create(&data).Error
}
