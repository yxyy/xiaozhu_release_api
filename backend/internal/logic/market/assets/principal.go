package assets

import (
	"context"
	"errors"
	"xiaozhu/backend/internal/model/common"
	"xiaozhu/backend/internal/model/market/assets"
	"xiaozhu/backend/utils"
)

type PrincipalLogic struct {
	ctx context.Context
	assets.Principal
	common.Params
}

type PrincipalListResponse struct {
	List  []*assets.Principal `json:"list"`
	Total int64               `json:"total"`
}

func NewPrincipalLogic(ctx context.Context) *PrincipalLogic {
	return &PrincipalLogic{ctx: ctx}
}

func (l *PrincipalLogic) GetParams() *common.Params {
	l.Params.Verify()
	return &l.Params
}

func (l *PrincipalLogic) Create() error {
	if l.Name == "" {
		return errors.New("名称不能为空")
	}
	if l.Code == "" {
		return errors.New("标识不能为空")
	}

	l.OptUser = l.ctx.Value("userId").(int)

	return l.Principal.Create(l.ctx)
}

func (l *PrincipalLogic) Update() error {
	if l.Id <= 0 {
		return errors.New("id无效")
	}
	l.OptUser = l.ctx.Value("userId").(int)

	return l.Principal.Update(l.ctx)
}

func (l *PrincipalLogic) List() (resp *PrincipalListResponse, err error) {

	list, total, err := l.Principal.List(l.ctx, l.GetParams())
	if err != nil {
		return nil, err
	}

	resp = new(PrincipalListResponse)
	resp.List = list
	resp.Total = total

	return
}

func (l *PrincipalLogic) ListAll() (list map[int]*common.IdName, err error) {

	resp, err := l.Principal.GetAll(l.ctx)
	if err != nil {
		return nil, err
	}
	return utils.ConvertIdNameMapById(resp), nil
}
