package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xiaozhu/api/internal/handler/game"
	"xiaozhu/api/internal/handler/user"
	"xiaozhu/api/internal/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	// 设置静态资源
	r.StaticFS("/storage/data", http.Dir("./storage/data"))
	r.StaticFS("/storage/uploads", http.Dir("./storage/uploads"))

	r.Use(middleware.Log)

	// 游戏初始化
	r.POST("/v1/user/init", game.Init)

	r.POST("/v1/auth/login", user.Login)

	r.POST("/v1/auth/register", user.Register)

	r.POST("/v1/auth/captcha", user.Captcha)

	// r.POST("/v1/auth/logout", user.Logout)

	// 加载系统路由
	InitUserRouter(r)

	return r
}
