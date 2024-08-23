package repo

import (
	"blog-server/internal/constantx"
	"blog-server/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type ITagRepo interface {
	Info(uint) (*entity.Tag, error)
	List() (*[]entity.Tag, error)
	Add(entity.Tag) error
	Delete(uint) error
}

type TagRepo struct {
	db *gorm.DB
}

func NewTagRepo() ITagRepo {
	return TagRepo{
		constantx.Db,
	}
}

func (tag TagRepo) Info(id uint) (*entity.Tag, error) {
	var res entity.Tag
	if result := tag.db.Model(&entity.Tag{}).Preload("Essays").First(&res, id); result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &res, nil
}

func (tag TagRepo) List() (*[]entity.Tag, error) {
	var res []entity.Tag
	if result := tag.db.Model(&entity.Tag{}).Find(&res); result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &res, nil
}

func (tag TagRepo) Add(paramss entity.Tag) error {
	if result := tag.db.Save(&paramss); result.Error != nil {
		return result.Error
	}
	return nil
}

func (tag TagRepo) Delete(id uint) error {
	t, err := tag.Info(id)
	if err != nil {
		return err
	}
	if result := tag.db.Delete(t); result.Error != nil {
		return result.Error
	}
	return nil
}
