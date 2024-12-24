package pay

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

// MidasLogic https://docs.qq.com/doc/DR1hhWlpnQXJXWHRh
// pay_event_sig = to_hex(hmac_sha256(app_key, event + '&' + payload))
type MidasLogic struct {
	ctx          context.Context
	MidasRequest MidasRequest
}

/*
{
    "ToUserName": "$your_username",
    "FromUserName": "oUrsf0TSXNtiZjP7JL9UUFiGJzmQ",
    "CreateTime": "$current_timestamp",
	"MsgType": "event",
	"Event": "minigame_game_pay_goods_deliver_notify",
	"MiniGame": {
		"Payload": "{\"OpenId\":\"example_open_id\",\"OutTradeNo\":\"example_out_trade_no\",\"WeChatPayInfo\":{\"MchOrderNo\":\"example_mch_order_no\",\"TransactionId\":\"example_transaction_id\"},\"Env\":1,\"GoodsInfo\":{\"ProductId\":\"example_product_id\",\"Quantity\":1,\"ZoneId\":\"example_zone_id\",\"OrigPrice\":1000,\"ActualPrice\":1000,\"Attach\":\"example_attach_data\",\"OrderSource\":1}}",
		"PayEventSig": "$your_pay_event_sig",
        "IsMock":true
	}
}

*/

type MidasRequest struct {
	ToUserName   string   `json:"ToUserName"`
	FromUserName string   `json:"FromUserName"`
	CreateTime   int64    `json:"CreateTime"`
	MsgType      string   `json:"MsgType"`
	Event        string   `json:"Event" binding:"required"`
	MiniGame     MiniGame `json:"MiniGame"`
}

type MiniGame struct {
	PayloadOrigin string `json:"Payload"  binding:"required"`
	Payload       Payload
	PayEventSig   string `json:"PayEventSig"  binding:"required"`
	IsMock        bool   `json:"IsMock"`
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

func NewMidasLogic(ctx context.Context) *MidasLogic {
	return &MidasLogic{ctx: ctx}
}

func (l *MidasLogic) ValidateSignature() error {
	key := viper.GetString("midasKey")
	hash := hmac.New(sha256.New, []byte(key))
	msg := l.MidasRequest.Event + "&" + l.MidasRequest.MiniGame.PayloadOrigin
	hash.Write([]byte(msg))
	signature := fmt.Sprintf("%x", hash.Sum(nil))

	fmt.Println(signature, l.MidasRequest.MiniGame.PayEventSig)
	if signature != l.MidasRequest.MiniGame.PayEventSig {
		return errors.New("签名不正确")
	}

	return nil
}

func (l *MidasLogic) ValidateRequest() error {
	if l.MidasRequest.MiniGame.Payload.OutTradeNo == "" {
		return errors.New("不存在的订单")
	}

	if l.MidasRequest.MiniGame.Payload.GoodsInfo.OrigPrice <= 0 {
		return errors.New("订单金额异常")
	}

	return nil
}

func (l *MidasLogic) GetOrderNum() string {
	return l.MidasRequest.MiniGame.Payload.OutTradeNo
}

func (l *MidasLogic) GetContext() context.Context {
	return l.ctx
}

type MidasResponse struct {
	ErrCode int    `json:"ErrCode"`
	ErrMsg  string `json:"ErrMsg"`
}

func (r *MidasResponse) Error(err error) *MidasResponse {
	r.ErrCode = -1
	r.ErrMsg = err.Error()
	return r
}

func (r *MidasResponse) Success() *MidasResponse {
	r.ErrCode = 0
	r.ErrMsg = "Success"
	return r
}
