package entity

import (
	"gorm.io/gorm"
)

type Essay struct {
	gorm.Model
	Title string `json:"title" gorm:"type:varchar(256);not null"`
	Post  string `json:"post" gorm:"type:text;not null"`
	Tags  []*Tag `gorm:"many2many:essay_tags;"`
}

type Tag struct {
	gorm.Model
	Name   string   `json:"name" gorm:"type:varchar(128);not null"`
	Essays []*Essay `gorm:"many2many:essay_tags;"`
}
