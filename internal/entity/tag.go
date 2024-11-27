package entity

type Tag struct {
	Entity
	Name   string  `json:"name,omitempty" gorm:"type:varchar(128); null; unique;"`
	Essays []Essay `json:"essays,omitempty" gorm:"many2many:essay_tags;"`
}
