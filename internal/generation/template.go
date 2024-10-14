package generation

import (
	"blog-server/public/utils"
	"io"
	"os"
	"strings"
)

func ApplayTemplate(templatePath, tagTemplatePath, title string, tag []string, postHtmlStr string) (string, error) {
	// 加载模板
	templateFile, err := os.Open(templatePath)

	if err != nil {
		return "", err
	}
	defer templateFile.Close()

	templateByteArray, err := io.ReadAll(templateFile)
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
	htmlStr := strings.Replace(string(templateByteArray), "${{title}}", title, 1)

	htmlStr = strings.Replace(htmlStr, "${{tag}}", tagStr, 1)

	// essay
	htmlStr = strings.Replace(htmlStr, "${{essay}}", postHtmlStr, 1)
	return htmlStr, nil
}
