package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//r.GET("/test13", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "https://www.google.com/")
	//})

	//r.POST("/test", func(c *gin.Context) {
	//	c.Redirect(http.StatusFound, "https://www.google.com/")
	//})

	r.GET("/test3", func(c *gin.Context) {
		//c.Redirect(http.StatusMovedPermanently, "/test2")
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})
	r.Run(":8080")
}
