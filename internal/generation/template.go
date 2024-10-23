package generation

import (
	"blog-server/internal/entity"
	"blog-server/public/utils"
	"bytes"
	"io/ioutil"
	"os"
	"text/template"
)

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

func GenerationSearch(searchTemplatePath string) (string, error) {
	// 使用 os 包打开文件
	searchTemplate, err := os.Open(searchTemplatePath)
	if err != nil {
		return "", err
	}
	defer searchTemplate.Close()
	searchhtml, err := ioutil.ReadAll(searchTemplate)
	if err != nil {
		return "", err
	}
	return string(searchhtml), nil
}
