package generation

import (
	"blog-server/public/utils"
	"io"
	"os"
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
		return strings.Replace(string(tagTemplateByteArray), "${{tag}}", t, 1), nil
	})

	tagStr := utils.Sum(tags, func(sum string, item string) string {
		return sum + item
	})

	// 替换title
	htmlStr := strings.Replace(string(postTemplateByteArray), "${{title}}", title, 1)

	htmlStr = strings.Replace(htmlStr, "${{tag}}", tagStr, 1)

	// essay
	htmlStr = strings.Replace(htmlStr, "${{essay}}", postHtmlStr, 1)
	return htmlStr, nil
}
