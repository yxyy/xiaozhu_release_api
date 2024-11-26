package assets

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"strconv"
	"sync"
	"xiaozhu/cron/internal/model/key"
	"xiaozhu/cron/utils"
)

var lockMap sync.Map

type Game struct {
	Id                int     `json:"game_id"`
	AppId             int     `json:"app_id" gorm:"app_id"`                             // 应用ID
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

type AppGame struct {
	Game
	AppName string `json:"app_name"`
}

func (*Game) TableName() string {
	return "games"
}

func NewAppGame() *AppGame {
	return &AppGame{}
}

func GetAppGameInfo(ctx context.Context, gameId int) (appGame *AppGame, err error) {
	if gameId <= 0 {
		return nil, errors.New("无效的游戏")
	}

	appGame = new(AppGame)
	gameIdKey := strconv.Itoa(gameId)
	keys := key.GameInfoPrefix + gameIdKey
	result, err := utils.RedisClient.Get(ctx, keys).Result()
	if err == nil {
		if result == key.CacheNotFound {
			return nil, errors.New("无效的游戏")
		}

		if err = json.Unmarshal([]byte(result), &appGame); err != nil {
			return nil, fmt.Errorf("缓存解析失败：%v", err)
		}
		return appGame, nil
	}

	if !errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("缓存读取失败: %v", err)
	}

	mu := getLock(gameId)
	mu.Lock()
	defer mu.Unlock()

	err = utils.MysqlDefaultDb.Table("games").
		WithContext(ctx).
		Select("games.*", "apps.app_name").
		Joins("left join apps on apps.id = games.app_id").
		Where("games.id", gameId).
		First(&appGame).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err = utils.RedisClient.Set(ctx, keys, key.CacheNotFound, key.GameInfoExpress).Err(); err != nil {
			return nil, fmt.Errorf("缓存设置失败: %v", err)
		}
		return nil, errors.New("无效的游戏")
	}

	if err != nil {
		return nil, fmt.Errorf("查询数据库失败: %v", err)
	}

	marshal, err := json.Marshal(&appGame)
	if err != nil {
		return nil, fmt.Errorf("数据序列化失败: %v", err)
	}

	err = utils.RedisClient.Set(ctx, keys, string(marshal), key.GameInfoExpress).Err()
	if err != nil {
		return nil, fmt.Errorf("缓存写入失败: %v", err)
	}

	return appGame, nil
}

func getLock(gameId int) *sync.Mutex {
	lock, _ := lockMap.LoadOrStore(gameId, &sync.Mutex{})
	return lock.(*sync.Mutex)
}
