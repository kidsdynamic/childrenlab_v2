package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/controller"
)

func initDeviceRouter(r *gin.Engine) {
	deviceAPI := r.Group("/v1/device")
	deviceAPI.GET("/whoRegisteredMacID", controller.WhoRegisteredMacID)

}
