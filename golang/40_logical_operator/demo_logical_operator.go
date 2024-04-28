package main

import "fmt"

// 【1】关系运算符:==,!=,>,<,> =，<=
//
//	关系运算符的结果都是bool型，也就是要么是true，要么是false
//
// 【2】关系表达式经常用在流程控制中
func main() {
	//与逻辑：&& :两个数值/表达式只要有一侧是false，结果一定为false
	//也叫短路与：只要第一个数值/表达式的结果是false，那么后面的表达式等就不用运算了，直接结果就是false  -->提高运算效率
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	//或逻辑：||:两个数值/表达式只要有一侧是true，结果一定为true
	//也叫短路或：只要第一个数值/表达式的结果是true，那么后面的表达式等就不用运算了，直接结果就是true -->提高运算效率
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	//非逻辑：取相反的结果：
	fmt.Println(!true)
	fmt.Println(!false)
}
