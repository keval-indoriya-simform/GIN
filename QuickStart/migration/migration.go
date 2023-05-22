package main

import (
	"github.com/keval-indoriya-simform/GIN/QuickStart/connection"
	"github.com/keval-indoriya-simform/GIN/QuickStart/models"
)

func main() {
	db := connection.GetConnection()
	defer connection.CloseConnection(db)
	db.AutoMigrate(&models.Video{})
}
