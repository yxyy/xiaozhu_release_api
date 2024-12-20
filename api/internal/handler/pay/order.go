package pay

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/logic/common"
	"xiaozhu/internal/logic/pay"
)

func Order(c *gin.Context) {
	response := common.NewResponse(c)

	l := pay.NewOrderLogic(c.Request.Context())
	if err := c.ShouldBind(&l.Request); err != nil {
		response.Error(err)
		return
	}
	l.Request.Ip = c.ClientIP()

	data, err := l.Create()
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(data)
}
