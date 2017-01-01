package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/controller"
)

func initSubHostRouter(r *gin.Engine) {
	v1 := r.Group("/v1/subHost")
	v1.Use(controller.Auth)
	v1.POST("/add", controller.RequestSubHost)
	v1.POST("/accept", controller.AcceptRequest)
	v1.POST("/deny", controller.DenyRequest)
	v1.GET("/list", controller.SubHostList)
}