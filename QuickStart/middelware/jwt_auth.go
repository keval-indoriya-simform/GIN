package middelware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keval-indoriya-simform/GIN/QuickStart/service"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := context.GetHeader("Authorization")
		tokenString := authHeader[len(BearerSchema):]

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if !token.Valid {
			log.Println(err)
			context.AbortWithStatus(http.StatusForbidden)
		}
	}
}
