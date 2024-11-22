package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/api/internal/handler/user"
	"xiaozhu/api/internal/middleware"
)

func InitUserRouter(r *gin.Engine) {

	userRouter := r.Group("v1/user").Use(middleware.Auth)
	{
		// 忘记密码
		userRouter.POST("/forget", user.Login)

	}

}
