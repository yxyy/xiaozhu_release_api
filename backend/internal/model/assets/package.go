package assets

import (
	"context"
	"xiaozhu/backend/internal/model/common"
	"xiaozhu/backend/utils"
)

type Package struct {
	common.Model
	Name              string  `json:"name" gorm:"name"`                                 // 名称
	AppId             int64   `json:"app_id" gorm:"app_id"`                             // 应用id
	GameId            int64   `json:"game_id" gorm:"game_id"`                           // 游戏id
	ChannelId         int64   `json:"channel_id" gorm:"channel_id"`                     // 渠道id
	CampaignId        int64   `json:"campaign_id" gorm:"campaign_id"`                   // 默认生成的自然量广告id
	OriginalPackageId int64   `json:"original_package_id" gorm:"original_package_id"`   // 母包id
	SdkId             int64   `json:"sdk_id" gorm:"sdk_id"`                             // 游戏id
	SkinId            int64   `json:"skin_id" gorm:"skin_id"`                           // sdk——id
	PackageName       string  `json:"package_name" gorm:"package_name"`                 // 包名称
	Status            int8    `json:"status" gorm:"status"`                             // 打包状态
	PackStatus        int8    `json:"pack_status" gorm:"pack_status"`                   // 打包状态
	LastPackageTime   int64   `json:"last_package_time" gorm:"last_package_time"`       // 最新打包时间
	CpCallbackUrl     string  `json:"cp_callback_url" gorm:"cp_callback_url"`           // 发货地址
	CpCallbackTestUrl string  `json:"cp_callback_test_url" gorm:"cp_callback_test_url"` // 发货测试地址
	IsChangePay       int8    `json:"is_change_pay" gorm:"is_change_pay"`               // 是否切支付 0-否 1-是
	IsSdkFloatOn      int8    `json:"is_sdk_float_on" gorm:"is_sdk_float_on"`           // SDK浮点0隐藏，1显示
	IsUserFloatOn     int8    `json:"is_user_float_on" gorm:"is_user_float_on"`         // 用户浮点，0隐藏，1显示
	IsRegLoginOn      int8    `json:"is_reg_login_on" gorm:"is_reg_login_on"`           // 是否预下载 0关闭，1正常
	IsVisitorOn       int8    `json:"is_visitor_on" gorm:"is_visitor_on"`               // 是否开启游客模式，0不开，1开
	IsAutoLoginOn     int8    `json:"is_auto_login_on" gorm:"is_auto_login_on"`         // 是否开启自动登录，0关闭，1开启
	IsLogOn           int8    `json:"is_log_on" gorm:"is_log_on"`                       // 是否开启错误日志上报，0不开，1开
	IsShm             int8    `json:"is_shm" gorm:"is_shm"`                             // 是否切换审核服  0不是，1是
	SwitchLogin       int8    `json:"switch_login" gorm:"switch_login"`                 // 是否切渠道登录  0不切，1切
	OnlineTime        int64   `json:"online_time" gorm:"online_time"`                   // 上线时间
	Scheme            string  `json:"scheme" gorm:"scheme"`                             // urlscheme,支付后跳转url
	Privacy           string  `json:"privacy" gorm:"privacy"`                           // 隐私协议
	Rate              float64 `json:"rate" gorm:"rate"`                                 // 游戏币比例1rmb=10yxb
	IdCardVerify      int8    `json:"id_card_verify" gorm:"id_card_verify"`             // 实名认证弹窗( 0关闭 1开启可关闭 2开启不可关闭， 3绝对不开)
	IsLimitMinor      int8    `json:"is_limit_minor" gorm:"is_limit_minor"`             // 是否限制未成年人游戏和消费 0否 1是 2 绝对不开
	DenyReg           int8    `json:"deny_reg" gorm:"deny_reg"`                         // 禁止注册新用户，0否1是
	Remarks           string  `json:"remarks" gorm:"remarks"`                           // 备注说明
}

type PackageList struct {
	Package
	Os int `json:"os"`
	// AppId    int    `json:"app_id"`
	AppName     string `json:"app_name"`
	GameName    string `json:"game_name"`
	ChannelName string `json:"channel_name"`
}

func (p *Package) List(ctx context.Context, in *common.Params) (list []*PackageList, total int64, err error) {
	tx := utils.MysqlDb.Model(&p).WithContext(ctx).
		Select("packages.*,os,game_name,app_name,channels.name as channel_name").
		Joins("left join games on games.id = packages.game_id").
		Joins("left join apps on games.app_id = apps.id").
		Joins("left join channels on packages.channel_id = channels.id")
	if p.ChannelId > 0 {
		tx = tx.Where("channel_id", p.ChannelId)
	}
	if p.GameId > 0 {
		tx = tx.Where("game_id", p.GameId)
	}
	if p.CampaignId > 0 {
		tx = tx.Where("campaign", p.CampaignId)
	}
	if p.Status > 0 {
		tx = tx.Where("status", p.Status)
	}

	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err = tx.Offset(in.Offset).Limit(in.Limit).Scan(&list).Error

	return

}

func (p *Package) Create(ctx context.Context) error {
	return utils.MysqlDb.Model(&p).WithContext(ctx).Create(&p).Error
}

func (p *Package) Update(ctx context.Context) error {
	return utils.MysqlDb.Model(&p).WithContext(ctx).Where("id", p.Id).Updates(&p).Error
}
