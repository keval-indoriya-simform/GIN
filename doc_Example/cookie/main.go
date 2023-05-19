package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("Cookies")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("Cookies", "test", 60, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})
	router.Run("localhost:8080")
}
