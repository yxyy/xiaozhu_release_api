package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/handler/assets/app"
	"xiaozhu/internal/handler/assets/app_type"
	"xiaozhu/internal/handler/assets/channel"
	"xiaozhu/internal/handler/assets/company"
	"xiaozhu/internal/handler/assets/game"
	"xiaozhu/internal/handler/common/images"
)

func InitCommonRouter(r *gin.Engine) {

	common := r.Group("/common")
	{
		// 研发公司下拉框
		common.GET("/company", company.Lists)
		// 应用类型下拉
		common.GET("/app_type", app_type.Lists)
		// 应用下拉
		common.GET("/app", app.Lists)
		// 游戏下拉
		common.GET("/game", game.Lists)
		// 渠道
		common.GET("/channel", channel.Lists)

	}

	uploads := r.Group("uploads")
	{
		uploads.POST("/images", images.Uploads)
	}
}
