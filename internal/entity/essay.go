package entity

import (
	"gorm.io/gorm"
)

type Essay struct {
	gorm.Model
	Title string `json:"title" gorm:"type:varchar(256);not null;"`
	Post  string `json:"post" gorm:"type:text;not null;"`
	Hide  bool   `json:"hide" gorm:"type:tinyint(1);not null;default:false;"`
	Tags  []*Tag `gorm:"many2many:essay_tags;"`
}
