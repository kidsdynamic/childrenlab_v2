package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/controller"
)

func initKidRouter(r *gin.Engine) {
	v1 := r.Group("/v1/kids")
	v1.GET("/list", controller.GetKidList)

	kidsAPI := r.Group("/v1/kids")
	kidsAPI.Use(controller.Auth)
	kidsAPI.POST("/add", controller.AddKid)
	kidsAPI.PUT("/update", controller.UpdateKid)
	kidsAPI.DELETE("/delete", controller.DeleteKid)
}
