package entity

type Essay struct {
	Entity
	Title string `json:"title,omitempty" gorm:"type:varchar(256);not null;"`
	Post  string `json:"post,omitempty" gorm:"type:text;not null;"`
	Hide  bool   `json:"hide,omitempty" gorm:"type:tinyint(1);not null;default:true;"`
	Tags  []Tag  `json:"tags,omitempty" gorm:"many2many:essay_tags;"`
	Files []File `json:"files,omitempty"`
}
