package router

import (
	"github.com/gin-gonic/gin"
	"github.com/keval-indoriya-simform/GIN/QuickStart/controller"
	"github.com/keval-indoriya-simform/GIN/QuickStart/middelware"
	"github.com/keval-indoriya-simform/GIN/QuickStart/service"
	"net/http"
)

var (
	Server                                     = gin.Default()
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController()
)

func init() { // Server.Static("/css", "QuickStart/templates/css")
	Server.LoadHTMLGlob("templates/*.html")

	Server.POST("/login", func(context *gin.Context) {
		token := loginController.Login(context)
		if token != "" {
			context.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			context.JSON(http.StatusForbidden, nil)
		}
	})
	apiRoutes := Server.Group("/api", middelware.AuthorizeJWT())
	apiRoutes.GET("/videos", func(context *gin.Context) {
		context.JSON(http.StatusOK, videoController.FindAll())
	})
	apiRoutes.POST("/videos", func(context *gin.Context) {
		context.JSON(http.StatusOK, videoController.Save(context))
	})
	viewRoutes := Server.Group("/view")
	viewRoutes.GET("/videos", videoController.ShowAll)
}
