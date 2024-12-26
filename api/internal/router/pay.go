package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/handler/pay"
	"xiaozhu/internal/middleware"
)

func InitPayRouter(r *gin.Engine) {

	router := r.Group("v1/pay").Use(middleware.Auth)
	{
		// 下单
		router.POST("/order", pay.Order)
	}

	notify := r.Group("v1/notify")
	{
		// 米大师直购
		notify.POST("/midas", pay.Midas)

		// 苹果内购
		notify.POST("/apple", pay.Apple)

	}

}
