package main

import (
	"sync"
	"time"
)

// ChannelSync 采用同步channel，使用两个goroutine完成发送和接收。每次发送和接收的时间间隔不同。我们分别打印发送和接收的值和时间
// 注意结果：
// - 两个协程发送和接收时间一致
// - 发送和接受的时间间隔以长的为准，可见发送和接收操作为同步操作
// 没有定义容量的channel，为同步channel，就像两个协程之间的管道，需要发送方和接受方都同时准备好，再同步的执行
func ChannelSync() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i
			println("Send ", i, ".\tNow:", time.Now().Format("15:04:05.999999999"))
			// 通过睡眠，模拟其它代码的执行时间
			time.Sleep(1 * time.Second)
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			println("Received ", v, ".\tNow:", time.Now().Format("15:04:05.999999999"))
			// 通过睡眠，模拟其它代码的执行时间
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Wait()
}

// ChannelASync 与同步channel一致，只是采用了容量为5的缓冲channel，使用两个goroutine完成发送和接收。每次发送和接收的时间间隔不同。我们分别打印发送和接收的值和时间。
// 注意结果：
// - 两个协程发送和接收时间不同
// - 发送方和接收方操作不会阻塞，可见发送和接收操作为异步操作
func ChannelASync() {
	// 初始化数据
	ch := make(chan int, 5)
	wg := sync.WaitGroup{}

	// 间隔发送
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i
			println("Send ", i, ".\tNow:", time.Now().Format("15:04:05.999999999"))
			// 间隔时间
			time.Sleep(1 * time.Second)
		}
		close(ch)
	}()

	// 间隔接收
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			println("Received ", v, ".\tNow:", time.Now().Format("15:04:05.999999999"))
			// 间隔时间，注意与send的间隔时间不同
			time.Sleep(3 * time.Second)
		}
	}()

	wg.Wait()
}
