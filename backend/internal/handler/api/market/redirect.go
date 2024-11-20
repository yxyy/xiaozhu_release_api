package market

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/logic/api"
	"xiaozhu/internal/model/common"
)

// BmOAuthRedirect 巨量授权回调
func BmOAuthRedirect(c *gin.Context) {
	response := common.NewResponse(c)

	l := api.NewBmLogic(c)

	if err := c.ShouldBind(&l); err != nil {
		response.Error(err)
		return
	}

	if err := l.Do(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}
