package service

import (
	"blog-server/internal/conf"
	"blog-server/internal/entity"
	"blog-server/internal/entity/dto"
	"blog-server/internal/generation"
	"blog-server/internal/repo"
	"fmt"
	"log"
	"os"
	"path"
)

type IEssayService interface {
	Info(uint) (*entity.Essay, error)
	List(dto.EssayDto) (*[]entity.Essay, error)
	Add(entity.Essay) error
	Update(entity.Essay) error
	Hide(uint) error
	Delete(uint) error
	Publish([]uint) error
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

func (essay EssayService) List(params dto.EssayDto) (*[]entity.Essay, error) {
	
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
func (essay EssayService) Publish(ids []uint) error {
	for _, id := range ids {
		// 获取文章数据转为html
		e, err := essay.Info(id)
		if err != nil {
			return err
		}
		html, err := generation.Md2html([]byte(e.Post))
		if err != nil {
			return err
		}
		// 写入文件系统
		filePath := path.Clean(fmt.Sprintf("%s/essay/", conf.GetConfig().StaticPath))
		err = os.MkdirAll(filePath, 0755)
		if err != nil {
			return err
		}
		filePath = path.Clean(fmt.Sprintf("%s/%s.html", filePath, e.Title))
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		bytesWritten, err := file.Write(html)
		if err != nil {
			return err
		}
		log.Printf("Successfully wrote %d bytes to %s\n", bytesWritten, filePath)
	}

	// 编制路由
	// 生成索引
	// TODO
	return nil
}
