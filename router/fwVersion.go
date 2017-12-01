package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/controller"
)

func initFWRouter(r *gin.Engine) {
	fwAPI := r.Group("/v1/fw")
	fwAPI.Use(controller.Auth)
	fwAPI.GET("/currentVersion/:macId", controller.GetCurrentFWVersionAndLink)
	fwAPI.PUT("/firmwareVersion", controller.UpdateDeviceFWVersion)
}
