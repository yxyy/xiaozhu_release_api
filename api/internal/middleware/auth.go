package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"xiaozhu/api/internal/logic/common"
	"xiaozhu/api/internal/model/user"
	"xiaozhu/api/utils"
)

func Auth(c *gin.Context) {
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

	result, err := utils.RedisClient.Get(c.Request.Context(), tokenT[1]).Result()
	if err != nil {
		response.SetServerError(fmt.Sprintf("服务器开小差了：%s", err))
		c.Abort()
		return
	}

	memberInfo := user.NewMemberInfo()
	if err = json.Unmarshal([]byte(result), &memberInfo); err != nil {
		response.SetServerError(fmt.Sprintf("服务器开小差了：%s", err))
		c.Abort()
		return
	}

	// c.Set("userId", memberInfo.Id)

	withValue := context.WithValue(c.Request.Context(), "userId", memberInfo.Id)
	withValue = context.WithValue(withValue, "account", memberInfo.Account)
	withValue = context.WithValue(withValue, "nickname", memberInfo.Nickname)
	c.Request = c.Request.WithContext(withValue)

	c.Next()
}
