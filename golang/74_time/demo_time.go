package main

import (
	"fmt"
	"time"
)

func main() {
	println("获取当前时间")
	now := time.Now()
	fmt.Printf("%v 对应类型：%T", now, now)

	fmt.Printf("年：%v \n", now.Year())
	fmt.Printf("月：%v \n", int(now.Month()))
	fmt.Printf("日：%v \n", now.Day())
	fmt.Printf("时：%v \n", now.Hour())
	fmt.Printf("分：%v \n", now.Minute())
	fmt.Printf("秒：%v \n", now.Second())

	println("\n日期的格式化")
	println("	（1）将日期以年月日时分秒按照格式输出为字符串")
	println("		Printf将字符串直接输出")
	fmt.Printf("当前年月日： %d-%d-%d 时分秒：%d:%d:%d  \n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	println("		Sprintf可以得到这个字符串，以便后续使用")
	dateStr := fmt.Sprintf("当前年月日： %d-%d-%d 时分秒：%d:%d:%d  \n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(dateStr)

	println("	（2）按照指定格式输出为字符串")
	// now.Format中传入的数字是固定的
	dateStr2 := now.Format("2006/01/02 15/04/05")
	fmt.Println(dateStr2)
	dateStr3 := now.Format("2006 15:04")
	fmt.Println(dateStr3)
}
