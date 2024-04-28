package main

import "fmt"

func main() {
	fmt.Println("实现一个功能：求和： 1+2+3+4+5")
	var sum int = 0

	// 注意：for的初始表达式 不能用var定义变量的形式，要用:=
	// 注意：for循环实际就是让程序员写代码的效率高了，但是底层该怎么执行还是怎么执行的，底层效率没有提高，只是程序员写代码简洁了而已
	for i := 1; i <= 5; i++ {
		sum += i
	}

	fmt.Println("1+2+3+4+5 = ", sum)

	//for range 结构是Go语言特有的一种的迭代结构，在许多情况下都非常有用，
	//for range 可以遍历数组、切片、字符串、map 及通道，
	//for range 语法上类似于其它语言中的 foreach 语句

	fmt.Println("\n 对str进行遍历，遍历的每个结果的索引值被i接收，每个结果的具体数值被value接收")
	var str string = "hello golang你好"
	for i, value := range str {
		fmt.Printf("索引为：%d,具体的值为：%c \n", i, value)
	}

	fmt.Println("\n break可以结束单个正在执行的循环")
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			fmt.Printf("i: %v, j: %v \n", i, j)
			if i == 2 && j == 2 {
				break
			}
		}
	}
	fmt.Println("-----end")

	fmt.Println("\n break结合label 结束label标记的循环")
label2:
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			fmt.Printf("i: %v, j: %v \n", i, j)
			if i == 2 && j == 2 {
				break label2
			}
		}
	}
	fmt.Println("-----end")

	fmt.Println("\n 使用 continue 结束本次循环，继续下一次循环")
	fmt.Println("   -- 求100以内能被6整除的数")
	for i := 1; i <= 100; i++ {
		if i%6 != 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("\n continue的作用是结束离它近的那个循环，继续离它近的那个循环，通过双重循环来验证")
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			if i == 2 && j == 2 {
				fmt.Println("不打印 i: 2, j: 2")
				continue
			}
			fmt.Printf("i: %v, j: %v \n", i, j)
		}
	}

	fmt.Println("\n continue结合label，继续label标记的循环，通过双重循环来验证")
label:
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			if i == 2 && j == 2 {
				continue label
			}
			fmt.Printf("i: %v, j: %v \n", i, j)
		}
	}
}
