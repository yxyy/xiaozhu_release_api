package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/model/system"
	"xiaozhu/utils"
)

func Auto(c *gin.Context) {
	if err := utils.MysqlDb.AutoMigrate(
		&system.User{},
		&system.SysRole{},
		&system.SysMenus{},
	); err != nil {
		fmt.Println(err)
	}
	c.Next()
}
