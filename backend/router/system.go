package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/handler/system/dept"
	"xiaozhu/internal/handler/system/logs"
	"xiaozhu/internal/handler/system/menu"
	"xiaozhu/internal/handler/system/role"
	"xiaozhu/internal/handler/system/user"
)

func InitSystemRouter(r *gin.Engine) {

	system := r.Group("system/v1")
	{
		// 菜单
		system.POST("/menu/create", menu.Create)
		system.POST("/menu/update", menu.Update)
		system.POST("/menu/list", menu.List)
		system.GET("/menu/list-tree", menu.ListTree)
		system.POST("/menu/list-all", menu.ListAll)

		// 部门
		system.POST("/dept/list", dept.List)
		system.POST("/dept/create", dept.Create)
		system.POST("/dept/update", dept.Update)

		// 系统账号
		system.POST("/user/list", user.List)
		system.GET("/user/userInfo", user.Info)
		system.POST("/user/create", user.Create)
		system.POST("/user/update", user.Update)
		system.POST("/user/save-role", user.SaveRole)
		system.POST("/user/remove", user.Remove)
		system.GET("/user/list-all", user.ListAll)

		// 系统角色
		system.POST("/role/create", role.Create)
		system.POST("/role/update", role.Update)
		system.POST("/role/list", role.List)
		system.GET("/role/list-all", role.ListAll)
		system.POST("/role/save-menu", role.UpdateMenu)

		// 操作日志
		system.POST("/operation-log/list", logs.List)

	}

}
