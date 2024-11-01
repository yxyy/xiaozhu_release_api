package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/handler/assets/app"
	"xiaozhu/internal/handler/assets/app_type"
	"xiaozhu/internal/handler/assets/channel"
	"xiaozhu/internal/handler/assets/company"
	"xiaozhu/internal/handler/assets/game"
	packages "xiaozhu/internal/handler/assets/package"
)

func InitAssetsRouter(r *gin.Engine) {

	assets := r.Group("/assets/v1")
	{
		// 研发公司
		assets.POST("/company/create", company.Create)
		assets.POST("/company/update", company.Update)
		assets.GET("/company/list", company.List)

		// 应用类型
		assets.POST("/app_type/create", app_type.Create)
		assets.POST("/app_type/update", app_type.Update)
		assets.GET("/app_type/list", app_type.List)

		// 应用
		assets.POST("/app/create", app.Create)
		assets.POST("/app/update", app.Update)
		assets.GET("/app/list", app.List)

		// 游戏
		assets.POST("/game/create", game.Create)
		assets.POST("/game/update", game.Update)
		assets.GET("/game/list", game.List)

		// 游戏
		assets.POST("/channel/create", channel.Create)
		assets.POST("/channel/update", channel.Update)
		assets.GET("/channel/list", channel.List)

		// 渠道包
		assets.POST("/package/create", packages.Create)
		assets.POST("/package/update", packages.Update)
		assets.GET("/package/list", packages.List)
	}
}
