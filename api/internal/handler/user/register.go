package user

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/logic/common"
	logic "xiaozhu/internal/logic/user/auth"
)

func Register(c *gin.Context) {

	response := common.NewResponse(c)
	l := logic.NewAuthLogic(c.Request.Context())

	if err := c.ShouldBind(&l); err != nil {
		response.Error(err)
		return
	}

	l.Ip = c.ClientIP()

	// todo 黑名单校验
	auther, err := logic.NewAuther(l)
	if err != nil {
		response.Error(err)
		return
	}

	// 账号信息
	resp, err := l.Register(auther)
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(resp)
}
