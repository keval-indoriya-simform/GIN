package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/keval-indoriya-simform/GIN/QuickStart/models"
	"github.com/keval-indoriya-simform/GIN/QuickStart/service"
	"net/http"
)

type VideoController interface {
	FindAll() []models.Video
	Save(context *gin.Context) models.Video
	ShowAll(context *gin.Context)
}

type Controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &Controller{
		service: service,
	}
}

func (c *Controller) FindAll() []models.Video {
	return c.service.FindAll()
}

func (c *Controller) Save(context *gin.Context) models.Video {
	var video models.Video
	context.BindJSON(&video)
	c.service.Save(video)
	return video
}

func (c *Controller) ShowAll(context *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	context.HTML(http.StatusOK, "index.html", data)
}
