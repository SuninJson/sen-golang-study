package main

import "fmt"

func main() {
	//切片(slice)是golang中一种特有的数据类型
	//数组有特定的用处，但是数组长度固定不可变，
	//切片构建在数组之上并且提供了更加便捷的使用方式
	//所以切片比数组更为常用

	//切片(slice)是对数组一个连续片段的引用，所以切片是一个引用类型。
	//这个片段可以是整个数组，或者是由起始和终止索引标识的一些项的子集。
	//需要注意的是，终止索引标识的项不包括在切片内。
	//切片提供了数组的动态窗口

	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("数组：", arr)

	fmt.Println("通过 切片名 := [数组片段的起始索引:数组片段的终止索引] 来定义切片")
	// 下面代码中[1:3]标志从数组下标1开始到下标3结束，但不包含下标为3的元素
	slice := arr[1:3]
	fmt.Printf("切片：%v，切片的元素个数：%d，切片的容量是数组的起始元素到数组结束元素的个数：%d\n",
		slice, len(slice), cap(slice))
	println("\n切片定义的简写方式")
	//简写方式：
	//1) slice := arr[0:end]  ----> slice := arr[:end]
	abbrSlice1 := arr[:3]
	fmt.Printf("abbrSlice1 := arr[:3] 切片：%v，切片的元素个数：%d，切片的容量：%d\n",
		abbrSlice1, len(abbrSlice1), cap(abbrSlice1))
	//2) slice := arr[start:len(arr)]  ----> slice := arr[start:]
	abbrSlice2 := arr[1:]
	fmt.Printf("abbrSlice2 := arr[1:] 切片：%v，切片的元素个数：%d，切片的容量：%d\n",
		abbrSlice2, len(abbrSlice2), cap(abbrSlice2))
	//3) slice := arr[0:len(arr)]   ----> slice := arr[:]
	abbrSlice3 := arr[:]
	fmt.Printf("abbrSlice3 := arr[:] 切片：%v，切片的元素个数：%d，切片的容量：%d\n",
		abbrSlice3, len(abbrSlice3), cap(abbrSlice3))

	println("\n验证切片中的元素与相关数组中元素的内存地址一致")
	fmt.Println("arr[1]的内存地址：", &arr[1])
	fmt.Println("slice[0]的内存地址：", &slice[0])

	println("\n通过make内置函数来创建切片。基本语法: 切片名 := make([]数据类型, 长度, 容量)")
	testMakeSlice := make([]int, 4, 20)
	fmt.Printf("通过make函数创建的切片的默认值：%v，长度：%d，容量：%d\n",
		testMakeSlice, len(testMakeSlice), cap(testMakeSlice))
	// make函数会创建一个数组，但这个数组是对外不可见的，然后通过切片来间接来访问这个数组进行访问

	println("\n定一个切片，直接指定具体数组")
	slice2 := []int{11, 22, 33, 44, 55, 66}
	fmt.Printf("直接就指定具体数组的切片：%v，元素个数：%d，切片的容量：%d\n",
		slice2, len(slice2), cap(slice2))

	println("\n遍历切片")
	println("方式1：for循环常规方式遍历")
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v \n", i, slice[i])
	}
	println("方式2：for-range 结构遍历切片")
	for i, v := range slice {
		fmt.Printf("slice[%v] = %v \n", i, v)
	}

	println("\n切片使用不能越界")
	//fmt.Printf("超出切片长度的元素：%d\n", slice[len(slice)])

	println("\n切片可以继续切片")
	againSlice := slice2[1:2]
	fmt.Printf("againSlice := slice2[1:2] slice2:%v\n, againSlice:%v\n",
		slice2, againSlice)

	println("\n切片可以动态增长")
	//底层原理：
	//1.底层追加元素的时候对数组进行扩容，老数组扩容为新数组：
	//2.创建一个新数组，将老数组中的4,7,3复制到新数组中，在新数组中追加88,50
	//3.slice2 底层数组的指向 指向的是新数组
	//4.往往我们在使用追加的时候其实想要做的效果给slice追加
	appendSlice := append(slice, 88, 50)
	fmt.Println("slice:", slice)
	fmt.Println("appendSlice := append(slice, 88, 50) \nappendSlice:", appendSlice)
	fmt.Println("slice[0]的内存地址：", &slice[0])
	fmt.Println("appendSlice[0]的内存地址：", &appendSlice[0])

	println("\n可以通过append函数将切片追加给切片")
	fmt.Println("slice:", slice)
	fmt.Println("slice2:", slice2)
	slice = append(slice, slice2...)
	fmt.Println("追加之后的slice：", slice)

	println("\n切片的拷贝")
	a := []int{1, 4, 7, 3, 6, 9}
	b := make([]int, 10)
	fmt.Println("a slice:", a)
	fmt.Println("b slice:", b)
	copy(b, a)
	fmt.Println("将a中对应数组中元素内容复制到b中对应的数组中，b slice:", b)
}
