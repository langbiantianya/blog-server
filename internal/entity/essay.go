package entity

import (
	"gorm.io/gorm"
)

type Essay struct {
	gorm.Model
	Title string   `json:"title"`
	Tags  []string `json:"tags"`
	Post  string   `json:"post"`
}
