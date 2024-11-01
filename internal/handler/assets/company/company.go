package company

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/logic/assets"
	"xiaozhu/internal/model/common"
)

func List(c *gin.Context) {
	response := common.NewResponse(c)
	serviceCompany := assets.NewServiceCompany()
	params := common.NewParams()

	if err := c.ShouldBind(&serviceCompany); err != nil {
		response.Error(err)
	}

	if err := c.ShouldBind(&params); err != nil {
		response.Error(err)
	}

	sc, total, err := serviceCompany.List(params)
	if err != nil {
		response.Error(err)
	}

	data := make(map[string]interface{})
	data["rows"] = sc
	data["total"] = total

	response.SuccessData(data)
}

func Create(c *gin.Context) {
	response := common.NewResponse(c)
	serviceCompany := assets.NewServiceCompany()

	if err := c.ShouldBind(&serviceCompany); err != nil {
		response.Error(err)
	}

	if err := serviceCompany.Create(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Update(c *gin.Context) {
	response := common.NewResponse(c)
	serviceCompany := assets.NewServiceCompany()

	if err := c.ShouldBind(&serviceCompany); err != nil {
		response.Error(err)
	}

	if err := serviceCompany.Update(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Lists(c *gin.Context) {
	response := common.NewResponse(c)
	serviceCompany := assets.NewServiceCompany()

	if err := c.ShouldBind(&serviceCompany); err != nil {
		response.Error(err)
	}

	data, err := serviceCompany.Lists()
	if err != nil {
		response.Error(err)
	}

	response.SuccessData(data)
}
