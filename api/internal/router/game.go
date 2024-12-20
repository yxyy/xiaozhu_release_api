package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/handler/game"
	"xiaozhu/internal/middleware"
)

func InitGameRouter(r *gin.Engine) {

	// 游戏初始化
	r.POST("/v1/game/init", game.Init)

	router := r.Group("v1/game").Use(middleware.Auth)
	{
		// 角色上报
		router.POST("/report", game.Report)

	}

}
