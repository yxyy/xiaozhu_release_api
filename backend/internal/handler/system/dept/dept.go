package dept

import (
	"github.com/gin-gonic/gin"
	logic "xiaozhu/backend/internal/logic/system"
	"xiaozhu/backend/internal/model/common"
	"xiaozhu/backend/internal/model/system"
)

func List(c *gin.Context) {
	response := common.NewResponse(c)

	request := system.SysDeptListRequest{}
	if err := c.ShouldBind(&request); err != nil {
		response.Error(err)
		return
	}

	l := logic.NewDeptLogic(c.Request.Context())
	list, err := l.List(request)
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(list)
}

func Create(c *gin.Context) {
	response := common.NewResponse(c)

	l := logic.NewDeptLogic(c.Request.Context())
	if err := c.ShouldBind(&l.SysDept); err != nil {
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

	l := logic.NewDeptLogic(c.Request.Context())
	if err := c.ShouldBind(&l.SysDept); err != nil {
		response.Error(err)
		return
	}

	if err := l.Update(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}
