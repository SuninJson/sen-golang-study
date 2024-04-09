package main

import (
	"fmt"
	"unsafe"
)

func main() {
	n1 := 3

	// 查看变量类型
	fmt.Printf("n1的数据类型：%T \n", n1)

	// 查看变量所占用字节
	fmt.Println("n1所占用字节：", unsafe.Sizeof(n1))
}
