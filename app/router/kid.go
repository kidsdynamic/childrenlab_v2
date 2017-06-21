package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/app/controller"
)

func initKidRouter(r *gin.Engine) {
	kidsAPI := r.Group("/v1/kids")
	kidsAPI.Use(controller.Auth)
	kidsAPI.POST("/add", controller.AddKid)
	kidsAPI.PUT("/update", controller.UpdateKid)
	kidsAPI.DELETE("/delete", controller.DeleteKid)
	kidsAPI.GET("/list", controller.GetKidList)
	kidsAPI.GET("/whoRegisteredMacID", controller.WhoRegisteredMacID)
	kidsAPI.POST("/batteryStatus", controller.UpdateBatteryStatus)
}
