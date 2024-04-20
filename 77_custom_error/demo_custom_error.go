package main

import (
	"errors"
	"fmt"
)

func main() {
	err := testError(1, 0)
	if err != nil {
		fmt.Println("自定义错误：", err)
		// 若程序出现错误以后，希望程序中断，退出程序，可以通过内置函数builtin.panic实现
		panic(err)
	}
	println("上面的函数执行成功")
	println("正常执行之后的逻辑")
}

func testError(n1 int, n2 int) error {
	if n2 == 0 {
		return errors.New("除数不可为0")
	} else {
		result := n1 / n2
		println(result)
		return nil
	}
}
