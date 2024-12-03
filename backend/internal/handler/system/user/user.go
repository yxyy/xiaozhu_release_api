package user

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/logic/system"
	"xiaozhu/internal/model/common"
)

func List(c *gin.Context) {

	logic := system.NewUserLogic(c.Request.Context())
	response := common.NewResponse(c)

	if err := c.ShouldBind(&logic.SysUser); err != nil {
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

func Info(c *gin.Context) {
	user := system.NewUserLogic(c.Request.Context())
	response := common.NewResponse(c)
	// user.Id = c.GetInt("userId")

	data, err := user.UserInfos()
	if err != nil {
		response.Error(err)
	}

	response.SuccessData(data)
}

func ListAll(c *gin.Context) {
	response := common.NewResponse(c)
	l := system.NewUserLogic(c.Request.Context())

	list, err := l.ListAll()
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(list)
}

func Create(c *gin.Context) {
	logic := system.NewUserLogic(c.Request.Context())
	response := common.NewResponse(c)
	if err := c.ShouldBind(&logic.SysUser); err != nil {
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
	logic := system.NewUserLogic(c.Request.Context())
	response := common.NewResponse(c)
	if err := c.ShouldBind(&logic.SysUser); err != nil {
		response.Error(err)
		return
	}
	if err := logic.Update(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}

func SaveRole(c *gin.Context) {
	logic := system.NewUserLogic(c.Request.Context())
	response := common.NewResponse(c)
	if err := c.ShouldBind(&logic.SysUser); err != nil {
		response.Error(err)
		return
	}
	if err := logic.SaveRole(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}

func Remove(c *gin.Context) {
	user := system.NewUserLogic(c.Request.Context())
	response := common.NewResponse(c)
	if err := c.ShouldBind(user); err != nil {
		response.Error(err)
		return
	}
	if err := user.Remove(); err != nil {
		response.Error(err)
		return
	}
	response.Success()
}
