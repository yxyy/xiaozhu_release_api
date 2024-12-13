package queue

import (
	"encoding/json"
	"time"
	"xiaozhu/internal/model/assets"
	"xiaozhu/internal/model/common"
	"xiaozhu/internal/model/key"
	logMod "xiaozhu/internal/model/log"
	"xiaozhu/utils/queue"
)

type LoginQueue struct {
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

func NewLoginQueue() *queue.Queue {
	return queue.NewBatchQueue(key.LoginQueue, &LoginQueue{}, 20)
}

func (l *LoginQueue) Run(q *queue.Queue, msg []string) (fail []string, err error) {
	if len(msg) == 0 {
		return nil, nil
	}

	var in LoginQueue
	var data []*logMod.Login
	createdAt := time.Now().Unix()
	for _, v := range msg {
		if err = json.Unmarshal([]byte(v), &in); err != nil {
			q.Log.Errorf("消息：%s,序列化数据失败:%s", v, err)
			fail = append(fail, v)
			continue
		}

		logs := q.Log.WithField("request_id", in.Message.RequestId) // 后面使用新的logs，防止污染全局的q.log
		days, ts, err := parseTimestamp(in.Message.Ts)
		if err != nil {
			q.Log.Warn(err)
		}

		gameInfo, err := assets.GetAppGameInfo(q.Ctx, in.Message.GameId)
		if err != nil {
			logs.Errorf("获取游戏信息失败:%s", err)
			fail = append(fail, v)
			continue
		}

		data = append(data, &logMod.Login{
			Id:          0,
			PromoteCode: "",
			AppId:       gameInfo.AppId,
			GameId:      in.Message.GameId,
			AppChannel:  0,
			Os:          in.Message.Os,
			// Cause:       "",
			UserId:    in.Message.UserId,
			Account:   in.Message.Account,
			LoginWay:  in.loginWay(),
			DeviceId:  in.Message.DeviceId,
			Ip:        in.Message.Ip,
			AreaCode:  "",
			Area:      "",
			Ts:        ts,
			CreatedAt: createdAt,
			Days:      days,
			RequestId: in.Message.RequestId,
		})
	}

	if len(data) == 0 {
		return
	}

	login := logMod.NewLogin()
	if err = login.BatchCreate(q.Ctx, data); err != nil {
		return msg, err
	}

	return
}

func (l *LoginQueue) loginWay() int {
	switch {
	case l.Message.Account != "": // 账号登录
		return 1
	case l.Message.Phone != 0: //  手机登录
		return 2
	case l.Message.Email != "": // 邮箱登录
		return 3
	case l.Message.WxCode != "": // 微信登录
		return 4
	default:
		return 0
	}
}
