package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/controller"
)

func initUserRouter(r *gin.Engine) {
	v1 := r.Group("/v1/user")

	v1.POST("/login", controller.Login)
	v1.POST("/register", controller.Register)
	v1.GET("/isTokenValid", controller.IsTokenValid)
	v1.GET("/isEmailAvailableToRegister", controller.IsEmailAvailableToRegister)
	v1.GET("/findByEmail", controller.FindUserByEmail)

	authAPI := r.Group("/v1/user")

	authAPI.Use(controller.Auth)
	authAPI.PUT("/updateProfile", controller.UpdateProfile)
	authAPI.GET("/retrieveUserProfile", controller.UserProfile)
	authAPI.PUT("/updateIOSRegistrationId", controller.UpdateIOSRegistrationId)

}
