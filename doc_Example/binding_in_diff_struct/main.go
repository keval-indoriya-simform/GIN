package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type FormA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type FormB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/json", func(ctx *gin.Context) {
		if ctx.ShouldBindBodyWith(&FormA{}, binding.JSON) == nil {
			fmt.Println("form a")
			ctx.String(http.StatusOK, "Body should be FormA, JSON")
		} else if ctx.ShouldBindBodyWith(&FormB{}, binding.JSON) == nil {
			fmt.Println("form b")
			ctx.String(http.StatusOK, "Body should be FormB, JSON")
		} else if ctx.ShouldBindBodyWith(&FormA{}, binding.XML) == nil {
			fmt.Println("form a xml")
			ctx.String(http.StatusOK, "Body should be FormA, XML")
		} else {
			ctx.String(http.StatusBadRequest, "Error")
		}
	})
	r.Run()
}
