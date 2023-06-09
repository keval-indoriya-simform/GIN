package main

import "github.com/gin-gonic/gin"

func main() {
	//gin.DisableConsoleColor()
	gin.ForceConsoleColor()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
