package entity

import (
	"gorm.io/gorm"
)

type Essay struct {
	gorm.Model
	Title string `json:"title,omitempty" gorm:"type:varchar(256);not null;"`
	Post  string `json:"post,omitempty" gorm:"type:text;not null;"`
	Hide  bool   `json:"hide,omitempty" gorm:"type:tinyint(1);not null;default:false;"`
	Tags  []*Tag `json:"tags,omitempty" gorm:"many2many:essay_tags;"`
}
