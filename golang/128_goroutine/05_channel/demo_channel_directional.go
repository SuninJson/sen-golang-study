package main

import "sync"

// ChannelDirectional 单向Channel，指的是仅支持接收或仅支持发送操作的Channel。语法上：
// - `chan<- T` 仅发送Channel
// - `<-chan T` 仅接收Channel
// 单向Channel的意义在于约束Channel的使用方式。
// 仅使用单向Channel通常没有实际意义，单向Channel最典型的使用方式是：
// 使用单向通道约束双向通道的操作。
// 语法上来说，就是我们会将双向Channel转换为单向Channel来使用。典型使用在函数参数或返回值类型中。
func ChannelDirectional() {
	// 初始化数据
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	// send and receive
	wg.Add(2)
	go setElement(ch, 42, wg)
	go getElement(ch, wg)

	wg.Wait()
}

// only receive channel
func getElement(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	println("received from ch, element is ", <-ch)
}

// only send channel
func setElement(ch chan<- int, v int, wg *sync.WaitGroup) {
	defer wg.Done()

	ch <- v
	println("send to ch, element is ", v)
}
