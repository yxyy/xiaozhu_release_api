package pay

import (
	"context"
	"github.com/spf13/viper"
	"strconv"
	"xiaozhu/internal/logic/common"
	"xiaozhu/internal/model/assets"
	"xiaozhu/internal/model/pay"
	"xiaozhu/utils"
)

type OrderLogic struct {
	ctx     context.Context
	Request Request
}

type Request struct {
	common.RequestForm
	PayChannel int    `json:"pay_channel"  binding:"oneof=1 2 3 4 5"` // 1-支付宝 2-微信 3-米大师 4-谷歌 5-ios
	ZoneId     int    `json:"zone_id" binding:"required"`
	ZoneName   string `json:"zone_name" binding:"required"`
	RoleId     int    `json:"role_id" binding:"required"`
	RoleName   string `json:"role_name" binding:"required"`
	RoleLevel  int    `json:"role_level" binding:"required"`
	GoodsId    string `json:"goods_id"  binding:"required"`
	CpOrderNum string `json:"cp_order_num"  binding:"required"`
	ExtData    string `json:"ext_data"`
	SandBox    int8   `json:"sand_box"`
}

type Response struct {
	OrderNum string `json:"order_num"`
}

func NewOrderLogic(ctx context.Context) *OrderLogic {
	return &OrderLogic{ctx: ctx}
}

func (l *OrderLogic) Create() (*Response, error) {

	goods := pay.Goods{GameId: l.Request.GameId, GoodsId: l.Request.GoodsId}
	if err := goods.Find(l.ctx); err != nil {
		return nil, err
	}

	game, err := assets.GetAppGameInfo(l.ctx, l.Request.GameId)
	if err != nil {
		return nil, err
	}

	order := pay.Order{
		MerchantId:    viper.GetInt("pay." + strconv.Itoa(l.Request.PayChannel) + ".MerchantId"),
		AppId:         game.AppId,
		GameId:        game.Id,
		PayChannel:    l.Request.PayChannel,
		UserId:        l.ctx.Value("userId").(int),
		Account:       l.ctx.Value("account").(string),
		PromoteCode:   "",
		Openid:        "",
		GoodsId:       l.Request.GoodsId,
		GoodsDesc:     goods.GoodsDesc,
		GoodsCurrency: goods.Currency,
		OrderNum:      utils.Uuid(),
		OrderType:     0,
		OrderPrice:    goods.Amount,
		PayMoney:      goods.ActualAmount,
		PayCurrency:   goods.Currency,
		PayStatus:     0,
		PayAt:         0,
		CpOrderNum:    l.Request.CpOrderNum,
		CpOrderStatus: 0,
		CpOrderAt:     0,
		ZoneId:        l.Request.ZoneId,
		ZoneName:      l.Request.ZoneName,
		RoleId:        l.Request.RoleId,
		RoleName:      l.Request.RoleName,
		RoleLevel:     l.Request.RoleLevel,
		Os:            l.Request.Os,
		DeviceId:      l.Request.DeviceId,
		Ip:            l.Request.Ip,
		ExtData:       l.Request.ExtData,
		SandBox:       l.Request.SandBox,
		Remarks:       "",
	}

	if err = order.Create(l.ctx); err != nil {
		return nil, err
	}

	return &Response{OrderNum: order.OrderNum}, nil
}
