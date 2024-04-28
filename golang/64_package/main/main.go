// 1.通过package关键字对包进行声明，建议包的声明和所在文件夹同名
// 2.main包是程序的入口包，通常main函数会放在main包下
package main

// 3.包名是从 $GOPATH/src/后开始计算，使用 ‘/’进行路径分割
// 4.多个包建议使用 import() 一次性导入
// 5.同级别的源文件的包的声明必须一致 ，一个目录/包下不能有重复的函数
// 6.包名和文件夹的名字，可以不一样
// 7.包到底是什么：
// （1）在程序层面，所有使用相同  package 包名  的源文件组成的代码模块
// （2）在源文件层面就是一个文件夹
import (
	"fmt"
	utils2 "sen-golang-study/golang/64_package/utils"
)

func main() {
	fmt.Println("main函数")
	utils2.ExchangeNum(10, 20)
	utils2.GetConn()
}
