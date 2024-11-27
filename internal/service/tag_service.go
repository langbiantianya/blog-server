package service

import (
	"blog-server/internal/entity"
	"blog-server/internal/repo"
)

type ITagService interface {
	List() (*[]entity.Tag, error)
}

type TagService struct {
	tagRepo repo.ITagRepo
}

func NewTagService(tagRepo repo.ITagRepo) ITagService {
	return TagService{
		tagRepo,
	}
}

func (t TagService) List() (*[]entity.Tag, error) {
	res, err := t.tagRepo.List(nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
