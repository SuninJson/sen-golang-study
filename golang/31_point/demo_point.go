package main

import "fmt"

func main() {
	var age int = 18
	//&符号+变量 就可以获取这个变量内存的地址
	fmt.Println("age变量的内存地址：", &age) //0xc0000a2058

	//定义一个指针变量：
	//var代表要声明一个变量
	//ptr 指针变量的名字
	//ptr对应的类型是：*int 是一个指针类型 （可以理解为 指向int类型的指针）
	var ptr *int = &age
	fmt.Println("ptr的值：", ptr)
	fmt.Println("ptr本身这个存储空间的地址为：", &ptr)

	//想获取ptr这个指针或者这个地址指向的那个数据，使用 * 符号来根据地址取值
	fmt.Printf("ptr指向的数值为：%v", *ptr) //ptr指向的数值为：18

	//通过指针改变指向值
	*ptr = 20
	fmt.Println("通过指针改变指向值后，age变量的值：", age)

	//指针变量所设定的值必须为地址值 下面语句会提示Cannot use 'age' (type int) as the type *int
	//ptr = age

	//指针变量的类型与所指向的值类型需要匹配 下面语句会提示Cannot use '&age' (type *int) as the type *float32
	//var floatP1 *float32 = &age
}
