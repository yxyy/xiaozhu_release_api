package assets

import (
	"context"
	"xiaozhu/internal/model/common"
	"xiaozhu/utils"
)

type Game struct {
	common.Model
	AppId             int64   `json:"app_id" gorm:"app_id"`                             // 应用ID
	PkgName           string  `json:"pkg_name" gorm:"pkg_name"`                         // 包名全局唯一
	GameName          string  `json:"game_name" gorm:"game_name"`                       // 游戏名
	CpCallbackUrl     string  `json:"cp_callback_url" gorm:"cp_callback_url"`           // CP 发货正式接口
	CpTestCallbackUrl string  `json:"cp_test_callback_url" gorm:"cp_test_callback_url"` // CP 发货测试接口
	AppKey            string  `json:"app_key" gorm:"app_key"`                           // 发货key
	ServerKey         string  `json:"server_key" gorm:"server_key"`                     // 服务端key
	Os                int8    `json:"os" gorm:"os"`                                     // 应用类型：1-Android 2-IOS 3-H5 4-小程序
	LinkH5            string  `json:"link_h5" gorm:"link_h5"`                           // H5 链接
	DownloadUrl       string  `json:"download_url" gorm:"download_url"`                 // 游戏下载地址
	Status            int8    `json:"status" gorm:"status"`                             // 状态: 对接中(0)、已上线 (1) 、已下线(2)
	Conversion        float32 `json:"conversion" gorm:"conversion"`                     // 人民币和游戏币转换倍率，人民币是 1
	Icon              string  `json:"icon" gorm:"icon"`                                 // icon
	Remark            string  `json:"remark" gorm:"remark"`                             // 备注
	PublishAt         int64   `json:"publish_at" gorm:"publish_at"`                     // 发布时间
	IsAuthRealName    int8    `json:"is_auth_real_name" gorm:"is_auth_real_name"`       // 是否需要实名认证  0-是 1-否
	IsLimitUnderage   int8    `json:"is_limit_underage" gorm:"is_limit_underage"`       // 是否限制未成年 0-是 1-否
	Signature         int8    `json:"signature" gorm:"signature"`                       // 签名方式 0-md5
}

type GameList struct {
	Game
	AppName string `json:"app_name"`
}

func (g *Game) Create(ctx context.Context) error {
	return utils.MysqlDb.Model(&g).WithContext(ctx).Create(&g).Error
}

func (g *Game) Update(ctx context.Context) error {
	return utils.MysqlDb.Model(&g).WithContext(ctx).Where("id", g.Id).Updates(&g).Error
}

func (g *Game) List(ctx context.Context, params *common.Params) (resp []*GameList, total int64, err error) {
	tx := utils.MysqlDb.Model(&g).WithContext(ctx).
		Select("games.*,apps.app_name AS app_name").
		Joins("left join apps on games.app_id = apps.id")
	if g.Id > 0 {
		tx = tx.Where("games.id", g.Id)
	}
	if g.AppId > 0 {
		tx = tx.Where("games.app_id", g.AppId)
	}
	if g.GameName != "" {
		tx = tx.Where("games.game_name like ?", "%"+g.GameName+"%")
	}

	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err = tx.Offset(params.Offset).Limit(params.Limit).Scan(&resp).Error

	return

}

type ListAllResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name" gorm:"game_name"`
	Os   int8   `json:"os" gorm:"os"`
}

func (g *Game) GetAll(ctx context.Context) (resp []*ListAllResponse, err error) {

	err = utils.MysqlDb.Model(&g).WithContext(ctx).Find(&resp).Error
	return
}
