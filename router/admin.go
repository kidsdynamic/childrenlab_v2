package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/controller"
)

func initAdminRouter(r *gin.Engine) {
	v1 := r.Group("/v1/admin/")
	v1.GET("/kids/list", controller.GetAllKidList)
	v1.GET("/activity/raw/:macId", controller.GetActivityRaw)
}
