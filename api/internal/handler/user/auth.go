package user

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/api/internal/logic/common"
	logic "xiaozhu/api/internal/logic/user/auth"
)

func Login(c *gin.Context) {

	response := common.NewResponse(c)
	l := logic.NewAuthLogic(c.Request.Context())

	if err := c.ShouldBind(&l); err != nil {
		response.Error(err)
		return
	}

	// todo 黑名单校验
	auther, err := logic.NewAuther(l)
	if err != nil {
		response.Error(err)
		return
	}

	// 账号信息
	resp, err := l.Login(auther)
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(resp)
}
