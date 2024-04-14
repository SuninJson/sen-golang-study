package main

import (
	"fmt"
	"unsafe"
)

func main() {
	n1 := 3

	// 查看变量类型
	fmt.Printf("n1的数据类型：%T \n", n1)

	// 查看变量所占用字节
	fmt.Println("n1所占用字节：", unsafe.Sizeof(n1))

	// int类型转换
	// golang中无法进行自动转换，比如显式转换
	//var n2 float32 = n1
	fmt.Println(n1)
	var n2 float32 = float32(n1)
	fmt.Println(n2)
	//注意：n1的类型其实还是int类型，只是将n1的值100转为了float32而已，n1还是int的类型
	fmt.Printf("%T\n", n1) //int

	//将int64转为int8的时候，编译不会出错的，但是会数据的溢出
	var n3 int64 = 888888
	var n4 int8 = int8(n3)
	fmt.Println(n4) //56
	var n5 int32 = 12
	var n6 int64 = int64(n5) + 30
	//一定要匹配=左右的数据类型
	fmt.Println(n5)
	fmt.Println(n6)
	var n7 int64 = 12
	//编译通过，但是结果可能会溢出
	var n8 int8 = int8(n7) + 127
	//编译不会通过
	//var n9 int8 = int8(n7) + 128
	fmt.Println(n8)
}
