package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "118_file/02_read_file/demo_read_file.go"

	println("\n使用os.ReadFile读取文件的内容并显示在终端")
	// os.ReadFile会一次将整个文件读入到内存中，所以这种方式适用于文件不大的情况。
	// 因为文件的打开和关闭操作被封装在os.ReadFile函数内部
	// 所以使用os.ReadFile不需要打开和关闭文件
	content, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("读取出错，错误为：", err)
	}

	// 如果读取成功，将内容显示在终端
	fmt.Printf("%v", string(content))

	println("\n通过缓冲区读取文件的内容")
	// 缓冲区读取文件，适合读取比较大的文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("文件打开失败，err=", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败，err=", err)
		}
	}(file)

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		// 下面语句 io.EOF 表示已经读取到文件的结尾
		if err == io.EOF {
			break
		}
	}
	fmt.Println("文件已被全部读取")
}
