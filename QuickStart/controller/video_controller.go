package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/keval-indoriya-simform/GIN/QuickStart/models"
	"github.com/keval-indoriya-simform/GIN/QuickStart/service"
	"net/http"
)

type VideoController interface {
	FindAll() []models.Video
	Save(context *gin.Context) models.Video
	ShowAll(context *gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(serv service.VideoService) VideoController {
	return &controller{
		service: serv,
	}
}

func (c *controller) FindAll() []models.Video {
	return c.service.FindAll()
}

func (c *controller) Save(context *gin.Context) models.Video {
	var video models.Video
	err := context.BindJSON(&video)
	if err != nil {
		panic(err)
	}
	validate := validator.New()
	err = validate.Struct(&video)
	if err != nil {
		panic(err)
	}

	c.service.Save(video)
	return video
}

func (c *controller) ShowAll(context *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	context.HTML(http.StatusOK, "index.html", data)
}
