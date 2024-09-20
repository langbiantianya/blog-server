package entity

import "time"

type Tag struct {
	ID        uint      `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt,omitempty" `
	UpdatedAt time.Time `json:"updatedAt,omitempty" `
	Name      string    `json:"name,omitempty" gorm:"type:varchar(128);not null; unique;"`
	Essays    []Essay   `json:"essays,omitempty" gorm:"many2many:essay_tags;"`
}
