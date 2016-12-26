package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/controller"
)

func initActivityRouter(r *gin.Engine) {
	activityAPI := r.Group("/v1/activity")
	activityAPI.Use(controller.Auth)
	activityAPI.POST("/uploadRawData", controller.UploadRawActivityData)
	activityAPI.GET("/retrieveData", controller.GetActivity)
}
