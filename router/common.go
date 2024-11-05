package router

import (
	"github.com/gin-gonic/gin"
	"xiaozhu/internal/handler/common/images"
)

func InitCommonRouter(r *gin.Engine) {

	uploads := r.Group("uploads")
	{
		uploads.POST("/images", images.Uploads)
	}
}
