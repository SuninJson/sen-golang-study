package main

import "fmt"

// 【1】为什么要使用函数：
// 提高代码的复用型，减少代码的冗余，代码的维护性也提高了
// 【2】函数的定义：
// 为完成某一功能的程序指令(语句)的集合,称为函数。
func main() {
	fmt.Println("通过函数实现两数相加")
	//实际参数：实际传入的数据
	sum1 := sum(10, 20)
	fmt.Println("10 + 20 = ", sum1)

	var n1, n2 = 30, 50
	sum2 := sum(n1, n2)
	fmt.Printf("%d + %d = %d \n", n1, n2, sum2)

	//函数和函数是并列的关系，所以我们定义的函数不能写到main函数中

	sum3, sub1 := sumAndSub(n1, n2)
	fmt.Printf("%d + %d = %d,%d - %d = %d\n", n1, n2, sum3, n1, n2, sub1)

	// 如果有返回值不想接受，可是使用‘_’进行忽略
	sum4, _ := sumAndSub(n1, n2)
	fmt.Printf("%d + %d = %d \n", n1, n2, sum4)

	fmt.Println("\n 基本数据类型和数组默认都是值传递的，即进行值拷贝。在函数内修改，不会影响到原来的值。")
	var num1 int = 10
	var num2 int = 20
	fmt.Printf("交换前的两个数： num1 = %v,num2 = %v \n", num1, num2)
	exchangeNum(num1, num2)
	fmt.Printf("交换后的两个数： num1 = %v,num2 = %v \n", num1, num2)

	fmt.Println("\n 验证可变参数")
	validateVariableParameter(1, 2, 3, 4, 5)
	validateVariableParameter(1, 2, 3, 4, 5, 6, 7, 8)

	fmt.Println("\n Go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用")
	funcVar := exchangeNum
	fmt.Printf("funcVar的类型是：%T,test函数的类型是：%T \n", funcVar, exchangeNum)

	fmt.Println("\n 可通过函数类型的变量调用函数")
	funcVar(10, 20)

	fmt.Println("\n Go中，函数可以作为形参，并且调用")
	testFuncInParam(10, 20, funcVar)

	fmt.Println("\n 为了简化数据类型定义,Go支持自定义数据类型")
	type myInt int
	var myIntNum1 myInt = 30
	fmt.Println("myIntNum1:", myIntNum1)

	fmt.Println("\n 虽然我们定义了myInt是int的别名，但在Golang中编译识别时认为它们不是同一种数据类型")
	//下面的语句会提示 Cannot use 'myIntNum1' (type myInt) as the type int
	//var intNum2 int = myIntNum1

	fmt.Println("\n Golang支持对函数返回值命名，返回值顺序不用对应")
	sum5, sub5 := testReturnVarName(n1, n2)
	fmt.Printf("%d + %d = %d,%d - %d = %d\n", n1, n2, sum5, n1, n2, sub5)
}

func testReturnVarName(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func testFuncInParam(n1 int, n2 int, funcVar func(num1 int, num2 int)) {
	funcVar(n1, n2)
}

//	func   函数名（形参列表)（返回值类型列表）{
//	 		执行语句..
//			return + 返回值列表
//	}
//
// （1）函数名：
// 遵循标识符命名规范:见名知意 addNum,驼峰命名addNum
// 首字母不能是数字
// 首字母大写该函数可以被本包文件和其它包文件使用(类似public)
// 首学母小写只能被本包文件使用，其它包文件不能使用(类似private)
// （2）形参列表：
// 形参列表：个数：可以是一个参数，可以是n个参数，可以是0个参数
// 形式参数列表：作用：接收外来的数据
// （3）返回值类型列表：函数的返回值对应的类型应该写在这个列表中，可没有返回值，也可返回一个或多个
func sum(n1 int, n2 int) int {
	return n1 + n2
}

// 返回两数的和与差
func sumAndSub(n1 int, n2 int) (int, int) {
	return n1 + n2, n1 - n2
}

// 自定义函数：功能：交换两个数
func exchangeNum(num1 int, num2 int) {
	fmt.Println("exchangeNum function...")
	var t int
	t = num1
	num1 = num2
	num2 = t
}

//Golang中函数不支持重载 下面代码会提示 'exchangeNum' redeclared in this package
//func exchangeNum(num1 int) {
//	return
//}

// Golang中支持可变参数
func validateVariableParameter(args ...int) {
	fmt.Println("函数内部处理可变参数的时候，将可变参数当做切片来处理")
	for i := 0; i < len(args); i++ {
		if i == len(args)-1 {
			fmt.Println(args[i])
		} else {
			fmt.Printf("%d ", args[i])
		}
	}
}
