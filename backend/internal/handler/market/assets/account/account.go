package account

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/logic/market/assets"
	"xiaozhu/internal/model/common"
)

func List(c *gin.Context) {
	response := common.NewResponse(c)
	l := assets.NewAccountLogic(c.Request.Context())
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

func ListAll(c *gin.Context) {
	response := common.NewResponse(c)
	l := assets.NewAccountLogic(c.Request.Context())

	list, err := l.ListAll()
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(list)
}

func Create(c *gin.Context) {
	response := common.NewResponse(c)
	l := assets.NewAccountLogic(c.Request.Context())

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

func Uploads(c *gin.Context) {
	response := common.NewResponse(c)
	l := assets.NewAccountLogic(c.Request.Context())

	if err := c.ShouldBind(&l); err != nil {
		response.Error(err)
		return
	}

	if err := l.BatchCreate(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}

func Update(c *gin.Context) {
	response := common.NewResponse(c)
	l := assets.NewAccountLogic(c.Request.Context())
	if err := c.ShouldBind(&l); err != nil {
		response.Error(err)
		return
	}

	if err := l.Update(); err != nil {
		response.Error(err)
	}

	response.Success()
}
