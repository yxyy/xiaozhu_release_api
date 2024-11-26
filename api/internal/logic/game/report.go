package game

import (
	"context"
	"encoding/json"
	"xiaozhu/api/internal/logic/common"
	"xiaozhu/api/internal/model/key"
	"xiaozhu/api/utils"
)

type ReportLogic struct {
	ctx           context.Context
	ReportRequest ReportRequest
}

type ReportRequest struct {
	common.RequestForm
	ZoneId    int    `json:"zone_id" binding:"required"`
	ZoneName  string `json:"zone_name" binding:"required"`
	RoleId    int    `json:"role_id" binding:"required"`
	RoleName  string `json:"role_name" binding:"required"`
	RoleLevel int    `json:"role_level" binding:"required"`
	Behavior  int    `json:"behavior" binding:"oneof=0 1 2 3"` // 行为: 0-create、1-enter、2-upgrade、3-online
	ExtData   string `json:"ext_data"`
}

func NewReportLogic(ctx context.Context) *ReportLogic {
	return &ReportLogic{ctx: ctx}
}

func (l *ReportLogic) Report() error {

	l.ReportRequest.RequestId = l.ctx.Value("request_id").(string)

	marshal, err := json.Marshal(&l.ReportRequest)
	if err != nil {
		return err
	}

	if err = utils.RedisClient.LPush(l.ctx, key.RoleEventQueue, marshal).Err(); err != nil {
		return err
	}

	return nil
}
