package packages

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/logic/assets"
	"xiaozhu/internal/model/common"
)

func List(c *gin.Context) {
	response := common.NewResponse(c)
	l := assets.NewPackageLogic(c.Request.Context())

	if err := c.ShouldBind(&l); err != nil {
		response.Error(err)
		return
	}

	resp, err := l.List()
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(resp)
}

func Create(c *gin.Context) {
	response := common.NewResponse(c)
	l := assets.NewPackageLogic(c.Request.Context())

	if err := c.ShouldBind(&l); err != nil {
		response.Error(err)
		return
	}

	if err := l.Create(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}

func Update(c *gin.Context) {
	response := common.NewResponse(c)
	l := assets.NewPackageLogic(c.Request.Context())

	if err := c.ShouldBind(&l); err != nil {
		response.Error(err)
		return
	}

	if err := l.Update(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}
