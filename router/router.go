package router

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	initUserRouter(r)

	return r
}
