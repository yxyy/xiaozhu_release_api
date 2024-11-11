package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/handler/market/assets/principal"
	"xiaozhu/internal/handler/market/assets/proxy_company"
	"xiaozhu/internal/handler/market/assets/proxy_project"
)

func InitMarketRouter(r *gin.Engine) {

	assets := r.Group("/market/assets")
	{
		// 开户主体
		assets.POST("/principal/create", principal.Create)
		assets.POST("/principal/update", principal.Update)
		assets.POST("/principal/list", principal.List)
		assets.GET("/principal/list-all", principal.ListAll)

		// 代理公司
		assets.POST("/proxy-company/create", proxy_company.Create)
		assets.POST("/proxy-company/update", proxy_company.Update)
		assets.POST("/proxy-company/list", proxy_company.List)
		assets.GET("/proxy-company/list-all", proxy_company.ListAll)

		// 代理项目
		assets.POST("/proxy-project/create", proxy_project.Create)
		assets.POST("/proxy-project/update", proxy_project.Update)
		assets.POST("/proxy-project/list", proxy_project.List)
		assets.GET("/proxy-project/list-all", proxy_project.ListAll)

	}
}
