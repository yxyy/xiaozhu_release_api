package game

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/api/internal/logic/common"
	logic "xiaozhu/api/internal/logic/game"
)

func Report(c *gin.Context) {
	response := common.NewResponse(c)

	l := logic.NewReportLogic(c.Request.Context())
	if err := c.ShouldBind(&l.ReportRequest); err != nil {
		response.Error(err)
		return
	}

	if err := l.Report(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}
