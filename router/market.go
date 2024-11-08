package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/handler/market/assets/principal"
)

func InitMarketRouter(r *gin.Engine) {

	assets := r.Group("/assets/v1")
	{
		// 研发公司
		assets.POST("/principal/create", principal.Create)
		assets.POST("/principal/update", principal.Update)
		assets.POST("/principal/list", principal.List)

	}
}
