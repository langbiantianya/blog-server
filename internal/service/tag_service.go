package service

import "blog-server/internal/entity"

type ITagService interface {
	List(entity.Tag) (*[]entity.Tag, error)
}
