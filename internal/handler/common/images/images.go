package images

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
	"xiaozhu/internal/model/common"
	"xiaozhu/utils"
)

func Uploads(c *gin.Context) {

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
