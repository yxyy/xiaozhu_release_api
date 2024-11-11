package assets

import (
	"context"
	"errors"
	"xiaozhu/internal/model/common"
	"xiaozhu/internal/model/market/assets"
	"xiaozhu/utils"
)

type ProxyCompanyLogic struct {
	ctx context.Context
	assets.ProxyCompany
	common.Params
}

type ProxyCompanyListResponse struct {
	List  []*assets.ProxyCompany `json:"list"`
	Total int64                  `json:"total"`
}

func NewProxyCompanyLogic(ctx context.Context) *ProxyCompanyLogic {
	return &ProxyCompanyLogic{ctx: ctx}
}

func (l *ProxyCompanyLogic) GetParams() *common.Params {
	l.Params.Verify()
	return &l.Params
}

func (l *ProxyCompanyLogic) Create() error {
	if l.Name == "" {
		return errors.New("名称不能为空")
	}
	if l.Code == "" {
		return errors.New("标识不能为空")
	}

	l.OptUser = l.ctx.Value("userId").(int)

	return l.ProxyCompany.Create(l.ctx)
}

func (l *ProxyCompanyLogic) Update() error {
	if l.Id <= 0 {
		return errors.New("id无效")
	}
	l.OptUser = l.ctx.Value("userId").(int)

	return l.ProxyCompany.Update(l.ctx)
}

func (l *ProxyCompanyLogic) List() (resp *ProxyCompanyListResponse, err error) {

	list, total, err := l.ProxyCompany.List(l.ctx, l.GetParams())
	if err != nil {
		return nil, err
	}

	resp = new(ProxyCompanyListResponse)
	resp.List = list
	resp.Total = total

	return
}

func (l *ProxyCompanyLogic) ListAll() (list map[int]*common.IdName, err error) {

	resp, err := l.ProxyCompany.GetAll(l.ctx)
	if err != nil {
		return nil, err
	}
	return utils.ConvertIdNameMapById(resp), nil
}
