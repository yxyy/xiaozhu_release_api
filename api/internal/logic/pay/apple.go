package pay

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"xiaozhu/internal/logic/common"
	"xiaozhu/internal/model/pay"
)

//  App Store Server API https://developer.apple.com/documentation/appstoreserverapi/get-v1-transactions-_transactionid_
// 旧版（废弃） https://developer.apple.com/cn/documentation/storekit/in-app_purchase/validating_receipts_with_the_app_store/

type AppleLogic struct {
	ctx context.Context
	common.RequestForm
	OrderNum      string `json:"order_num"  binding:"required"`
	TransactionId string `json:"transactionId"  binding:"required"`
}

func NewAppleLogic(ctx context.Context) *AppleLogic {
	return &AppleLogic{ctx: ctx}
}

func (l *AppleLogic) Context() context.Context {
	return l.ctx
}

func (l *AppleLogic) Validate() (*pay.Order, error) {
	order := &pay.Order{OrderNum: l.OrderNum}
	if err := order.GetOrderInfo(l.ctx); err != nil {
		return nil, err
	}

	if order.GameId != l.GameId {
		return nil, fmt.Errorf("本游戏不存在该订单")
	}

	if order.OrderPrice <= 0 {
		return nil, fmt.Errorf("该订单金额异常")
	}

	transactionInfo, err := l.GetTransactionInfo(l.TransactionId)
	if err != nil {
		return nil, err
	}

	if transactionInfo.Price != int64(order.OrderPrice) || transactionInfo.Price <= 0 {
		return nil, fmt.Errorf("该订单金额异常")
	}

	order.TransactionId = l.TransactionId

	return order, nil
}

type AppStoreServerAPIResponse struct {
	SignedTransactionInfo string `json:"signedTransactionInfo"`
}

func (l *AppleLogic) GetTransactionInfo(transactionId string) (*common.JWSTransactionDecodedPayload, error) {
	token, err := common.GetAppleToken(l.ctx)
	if err != nil {
		return nil, err
	}
	urls := fmt.Sprintf("https://api.storekit.itunes.apple.com/inApps/v1/transactions/%s", transactionId)
	request, err := http.NewRequestWithContext(l.ctx, "GET", urls, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("apple validation failed with status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	appStoreServerAPIResponse := AppStoreServerAPIResponse{}
	if err = json.Unmarshal(body, &appStoreServerAPIResponse); err != nil {
		return nil, fmt.Errorf("反序列化失败: %w", err)
	}

	jwsTransactionDecodedPayload, err := common.ParseAppStoreServerAPIToken(appStoreServerAPIResponse.SignedTransactionInfo)
	if err != nil {
		return nil, err
	}

	return jwsTransactionDecodedPayload, nil
}
