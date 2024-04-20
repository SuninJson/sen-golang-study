package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := "Hello Golang"

	println("统计字符串的长度,按字节进行统计：")
	println("字符串长度：", len(str1))

	println("\n 字符串遍历方式1，for-range键值循环：")
	for i, value := range str1 {
		fmt.Printf("索引：%d，值：%c \n", i, value)
	}

	println("\n 字符串遍历方式2，转换为切片后循环：")
	strSlice := []rune(str1)
	for i := 0; i < len(strSlice); i++ {
		fmt.Printf("索引：%d，值：%c \n", i, strSlice[i])
	}

	println("\n 字符串转整数")
	n1, err := strconv.Atoi("66")
	println(n1, err)

	n2, err := strconv.Atoi("一")
	println(n2, err.Error())

	println("\n 整数转字符串")
	str1 = strconv.Itoa(6887)
	println(str1)

	println("\n 查找子串是否在指定的字符串中")
	isContain := strings.Contains("javaandgolang", "go")
	println(isContain)

	println("\n 统计一个字符串有几个指定的子串")
	subStrCount := strings.Count("go go go", "go")
	println(subStrCount)

	println("\n 字符串比较")
	testIsEqualS1 := "go"
	testIsEqualS2 := "Go"
	isEqual := testIsEqualS1 == testIsEqualS2
	fmt.Printf("区分大小写的字符串比较，str1 = %s，str2 = %s，isEqual = %t \n",
		testIsEqualS1, testIsEqualS2, isEqual)

	isEqual = strings.EqualFold(testIsEqualS1, testIsEqualS2)
	fmt.Printf("不区分大小写的字符串比较，str1 = %s，str2 = %s，isEqual = %t \n",
		testIsEqualS1, testIsEqualS2, isEqual)

	println("\n 返回子串在字符串第一次出现的索引值，如果没有返回-1 ")
	println(strings.Index("go go go", "go"))

	println("\n 字符串的替换")
	//最后一个参数为-1时表示全部替换，正数表示替换多少个
	testReplace1 := strings.Replace("go and java go go", "go", "golang", -1)
	testReplace2 := strings.Replace("go and java go go", "go", "golang", 1)
	println(testReplace1)
	println(testReplace2)

	println("\n 按照指定的某个字符，为分割标识，将一个学符串拆分成字符串数组")
	testSplitArr := strings.Split("go-python-java", "-")
	for _, value := range testSplitArr {
		println(value)
	}

	println("\n 将字符串的字母进行大小写的转换")
	println(strings.ToLower("Go"))
	println(strings.ToUpper("go"))

	println("\n 将字符串左右两边的空格去掉")
	println(strings.TrimSpace("     go and java    "))

	println("\n 将字符串左右两边指定的字符去掉")
	println(strings.Trim("~golang~ ", " ~"))

	println("\n 将字符串左边指定的字符去掉")
	println(strings.Trim("~golang~ ", " ~"))

	println("\n 将字符串右边指定的字符去掉")
	println(strings.TrimRight("~golang~", "~"))

	println("\n 判断字符串是否以指定的字符串开头")
	println(strings.HasPrefix("http://java.sun.com/jsp/jstl/fmt", "http"))

	println("\n 判断字符串是否以指定的字符串结束")
	println(strings.HasSuffix("demo.png", ".png"))

}
