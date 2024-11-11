package assets

import (
	"context"
	"errors"
	"xiaozhu/internal/model/common"
	"xiaozhu/internal/model/market/assets"
	"xiaozhu/utils"
)

type ProxyProjectLogic struct {
	ctx context.Context
	assets.ProxyProject
	common.Params
}

type ProxyProjectListResponse struct {
	List  []*assets.ProxyProjectList `json:"list"`
	Total int64                      `json:"total"`
}

func NewProxyProjectLogic(ctx context.Context) *ProxyProjectLogic {
	return &ProxyProjectLogic{ctx: ctx}
}

func (l *ProxyProjectLogic) GetParams() *common.Params {
	l.Params.Verify()
	return &l.Params
}

func (l *ProxyProjectLogic) Create() error {
	if l.Name == "" {
		return errors.New("名称不能为空")
	}
	if l.ChannelId < 1 {
		return errors.New("媒体渠道不能为空")
	}
	if l.PrincipalId < 1 {
		return errors.New("开户主体不能为空")
	}
	if l.ProxyCompanyId < 1 {
		return errors.New("代理商不能为空")
	}

	l.OptUser = l.ctx.Value("userId").(int)

	return l.ProxyProject.Create(l.ctx)
}

func (l *ProxyProjectLogic) Update() error {
	if l.Id <= 0 {
		return errors.New("id无效")
	}
	l.OptUser = l.ctx.Value("userId").(int)

	return l.ProxyProject.Update(l.ctx)
}

func (l *ProxyProjectLogic) List() (resp *ProxyProjectListResponse, err error) {

	list, total, err := l.ProxyProject.List(l.ctx, l.GetParams())
	if err != nil {
		return nil, err
	}

	resp = new(ProxyProjectListResponse)
	resp.List = list
	resp.Total = total

	return
}

func (l *ProxyProjectLogic) ListAll() (list map[int]*common.IdName, err error) {

	resp, err := l.ProxyProject.GetAll(l.ctx)
	if err != nil {
		return nil, err
	}
	return utils.ConvertIdNameMapById(resp), nil
}
