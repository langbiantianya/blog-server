package dto

import "time"

type EssayDTO struct {
	PaginationDTO
	CreatedAt time.Time `form:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt time.Time `form:"updatedAt" json:"updatedAt,omitempty"`
	Title     string    `form:"title" json:"title,omitempty"`
	Post      string    `form:"post" json:"post,omitempty"`
	Hide      bool      `form:"hide" json:"hide,omitempty"`
	Tags      []string  `form:"tags" json:"tags,omitempty"`
}
