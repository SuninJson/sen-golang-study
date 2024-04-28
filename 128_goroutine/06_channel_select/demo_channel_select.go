package main

import (
	"fmt"
)

// 【1】select功能：解决多个管道的选择问题，也可以叫做多路复用，可以从多个管道中随机公平地选择一个来执行
// PS：case后面必须进行的是io操作，不能是等值，随机去选择一个io操作
// PS：default防止select被阻塞住，加入default
func main() {
	intChan := make(chan int, 1)
	go func() {
		//time.Sleep(time.Second * 2)
		intChan <- 1
	}()

	stringChan := make(chan string, 1)
	go func() {
		//time.Sleep(time.Second * 1)
		stringChan <- "a"
	}()

	select {
	case v := <-intChan:
		fmt.Println("intChan:", v)
	case v := <-stringChan:
		fmt.Println("stringChan:", v)
	default:
		fmt.Println("防止select被阻塞")
	}
}
