package main

import (
	"fmt"
	"runtime"
	"time"
)

func GoroutineNumCtl() {
	// 单独使用一个协程来打印当前协程的数量
	go func() {
		for {
			fmt.Println("当前系统Goroutine的数量：", runtime.NumGoroutine())
			time.Sleep(10 * time.Second)
		}
	}()

	// 通过定义一个channel，来作为协程数量的计数器
	ch := make(chan struct{}, 1024)

	// 并发开启多个goroutine
	for {
		// 启动协程前，向channel中添加一个元素，来表示启动了一个协程，若channel满了，下行代码会阻塞
		ch <- struct{}{}
		go func() {
			println("协程", len(ch), "执行")
			// 通过Sleep模拟其它代码的执行所需时间
			time.Sleep(10 * time.Second)

			// 协程执行结束时，通过接受ch中的一个元素，来表示减少了一个协程的数量
			<-ch
		}()
	}
}
