package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.String(200, "pong")
}

func PostReq(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"page": page,
	})
	//fmt.Printf("id: %s; page: %s;", id, page)
}
