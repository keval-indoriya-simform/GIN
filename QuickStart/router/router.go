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

func init() {
	Server.Static("/css", "./templates/css")
	Server.LoadHTMLGlob("templates/*.html")

	Server.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", nil)
	})
	Server.POST("/login", func(context *gin.Context) {
		token := loginController.Login(context)
		if token != "" {
			context.SetCookie("Cookies", token, 60, "/", "localhost", false, true)
			context.Redirect(http.StatusOK, "/view/videos")
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
