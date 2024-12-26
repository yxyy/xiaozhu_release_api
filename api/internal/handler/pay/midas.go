package pay

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/logic/common"
	"xiaozhu/internal/logic/pay"
)

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
