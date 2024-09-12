package repo

import (
	"blog-server/internal/entity"
	"blog-server/public/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type IEssayRepo interface {
	Info(uint) (*entity.Essay, error)
	Find(entity.Essay) (*[]entity.Essay, error)
	Add(entity.Essay) error
	Update(entity.Essay) error
	Hide(uint) error
	Delete(uint) error
}

type EssayRepo struct {
	db *gorm.DB
}

func NewEssayRepo(db *gorm.DB) IEssayRepo {
	return &EssayRepo{db: db}
}

func (essay EssayRepo) Info(id uint) (*entity.Essay, error) {
	var res entity.Essay
	if result := essay.db.Model(&entity.Essay{}).Preload("Tags").First(&res, id); result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &res, nil
}

func (essay EssayRepo) Find(params entity.Essay) (*[]entity.Essay, error) {
	var res []entity.Essay
	query := essay.db.Model(&entity.Essay{}).Preload("Tags")

	if params.Title != "" {
		query.Where("essay.title LIKE ?", fmt.Sprintf("%%%s%%", params.Title))
	}

	if params.Post != "" {
		query.Where("essay.post LIKE ?", fmt.Sprintf("%%%s%%", params.Post))
	}

	if len(params.Tags) != 0 {
		query.Where("tag.id in ?", utils.Map(params.Tags, func(index int, item *entity.Tag) (uint, error) {
			if item != nil {
				return item.ID, nil
			} else {
				return 0, fmt.Errorf("item is nil")
			}
		}))
	}

	if result := query.Find(&res); result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &res, nil
}

func (essay EssayRepo) Add(params entity.Essay) error {
	if result := essay.db.Save(&params); result.Error != nil {
		return result.Error
	}
	return nil
}

func (essay EssayRepo) Update(params entity.Essay) error {
	if result := essay.db.Model(&entity.Essay{}).Where("id", params.ID).Updates(&params); result.Error != nil {
		return result.Error
	}
	return nil

}

func (essay EssayRepo) Hide(id uint) error {
	e := &entity.Essay{}
	e.ID = id
	if err := essay.Update(*e); err != nil {
		return err
	}
	return nil
}

func (essay EssayRepo) Delete(id uint) error {
	e, err := essay.Info(id)
	if err != nil {
		return err
	}
	if result := essay.db.Delete(e); result.Error != nil {
		return result.Error
	}
	return nil
}
