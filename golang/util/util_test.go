package demo_util

import (
	"fmt"
	"testing"
)

func TestReplaceStringInFilesAndDirs(t *testing.T) {
	err := ReplaceStringInFilesAndDirs("D:\\workspace\\flby\\flby-manager-system", "yudao", "flby")
	if err != nil {
		fmt.Println("操作过程中发生错误：", err)
	} else {
		fmt.Println("操作完成")
	}
}

func TestRenameFilesSequentially(t *testing.T) {
	RenameFilesSequentially("C:\\Users\\huang\\Pictures\\峰林布依\\官网图\\篝火晚会")
}
