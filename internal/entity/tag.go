package entity

import "time"

type Tag struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string   `json:"name" gorm:"type:varchar(128);not null; unique;"`
	Essays    []Essay `gorm:"many2many:essay_tags;"`
}
