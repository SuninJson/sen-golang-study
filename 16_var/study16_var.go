package main

import "fmt"

// 变量的使用方式
// 一、全局变量：定义在函数外的变量
//  01：单行声明
var static01 = "static01"

//  02：一次性声明
var (
	static02 = 6.6
	static03 = 99
)

func main() {
	fmt.Println(static01)
	fmt.Println(static02)
	fmt.Println(static03)

	// 二、局部变量
	//  01：指定变量的类型，并赋值
	var num int32 = 10
	fmt.Println(num)

	//  02：指定变量类型，但不赋值，变量会使用默认值，int类型的默认值为0
	var num2 int
	var s1 string
	fmt.Println(num2)
	fmt.Println("s1的值:", s1, ";")

	//  03：若没有指定变量类型，Golang会根据变量的值推断类型
	var s2 = "Hello"
	fmt.Println(s2)

	//  04：省略var关键字
	s3 := "省略var关键字"
	fmt.Println(s3)

	//  05：声明多个变量
	var n3, n4, n5 int
	fmt.Println(n3, n4, n5)

	var n6, name, n7 = 10, "jack", 8.8
	fmt.Println(n6, name, n7)

	n8, s4 := 9.9, "s4"
	fmt.Println(n8, s4)
}
