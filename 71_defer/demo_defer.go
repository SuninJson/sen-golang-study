package main

import "fmt"

// 在函数中，我们经常需要创建资源，为了在函数执行完毕后，及时的释放资源，Golang的设计者提供了defer关键字
func main() {
	n1 := 1
	n2 := 2

	// defer关键字标记的语句不会被立刻执行，而是暂时将defer后的语句压入到栈中，先执行完后面的语句
	// 栈的特点是先进后出，所以下面的程序会先打印n2，然后再打印n1
	defer println("defer n1 = ", n1)
	defer println("defer n2 = ", n2)

	n1 = 10
	n2 = 20

	println("\n defer关键字，会将相关的值拷贝入栈中，不会随着函数后面的变化而变化")
	fmt.Printf("n1 = %d,n2 = %d, sum = %d \n", n1, n2, n1+n2)
}
