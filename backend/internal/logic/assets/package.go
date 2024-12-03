package assets

import (
	"context"
	"errors"

	"xiaozhu/internal/model/assets"
	"xiaozhu/internal/model/common"
)

type PackageLogic struct {
	ctx context.Context
	assets.Package
	common.Params
}

func NewPackageLogic(ctx context.Context) *PackageLogic {
	return &PackageLogic{ctx: ctx}
}

func (l *PackageLogic) GetParams() *common.Params {
	l.Params.Verify()
	return &l.Params
}

type ListResponse struct {
	List  []*assets.PackageList `json:"list"`
	Total int64                 `json:"total"`
}

func (l *PackageLogic) List() (resp *ListResponse, err error) {

	list, t, err := l.Package.List(l.ctx, l.GetParams())
	if err != nil {
		return nil, err
	}

	resp = new(ListResponse)
	resp.List = list
	resp.Total = t

	return

}

func (l *PackageLogic) Create() error {
	if l.Name == "" {
		return errors.New("渠道包名称不能为空")
	}

	if l.ChannelId <= 0 {
		return errors.New("渠道不能为空")
	}

	if l.GameId <= 0 {
		return errors.New("游戏不能为空")
	}

	if l.Status < 0 {
		return errors.New("状态无效")
	}

	// l.CampaignId = utils.Random(8)
	// if l.CampaignId == "" {
	// 	return errors.New("渠道包标识生成错误")
	// }

	return l.Package.Create(l.ctx)
}

func (l *PackageLogic) Update() error {

	if l.Id <= 0 {
		return errors.New("id不能为空")
	}

	if l.Status < 0 {
		return errors.New("状态无效")
	}

	return l.Package.Update(l.ctx)
}
