package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xiaozhu/internal/handler/system/auth"
	"xiaozhu/internal/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	// 设置静态资源
	r.StaticFS("/storage/uploads", http.Dir("./storage/uploads"))

	r.Use(middleware.Log)
	// r.Use(middleware.Auto)

	r.POST("/system/v1/auth/login", auth.Login)

	r.POST("/system/v1/auth/refresh", auth.Refresh)

	r.Use(middleware.Auth)
	// r.Use(middleware.SysMenu)

	r.POST("/system/v1/auth/logout", auth.Logout)

	r.GET("/home", func(c *gin.Context) {

		fmt.Println(666666)
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "6666",
		})
	})

	// 加载系统路由
	InitSystemRouter(r)

	// 公共路由
	InitCommonRouter(r)

	// 加载运营路由
	InitAssetsRouter(r)

	return r
}
