package game

import (
	"context"
	"xiaozhu/internal/logic/common"
	"xiaozhu/internal/model/key"
	"xiaozhu/utils/queue"
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
	Event     int    `json:"event" binding:"oneof=0 1 2 3"` // 事件类型： 0-upgrade、1-create、2-enter、3-online
	ExtData   string `json:"ext_data"`
}

func NewReportLogic(ctx context.Context) *ReportLogic {
	return &ReportLogic{ctx: ctx}
}

func (l *ReportLogic) Report() error {

	l.ReportRequest.RequestId = l.ctx.Value("request_id").(string)
	l.ReportRequest.UserId = l.ctx.Value("userId").(int)

	return queue.Push(l.ctx, key.RoleEventQueue, l.ReportRequest)

}
