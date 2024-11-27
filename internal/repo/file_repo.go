package repo

import (
	"blog-server/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type IFileRepo interface {
	IRepo[entity.File]
}

type FileRepo struct {
	db *gorm.DB
}

func NewFileRepo(db *gorm.DB) IFileRepo {
	return FileRepo{
		db,
	}
}

func (file FileRepo) Info(id uint) (*entity.File, error) {
	var res entity.File
	if result := file.db.Model(&entity.File{}).Preload("Essays").First(&res, id); result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &res, nil
}

func (file FileRepo) List(any) (*[]entity.File, error) {
	var res []entity.File
	if result := file.db.Model(&entity.File{}).Find(&res); result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &res, nil
}

func (file FileRepo) Add(paramss entity.File) error {
	if result := file.db.Save(&paramss); result.Error != nil {
		return result.Error
	}
	return nil
}

func (file FileRepo) Delete(id uint) error {
	t, err := file.Info(id)
	if err != nil {
		return err
	}
	if result := file.db.Delete(t); result.Error != nil {
		return result.Error
	}
	return nil
}
