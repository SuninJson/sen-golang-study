package test_init

import "fmt"

var TestVar string

func init() {
	fmt.Println("test_init中的init函数被执行")
	TestVar = "Hello"
}
