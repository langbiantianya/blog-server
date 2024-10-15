package generation

import (
	"blog-server/internal/entity"
	"blog-server/public/utils"
	"io"
	"os"
	"strconv"
	"strings"
)

func ApplayTemplate(postTemplatePath, tagTemplatePath, title string, tag []string, postHtmlStr string) (string, error) {
	// 加载模板
	postTemplateFile, err := os.Open(postTemplatePath)

	if err != nil {
		return "", err
	}
	defer postTemplateFile.Close()

	postTemplateByteArray, err := io.ReadAll(postTemplateFile)
	if err != nil {
		return "", err
	}

	tagTemplateFile, err := os.Open(tagTemplatePath)

	if err != nil {
		return "", err
	}
	defer tagTemplateFile.Close()

	tagTemplateByteArray, err := io.ReadAll(tagTemplateFile)

	if err != nil {
		return "", err
	}
	// 生成tag
	tags := utils.Map(tag, func(index int, t string) (string, error) {
		return strings.ReplaceAll(string(tagTemplateByteArray), "${{tag}}", t), nil
	})

	tagStr := utils.Sum(tags, func(sum string, item string) string {
		return sum + item
	})

	// 替换title
	htmlStr := strings.ReplaceAll(string(postTemplateByteArray), "${{title}}", title)

	htmlStr = strings.ReplaceAll(htmlStr, "${{tag}}", tagStr)

	// essay
	htmlStr = strings.ReplaceAll(htmlStr, "${{essay}}", postHtmlStr)
	return htmlStr, nil
}

// 生成主页
func GenerationHomePage(homePageTemplatePath, pageItemTemplatePath, tagTemplatePath string, essays []entity.Essay) (string, error) {
	// 加载模板
	homePageTemplateFile, err := os.Open(homePageTemplatePath)

	if err != nil {
		return "", err
	}
	defer homePageTemplateFile.Close()

	homePageTemplateByteArray, err := io.ReadAll(homePageTemplateFile)
	if err != nil {
		return "", err
	}

	tagTemplateFile, err := os.Open(tagTemplatePath)

	if err != nil {
		return "", err
	}
	defer homePageTemplateFile.Close()

	tagTemplateByteArray, err := io.ReadAll(tagTemplateFile)
	if err != nil {
		return "", err
	}

	pageItemTemplateFile, err := os.Open(pageItemTemplatePath)

	if err != nil {
		return "", err
	}
	defer homePageTemplateFile.Close()

	pageItemTemplateByteArray, err := io.ReadAll(pageItemTemplateFile)
	if err != nil {
		return "", err
	}
	// 添加替换模板
	htmlStr := string(homePageTemplateByteArray)
	pageItemStr := ""
	for index, essay := range essays {
		if index == 5 {
			break
		}
		_pageItemStr := strings.ReplaceAll(string(pageItemTemplateByteArray), "${{title}}", essay.Title)
		_pageItemStr = strings.ReplaceAll(_pageItemStr, "${{post}}", Md2html(essay.Title, essay.Post))
		_pageItemStr = strings.ReplaceAll(_pageItemStr, "${{id}}", strconv.FormatUint(uint64(essay.ID), 10))
		_tagStr := ""
		for _, tag := range essay.Tags {
			_tagStr = _tagStr + strings.ReplaceAll(string(tagTemplateByteArray), "${{tag}}", tag.Name)
		}
		pageItemStr = pageItemStr + strings.ReplaceAll(_pageItemStr, "${{tag}}", _tagStr)
	}
	htmlStr = strings.ReplaceAll(htmlStr, "${{essayPage}}", pageItemStr)
	return htmlStr, nil
}
