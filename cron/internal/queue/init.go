package queue

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
	"xiaozhu/cron/internal/model/assets"
	"xiaozhu/cron/internal/model/common"
	"xiaozhu/cron/internal/model/key"
	logMod "xiaozhu/cron/internal/model/log"
	"xiaozhu/cron/utils"
)

const InitRetryLimit = 3

type InitRequest struct {
	common.RequestForm
	ProductKey string `json:"product_key"`
	Lang       string `json:"lang"`
	Debug      string `json:"debug" `
}

// 使用缓存chan,防止暴增
var jobChan = make(chan struct{}, 100)

func Init(ctx context.Context) {
	for {
		result, err := utils.RedisClient.BRPop(ctx, time.Second*10, key.InitQueue).Result()
		if err != nil {
			log.Errorf("获取队列:%s数据失败：%s", key.InitQueue, err)
			continue
		}
		jobChan <- struct{}{}
		go runInit(ctx, result[1])
	}
}

func runInit(ctx context.Context, result string) {
	defer func() {
		<-jobChan
	}()
	var initRequest InitRequest
	err := json.Unmarshal([]byte(result), &initRequest)
	if err != nil {
		log.Errorf("序列化数据失败:%s", result)
		return
	}

	logs := log.WithField("request_id", initRequest.RequestId)
	days := parseTimestamp(initRequest.Ts, logs)
	ts := initRequest.Ts
	if ts <= 0 {
		ts = time.Now().Unix()
	}

	gameInfo, err := assets.GetAppGameInfo(ctx, initRequest.GameId)
	if err != nil {
		log.Errorf("获取游戏信息失败:%s", err)
		return
	}

	active := logMod.Active{
		AppId:      gameInfo.AppId,
		GameId:     initRequest.GameId,
		AppChannel: 0,
		Os:         initRequest.Os,
		Cause:      "",
		DeviceId:   initRequest.DeviceId,
		Ip:         initRequest.Ip,
		AreaCode:   "",
		Area:       "",
		Ts:         ts,
		CreatedAt:  time.Now().Unix(),
		Days:       days,
		RequestId:  initRequest.RequestId,
	}

	device := logMod.Device{
		Id:          0,
		PromoteCode: "",
		Adid:        "",
		AppId:       gameInfo.AppId,
		GameId:      initRequest.GameId,
		AppChannel:  0,
		AreaCode:    "",
		Area:        "",
		Os:          initRequest.Os,
		Cause:       "",
		DeviceId:    initRequest.DeviceId,
		Ip:          initRequest.Ip,
		ChannelId:   0,
		CreatedAt:   0,
		Ts:          ts,
		Days:        days,
		RequestId:   initRequest.RequestId,
	}

	err = utils.MysqlLogDb.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&active).WithContext(ctx).Create(&active).Error
		if err != nil {
			return fmt.Errorf("插入 Active 失败: %w", err)
		}
		err = tx.Model(&device).WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&device).Error
		if err != nil {
			return fmt.Errorf("插入 Device 失败: %w", err)
		}
		return nil
	})

	if err != nil {
		logs.Errorf("数据插入失败:%s", err)
		handleInitRetry(ctx, initRequest, logs)
	}
}

func handleInitRetry(ctx context.Context, req InitRequest, logs *log.Entry) {
	if req.ReTry >= InitRetryLimit {
		logs.Warn("达到最大重试次数，不再重试")
		return
	}
	req.ReTry++
	data, err := json.Marshal(req)
	if err != nil {
		logs.Errorf("重试任务序列化失败: %v", err)
		return
	}
	if err = utils.RedisClient.LPush(ctx, key.InitQueue, data).Err(); err != nil {
		logs.Errorf("任务重新入队失败: %v", err)
	}
}

func parseTimestamp(ts int64, logs *log.Entry) int {
	format := time.Unix(ts, 0).Format("20060402")
	days, err := strconv.Atoi(format)
	if err != nil {
		logs.Errorf("时间转换失败: %v", err)
		return int(time.Now().Unix())
	}
	return days
}
