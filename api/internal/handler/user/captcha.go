package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/logic/common"
	logic "xiaozhu/internal/logic/user"
)

func Captcha(c *gin.Context) {

	l := logic.NewCaptchaLogic(c.Request.Context())
	response := common.NewResponse(c)
	if err := c.ShouldBind(&l); err != nil {
		response.Error(err)
		return
	}

	if err := l.Captcha(); err != nil {
		fmt.Printf("错误信息：%s", err)
		response.Error(err)
		return
	}

	response.Success()
}
