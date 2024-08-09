package demo_util

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
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

func RenameFilesSequentially(dirPath string) {

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Read directory error:", err)
		return
	}
	fileInfos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Println("Read directory error:", err)
			return
		}
		fileInfos = append(fileInfos, info)
	}

	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].Name() < fileInfos[j].Name()
	})

	for i, fileInfo := range fileInfos {
		oldPath := filepath.Join(dirPath, fileInfo.Name())
		newName := strconv.Itoa(i+1) + filepath.Ext(fileInfo.Name())
		newPath := filepath.Join(dirPath, newName)

		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Println("Failed to rename file:", err)
		} else {
			fmt.Printf("File renamed from %s to %s\n", oldPath, newPath)
		}
	}
}
