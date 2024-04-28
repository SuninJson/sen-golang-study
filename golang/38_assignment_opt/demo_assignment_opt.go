package main

import "fmt"

func main() {
	var num1 int = 10
	fmt.Println(num1)
	//=右侧的值运算清楚后，再赋值给=的左侧
	var num2 int = (10+20)%3 + 3 - 7
	fmt.Println(num2)

	//等价num3 = num3 + 20;
	var num3 int = 10
	num3 += 20
	fmt.Println(num3)

	// 交换
	var a int = 8
	var b int = 4
	fmt.Printf("交换前的值：a = %v,b = %v\n", a, b)
	//引入一个中间变量：
	var t int
	t = a
	a = b
	b = t
	fmt.Printf("交换后的值：a = %v,b = %v", a, b)
}
