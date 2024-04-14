package main

import "fmt"

//【1】Golang的 goto 语句可以无条件地转移到程序中指定的行。
//【2】goto语句通常与条件语句配合使用。可用来实现条件转移.
//【3】在Go程序设计中一般不建议使用goto语句，以免造成程序流程的混乱。
func main() {
	fmt.Println("通过 return 关键字结束当前的函数")
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		if i == 14 {
			return
		}
	}
	fmt.Println("hello golang")
}
