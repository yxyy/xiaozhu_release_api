package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/config/mysql"
	"xiaozhu/internal/model/assets"
)

func Auto(c *gin.Context) {
	if err := mysql.PlatformDB.AutoMigrate(
		// &system.User{},
		// &system.SysRole{},
		// &system.SysMenus{},
		&assets.Package{},
	); err != nil {
		fmt.Println(err)
	}
	c.Next()
}
