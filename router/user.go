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
	v1.GET("/resetPasswordPage", controller.ResetPasswordPage)
	v1.POST("/resetPassword", controller.ResetPassword)
	v1.POST("/sendResetPasswordEmail", controller.SendResetPasswordEmail)
	authAPI := r.Group("/v1/user")

	authAPI.Use(controller.Auth)
	authAPI.PUT("/updateProfile", controller.UpdateProfile)
	authAPI.GET("/retrieveUserProfile", controller.UserProfile)
	authAPI.PUT("/updateIOSRegistrationId", controller.UpdateIOSRegistrationId)
	authAPI.PUT("/updateAndroidRegistrationId", controller.UpdateAndroidRegistrationId)
	authAPI.POST("/updateLanguage", controller.UpdateLanguage)
	authAPI.GET("/getUserByEmail", controller.GetUserByEmail)
	authAPI.POST("/updatePassword", controller.UpdatePassword)

}
