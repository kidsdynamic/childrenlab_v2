package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/controller"
)

func initAdminRouter(r *gin.Engine) {
	adminAuthAPI := r.Group("/admin")
	adminAuthAPI.Use(controller.AdminAuth)
	r.POST("/admin/login", controller.AdminLogin)
	adminAuthAPI.GET("/userList", controller.GetAllUser)
	adminAuthAPI.GET("/kidList", controller.GetAllKidList)
	adminAuthAPI.GET("/activityList/:kidId", controller.GetActivityListForAdmin)
	adminAuthAPI.GET("/activityRawList/:macId", controller.GetActivityRawForAdmin)
}
