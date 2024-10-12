package service

import (
	"blog-server/internal/conf"
	"blog-server/internal/entity"
	"blog-server/internal/entity/dto"
	"blog-server/internal/entity/vo"
	"blog-server/internal/generation"
	"blog-server/internal/repo"
	"blog-server/public/utils"
	"errors"
	"fmt"
	"os"
)

type IEssayService interface {
	Info(uint) (*entity.Essay, error)
	List(dto.EssayDTO) (*vo.PaginationVO[[]entity.Essay], error)
	Add(entity.Essay) error
	Update(entity.Essay) error
	Hide(uint) error
	Delete(uint) error
	Publish(uint) error
}

type EssayService struct {
	essayRepo repo.IEssayRepo
	tagRepo   repo.ITagRepo
}

func NewEssayService(essayRepo repo.IEssayRepo, tagRepo repo.ITagRepo) IEssayService {
	return &EssayService{
		essayRepo: essayRepo,
		tagRepo:   tagRepo,
	}
}

func (essay EssayService) Info(id uint) (*entity.Essay, error) {
	return essay.essayRepo.Info(id)
}

func (essay EssayService) List(params dto.EssayDTO) (*vo.PaginationVO[[]entity.Essay], error) {

	return essay.essayRepo.Find(params)
}

func (essay EssayService) Add(params entity.Essay) error {
	// TODO 保留一份到文件系统中
	return essay.essayRepo.Add(params)
}

func (essay EssayService) Update(params entity.Essay) error {
	return essay.essayRepo.Update(params)
}

func (essay EssayService) Hide(id uint) error {
	return essay.essayRepo.Hide(id)
}

func (essay EssayService) Delete(id uint) error {
	return essay.essayRepo.Delete(id)
}

// TODO: 考虑异步处理
func (essay EssayService) Publish(id uint) error {
	// 生成转换html文本
	res, err := essay.essayRepo.Info(id)
	if err != nil {
		return err
	}
	var md2htmlStr string
	if res != nil {
		md2htmlStr = generation.Md2html(res.Title, res.Post)
	} else {
		return errors.New("没有找到文章")
	}

	// 获取模板
	staticPath := conf.GetConfig().StaticPath
	defaultTemplatePath := staticPath + "/template/default.html"
	customizedTeplatePath := staticPath + "/template/index.html"
	_, customizedErr := os.Stat(customizedTeplatePath)
	_, defaultErr := os.Stat(defaultTemplatePath)

	var templatePath string

	if customizedErr == nil {
		templatePath = customizedTeplatePath
	} else if defaultErr == nil {
		templatePath = defaultTemplatePath
	} else {
		return errors.Join(defaultErr, customizedErr)
	}

	// 生成页面文件
	htmlStr, err := generation.ApplayTemplate(res.Title, templatePath, md2htmlStr)

	if err != nil {
		return nil
	}
	// 写入文件
	err = generation.WireStr2File(fmt.Sprintf("%s/post/%d/index.html", conf.GetConfig().StaticOutPath, res.ID), htmlStr)
	if err != nil {
		return err
	}
	// 编制tags
	essays, err := essay.essayRepo.Find(
		dto.EssayDTO{
			Hide: false,
		})
	if err != nil {
		return err
	}
	// 取出已经发布的tag去重
	tags := utils.Flatten(
		utils.Map(essays.Data, func(index int, item entity.Essay) ([]string, error) {
			return utils.Map(item.Tags, func(_ int, tag entity.Tag) (string, error) {
				return tag.Name, nil
			}), err
		}))
	tags = utils.DistinctBy(tags, func(item string) any {
		return item
	})
	
	// 生成索引
	// TODO

	err = essay.essayRepo.Publish(id)
	if err != nil {
		return err
	}

	return nil
}
