package service

import "github.com/keval-indoriya-simform/GIN/QuickStart/models"

type VideoService interface {
	Save(video models.Video) models.Video
	FindAll() []models.Video
}

type videoService struct {
	videos []models.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video models.Video) models.Video {
	models.InsertVideoData(&video)
	return video
}
func (service *videoService) FindAll() []models.Video {
	return models.GetAllVideo()
}
