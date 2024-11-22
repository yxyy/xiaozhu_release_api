package user

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/api/internal/logic/common"
	logic "xiaozhu/api/internal/logic/user"
)

func Init(c *gin.Context) {
	response := common.NewResponse(c)

	l := logic.NewInitLogic(c.Request.Context())
	if err := c.ShouldBind(&l.InitRequest); err != nil {
		response.Error(err)
		return
	}

	data, err := l.Init()

	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(data)
}
