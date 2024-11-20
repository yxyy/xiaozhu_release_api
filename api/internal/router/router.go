package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xiaozhu/api/internal/handler/user/auth"
	"xiaozhu/api/internal/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	// 设置静态资源
	r.StaticFS("/storage/data", http.Dir("./storage/data"))
	r.StaticFS("/storage/uploads", http.Dir("./storage/uploads"))

	r.Use(middleware.Log)

	r.POST("/system/v1/auth/login", auth.Login)

	r.POST("/system/v1/auth/refresh", auth.Refresh)

	r.Use(middleware.Auth)
	// r.Use(middleware.Auto)

	r.POST("/system/v1/auth/logout", auth.Logout)

	// 加载系统路由
	InitUserRouter(r)

	return r
}
