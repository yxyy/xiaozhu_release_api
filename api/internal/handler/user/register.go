package user

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/api/internal/logic/common"
	logic "xiaozhu/api/internal/logic/user/auth"
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
	var login logic.Auther
	switch true {
	// case l.Mobile.Phone != 0:
	// 	// TODO 手机登录
	// 	login = l.Mobile
	case l.Email.Email != "":
		// 邮箱登录
		login = l.Email
	case l.WeChat.WxCode != "":
		// 微信登录
		login = l.WeChat
	case l.Account.Account != "":
		// 账号登录
		login = l.Account
	default:
		response.Fail("无效的登录方式")
		return
	}

	// 账号信息
	resp, err := l.Register(login)
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(resp)
}
