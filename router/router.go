package router

import (
	"fmt"
	"log"
	"os"
	"strings"

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
	if strings.Contains(rootPath, "src/github.com/kidsdynamic") {
		r.LoadHTMLGlob(fmt.Sprintf("%s/templates/*.html", rootPath))
	} else {
		r.LoadHTMLGlob(fmt.Sprintf("%s/src/github.com/kidsdynamic/childrenlab_v2/templates/*.html", rootPath))
	}

	initFWRouter(r)
	initUserRouter(r)
	initKidRouter(r)
	initActivityRouter(r)
	initEventRouter(r)
	initSubHostRouter(r)

	return r
}
