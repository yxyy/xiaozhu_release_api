package logs

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/logic/system"
	"xiaozhu/internal/model/common"
)

func List(c *gin.Context) {

	logic := system.NewSysUserLogLogic(c.Request.Context())
	response := common.NewResponse(c)
	if err := c.ShouldBind(&logic.SysUserLog); err != nil {
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
