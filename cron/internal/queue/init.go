package queue

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"time"
	"xiaozhu/cron/internal/model/common"
	"xiaozhu/cron/internal/model/key"
	"xiaozhu/cron/utils"
)

type InitRequest struct {
	common.RequestForm
	ProductKey string `json:"product_key" binding:"required"`
	Lang       string `json:"lang"`
	Debug      string `json:"debug" `
}

func Init(ctx context.Context) {
	for {
		result, err := utils.RedisClient.BRPop(ctx, time.Second*2, key.InitQueue).Result()
		if err != nil {
			log.Fatalf("获取队列:%s数据失败：%s", key.InitQueue, err)
			return
		}

		var initRequest InitRequest
		err = json.Unmarshal([]byte(result[1]), &initRequest)
		if err != nil {
			log.Errorf("序列化数据失败:%s", result[1])
			continue
		}
	}

}
