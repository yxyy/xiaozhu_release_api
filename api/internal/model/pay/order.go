package pay

import (
	"context"
	"errors"
	"xiaozhu/internal/config/mysql"
)

type Order struct {
	Id            int    `json:"id"`
	AppId         int    `json:"app_id"`
	GameId        int    `json:"game_id"`
	MerchantId    int    `json:"merchant_id"`
	PayChannel    int    `json:"pay_channel"`
	UserId        int    `json:"user_id"`
	Account       string `json:"account"`
	PromoteCode   string `json:"promote_code"`
	Openid        string `json:"openid"`
	GoodsId       string `json:"goods_id"`
	GoodsDesc     string `json:"goods_desc"`
	GoodsCurrency string `json:"goods_currency"`
	OrderNum      string `json:"order_num"`
	OrderType     int8   `json:"order_type"`
	OrderPrice    int    `json:"order_price"`
	PayMoney      int    `json:"pay_money"`
	PayCurrency   string `json:"pay_currency"`
	PayStatus     int8   `json:"pay_status"`
	PayAt         int64  `json:"pay_at"`
	CpOrderNum    string `json:"cp_order_num"`
	CpOrderStatus int8   `json:"cp_order_status"`
	CpOrderAt     int64  `json:"cp_order_at"`
	ZoneId        int    `json:"zone_id"`
	ZoneName      string `json:"zone_name"`
	RoleId        int    `json:"role_id"`
	RoleName      string `json:"role_name"`
	RoleLevel     int    `json:"role_level"`
	Os            string `json:"os"`
	DeviceId      string `json:"device_id"`
	Ip            string `json:"ip"`
	ExtData       string `json:"ext_data"`
	SandBox       int8   `json:"sand_box"`
	Remarks       string `json:"remarks"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
}

func (o *Order) TableName() string {
	return "pay_orders"
}

func NewOrder() *Order {
	return &Order{}
}

func (o *Order) Create(ctx context.Context) error {
	return mysql.PlatformDB.Model(&o).WithContext(ctx).Create(o).Error
}

func (o *Order) GetOrderInfo(ctx context.Context) error {
	if o.OrderNum == "" {
		return errors.New("订单不能为空")
	}
	return mysql.PlatformDB.Model(&o).WithContext(ctx).Where("orderNum", o.OrderNum).First(&o).Error
}
