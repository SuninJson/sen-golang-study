package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	linuxAuthority := os.FileMode(0666).String()
	fmt.Println("0666在Linux中表示的文件权限：", linuxAuthority)

	file, err := os.OpenFile("118_file/03_write_file/test_write.txt",
		os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败", err)
		}
	}(file)

	// 写入文件操作：---> IO流 ---> 缓冲输出流(带缓冲区)
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		_, err := writer.WriteString("Hello Golang\n")
		if err != nil {
			fmt.Println("写入文件失败", err)
			return
		}
	}

	// 将缓冲区数据刷新到文件中
	err = writer.Flush()
	if err != nil {
		fmt.Println("将缓冲区数据写入文件失败", err)
		return
	}

}
