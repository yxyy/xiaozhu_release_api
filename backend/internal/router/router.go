package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xiaozhu/backend/internal/handler/system/auth"
	"xiaozhu/backend/internal/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	// 设置静态资源
	r.StaticFS("/storage/data", http.Dir("./storage/data"))
	r.StaticFS("/storage/uploads", http.Dir("./storage/uploads"))

	r.Use(middleware.Log)
	// r.Use(middleware.Auto)

	// 注册api路由
	InitApiRouter(r)

	r.POST("/system/v1/auth/login", auth.Login)

	r.POST("/system/v1/auth/refresh", auth.Refresh)

	r.Use(middleware.Auth)
	// r.Use(middleware.Auto)

	r.POST("/system/v1/auth/logout", auth.Logout)

	r.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "6666",
		})
	})

	// 加载系统路由
	InitSystemRouter(r)

	// 公共路由
	InitCommonRouter(r)

	// 加载资产路由
	InitAssetsRouter(r)

	// 加载市场路由
	InitMarketRouter(r)

	return r
}
