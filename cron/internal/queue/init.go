package queue

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
	"xiaozhu/internal/model/assets"
	"xiaozhu/internal/model/common"
	"xiaozhu/internal/model/key"
	logMod "xiaozhu/internal/model/log"
	"xiaozhu/utils"
	"xiaozhu/utils/queue"
)

type InitQueue struct {
	common.RequestForm
	ProductKey string `json:"product_key"`
	Lang       string `json:"lang"`
	Debug      string `json:"debug" `
}

func NewInit() *InitQueue {
	return &InitQueue{}
}

func NewInitQueue() *queue.Queue {
	return queue.NewQueue(key.InitQueue, NewInit())
}

func (l *InitQueue) Run(q *queue.Queue, msg string) error {
	if msg == "" {
		q.Log.Errorf("消息为空的")
		return errors.New("消息为空的")
	}

	err := json.Unmarshal([]byte(msg), &l)
	if err != nil {
		q.Log.Errorf("消息：%s,序列化数据失败:%s", msg, err)
		return err
	}

	createdAt := time.Now().Unix()
	logs := q.Log.WithField("request_id", l.RequestId) // 后面使用新的logs，防止污染全局的q.log
	days := parseTimestamp(l.Ts, logs)
	ts := l.Ts
	if ts <= 0 {
		ts = time.Now().Unix()
	}

	gameInfo, err := assets.GetAppGameInfo(q.Ctx, l.GameId)
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

	err = utils.MysqlLogDb.WithContext(q.Ctx).Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&active).WithContext(q.Ctx).Create(&active).Error
		if err != nil {
			return fmt.Errorf("插入 Active 失败: %w", err)
		}
		err = tx.Model(&device).WithContext(q.Ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&device).Error
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
