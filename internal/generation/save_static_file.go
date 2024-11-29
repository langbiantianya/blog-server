package generation

import (
	"os"
	"path"
	"strings"
)

func WireStr2File(filePath string, content string) error {
	// 规范文件路径
	filePath = path.Clean(filePath)
	// 判断路径是否存在
	dir := filePath[:strings.LastIndex(filePath, "/")]
	_, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	// 写入文件内容
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
