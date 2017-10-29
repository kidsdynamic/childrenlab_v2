package router

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	/*
		if strings.Contains(rootPath, "src/github.com/kidsdynamic") {
			r.LoadHTMLGlob(fmt.Sprintf("%s/view/template/*", rootPath))
		} else {
			r.LoadHTMLGlob(fmt.Sprintf("%s/src/github.com/kidsdynamic/childrenlab_v2/app/view/template/*", rootPath))
		}
	*/

	//r.Static("/server/assets", fmt.Sprintf("%s/view/assets", rootPath))

	initAdminRouter(r)
	initFWRouter(r)
	initUserRouter(r)
	initKidRouter(r)
	initActivityRouter(r)
	initEventRouter(r)
	initSubHostRouter(r)

	return r
}
