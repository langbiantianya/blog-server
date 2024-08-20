package service

import (
	"blog-server/internal/entity"
	"blog-server/internal/repo"
)

type IEssayService interface {
	Info(id uint) (entity.Essay, error)
	List(params any) (*[]entity.Essay, error)
	Add(entity.Essay) error
	Update(entity.Essay) error
	Hide(id uint) error
	Delete(id uint) error
}

type EssayService struct {
	essayRepo repo.IEssayRepo
}

func NewEssayRepo(db *repo.IEssayRepo) IEssayService {
	return &EssayService{
		essayRepo: *db,
	}
}

func (essay EssayService) Info(id uint) (entity.Essay, error) {
	panic("TODO")
}
func (essay EssayService) List(params any) (*[]entity.Essay, error) {
	panic("TODO")
}
func (essay EssayService) Add(entity.Essay) error {
	panic("TODO")
}
func (essay EssayService) Update(entity.Essay) error {
	panic("TODO")
}
func (essay EssayService) Hide(id uint) error {
	panic("TODO")
}
func (essay EssayService) Delete(id uint) error {
	panic("TODO")
}
