package router

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	initUserRouter(r)
	initKidRouter(r)
	initAvatarRouter(r)
	initActivityRouter(r)
	initEventRouter(r)
	initSubHostRouter(r)

	return r
}
