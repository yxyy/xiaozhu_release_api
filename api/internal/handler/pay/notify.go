package pay

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/logic/common"
	"xiaozhu/internal/logic/pay"
)

// TODO 客户端支付完成后，
//  服务端需要校验，
//  然后再通知cp发货
//  支付宝、微信、米大师直购是回调模型（第三方支付）
//  谷歌、苹果是收据验证模型(iOS/Google Play 订阅和内购）

func Midas(c *gin.Context) {
	r := new(pay.MidasResponse)
	response := common.NewResponse(c)
	l := pay.NewMidasLogic(c.Request.Context())
	if err := c.ShouldBind(&l.MidasRequest); err != nil {
		response.OriginError(r.Error(err))
		return
	}

	if err := pay.Notify(l); err != nil {
		response.SetOriginServerError(r.Error(err))
		return
	}

	response.OriginSuccess(r.Success())

}
