package queue

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
	"xiaozhu/internal/model/assets"
	"xiaozhu/internal/model/common"
	logMod "xiaozhu/internal/model/log"
	"xiaozhu/utils"
)

type InitQueue struct {
	common.RequestForm
	ProductKey string `json:"product_key"`
	Lang       string `json:"lang"`
	Debug      string `json:"debug" `
}

func NewInitQueue() *InitQueue {
	return &InitQueue{}
}

func (l *InitQueue) Run(q *Queue, result string) error {

	err := json.Unmarshal([]byte(result), &l)
	if err != nil {
		q.log.Errorf("序列化数据失败:%s", result)
		return err
	}

	createdAt := time.Now().Unix()
	logs := q.log.WithField("request_id", l.RequestId) // 后面使用新的logs，防止污染全局的q.log
	days := parseTimestamp(l.Ts, logs)
	ts := l.Ts
	if ts <= 0 {
		ts = time.Now().Unix()
	}

	gameInfo, err := assets.GetAppGameInfo(q.ctx, l.GameId)
	if err != nil {
		logs.Errorf("获取游戏信息失败:%s", err)
		return err
	}

	active := logMod.Active{
		AppId:      gameInfo.AppId,
		GameId:     l.GameId,
		AppChannel: 0,
		Os:         l.Os,
		Cause:      "",
		DeviceId:   l.DeviceId,
		Ip:         l.Ip,
		AreaCode:   "",
		Area:       "",
		Ts:         ts,
		CreatedAt:  createdAt,
		Days:       days,
		RequestId:  l.RequestId,
	}

	device := logMod.Device{
		Id:          0,
		PromoteCode: "",
		Adid:        "",
		AppId:       gameInfo.AppId,
		GameId:      l.GameId,
		AppChannel:  0,
		AreaCode:    "",
		Area:        "",
		Os:          l.Os,
		Cause:       "",
		DeviceId:    l.DeviceId,
		Ip:          l.Ip,
		ChannelId:   0,
		CreatedAt:   createdAt,
		Ts:          ts,
		Days:        days,
		RequestId:   l.RequestId,
	}

	err = utils.MysqlLogDb.WithContext(q.ctx).Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&active).WithContext(q.ctx).Create(&active).Error
		if err != nil {
			return fmt.Errorf("插入 Active 失败: %w", err)
		}
		err = tx.Model(&device).WithContext(q.ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&device).Error
		if err != nil {
			return fmt.Errorf("插入 Device 失败: %w", err)
		}
		return nil
	})

	if err != nil {
		logs.Errorf("数据插入失败:%s", err)
		return err
	}
	return nil
}
