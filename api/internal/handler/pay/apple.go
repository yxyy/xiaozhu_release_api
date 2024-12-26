package pay

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/logic/common"
	"xiaozhu/internal/logic/pay"
)

func Apple(c *gin.Context) {
	response := common.NewResponse(c)
	l := pay.NewAppleLogic(c.Request.Context())
	if err := c.ShouldBind(&l); err != nil {
		response.Error(err)
		return
	}

	if err := pay.Invoice(l); err != nil {
		response.Error(err)
		return
	}

	response.Success()

}
