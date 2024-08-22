package repo

import (
	"blog-server/internal/constantx"
	"blog-server/internal/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type IEssayRepo interface {
	Info(id uint) (*entity.Essay, error)
	Find(params entity.Essay) (*[]entity.Essay, error)
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

func (essay EssayRepo) Info(id uint) (*entity.Essay, error) {
	var res entity.Essay
	if result := essay.db.Model(&entity.Essay{}).First(&res, id); result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &res, nil
}

func (essay EssayRepo) Find(params entity.Essay) (*[]entity.Essay, error) {
	var res []entity.Essay
	query := essay.db.Model(&entity.Essay{})

	if params.Title != "" {
		query.Where("title LIKE ?", fmt.Sprintf("%%%s%%", params.Title))
	}

	if params.Post != "" {
		query.Where("post LIKE ?", fmt.Sprintf("%%%s%%", params.Post))
	}

	if len(params.Tags) !=0 {
		query.Where("tags ")
	}

	if result := query.Find(&res); result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
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
