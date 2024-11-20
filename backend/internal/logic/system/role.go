package system

import (
	"context"
	"errors"
	"xiaozhu/internal/model/common"
	"xiaozhu/internal/model/system"
)

type SysRoleLogic struct {
	ctx     context.Context
	SysRole system.SysRole
	Params  common.Params
}

type RoleListResponse struct {
	List  []*system.SysRole `json:"list"`
	Total int64             `json:"total"`
}

func NewSysRoleLogic(ctx context.Context) *SysRoleLogic {
	return &SysRoleLogic{ctx: ctx}
}

func (l *SysRoleLogic) GetParams() *common.Params {
	l.Params.Verify()
	return &l.Params
}

func (l *SysRoleLogic) ListLogic() (resp *RoleListResponse, err error) {

	request := &system.RoleListRequest{
		Name:   l.SysRole.Name,
		Code:   l.SysRole.Code,
		Status: l.SysRole.Status,
		Params: l.GetParams(),
		Ids:    nil,
	}
	list, total, err := l.SysRole.List(l.ctx, request)
	if err != nil {
		return nil, err
	}

	resp = new(RoleListResponse)
	resp.List = list
	resp.Total = total

	return
}

func (l *SysRoleLogic) ListAllLogic() (resp []*common.IdName, err error) {

	l.Params.Offset = 0
	l.Params.Limit = 10000
	request := &system.RoleListRequest{
		Name:   l.SysRole.Name,
		Code:   l.SysRole.Code,
		Status: l.SysRole.Status,
		Params: &l.Params,
		Ids:    nil,
	}
	list, _, err := l.SysRole.List(l.ctx, request)
	if err != nil {
		return nil, err
	}

	for _, v := range list {
		resp = append(resp, &common.IdName{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	return
}

func (l *SysRoleLogic) Create() error {
	if l.SysRole.Name == "" {
		return errors.New("名称不能为空")
	}
	if l.SysRole.Code == "" {
		return errors.New("标识不能为空不能为空")
	}

	l.SysRole.Status = new(int)

	return l.SysRole.Create()
}

func (l *SysRoleLogic) Update() error {
	if l.SysRole.Id <= 0 {
		return errors.New("id 无效")
	}

	return l.SysRole.Update()
}

func (l *SysRoleLogic) UpdateMenu() error {
	if l.SysRole.Id <= 0 {
		return errors.New("id 无效")
	}

	if l.SysRole.DataScopeMenuIds == "" {
		return errors.New("菜单权限不能为空")
	}

	return l.SysRole.Update()
}
