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

}