package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"xiaozhu/api/internal/logic/common"
	"xiaozhu/api/internal/model/user"
)

func Auth(c *gin.Context) {
	accessToken := c.Request.Header.Get("Authorization")
	fmt.Println("----------------------666666")
	response := common.NewResponse(c)
	if accessToken == "" {
		response.SetResult(403, "Access-Token is empty", nil)
		c.Abort()
		return
	}

	tokenT := strings.Split(accessToken, "Bearer ")
	if len(tokenT) != 2 {
		response.SetResult(403, "Access-Token is invalid", nil)
		c.Abort()
		return
	}

	sysUser, err := user.ParseToken(tokenT[1], 1)
	if err != nil {
		response.SetResult(403, "Access-Token is invalid", nil)
		c.Abort()
		return
	}

	c.Set("userId", sysUser.Id)
	c.Set("RoleIds", sysUser.RoleIds)
	c.Set("nickname", sysUser.Nickname)
	c.Set("account", sysUser.Account)

	withValue := context.WithValue(c.Request.Context(), "userId", sysUser.Id)
	withValue = context.WithValue(withValue, "roleIds", sysUser.RoleIds)
	withValue = context.WithValue(withValue, "nickname", sysUser.Nickname)
	c.Request = c.Request.WithContext(withValue)

	c.Next()
}
