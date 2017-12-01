package router

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("/go/src/github.com/kidsdynamic/childrenlab_v2/templates/*.html")

	initFWRouter(r)
	initUserRouter(r)
	initKidRouter(r)
	initActivityRouter(r)
	initEventRouter(r)
	initSubHostRouter(r)

	return r
}
