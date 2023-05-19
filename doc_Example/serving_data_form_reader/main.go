package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//router.GET("/someDataFromReader1", func(c *gin.Context) {
	//	response, err := http.Get("https://mobcup.net/d/ig1kiy8m/mp3")
	//	if err != nil || response.StatusCode != http.StatusOK {
	//		c.Status(http.StatusServiceUnavailable)
	//		return
	//	}
	//
	//	reader := response.Body
	//	contentLength := response.ContentLength
	//	contentType := response.Header.Get("Content-Type")
	//
	//	extraHeaders := map[string]string{
	//		"Content-Disposition": `attachment;`,
	//	}
	//
	//	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	//})

	router.Static("/temp", "./templates")
	//router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/index", "./index.html")

	router.Run(":8080")
}
