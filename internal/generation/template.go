package generation

import (
	"io"
	"os"
	"strings"
)

func ApplayTemplate(templatePath, title string, tag []string, postHtmlStr string) (string, error) {
	// 加载模板
	file, err := os.Open(templatePath)

	if err != nil {
		return "", err
	}
	defer file.Close()
	templateByteArray, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	// 替换title
	htmlStr := strings.Replace(string(templateByteArray), "${{title}}", title, 1)
	// essay
	htmlStr = strings.Replace(htmlStr, "${{essay}}", postHtmlStr, 1)
	return htmlStr, nil
}
