package router

import (
	"github.com/gin-gonic/gin"
	"github.com/keval-indoriya-simform/GIN/QuickStart/controller"
	"github.com/keval-indoriya-simform/GIN/QuickStart/middelware"
	"github.com/keval-indoriya-simform/GIN/QuickStart/service"
)

var (
	Server                                     = gin.Default()
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func init() {
	Server.Use(middelware.BasicAuth())

	Server.Static("/css", "QuickStart/templates/css")
	Server.LoadHTMLGlob("templates/*.html")

	apiRoutes := Server.Group("/api")
	{
		apiRoutes.GET("/videos", func(context *gin.Context) {
			context.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(context *gin.Context) {
			context.JSON(200, videoController.Save(context))
		})
	}

	viewRoutes := Server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}
}
