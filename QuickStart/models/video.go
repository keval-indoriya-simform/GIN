package models

type Video struct {
	Title       string `json:"title,omitempty" binding:"min=2,max=10"`
	Description string `json:"description,omitempty" binding:"max=20"`
	URL         string `json:"url,omitempty" validate:"required,url"`
}
