package middelware

import (
	"github.com/gin-gonic/gin"
	"github.com/keval-indoriya-simform/GIN/QuickStart/service"
	"log"
	"net/http"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := context.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			//claims := token.Claims.(jwt.MapClaims)
			//log.Println("Claims[Name]", claims["name"])
			//log.Println("Claims[Admin]", claims["admin"])
			//log.Println("Claims[Issuer]", claims["iss"])
			//log.Println("Claims[IssuedAt]", claims["iat"])
			//log.Println("Claims[ExpiresAt]", claims["exp"])
		} else {
			log.Println(err)
			context.AbortWithStatus(http.StatusForbidden)
		}
	}
}
