package pay

import (
	"context"
)

// MidasLogic https://docs.qq.com/doc/DR1hhWlpnQXJXWHRh
// pay_event_sig = to_hex(hmac_sha256(app_key, event + '&' + payload))
type MidasLogic struct {
	ctx          context.Context
	MidasRequest MidasRequest
}

type MidasRequest struct {
	ToUserName   string   `json:"ToUserName"`
	FromUserName string   `json:"FromUserName"`
	CreateTime   int64    `json:"CreateTime"`
	MsgType      string   `json:"MsgType"`
	Event        string   `json:"Event"`
	MiniGame     MiniGame `json:"MiniGame"`
}

type MiniGame struct {
	Payload     string `json:"Payload"`
	PayEventSig string `json:"PayEventSig"`
	IsMock      bool   `json:"IsMock"`
}

type Payload struct {
	OpenId     string    `json:"OpenId"`
	Env        int       `json:"Env"`
	OutTradeNo string    `json:"OutTradeNo"`
	GoodsInfo  GoodsInfo `json:"GoodsInfo"`
}

type GoodsInfo struct {
	ProductId   string  `json:"ProductId"`
	Quantity    int     `json:"Quantity"`
	ZoneId      string  `json:"ZoneId"`
	OrigPrice   float64 `json:"OrigPrice"`
	ActualPrice float64 `json:"ActualPrice"`
	Attach      string  `json:"Attach"`
	OrderSource int     `json:"OrderSource"`
}

type MidasResponse struct {
	ErrCode int    `json:"ErrCode"`
	ErrMsg  string `json:"ErrMsg"`
}

func (l *MidasLogic) Notify() error {

	return nil
}
