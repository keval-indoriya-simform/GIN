package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/keval-indoriya-simform/GIN/QuickStart/service"
)

var (
	loginService = service.NewLoginService()
	jwtService   = service.NewJWTService()
	credential   = Credential{
		Username: "keval",
		Password: "root",
	}
)

type LoginController interface {
	Login(context *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

type Credential struct {
	Username string
	Password string
}

func NewLoginController() LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (controller *loginController) Login(context *gin.Context) string {
	err := context.ShouldBind(&credential)
	if err != nil {
		return ""
	}
	isAuthenticated := controller.loginService.Login(credential.Username, credential.Password)
	if isAuthenticated {
		return controller.jwtService.GenerateToken(credential.Username, true)
	}
	return ""
}
