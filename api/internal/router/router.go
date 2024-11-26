package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xiaozhu/api/internal/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	// 设置静态资源
	r.StaticFS("/storage/data", http.Dir("./storage/data"))
	r.StaticFS("/storage/uploads", http.Dir("./storage/uploads"))

	r.Use(middleware.Log)

	// 加载用户路由
	InitUserRouter(r)

	// 加载游戏路由
	InitGameRouter(r)

	return r
}
