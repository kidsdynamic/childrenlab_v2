package router

import (
	"fmt"
	"os"

	"log"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	r.LoadHTMLGlob(fmt.Sprintf("%s/src/github.com/kidsdynamic/childrenlab_v2/view/template/*", rootPath))
	r.Static("/server/assets", fmt.Sprintf("%s/view/assets", rootPath))

	initAdminRouter(r)

	initUserRouter(r)
	initKidRouter(r)
	initActivityRouter(r)
	initEventRouter(r)
	initSubHostRouter(r)

	return r
}
