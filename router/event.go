package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/controller"
)

func initEventRouter(r *gin.Engine) {
	calendarAPI := r.Group("/v1/event")
	calendarAPI.Use(controller.Auth)
	calendarAPI.POST("/add", controller.AddCalendarEvent)
	calendarAPI.PUT("/update", controller.UpdateCalendarEvent)
	calendarAPI.DELETE("/delete", controller.DeleteEvent)
	calendarAPI.GET("/retrieveEvents", controller.GetCalendarEvent)
	calendarAPI.GET("/retrieveAllEventsWithTodo", controller.RetrieveAllEventWithTodoByUser)

	todoAPI := r.Group("/v1/event/todo")
	todoAPI.Use(controller.Auth)
	todoAPI.PUT("/done", controller.TodoDone)

}
