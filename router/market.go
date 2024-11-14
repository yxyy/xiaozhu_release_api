package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/handler/market/assets/account"
	"xiaozhu/internal/handler/market/assets/appid"
	"xiaozhu/internal/handler/market/assets/principal"
	"xiaozhu/internal/handler/market/assets/project"
	"xiaozhu/internal/handler/market/assets/proxy_company"
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

		// 推广项目
		assets.POST("/proxy-project/create", project.Create)
		assets.POST("/proxy-project/update", project.Update)
		assets.POST("/proxy-project/list", project.List)
		assets.GET("/proxy-project/list-all", project.ListAll)

		// 广告账号
		assets.POST("/account/create", account.Create)
		assets.POST("/account/uploads", account.Uploads)
		assets.POST("/account/update", account.Update)
		assets.POST("/account/list", account.List)
		assets.GET("/account/list-all", account.ListAll)

		// 授权应用
		assets.POST("/appid/create", appid.Create)
		assets.POST("/appid/update", appid.Update)
		assets.POST("/appid/list", appid.List)

	}
}
