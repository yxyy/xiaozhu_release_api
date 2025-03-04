package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/handler/assets/app"
	"xiaozhu/internal/handler/assets/channel"
	"xiaozhu/internal/handler/assets/game"
	packages "xiaozhu/internal/handler/assets/package"
)

func InitAssetsRouter(r *gin.Engine) {

	assets := r.Group("/assets/v1")
	{
		// 渠道
		assets.POST("/channel/create", channel.Create)
		assets.POST("/channel/update", channel.Update)
		assets.POST("/channel/list", channel.List)
		assets.GET("/channel/list-all", channel.ListAll)

		// 应用
		assets.POST("/app/create", app.Create)
		assets.POST("/app/update", app.Update)
		assets.POST("/app/list", app.List)
		assets.GET("/app/list-all", app.ListAll)

		// 游戏
		assets.POST("/game/create", game.Create)
		assets.POST("/game/update", game.Update)
		assets.POST("/game/list", game.List)
		assets.GET("/game/list-all", game.ListAll)

		// 渠道包
		assets.POST("/package/create", packages.Create)
		assets.POST("/package/update", packages.Update)
		assets.POST("/package/list", packages.List)
		assets.GET("/package/list-all", packages.List)
	}
}
