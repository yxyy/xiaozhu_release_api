package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/handler/api/market"
)

func InitApiRouter(r *gin.Engine) {

	api := r.Group("/api")
	{
		// 巨量授权回调
		api.POST("/bm/redirect", market.BmOAuthRedirect)

	}
}
