package main

import "fmt"

func main() {
	//定义浮点类型的数据：
	var num1 float32 = 3.14
	fmt.Println(num1)
	//可以表示正浮点数，也可以表示负的浮点数
	var num2 float32 = -3.14
	fmt.Println(num2)
	//浮点数可以用十进制表示形式，也可以用科学计数法表示形式  E 大写小写都可以的
	var num3 float32 = 314e-2
	fmt.Println(num3)
	var num4 float32 = 314e+2
	fmt.Println(num4)
	var num5 float32 = 314e+2
	fmt.Println(num5)
	var num6 float64 = 314e+2
	fmt.Println(num6)
	//浮点类型底层存储：符号位+指数位+尾数位，所以尾数位只是存了 一个大概，很可能会出现精度的损失
	//因浮点数可能会有精度的损失，所以通常情况下，建议使用：float64
	var num7 float32 = 256.000000916
	fmt.Println(num7)
	var num8 float64 = 256.000000916
	fmt.Println(num8)
	//golang中默认的浮点类型为：float64
	var num9 = 3.17
	fmt.Printf("num9对应的默认的类型为：%T", num9)
}
