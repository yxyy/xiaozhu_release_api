package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/handler/common/uploads"
)

func InitCommonRouter(r *gin.Engine) {

	file := r.Group("common/uploads")
	{
		file.POST("/", uploads.Uploads)
		file.POST("/images", uploads.Image)
	}
}
