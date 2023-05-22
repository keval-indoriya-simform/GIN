package models

import "github.com/keval-indoriya-simform/GIN/QuickStart/connection"

type Video struct {
	Title       string `json:"title,omitempty" binding:"min=2,max=10"`
	Description string `json:"description,omitempty" binding:"max=20"`
	URL         string `json:"url,omitempty" validate:"required,url"`
}

func InsertVideoData(video *Video) {
	db := connection.GetConnection()
	defer connection.CloseConnection(db)
	db.Create(video)
}

func GetAllVideo() []Video {
	var videos []Video
	db := connection.GetConnection()
	defer connection.CloseConnection(db)
	db.Find(&videos)
	return videos
}
