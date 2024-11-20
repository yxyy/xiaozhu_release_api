package role

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/backend/internal/logic/system"
	"xiaozhu/backend/internal/model/common"
)

func Create(c *gin.Context) {

	logic := system.NewSysRoleLogic(c.Request.Context())
	response := common.NewResponse(c)

	if err := c.ShouldBind(&logic.SysRole); err != nil {
		response.Fail("参数格式错误，请检查")
		return
	}

	// group.OptUser = c.GetInt("userId")
	if err := logic.Create(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}

func Update(c *gin.Context) {

	logic := system.NewSysRoleLogic(c.Request.Context())
	response := common.NewResponse(c)

	if err := c.ShouldBind(&logic.SysRole); err != nil {
		response.Fail("参数格式错误，请检查")
		return
	}

	if err := logic.Update(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}

func List(c *gin.Context) {

	logic := system.NewSysRoleLogic(c.Request.Context())
	response := common.NewResponse(c)
	if err := c.ShouldBind(&logic.SysRole); err != nil {
		response.Error(err)
		return
	}

	list, err := logic.ListLogic()
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(list)
}

func ListAll(c *gin.Context) {

	logic := system.NewSysRoleLogic(c.Request.Context())
	response := common.NewResponse(c)
	list, err := logic.ListAllLogic()
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(list)
}

func UpdateMenu(c *gin.Context) {

	logic := system.NewSysRoleLogic(c.Request.Context())
	response := common.NewResponse(c)

	if err := c.ShouldBind(&logic.SysRole); err != nil {
		response.Fail("参数格式错误，请检查")
		return
	}

	if err := logic.UpdateMenu(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}
