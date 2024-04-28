package main

import (
	"fmt"
	"sync"
)

// 通过refer+recover捕获panic进行处理，防止协程出现问题，导致主线程受到影响
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		makeError()
	}()

	go func() {
		defer wg.Done()
		printNum()
	}()

	wg.Wait()
}

func printNum() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func makeError() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("makeError()出现错误：", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}
