package queue

import (
	"context"
	"encoding/json"
	"time"
	"xiaozhu/internal/model/assets"
	"xiaozhu/internal/model/common"
	"xiaozhu/internal/model/key"
	logMod "xiaozhu/internal/model/log"
	"xiaozhu/utils/queue"
)

type RegisterQueue struct {
	common.TopicExtra
	Message struct {
		common.RequestForm
		Account string `json:"account" form:"account" gorm:"account"`
		Phone   int    `json:"phone" form:"phone" gorm:"phone"`
		MbCode  int    `json:"mb_code" form:"mb_code"`
		WxCode  string `json:"wx_code" form:"wx_code" gorm:"wx_code"`
		Email   string `json:"email" form:"code"`
		EmCode  string `json:"em_code" form:"em_code"`
	}
}

func NewRegisterQueue() *queue.Queue {
	return queue.NewBatchQueue(key.RegisterQueue, &RegisterQueue{}, 20)
}

func (l *RegisterQueue) Run(q *queue.Queue, msg []string) (fail []string, err error) {
	if len(msg) == 0 {
		return nil, nil
	}

	var in RegisterQueue
	var data []*logMod.Register
	createdAt := time.Now().Unix()
	for _, v := range msg {
		if err = json.Unmarshal([]byte(v), &in); err != nil {
			q.Log.Errorf("消息：%s,序列化数据失败:%s", v, err)
			fail = append(fail, v)
			continue
		}

		ctx := context.WithValue(q.Ctx, "request_id", in.Message.RequestId)
		days, ts, err := parseTimestamp(in.Message.Ts)
		if err != nil {
			q.Log.WithContext(ctx).Warn(err)
		}

		gameInfo, err := assets.GetAppGameInfo(q.Ctx, in.Message.GameId)
		if err != nil {
			q.Log.WithContext(ctx).Errorf("获取游戏信息失败:%s", err)
			fail = append(fail, v)
			continue
		}

		data = append(data, &logMod.Register{
			Id:          0,
			PromoteCode: "",
			AppId:       gameInfo.AppId,
			GameId:      in.Message.GameId,
			Os:          in.Message.Os,
			UserId:      in.Message.UserId,
			Account:     in.Message.Account,
			DeviceId:    in.Message.DeviceId,
			Ip:          in.Message.Ip,
			AreaCode:    "",
			Area:        "",
			Ts:          ts,
			CreatedAt:   createdAt,
			Days:        days,
			RequestId:   in.Message.RequestId,
		})
	}

	if len(data) == 0 {
		return
	}

	login := logMod.NewRegister()
	if err = login.BatchCreate(q.Ctx, data); err != nil {
		return msg, err
	}

	return
}
