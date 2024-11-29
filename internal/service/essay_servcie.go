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
	"log"
	"os"
)

type IEssayService interface {
	InitBlog() error
	Info(uint) (*entity.Essay, error)
	List(dto.EssayDTO) (*vo.PaginationVO[[]entity.Essay], error)
	Add(entity.Essay) error
	Update(entity.Essay) error
	Hide(uint) error
	Delete(uint) error
	Publish(uint) error
	SaveFile(func() string) error
}

type EssayService struct {
	essayRepo repo.IEssayRepo
	tagRepo   repo.ITagRepo
	fileRepo  repo.IFileRepo
}

func NewEssayService(essayRepo repo.IEssayRepo, tagRepo repo.ITagRepo) IEssayService {
	service := EssayService{
		essayRepo: essayRepo,
		tagRepo:   tagRepo,
	}
	err := service.InitBlog()
	if err != nil {
		log.Fatal(err)
	}
	return &service
}

func (essay EssayService) InitBlog() error {
	staticPath := conf.GetConfig().StaticPath
	essays, err := essay.essayRepo.Find(dto.EssayDTO{
		PaginationDTO: dto.PaginationDTO{
			Limit: 5,
		},
		Hide: false,
	})
	if err != nil {
		return err
	}
	allEssays, err := essay.essayRepo.Find(dto.EssayDTO{
		Hide: false,
	})
	if err != nil {
		return err
	}
	// 生成索引
	indexJson, err := generation.Index(allEssays.Data)
	if err != nil {
		return err
	}
	err = generation.WireStr2File(fmt.Sprintf("%s/index.json", conf.GetConfig().StaticOutPath), indexJson)
	if err != nil {
		return err
	}
	log.Println("索引生成完成")
	// 生成文章
	postTemplatePath, err := utils.GetFilePath(staticPath+"/template/post.gohtml", staticPath+"/template/defaultPost.gohtml")
	if err != nil {
		return err
	}
	for _, essay := range allEssays.Data {
		go func() {
			log.Println("正在生成《" + essay.Title + "》")
			postHtml, err := generation.GenerationPostV2(postTemplatePath, essay)
			if err != nil {
				log.Fatal(err)
			}
			// 写入文件
			err = generation.WireStr2File(fmt.Sprintf("%s/post/%d/index.html", conf.GetConfig().StaticOutPath, essay.ID), postHtml)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	// 生成主页
	homeTemplatePath, err := utils.GetFilePath(staticPath+"/template/home.gohtml", staticPath+"/template/defaultHome.gohtml")
	if err != nil {
		return err
	}

	homeHtml, err := generation.GenerationHomePageV2(homeTemplatePath, essays.Data)
	if err != nil {
		return err
	}
	err = generation.WireStr2File(fmt.Sprintf("%s/index.html", conf.GetConfig().StaticOutPath), homeHtml)
	if err != nil {
		return err
	}
	log.Println("主页生成完成")
	// 生成搜索
	searchTemplatePath, err := utils.GetFilePath(staticPath+"/template/search.gohtml", staticPath+"/template/defaultSearch.gohtml")
	if err != nil {
		return err
	}
	searchHtml, err := generation.GenerationSearch(searchTemplatePath)
	if err != nil {
		return err
	}
	err = generation.WireStr2File(fmt.Sprintf("%s/search/index.html", conf.GetConfig().StaticOutPath), searchHtml)
	if err != nil {
		return err
	}
	log.Println("搜索生成完成")
	return nil
}

func (essay EssayService) Info(id uint) (*entity.Essay, error) {
	return essay.essayRepo.Info(id)
}

func (essay EssayService) List(params dto.EssayDTO) (*vo.PaginationVO[[]entity.Essay], error) {

	return essay.essayRepo.Find(params)
}

func (essay EssayService) Add(params entity.Essay) error {
	return essay.essayRepo.Add(params)
}

func (essay EssayService) Update(params entity.Essay) error {
	return essay.essayRepo.Update(params)
}

func (essay EssayService) Hide(id uint) error {
	// 删除已经生成的文件

	err := os.RemoveAll(fmt.Sprintf("%s/post/%d", conf.GetConfig().StaticOutPath, id))
	if err != nil {
		return err
	}

	// 重新生成索引
	essaysT, err := essay.essayRepo.Find(dto.EssayDTO{
		Hide: false,
	})

	if err != nil {
		return err
	}

	essays := utils.Filter(essaysT.Data, func(essayT entity.Essay) bool {
		return essayT.ID != id
	})

	// 生成索引
	indexJson, err := generation.Index(essays)
	if err != nil {
		return err
	}
	// 写入文件
	err = generation.WireStr2File(fmt.Sprintf("%s/index.json", conf.GetConfig().StaticOutPath), indexJson)
	if err != nil {
		return err
	}

	// 生成主页
	staticPath := conf.GetConfig().StaticPath
	homeTemplatePath, err := utils.GetFilePath(staticPath+"/template/home.gohtml", staticPath+"/template/defaultHome.gohtml")
	if err != nil {
		return err
	}

	homeHtml, err := generation.GenerationHomePageV2(homeTemplatePath, essays)
	if err != nil {
		return err
	}
	// 写入文件
	err = generation.WireStr2File(fmt.Sprintf("%s/index.html", conf.GetConfig().StaticOutPath), homeHtml)
	if err != nil {
		return err
	}
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

	if res != nil {
		res.Post = generation.Md2html(res.Title, res.Post)
	} else {
		return errors.New("没有找到文章")
	}

	// 获取模板
	staticPath := conf.GetConfig().StaticPath
	postTemplatePath, err := utils.GetFilePath(staticPath+"/template/post.gohtml", staticPath+"/template/defaultPost.gohtml")
	if err != nil {
		return err
	}

	// 生成页面文件
	htmlStr, err := generation.GenerationPostV2(postTemplatePath, *res)

	if err != nil {
		return nil
	}

	// 写入文件
	err = generation.WireStr2File(fmt.Sprintf("%s/post/%d/index.html", conf.GetConfig().StaticOutPath, res.ID), htmlStr)
	if err != nil {
		return err
	}

	// 生成索引文件
	essaysT, err := essay.essayRepo.Find(dto.EssayDTO{
		Hide: false,
	})
	if err != nil {
		return err
	}

	essays := make([]entity.Essay, len(essaysT.Data)+1)
	essays[0] = *res
	copy(essays[1:], essaysT.Data)

	indexJson, err := generation.Index(essays)

	if err != nil {
		return err
	}
	// 写入文件
	err = generation.WireStr2File(fmt.Sprintf("%s/index.json", conf.GetConfig().StaticOutPath), indexJson)
	if err != nil {
		return err
	}

	// 生成主页
	homeTemplatePath, err := utils.GetFilePath(staticPath+"/template/home.gohtml", staticPath+"/template/defaultHome.gohtml")
	if err != nil {
		return err
	}

	homeHtml, err := generation.GenerationHomePageV2(homeTemplatePath, essays)
	if err != nil {
		return err
	}
	// 写入文件
	err = generation.WireStr2File(fmt.Sprintf("%s/index.html", conf.GetConfig().StaticOutPath), homeHtml)
	if err != nil {
		return err
	}
	// 写入文件
	err = essay.essayRepo.Publish(id)
	if err != nil {
		return err
	}

	return nil
}

func (essay EssayService) SaveFile(filePath func() string) error {

	err := essay.fileRepo.Add(entity.File{
		Path: filePath(),
	})
	if err != nil {
		return err
	}
	return nil
}
