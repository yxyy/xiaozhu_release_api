package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"strings"
	"xiaozhu/internal/model/common"
	"xiaozhu/internal/model/system"
)

func Auth(c *gin.Context) {

	// for key, values := range c.Request.Header {
	// 	for _, value := range values {
	// 		fmt.Printf("%s: %s\n", key, value)
	// 	}
	// }

	accessToken := c.Request.Header.Get("Authorization")
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

	user, err := system.ParseToken(tokenT[1], 1)
	if err != nil {
		response.SetResult(403, "Access-Token is invalid", nil)
		c.Abort()
		return
	}

	c.Set("userId", user.Id)
	c.Set("RoleIds", user.RoleIds)
	c.Set("nickname", user.Nickname)
	c.Set("account", user.Account)

	withValue := context.WithValue(c.Request.Context(), "userId", user.Id)
	withValue = context.WithValue(withValue, "roleIds", user.RoleIds)
	withValue = context.WithValue(withValue, "nickname", user.Nickname)
	c.Request = c.Request.WithContext(withValue)

	c.Next()
}
