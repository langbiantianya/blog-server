package repo

import (
	"blog-server/internal/entity"
	"blog-server/internal/entity/dto"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type IEssayRepo interface {
	Info(uint) (*entity.Essay, error)
	Find(dto.EssayDto) (*[]entity.Essay, error)
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

func (essay EssayRepo) Find(params dto.EssayDto) (*[]entity.Essay, error) {
	var res []entity.Essay
	query := essay.db.Model(&entity.Essay{}).Distinct().Preload("Tags").Joins("left JOIN essay_tags on essay_tags.essay_id = essays.id").Joins("left JOIN tags ON tags.id = essay_tags.tag_id")

	if params.Title != "" {
		query.Where("essay.title LIKE ?", fmt.Sprintf("%%%s%%", params.Title))
	}

	if params.Post != "" {
		query.Where("essay.post LIKE ?", fmt.Sprintf("%%%s%%", params.Post))
	}

	if len(params.Tags) != 0 {
		query.Where("tags.name in ?", params.Tags)
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

	err := essay.db.Transaction(func(tx *gorm.DB) error {
		tags := make([]entity.Tag, 0)

		for _, tag := range params.Tags {
			if tag.ID == 0 {
				if res := tx.Model(&entity.Tag{}).Where("name = ?", tag.Name).First(&tag); res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
					return res.Error
				} else if errors.Is(res.Error, gorm.ErrRecordNotFound) {
					if res := tx.Model(&entity.Tag{}).Save(&tag); res.Error != nil {
						return res.Error
					}
				}
			} else {
				if res := tx.Model(&entity.Tag{}).First(&tag, tag.ID); res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
					return res.Error
				}
			}
			tags = append(tags, tag)
		}
		params.Tags = tags

		if res := tx.Exec("DELETE FROM essay_tags WHERE essay_id = ?", params.ID); res.Error != nil {
			return res.Error
		}
		if len(params.Tags) > 0 {
			valueArry := []byte("(?),")
			essayTags := make([]uint, 0)
			sql := []byte("INSERT INTO essay_tags (tag_id,essay_id)  VALUES ")

			for _, tag := range params.Tags {
				sql = append(sql, valueArry...)
				essayTags = append(essayTags, tag.ID, params.ID)
			}
			sqlstr := string(sql)
			sqlstr = sqlstr[:len(sqlstr)-1]
			if res := tx.Exec(sqlstr, essayTags); res.Error != nil {
				return res.Error
			}
		}

		if res := tx.Model(&entity.Essay{}).Where("id", params.ID).Updates(params); res.Error != nil {
			return res.Error
		}

		return nil
	})
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (essay EssayRepo) Hide(id uint) error {
	e := &entity.Essay{}
	e.ID = id
	e.Hide = true
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
	if result := essay.db.Model(&entity.Essay{}).Delete(e); result.Error != nil {
		return result.Error
	}
	return nil
}
