package dept

import (
	"github.com/gin-gonic/gin"
	"xiaozhu-api/internal/logic/common"
	logic "xiaozhu-api/internal/logic/user"
	"xiaozhu-api/internal/model/user"
)

func List(c *gin.Context) {
	response := common.NewResponse(c)

	request := user.SysDeptListRequest{}
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
