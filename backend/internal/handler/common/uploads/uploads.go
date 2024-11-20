package uploads

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
	logic "xiaozhu/backend/internal/logic/conmon"
	"xiaozhu/backend/internal/model/common"
	"xiaozhu/backend/utils"
)

func Uploads(c *gin.Context) {

	response := common.NewResponse(c)
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(err)
		return
	}

	l := logic.NewFileLogic(c.Request.Context(), file)
	finalFilePath, err := logic.SaveLogic(l)
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(finalFilePath)

}

func Image(c *gin.Context) {

	response := common.NewResponse(c)
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(err)
	}

	format := time.Now().Format("20060102")
	path := viper.GetString("oss.images") + "/" + format
	if err = utils.TidyDirectory("./" + path); err != nil {
		response.Error(err)
	}
	filepath := path + "/" + utils.Salt() + ".jpg"
	if err = c.SaveUploadedFile(file, "./"+filepath); err != nil {
		response.Error(err)
	}

	response.SuccessData(filepath)

}
