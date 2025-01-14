package game

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/logic/common"
	logic "xiaozhu/internal/logic/game"
)

func Init(c *gin.Context) {
	response := common.NewResponse(c)

	l := logic.NewInitLogic(c.Request.Context())
	if err := c.ShouldBind(&l.InitRequest); err != nil {
		response.Error(err)
		return
	}

	l.InitRequest.Ip = c.ClientIP()
	data, err := l.Init()

	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(data)
}
