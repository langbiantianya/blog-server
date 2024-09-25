package repo

import (
	"blog-server/internal/entity"
	"blog-server/internal/entity/dto"
	"blog-server/internal/entity/vo"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type IEssayRepo interface {
	Info(uint) (*entity.Essay, error)
	Find(dto.EssayDTO) (*vo.PaginationVO[[]entity.Essay], error)
	Add(entity.Essay) error
	Update(entity.Essay) error
	Hide(uint) error
	Delete(uint) error
	Publish(uint) error
}

type EssayRepo struct {
	db *gorm.DB
}

// NewEssayRepo 是一个工厂函数，用于创建并返回一个 IEssayRepo 接口的实现。
// 参数:
//   - db: 一个指向 gorm.DB 实例的指针，gorm.DB 是一个 Go 语言的 ORM 库。
//
// 返回值:
//   - IEssayRepo: 一个实现了 IEssayRepo 接口的结构体实例，用于操作数据库中的文章数据。
func NewEssayRepo(db *gorm.DB) IEssayRepo {
	return &EssayRepo{db: db}
}

// Info 根据给定的文章ID从数据库中查询并返回相应的文章信息。
// 参数:
//   - id: 要查询的文章的ID。
//
// 返回值:
//   - *entity.Essay: 查询到的文章信息，如果查询失败则返回nil。
//   - error: 如果查询过程中发生错误，返回错误信息；否则返回nil。
func (essay EssayRepo) Info(id uint) (*entity.Essay, error) {
	var res entity.Essay
	// 执行数据库查询操作，Preload方法用于预加载关联的Tags数据
	if result := essay.db.Model(&entity.Essay{}).Preload("Tags").First(&res, id); result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &res, nil
}

// Find 根据给定的查询参数从数据库中查找文章列表，并返回分页结果。
// 参数:
//   - params: 查询参数，包含标题、帖子内容和标签等条件。
//
// 返回值:
//   - *vo.PaginationVO[[]entity.Essay]: 分页结果，包含当前页码、每页限制条数、总记录数和查询到的文章列表。
//   - error: 如果查询过程中发生错误，返回错误信息；否则返回 nil。
func (essay EssayRepo) Find(params dto.EssayDTO) (*vo.PaginationVO[[]entity.Essay], error) {
	// 初始化查询条件和结果变量
	var res []entity.Essay
	query := essay.db.Model(&entity.Essay{}).Distinct().Preload("Tags").Joins("left JOIN essay_tags on essay_tags.essay_id = essays.id").Joins("left JOIN tags ON tags.id = essay_tags.tag_id")

	// 根据标题进行模糊查询
	if params.Title != "" {
		query.Where("essays.title LIKE ?", fmt.Sprintf("%%%s%%", params.Title))
	}

	// 根据帖子内容进行模糊查询
	if params.Post != "" {
		query.Where("essays.post LIKE ?", fmt.Sprintf("%%%s%%", params.Post))
	}

	query.Where("essays.hide = ?", params.Hide)

	// 根据标签进行查询
	if len(params.Tags) != 0 {
		query.Where("tags.name in ?", params.Tags)
	}

	// 计算总记录数
	var szie int64
	query.Group("essays.id").Count(&szie)

	// 执行分页查询
	if result := query.Order("essays.updated_at DESC").Limit(params.GetLimit()).Offset(params.GetOffset()).Find(&res); result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	// 返回分页结果
	return &vo.PaginationVO[[]entity.Essay]{
		Page:  params.GetPage(),
		Limit: params.GetLimit(),
		Size:  szie,
		Data:  res,
	}, nil
}

// Add 将给定的文章参数添加到数据库中。
// 参数:
//   - params: 要添加的文章实体。
//
// 返回值:
//   - error: 如果添加过程中发生错误，返回错误信息；否则返回 nil。
func (essay EssayRepo) Add(params entity.Essay) error {
	if result := essay.db.Save(&params); result.Error != nil {
		return result.Error
	}
	return nil
}

// Update 更新数据库中指定ID的文章信息。
// 该方法首先处理文章的标签，确保每个标签都存在于数据库中，
// 如果不存在则创建新的标签。然后，它会删除旧的标签关联，
// 并根据提供的文章ID和标签ID创建新的关联。
// 最后，它会更新文章的其他信息。
// 参数:
//   - params: 要更新的文章实体。
//
// 返回值:
//   - error: 如果更新过程中发生错误，返回错误信息；否则返回 nil。
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
	if result := essay.db.Updates(e); result.Error != nil {
		return result.Error
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

func (essay EssayRepo) Publish(id uint) error {
	e := entity.Essay{
		Hide: false,
	}
	if result := essay.db.Model(&entity.Essay{}).Select("hide").Where("id = ?", id).Updates(e); result.Error != nil {
		return result.Error
	}
	return nil
}
