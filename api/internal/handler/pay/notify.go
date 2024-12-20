package pay

import (
	"github.com/gin-gonic/gin"
)

// TODO 客户端支付完成后，
//  服务端需要校验，
//  然后再通知cp发货
//  支付宝、微信、米大师直购是回调模型（第三方支付）
//  谷歌、苹果是收据验证模型(iOS/Google Play 订阅和内购）

func Midas(c *gin.Context) {

}
