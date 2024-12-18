package queue

import (
	"encoding/json"
	"time"
	"xiaozhu/internal/model/assets"
	"xiaozhu/internal/model/common"
	"xiaozhu/internal/model/key"
	mod "xiaozhu/internal/model/log"
	"xiaozhu/utils/queue"
)

type RoleQueue struct {
	common.TopicExtra
	Message struct {
		common.RequestForm
		Account   string `json:"account" form:"account" gorm:"account"`
		Phone     int    `json:"phone" form:"phone" gorm:"phone"`
		MbCode    int    `json:"mb_code" form:"mb_code"`
		WxCode    string `json:"wx_code" form:"wx_code" gorm:"wx_code"`
		Email     string `json:"email" form:"code"`
		EmCode    string `json:"em_code" form:"em_code"`
		Event     int    `json:"event"`
		ZoneId    int    `json:"zone_id"`
		ZoneName  string `json:"zone_name"`
		RoleId    int    `json:"role_id"`
		RoleName  string `json:"role_name"`
		RoleLevel int    `json:"role_level"`
	}
}

func NewRoleQueue() *queue.Queue {
	return queue.NewBatchQueue(key.RoleEventQueue, &RoleQueue{}, 20)
}

func (l *RoleQueue) Run(q *queue.Queue, msg []string) (fail []string, err error) {
	if len(msg) == 0 {
		return nil, nil
	}

	var in RoleQueue
	var events []*mod.RoleEvent
	var memberGameRoles []*mod.MemberGameRole
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

		events = append(events, &mod.RoleEvent{
			RoleInfo: mod.RoleInfo{
				PromoteCode: "",
				AppId:       gameInfo.AppId,
				GameId:      in.Message.GameId,
				ChannelId:   0,
				Os:          in.Message.Os,
				UserId:      in.Message.UserId,
				Account:     in.Message.Account,
				ZoneId:      in.Message.ZoneId,
				ZoneName:    in.Message.ZoneName,
				RoleId:      in.Message.RoleId,
				RoleName:    in.Message.RoleName,
				RoleLevel:   in.Message.RoleLevel,
				DeviceId:    in.Message.DeviceId,
				Ip:          in.Message.Ip,
				AreaCode:    "",
				Area:        "",
				Ts:          ts,
				CreatedAt:   createdAt,
			},
			Event:     in.Message.Event,
			Days:      days,
			RequestId: in.Message.RequestId,
		})

		memberGameRoles = append(memberGameRoles, &mod.MemberGameRole{
			RoleInfo: mod.RoleInfo{
				PromoteCode: "",
				AppId:       gameInfo.AppId,
				GameId:      in.Message.GameId,
				ChannelId:   0,
				Os:          in.Message.Os,
				UserId:      in.Message.UserId,
				Account:     in.Message.Account,
				ZoneId:      in.Message.ZoneId,
				ZoneName:    in.Message.ZoneName,
				RoleId:      in.Message.RoleId,
				RoleName:    in.Message.RoleName,
				RoleLevel:   in.Message.RoleLevel,
				DeviceId:    in.Message.DeviceId,
				Ip:          in.Message.Ip,
				AreaCode:    "",
				Area:        "",
				Ts:          ts,
				CreatedAt:   createdAt,
			},
			LastIp: in.Message.Ip,
		})
	}

	if len(events) > 0 {
		event := mod.NewRoleEvent()
		if err = event.BatchCreate(q.Ctx, events); err != nil {
			return msg, err
		}
	}

	if len(memberGameRoles) > 0 {
		role := mod.NewMemberGameRole()
		if err = role.Save(q.Ctx, memberGameRoles); err != nil {
			return msg, err
		}
	}

	return

}

func (l *RoleQueue) loginWay() int {
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
