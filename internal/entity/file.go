package entity

type File struct {
	Entity
	Path    string `json:"path,omitempty" gorm:"type:varchar(512);not null;"`
	EssayId uint   `json:"essayId,omitempty" gorm:"null;"`
}
