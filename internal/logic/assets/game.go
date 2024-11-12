package assets

import (
	"context"
	"errors"
	"xiaozhu/internal/model/assets"
	"xiaozhu/internal/model/common"
	"xiaozhu/utils"
)

type GameLogic struct {
	ctx context.Context
	assets.Game
	common.Params
}

type GameListResponse struct {
	List  []*assets.GameList `json:"list"`
	Total int64              `json:"total"`
}

func NewGameLogic(ctx context.Context) *GameLogic {
	return &GameLogic{ctx: ctx}
}

func (l *GameLogic) GetParams() *common.Params {
	l.Params.Verify()
	return &l.Params
}

func (l *GameLogic) List() (resp *GameListResponse, err error) {

	list, total, err := l.Game.List(l.ctx, l.GetParams())
	if err != nil {
		return nil, err
	}

	resp = new(GameListResponse)
	resp.List = list
	resp.Total = total

	return
}

func (l *GameLogic) Create() error {
	if l.Game.GameName == "" {
		return errors.New("名称不能为空")
	}

	if l.Game.AppId <= 0 {
		return errors.New("应用不能为空")
	}

	if l.Game.Os <= 0 {
		return errors.New("操作系统不能为空")
	}

	if l.Game.Status < 0 || l.Game.Status > 1 {
		return errors.New("状态无效")
	}

	if l.Game.CpCallbackUrl == "" {
		return errors.New("回调地址不能为空")
	}
	if err := utils.ParseUrl(l.Game.CpCallbackUrl); err != nil {
		return errors.New("回调地址：" + err.Error())
	}

	if l.Game.CpTestCallbackUrl != "" {
		if err := utils.ParseUrl(l.Game.CpTestCallbackUrl); err != nil {
			return errors.New("测试回调地址：" + err.Error())
		}
	}

	l.Game.AppKey = utils.Md5(utils.Salt())
	l.Game.ServerKey = utils.Md5(utils.Salt())

	return l.Game.Create(l.ctx)
}

func (l *GameLogic) Update() error {
	if l.Game.Id <= 0 {
		return errors.New("id无效")
	}

	if l.Game.CpCallbackUrl != "" {
		if err := utils.ParseUrl(l.Game.CpCallbackUrl); err != nil {
			return errors.New("回调地址：" + err.Error())
		}
	}

	if l.Game.CpTestCallbackUrl != "" {
		if err := utils.ParseUrl(l.Game.CpTestCallbackUrl); err != nil {
			return errors.New("测试回调地址：" + err.Error())
		}
	}

	// 不更新key
	l.Game.AppKey = ""
	l.Game.ServerKey = ""

	return l.Game.Update(l.ctx)
}

func (l *GameLogic) ListAll() (resp []*assets.ListAllResponse, err error) {

	return l.Game.GetAll(l.ctx)
}
