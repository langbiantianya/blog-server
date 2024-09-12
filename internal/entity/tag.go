package entity

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name   string   `json:"name" gorm:"type:varchar(128);not null; unique;"`
	Essays []*Essay `gorm:"many2many:essay_tags;"`
}
