package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/controller"
)

func initAvatarRouter(r *gin.Engine) {
	avatarAPI := r.Group("/v1/user/avatar")
	avatarAPI.Use(controller.Auth)
	avatarAPI.POST("/upload", controller.UploadAvatar)
	avatarAPI.POST("/uploadKid", controller.UploadKidAvatar)
}
