package main

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

var (
	router = gin.Default()
	once   sync.Once
)

func Router() *gin.Engine {
	once.Do(setupRouter)
	return router
}

func setupRouter() {
	router.GET("/ping", Ping)
	router.POST("/post", PostReq)
	router.GET("/:name/:id", URL)
	router.POST("/map", Map)
	router.POST("/struct", StructBinding)
	router.POST("/upload", Upload)
}
