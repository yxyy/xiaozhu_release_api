package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/api/internal/handler/user"
	"xiaozhu/api/internal/middleware"
)

func InitUserRouter(r *gin.Engine) {

	r.POST("/v1/auth/login", user.Login)

	r.POST("/v1/auth/register", user.Register)

	r.POST("/v1/auth/captcha", user.Captcha)

	// r.POST("/v1/auth/logout", user.Logout)

	userRouter := r.Group("v1/user").Use(middleware.Auth)
	{
		// 我的信息
		userRouter.POST("/info", user.Info)

		// 忘记密码
		userRouter.POST("/forget", user.Login)

	}

}
