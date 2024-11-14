package assets

import (
	"context"
	"errors"
	"net/url"
	"xiaozhu/internal/logic/system"
	"xiaozhu/internal/model/common"
	"xiaozhu/internal/model/market/assets"
	"xiaozhu/utils"
)

type MarketAppIdLogic struct {
	ctx context.Context
	assets.MarketAppId
	common.Params
}

type MarketAppIdListResponse struct {
	List  []*assets.MarketAppIdList `json:"list"`
	Total int64                     `json:"total"`
}

func NewMarketAppIdLogic(ctx context.Context) *MarketAppIdLogic {
	return &MarketAppIdLogic{ctx: ctx}
}

func (l *MarketAppIdLogic) GetParams() *common.Params {
	l.Params.Verify()
	return &l.Params
}

func (l *MarketAppIdLogic) Create() error {
	if l.Name == "" {
		return errors.New("名称不能为空")
	}
	if l.ChannelId < 1 {
		return errors.New("媒体不能为空")
	}
	if l.AppId == "" {
		return errors.New("app_id不能为空")
	}
	if l.Secret == "" {
		return errors.New("secret 不能为空")
	}
	if l.Code == "" {
		return errors.New("标识 不能为空")
	}

	l.OptUser = l.ctx.Value("userId").(int)

	return l.MarketAppId.Create(l.ctx)
}

func (l *MarketAppIdLogic) Update() error {
	if l.Id <= 0 {
		return errors.New("id无效")
	}
	l.OptUser = l.ctx.Value("userId").(int)

	return l.MarketAppId.Update(l.ctx)
}

func (l *MarketAppIdLogic) List() (resp *MarketAppIdListResponse, err error) {

	list, total, err := l.MarketAppId.List(l.ctx, l.GetParams())
	if err != nil {
		return nil, err
	}

	resp = new(MarketAppIdListResponse)

	userLogic := system.NewUserLogic(l.ctx)
	userInfo, err := userLogic.ListAll()
	if err != nil {
		return nil, err
	}

	for _, v := range list {
		v.AuthUrl = oAuthUrl(v)
		v.OptUserName = userInfo[v.OptUser].Name
		resp.List = append(resp.List, v)
	}

	resp.Total = total

	return
}

func (l *MarketAppIdLogic) ListAll() (list map[int]*common.IdName, err error) {

	resp, err := l.MarketAppId.GetAll(l.ctx)
	if err != nil {
		return nil, err
	}
	return utils.ConvertIdNameMapById(resp), nil
}

func oAuthUrl(l *assets.MarketAppIdList) string {
	if l.ChannelCode == "bm" {
		return bmOAuthUrl(l)
	}

	if l.ChannelCode == "tx" {
		return txOAuthUrl(l)
	}

	return ""
}

// bmOAuthUrl  巨量授权地址
func bmOAuthUrl(l *assets.MarketAppIdList) string {
	var urls = make(url.Values)
	urls.Add("app_id", l.AppId)
	urls.Add("state", l.Code)
	urls.Add("redirect_uri", l.RedirectUri)
	return l.AuthUrl + "?" + urls.Encode()
}

// txOAuthUrl 腾讯授权地址
func txOAuthUrl(l *assets.MarketAppIdList) string {
	var urls = make(url.Values)
	urls.Add("client_id", l.AppId)
	urls.Add("state", l.Code)
	urls.Add("redirect_uri", l.RedirectUri)
	// urls.Add("scope", l.RedirectUri)

	// urls.Add("account_type", l.Params)

	return l.AuthUrl + "?" + urls.Encode()
}
