package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
	"xiaozhu-api/internal/logic/common"
	system2 "xiaozhu-api/internal/logic/user/auth"
	"xiaozhu-api/internal/model/user"
)

func Login(c *gin.Context) {

	params := system2.NewParams()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&params); err != nil {
		response.Error(err)
		return
	}

	var login system2.Loginer
	switch true {
	case params.Mobile != nil:
		// TODO 手机登录
		login = params.Mobile
	case params.Account != nil:
		// 账号登录
		login = params.Account
	case params.WeChat != nil:
		// 微信登录
		login = params.WeChat
	default:
		response.Fail("无效的登录方式")
	}

	// 账号信息
	userInfo, err := system2.Login(login)
	if err != nil {
		response.Error(err)
		return
	}

	// token信息
	token, err := user.GetAccessToken(userInfo)
	if err != nil {
		response.Error(err)
		return
	}

	refreshToken, err := user.GetRefreshToken(userInfo)
	if err != nil {
		response.Error(err)
		return
	}

	resp := user.NewLoginResponse()
	resp.AccessToken = token
	resp.RefreshToken = refreshToken
	resp.AccessExpire = viper.GetInt64("Auth.AccessExpire")
	resp.RefreshExpire = time.Now().Add(time.Second * time.Duration(viper.GetInt64("Auth.RefreshExpire"))).UnixMilli()
	resp.Role = []string{"admin"}
	resp.Username = "admin"

	userInfo.LastIp = c.ClientIP()
	userInfo.LastTime = time.Now().Unix()
	if err = userInfo.Update(c.Request.Context()); err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(resp)
}

func Logout(c *gin.Context) {

	response := common.NewResponse(c)
	request := user.NewRefreshRequest()
	if err := c.ShouldBind(&request); err != nil {
		response.Error(err)
	}

	if request.RefreshToken == "" {
		response.SetResult(403, "Access-Token is empty365656", nil)
		return
	}

	_, err := user.ParseToken(request.RefreshToken, 2)
	if err != nil {
		response.SetResult(403, "Access-Token is invalid", nil)
		c.Abort()
		return
	}

	response.Success()
}

func Refresh(c *gin.Context) {

	response := common.NewResponse(c)
	request := user.NewRefreshRequest()
	if err := c.ShouldBind(&request); err != nil {
		response.Error(err)
	}

	if request.RefreshToken == "" {
		response.SetResult(401, "Access-Token is empty365656", nil)
		return
	}

	userInfo, err := user.ParseToken(request.RefreshToken, 2)
	if err != nil {
		response.SetCodeError(401, err.Error())
		return
	}

	token, err := user.GetAccessToken(userInfo)
	if err != nil {
		response.Error(err)
		return
	}

	resp := user.NewLoginResponse()
	resp.AccessToken = token
	resp.AccessExpire = viper.GetInt64("Auth.AccessExpire")
	resp.RefreshToken = request.RefreshToken

	response.SuccessData(resp)
}

func RoleMenu(c *gin.Context) {

	response := common.NewResponse(c)
	menu, err := system2.RoleMenu()
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(menu)
}
