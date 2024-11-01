package system

import (
	"context"
	"xiaozhu/internal/model/common"
	"xiaozhu/internal/model/system"
)

type SysUserLogLogic struct {
	ctx        context.Context
	SysUserLog system.SysUserLog
	Params     common.Params
}

type SysUserLogListResponse struct {
	List  []*system.SysUserLog `json:"list"`
	Total int64                `json:"total"`
}

func NewSysUserLogLogic(ctx context.Context) *SysUserLogLogic {
	return &SysUserLogLogic{ctx: ctx}
}

func (l SysUserLogLogic) GetParams() *common.Params {
	l.Params.Verify()
	return &l.Params
}

func (l *SysUserLogLogic) ListLogic() (resp *SysUserLogListResponse, err error) {

	list, total, err := l.SysUserLog.List(l.ctx, l.GetParams())
	if err != nil {
		return nil, err
	}

	resp = new(SysUserLogListResponse)
	resp.List = list
	resp.Total = total

	return
}

//
// func (l *SysUserLogLogic) Create() error {
// 	if l.SysUserLog.AppName == "" {
// 		return errors.New("名称不能为空")
// 	}
// 	if l.SysRole.Code == "" {
// 		return errors.New("标识不能为空不能为空")
// 	}
//
// 	l.SysRole.Status = new(int)
//
// 	return l.SysRole.Create()
// }

// func (l *SysUserLogLogic) Update() error {
// 	if l.SysRole.Id <= 0 {
// 		return errors.New("id 无效")
// 	}
//
// 	return l.SysRole.Update()
// }
