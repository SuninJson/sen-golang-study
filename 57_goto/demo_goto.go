package main

import "fmt"

//【1】Golang的 goto 语句可以无条件地转移到程序中指定的行。
//【2】goto语句通常与条件语句配合使用。可用来实现条件转移.
//【3】在Go程序设计中一般不建议使用goto语句，以免造成程序流程的混乱。
func main() {
	fmt.Println("hello golang1")
	fmt.Println("hello golang2")
	if 1 == 1 {
		goto label1
	}
	fmt.Println("hello golang3")
	fmt.Println("hello golang4")
	fmt.Println("hello golang5")
	fmt.Println("hello golang6")
label1:
	fmt.Println("hello golang7")
	fmt.Println("hello golang8")
	fmt.Println("hello golang9")
}
