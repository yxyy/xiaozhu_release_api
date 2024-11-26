package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/api/internal/handler/game"
	"xiaozhu/api/internal/middleware"
)

func InitGameRouter(r *gin.Engine) {

	// 游戏初始化
	r.POST("/v1/game/init", game.Init)

	gameRouter := r.Group("v1/game").Use(middleware.Auth)
	{
		// 角色上报
		gameRouter.POST("/report", game.Report)

	}

}
