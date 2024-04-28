package main

import (
	"fmt"
	"os"
)

func main() {
	println("\n打开文件")
	file, err := os.Open("118_file/01_open_close_file/demo_open_close_file.go")
	if err != nil {
		fmt.Println("文件打开出错：", err)
	}

	fmt.Printf("文件=%v", file)

	println("\n关闭文件")
	err2 := file.Close()
	if err2 != nil {
		fmt.Println("关闭失败")
	}
}
