package pay

import (
	"context"
	"fmt"
	"google.golang.org/api/androidpublisher/v3"
	"google.golang.org/api/option"
	"xiaozhu/internal/config"
	"xiaozhu/internal/logic/common"
	"xiaozhu/internal/model/pay"
)

// https://developers.google.com/android-publisher/api-ref/rest/v3/purchases.products/get?hl=zh-cn

type GoogleLogic struct {
	ctx context.Context
	common.RequestForm
	OrderNum string `json:"order_num"  binding:"required"`
	Token    string `json:"token"  binding:"required"`
}

type ProductPurchase struct {
	Kind                        string `json:"kind"`
	PurchaseTimeMillis          int64  `json:"purchaseTimeMillis"`
	PurchaseState               int    `json:"purchaseState"`
	ConsumptionState            int    `json:"consumptionState"`
	DeveloperPayload            string `json:"developerPayload"`
	OrderId                     string `json:"orderId"`
	PurchaseType                int    `json:"purchaseType"`
	AcknowledgementState        int    `json:"acknowledgementState"`
	PurchaseToken               string `json:"purchaseToken"`
	ProductId                   string `json:"productId"`
	Quantity                    int    `json:"quantity"`
	ObfuscatedExternalAccountId string `json:"obfuscatedExternalAccountId"`
	ObfuscatedExternalProfileId string `json:"obfuscatedExternalProfileId"`
	RegionCode                  string `json:"regionCode"`
	RefundableQuantity          int    `json:"refundableQuantity"`
}

func NewGoogleLogic(ctx context.Context) *GoogleLogic {
	return &GoogleLogic{ctx: ctx}
}

func (l *GoogleLogic) Context() context.Context {
	return l.ctx
}

func (l *GoogleLogic) Validate() (*pay.Order, error) {
	// order := &pay.Order{OrderNum: l.OrderNum}
	// if err := order.GetOrderInfo(l.ctx); err != nil {
	// 	return nil, err
	// }
	//
	// if order.GameId != l.GameId {
	// 	return nil, fmt.Errorf("本游戏不存在该订单")
	// }
	//
	// if order.OrderPrice <= 0 {
	// 	return nil, fmt.Errorf("该订单金额异常")
	// }

	// todo 谷歌校验
	// option.WithCredentialsJSON()
	service, err := androidpublisher.NewService(l.ctx, option.WithCredentialsFile(config.RootDir+"/etc/xxx.json"))
	if err != nil {
		return nil, err
	}

	get := service.Purchases.Products.Get("666", "7777", "8888")
	do, err := get.Do()
	fmt.Println(do, err)

	return nil, nil
}
