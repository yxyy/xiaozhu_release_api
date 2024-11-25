package user

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/api/internal/logic/common"
	logic "xiaozhu/api/internal/logic/user"
)

func Info(c *gin.Context) {

	l := logic.NewMemberLogic(c.Request.Context())
	response := common.NewResponse(c)
	if err := c.ShouldBind(&l); err != nil {
		response.Error(err)
		return
	}

	info, err := l.Info()
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(info)
}
