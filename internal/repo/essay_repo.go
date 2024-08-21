package repo

import (
	"blog-server/internal/constantx"
	"blog-server/internal/entity"

	"gorm.io/gorm"
)

type IEssayRepo interface {
	Info(id uint) (entity.Essay, error)
	FindOne(params any) (entity.Essay, error)
	Find(params any) (*[]entity.Essay, error)
	Add(entity.Essay) error
	Update(entity.Essay) error
	Hide(id uint) error
	Delete(id uint) error
}

type EssayRepo struct {
	db *gorm.DB
}

func NewEssayRepo() IEssayRepo {
	return &EssayRepo{db: constantx.Db}
}

func (essay EssayRepo) Info(id uint) (entity.Essay, error) {
	panic("TODO")
}

func (essay EssayRepo) FindOne(params any) (entity.Essay, error) {
	panic("TODO")

}

func (essay EssayRepo) Find(params any) (*[]entity.Essay, error) {
	panic("TODO")

}

func (essay EssayRepo) Add(entity.Essay) error {
	panic("TODO")

}

func (essay EssayRepo) Update(entity.Essay) error {
	panic("TODO")

}

func (essay EssayRepo) Hide(id uint) error {
	panic("TODO")

}

func (essay EssayRepo) Delete(id uint) error {
	panic("TODO")
}
