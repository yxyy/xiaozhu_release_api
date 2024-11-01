package packages

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/logic/assets"
	"xiaozhu/internal/model/common"
)

func List(c *gin.Context) {
	response := common.NewResponse(c)
	servicePackage := assets.NewServicePackage()

	if err := c.ShouldBind(&servicePackage); err != nil {
		response.Error(err)
	}

	if err := servicePackage.Create(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Create(c *gin.Context) {
	response := common.NewResponse(c)
	servicePackage := assets.NewServicePackage()

	if err := c.ShouldBind(&servicePackage); err != nil {
		response.Error(err)
	}

	servicePackage.OptUser = c.GetInt("userId")
	if err := servicePackage.Create(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Update(c *gin.Context) {
	response := common.NewResponse(c)
	servicePackage := assets.NewServicePackage()

	if err := c.ShouldBind(&servicePackage); err != nil {
		response.Error(err)
	}

	servicePackage.OptUser = c.GetInt("userId")
	if err := servicePackage.Update(); err != nil {
		response.Error(err)
	}

	response.Success()
}
