package main

// 闭包就是一个函数和与其相关的引用环境组合的一个整体
// 闭包本质依旧是一个匿名函数，只是这个函数引入外界的变量/参数
// 匿名函数+引用的变量/参数 = 闭包
// 闭包中使用的变量/参数会一直保存在内存中，所以闭包不可滥用（对内存消耗大）
func getSum() func(int) int {
	//匿名函数中引用的那个变量会一直保存在内存中，可以一直使用
	var sum int
	return func(n int) int {
		sum += n
		return sum
	}
}

func main() {
	f := getSum()
	println(f(1))
	println(f(2))
	println(f(3))
}
