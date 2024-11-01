package menu

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/logic/system"
	"xiaozhu/internal/model/common"
	system2 "xiaozhu/internal/model/system"
)

func Create(c *gin.Context) {
	response := common.NewResponse(c)
	logic := system.NewMenuLogic(c.Request.Context())

	if err := c.ShouldBind(&logic); err != nil {
		response.Error(err)
		return
	}

	if err := logic.Create(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}

func Update(c *gin.Context) {

	response := common.NewResponse(c)
	logic := system.NewMenuLogic(c.Request.Context())

	if err := c.ShouldBind(&logic); err != nil {
		response.Error(err)
		return
	}
	// servicesMenu.OptUser = c.GetInt("userId")
	if err := logic.Update(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}

func List(c *gin.Context) {

	response := common.NewResponse(c)
	logic := system.NewMenuLogic(c.Request.Context())
	var params = &system2.MenuListRequest{}
	if err := c.ShouldBind(&params); err != nil {
		response.Error(err)
		return
	}

	list, err := logic.List(params)
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(list)
}

func ListTree(c *gin.Context) {
	response := common.NewResponse(c)
	logic := system.NewMenuLogic(c.Request.Context())
	resp, err := logic.MenuTree()
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(resp)

}

func ListAll(c *gin.Context) {
	response := common.NewResponse(c)
	logic := system.NewMenuLogic(c.Request.Context())
	resp, err := logic.MenuAll()
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(resp)

}
