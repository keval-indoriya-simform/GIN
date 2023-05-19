package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.SecureJsonPrefix("kevaljson;")

	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"Keval", "foo"}
		c.SecureJSON(http.StatusOK, names)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
