package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu-api/internal/handler/user/dept"
)

func InitUserRouter(r *gin.Engine) {

	system := r.Group("v1/user")
	{
		// 部门
		system.POST("/dept/list", dept.List)
		system.POST("/dept/create", dept.Create)
		system.POST("/dept/update", dept.Update)

	}

}
