package main

import "fmt"

func main() {
	println("通过数组实现的功能：给出五个学生的成绩，求出成绩的总和，平均数：")

	//通过 ”var 数组名 [数组大小]数据类型“ 来定义数组
	var scores [5]int

	fmt.Println("数组中元素的默认值为0：", scores)

	scores[0] = 95
	scores[1] = 91
	scores[2] = 39
	scores[3] = 60
	scores[4] = 21

	//求和：
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}

	//平均数：
	avg := sum / len(scores)

	fmt.Printf("成绩的总和为：%v,成绩的平均数为：%v\n", sum, avg)

	// fmt.Printf 中 %p 表示为十六进制，并加上前导的0x
	fmt.Printf("\n证明数组中存储的是地址值,scores的地址为：%p\n", &scores)
	fmt.Printf("scores第1个元素的地址为：%p\n", &scores[0])
	fmt.Printf("scores第2个元素的地址为：%p\n", &scores[1])
	fmt.Printf("scores第3个元素的地址为：%p\n", &scores[2])

	println("所以数组中数组的地址和第一个元素的地址相同，" +
		"int类型默认为int32,每个元素占4个字节，数组中各元素的地址是连续的")

	println("\n展示一下班级的每个学生的成绩：（数组进行遍历）")
	println("方式1：通过普通for循环对数组进行遍历")
	for i := 0; i < len(scores); i++ {
		fmt.Printf("第%d个学生的成绩为：%d\n", i+1, scores[i])
	}

	println("方式2：通过for-range循环对数组进行遍历")
	for key, value := range scores {
		fmt.Printf("第%d个学生的成绩为：%d\n", key+1, value)
	}

	println("\n初始化数组的方式")
	var arr1 = [3]int{1, 2, 3}
	fmt.Printf("通过指定数组的大小初始化数组，数组：%v，容量：%d，长度：%d\n",
		arr1, cap(arr1), len(arr1))
	var arr2 = [...]int{3, 4, 5, 6}
	fmt.Printf("不指定数组大小初始化数组，数组：%v，容量：%d，长度：%d\n",
		arr2, cap(arr2), len(arr2))
	var arr3 = [...]int{2: 9, 1: 8, 0: 7}
	fmt.Printf("指定数组中元素的下标初始化数组，数组：%v，容量：%d，长度：%d\n",
		arr3, cap(arr3), len(arr3))

	println("\nGo中数组属值类型，在默认情况下是值传递")
	fmt.Println("执行testPass函数前，数组为：", arr1)
	testPass(arr1)
	fmt.Println("执行testPass函数后，数组为：", arr1)

	println("\n若想在其它函数中，去修改原来的数组，可以使用引用传递(指针方式)。 ")
	fmt.Println("执行testInParam函数前，数组为：", arr1)
	// 使用 & 符号来取地址
	testPassByReference(&arr1)
	fmt.Println("执行testPass函数后，数组为：", arr1)

	println("\n二维数组")
	var twoDimensionArr [2][3]int8
	fmt.Println("二维数组的默认值：", twoDimensionArr)

	fmt.Printf("二维数组的内存地址：%p\n", &twoDimensionArr)
	fmt.Printf("二维数组第1个数组的内存地址：%p\n", &twoDimensionArr[0])
	fmt.Printf("二维数组第1个数组中第1个元素的内存地址：%p\n", &twoDimensionArr[0][0])
	fmt.Printf("二维数组第1个数组中第2个元素的内存地址：%p\n", &twoDimensionArr[0][1])
	fmt.Printf("二维数组第1个数组中第3个元素的内存地址：%p\n", &twoDimensionArr[0][2])
	fmt.Printf("二维数组第2个数组的内存地址：%p\n", &twoDimensionArr[1])
	fmt.Printf("二维数组第2个数组中第1个元素的内存地址：%p\n", &twoDimensionArr[1][0])
	fmt.Printf("二维数组第2个数组中第2个元素的内存地址：%p\n", &twoDimensionArr[1][1])
	fmt.Printf("二维数组第2个数组中第3个元素的内存地址：%p\n", &twoDimensionArr[1][2])

	println("\n对二维数组进行赋值")
	fmt.Println("赋值前的二维数组：", twoDimensionArr)
	twoDimensionArr[0][0] = 1
	twoDimensionArr[0][1] = 2
	twoDimensionArr[1][1] = 11
	fmt.Println("赋值后的二维数组：", twoDimensionArr)

	println("\n遍历二维数组")
	println("方式1：普通for循环")
	for i := 0; i < len(twoDimensionArr); i++ {
		for j := 0; j < len(twoDimensionArr[i]); j++ {
			fmt.Printf("arr[%v][%v]=%v\t", i, j, twoDimensionArr[i][j])
		}
		fmt.Println()
	}

	println("方式2：for range循环")
	for key, value := range twoDimensionArr {
		for k, v := range value {
			fmt.Printf("arr[%v][%v]=%v\t", key, k, v)
		}
		fmt.Println()
	}
}

func testPassByReference(arr *[3]int) {
	arr[0] = 11
}

func testPass(arr1 [3]int) {
	arr1[0] = 11
}
