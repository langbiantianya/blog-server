package generation

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

func Md2html(markdownText []byte) ([]byte, error) {
	// 创建一个新的 Goldmark 解析器
	parser := goldmark.New(
		// 可以添加一些选项来启用或配置特定的功能
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
			html.WithUnsafe(), // 允许渲染原始 HTML
		),
	)
	var htmlContent bytes.Buffer
	// 解析 Markdown 并渲染为 HTML
	if err := parser.Convert(markdownText, &htmlContent); err != nil {
		return nil, err
	}
	// 打印转换后的 HTML
	return htmlContent.Bytes(), nil
}
