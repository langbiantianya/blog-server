package entity

import "time"

type Essay struct {
	ID        uint      `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt,omitempty" `
	UpdatedAt time.Time `json:"updatedAt,omitempty" `
	Title     string    `json:"title,omitempty" gorm:"type:varchar(256);not null;"`
	Post      string    `json:"post,omitempty" gorm:"type:text;not null;"`
	Hide      bool      `json:"hide,omitempty" gorm:"type:tinyint(1);not null;default:false;"`
	Tags      []Tag     `json:"tags,omitempty" gorm:"many2many:essay_tags;"`
}
