package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"strings"
	"xiaozhu/backend/internal/model/common"
	"xiaozhu/backend/utils"
)

func Menu(c *gin.Context) {

	groupName := c.GetString("groupName")
	// 草鸡管理员直接放行
	if groupName == "supper" {
		c.Next()
		return
	}
	response := common.NewResponse(c)
	path := c.FullPath()
	index := strings.LastIndex(path, "/")
	if index <= 0 {
		response.SetResult(4004, "无效的请求路径", nil)
		c.Abort()
		return
	}

	result, err := utils.RedisClient.HGet(context.Background(), "menu_router:"+groupName, path[:index]).Result()
	if err != nil || result != "1" {
		response.SetResult(4003, "没有权限", nil)
		c.Abort()
		return
	}

	c.Next()
}
