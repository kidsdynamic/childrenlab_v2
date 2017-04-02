package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/controller"
)

func initSubHostRouter(r *gin.Engine) {
	v1 := r.Group("/v1/subHost")
	v1.Use(controller.Auth)
	v1.POST("/add", controller.RequestSubHostToUser)
	v1.PUT("/accept", controller.AcceptRequest)
	v1.PUT("/deny", controller.DenyRequest)
	v1.GET("/list", controller.SubHostList)
	v1.PUT("/removeKid", controller.RemoveSubHostKid)
	v1.DELETE("/delete", controller.DeleteRequest)
	//v1.GET("/permission", controller.HasPermission)
}
