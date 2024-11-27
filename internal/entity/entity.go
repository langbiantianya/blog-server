package entity

import "time"

type Entity struct {
	ID        uint      `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt,omitempty" `
	UpdatedAt time.Time `json:"updatedAt,omitempty" `
}

var MigeateEntity = []interface{}{
	new(Essay),
	new(Tag),
	new(File),
}
