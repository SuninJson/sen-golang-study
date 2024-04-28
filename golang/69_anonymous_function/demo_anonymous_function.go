package main

import "fmt"

func main() {
	fmt.Println("定义匿名函数")
	sub := func(n1 int, n2 int) int {
		return n1 - n2
	}
	fmt.Println("匿名函数sub调用的结果：", sub(1, 2))

	fmt.Println("\n 定义匿名函数同时进行调用")
	result := func(n1 int, n2 int) int {
		return n1 + n2
	}(1, 2)

	fmt.Println("定义匿名函数同时进行调用的结果：", result)
}
