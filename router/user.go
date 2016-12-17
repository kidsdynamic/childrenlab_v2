package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/controller"
)

func initUserRouter(r *gin.Engine) {
	v1 := r.Group("/v1/user")

	v1.POST("/login", controller.Login)
	v1.POST("/register", controller.Register)
	v1.POST("/isTokenValid", controller.IsTokenValid)
}
