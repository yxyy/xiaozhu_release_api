package user

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/api/internal/logic/common"
	logic "xiaozhu/api/internal/logic/user/auth"
)

func Login(c *gin.Context) {

	l := logic.NewLoginLogic(c.Request.Context())
	response := common.NewResponse(c)
	if err := c.ShouldBind(&l); err != nil {
		response.Error(err)
		return
	}

	// todo 黑名单校验
	var login logic.Loginer
	switch true {
	case l.Mobile.Phone != 0:
		// TODO 手机登录
		login = l.Mobile
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
	resp, err := l.Login(login)
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(resp)
}

// func Logout(c *gin.Context) {
//
// 	response := common.NewResponse(c)
// 	request := user.NewRefreshRequest()
// 	if err := c.ShouldBind(&request); err != nil {
// 		response.Error(err)
// 	}
//
// 	if request.RefreshToken == "" {
// 		response.SetResult(403, "Access-Token is empty365656", nil)
// 		return
// 	}
//
// 	_, err := user.ParseToken(request.RefreshToken, 2)
// 	if err != nil {
// 		response.SetResult(403, "Access-Token is invalid", nil)
// 		c.Abort()
// 		return
// 	}
//
// 	response.Success()
// }
