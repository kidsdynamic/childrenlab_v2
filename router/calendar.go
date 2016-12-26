package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/controller"
)

func initCalendarRouter(r *gin.Engine) {
	calendarAPI := r.Group("/v1/calendar")
	calendarAPI.Use(controller.Auth)
	calendarAPI.POST("/add", controller.AddCalendarEvent)
	calendarAPI.PUT("/edit", controller.UploadKidAvatar)
	calendarAPI.DELETE("/delete", nil)

	todoAPI := r.Group("/v1/calendar/todo")
	todoAPI.Use(controller.Auth)
	todoAPI.POST("/add", controller.UploadAvatar)
	todoAPI.PUT("/edit", controller.UploadKidAvatar)
	todoAPI.DELETE("/delete", nil)
}
