package system

import (
	"context"
	"errors"
	"xiaozhu/internal/model/system"
)

type DeptLogic struct {
	ctx     context.Context
	SysDept system.SysDept
}

func NewDeptLogic(ctx context.Context) *DeptLogic {
	return &DeptLogic{ctx: ctx}
}

func (l *DeptLogic) List(request system.SysDeptListRequest) (*system.SysDeptListResponse, error) {

	dept := system.NewSysDept()
	return dept.List(l.ctx, request)

}

func (l *DeptLogic) Create() error {
	if l.SysDept.Name == "" {
		return errors.New("部门名称不能为空")
	}

	return l.SysDept.Create(l.ctx)
}

func (l *DeptLogic) Update() error {
	if l.SysDept.Id < 1 {
		return errors.New("部门id不能为空")
	}

	return l.SysDept.Update(l.ctx)
}
