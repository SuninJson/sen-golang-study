package main

import "fmt"

// 【1】关系运算符:==,!=,>,<,> =，<=
//
//	关系运算符的结果都是bool型，也就是要么是true，要么是false
//
// 【2】关系表达式经常用在流程控制中
func main() {
	fmt.Println(5 == 9) //判断左右两侧的值是否相等，相等返回true，不相等返回的是false， ==不是=
	fmt.Println(5 != 9) //判断不等于
	fmt.Println(5 > 9)
	fmt.Println(5 < 9)
	fmt.Println(5 >= 9)
	fmt.Println(5 <= 9)
}
