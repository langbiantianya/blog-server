package entity

import "time"

type Essay struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string `json:"title,omitempty" gorm:"type:varchar(256);not null;"`
	Post      string `json:"post,omitempty" gorm:"type:text;not null;"`
	Hide      bool   `json:"hide,omitempty" gorm:"type:tinyint(1);not null;default:false;"`
	Tags      []*Tag `json:"tags,omitempty" gorm:"many2many:essay_tags;"`
}
