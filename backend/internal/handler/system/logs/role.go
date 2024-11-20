package logs

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/backend/internal/logic/system"
	"xiaozhu/backend/internal/model/common"
)

func List(c *gin.Context) {

	logic := system.NewSysUserLogLogic(c.Request.Context())
	response := common.NewResponse(c)
	if err := c.ShouldBind(&logic); err != nil {
		response.Error(err)
		return
	}

	list, err := logic.ListLogic()
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(list)
}
