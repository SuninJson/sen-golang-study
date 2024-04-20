package main

import "fmt"

func main() {
	// defer后通过在函数末尾加上"()"来对该匿名函数进行调用
	defer func() {
		// 调用recover内置函数，可以捕获错误
		err := recover()
		if err != nil {
			fmt.Println("err：", err)
		}
	}()

	n1 := 1
	n2 := 0

	n3 := n1 / n2
	fmt.Println(n3)
}
