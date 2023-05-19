package main

import (
	"github.com/gin-gonic/gin"
	"sync"
)

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
}
