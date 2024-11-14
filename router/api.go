package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/handler/assets/channel"
)

func InitApiRouter(r *gin.Engine) {

	api := r.Group("/api")
	{
		// 渠道
		api.POST("/channel/create", channel.Create)

	}
}
