package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/model/assets"
	"xiaozhu/utils"
)

func Auto(c *gin.Context) {
	if err := utils.MysqlDb.AutoMigrate(
		// &system.User{},
		// &system.SysRole{},
		// &system.SysMenus{},
		&assets.Package{},
	); err != nil {
		fmt.Println(err)
	}
	c.Next()
}
