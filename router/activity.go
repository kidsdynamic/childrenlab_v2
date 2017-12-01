package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/controller"
)

func initActivityRouter(r *gin.Engine) {
	v1 := r.Group("/v1/activity")
	v1.GET("/list/:kidId", controller.GetActivityList)

	activityAPI := r.Group("/v1/activity")
	activityAPI.Use(controller.Auth)
	activityAPI.POST("/uploadRawData", controller.UploadRawActivityData)
	activityAPI.GET("/retrieveData", controller.GetActivity)
	activityAPI.GET("/retrieveDataByTime", controller.GetActivityByTime)
	activityAPI.GET("/retrieveHourlyDataByTime", controller.GetTodayHourlyActivity)
}
