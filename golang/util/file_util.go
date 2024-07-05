package demo_util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReplaceStringInFilesAndDirs(rootDir string, oldStr string, newStr string) error {
	err := filepath.Walk(rootDir, doReplace(oldStr, newStr))
	for err != nil {
		err = filepath.Walk(rootDir, doReplace(oldStr, newStr))
	}

	return err
}

func doReplace(oldStr string, newStr string) func(path string, info os.FileInfo, err error) error {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// 获取文件夹名并替换字符串
			if path == "target" {
				return nil
			}
		} else {
			// 读取文件内容
			content, err := os.ReadFile(path)
			if err != nil {
				log.Println(err)
				return err
			}
			contentStr := string(content)
			if strings.Contains(contentStr, "hctx") {
				// 替换字符串
				newContent := strings.ReplaceAll(contentStr, oldStr, newStr)
				// 写入替换后的内容到文件
				err = os.WriteFile(path, []byte(newContent), info.Mode())
				if err != nil {
					return err
				}
				fmt.Printf("文件已修改：%s\n", path)
			}
		}

		if strings.Contains(path, oldStr) {
			newPath := strings.Replace(path, oldStr, newStr, -1)
			if err := os.Rename(path, newPath); err != nil {
				log.Println(err)
			}
			fmt.Printf("文件夹已重命名：%s -> %s\n", path, newPath)
		}
		return nil
	}
}
