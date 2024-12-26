package pay

import (
	"context"
	"errors"
	"xiaozhu/internal/model/key"
	"xiaozhu/internal/model/pay"
	"xiaozhu/utils/queue"
)

// TODO 客户端支付完成后，
//  服务端需要校验，
//  然后再通知cp发货
//  支付宝、微信、米大师直购是回调模型（第三方支付）
//  谷歌、苹果是收据验证模型(iOS/Google Play 订阅和内购）

// Notifier 回调模型
type Notifier interface {
	ValidateRequest() error
	ValidateSignature() error
	GetOrderNum() string
	GetContext() context.Context
}

// Invoicer 收据验证模型
type Invoicer interface {
	Validate() (*pay.Order, error)
	GetContext() context.Context
}

func Notify(notifier Notifier) error {
	if notifier == nil {
		return errors.New("无效订单")
	}

	if err := notifier.ValidateSignature(); err != nil {
		return err
	}

	if err := notifier.ValidateRequest(); err != nil {
		return err
	}

	return processOrder(notifier.GetContext(), notifier.GetOrderNum())
}

func Invoice(invoice Invoicer) error {
	if invoice == nil {
		return errors.New("无效订单")
	}

	order, err := invoice.Validate()
	if err != nil {
		return err
	}

	return queue.Push(invoice.GetContext(), key.OrderQueue, order)
}

func processOrder(ctx context.Context, orderNum string) error {
	order := &pay.Order{OrderNum: orderNum}
	if err := order.GetOrderInfo(ctx); err != nil {
		return err
	}

	return queue.Push(ctx, key.OrderQueue, order)
}
