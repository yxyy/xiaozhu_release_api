package assets

import (
	"context"
	"errors"
	"xiaozhu/internal/model/assets"
	"xiaozhu/internal/model/common"
	"xiaozhu/utils"
)

type ChannelLogic struct {
	ctx context.Context
	assets.Channel
	common.Params
}

type ChannelListResponse struct {
	List  []*assets.Channel `json:"list"`
	Total int64             `json:"total"`
}

func NewChannelLogic(ctx context.Context) *ChannelLogic {
	return &ChannelLogic{ctx: ctx}
}

func (l *ChannelLogic) GetParams() *common.Params {
	l.Params.Verify()
	return &l.Params
}

func (l *ChannelLogic) Create() error {
	if l.Name == "" {
		return errors.New("名称不能为空")
	}
	if l.Code == "" {
		return errors.New("标识不能为空")
	}

	l.OptUser = l.ctx.Value("userId").(int)

	return l.Channel.Create(l.ctx)
}

func (l *ChannelLogic) Update() error {
	if l.Id <= 0 {
		return errors.New("id无效")
	}
	l.OptUser = l.ctx.Value("userId").(int)

	return l.Channel.Update(l.ctx)
}

func (l *ChannelLogic) List() (resp *ChannelListResponse, err error) {

	list, total, err := l.Channel.List(l.ctx, l.GetParams())
	if err != nil {
		return nil, err
	}

	resp = new(ChannelListResponse)
	resp.List = list
	resp.Total = total

	return
}

func (l *ChannelLogic) ListAll() (list map[int]*common.IdName, err error) {

	resp, err := l.Channel.GetAll(l.ctx)
	if err != nil {
		return nil, err
	}
	return utils.ConvertIdNameMapById(resp), nil
}
