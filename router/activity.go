package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/controller"
)

func initActivityRouter(r *gin.Engine) {
	avatarAPI := r.Group("/v1/activity")
	avatarAPI.Use(controller.Auth)
	avatarAPI.POST("/uploadRawData", controller.UploadRawActivityData)
}
