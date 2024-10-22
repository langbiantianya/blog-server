package generation

import (
	"blog-server/internal/entity"
	"blog-server/public/utils"
	"bytes"
	"io"
	"os"
	"strings"
	"text/template"
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
	htmlStr = strings.ReplaceAll(htmlStr, "${{post}}", postHtmlStr)
	return htmlStr, nil
}

func GenerationPostV2(postTemplatePath string, essay entity.Essay) (string, error) {
	postTemplate, err := template.ParseFiles(postTemplatePath)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = postTemplate.Execute(&buf, essay)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func GenerationHomePageV2(homePageTemplatePath string, essays []entity.Essay) (string, error) {
	essays = utils.Map(essays, func(index int, item entity.Essay) (entity.Essay, error) {
		item.Post = Md2html(item.Title, item.Post)
		return item, nil
	})

	homePageTemplate, err := template.ParseFiles(homePageTemplatePath)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = homePageTemplate.Execute(&buf, essays)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
