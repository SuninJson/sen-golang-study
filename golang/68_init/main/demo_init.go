package main

import (
	"fmt"
	"sen-golang-study/golang/68_init"
)

var num = test()

func test() int {
	fmt.Println("全局变量的定义调用的函数被执行")
	return 1
}

// init函数：初始化函数，可以用来进行一些初始化的操作
// 每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被Go运行框架调用。
// 各函数的执行顺序：所依赖包中的init函数 -> 全局变量的定义 -> init函数 -> main函数
func init() {
	fmt.Println("init函数被执行")
}

func main() {
	fmt.Println("main函数被执行")
	fmt.Println("main文件中的测试变量值为：", num)
	fmt.Println("test_init中的测试变量值为：", test_init.test_init.TestVar)
}
