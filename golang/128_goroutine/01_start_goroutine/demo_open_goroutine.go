package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"log"
	"runtime"
	"strconv"
	"time"
)

// 【1】程序(program)
// 是为完成特定任务、用某种语言编写的一组指令的集合,是一段静态的代码。 （程序是静态的）
// 【2】进程(process)
// 是程序的一次执行过程。正在运行的一个程序，进程作为资源分配的单位，在内存中会为每个进程分配不同的内存区域。
// 进程是动态的，有其生命周期：有它自身的产生、存在和消亡的过程
// 【3】线程(thread)
// 进程可进一步细化为线程， 是一个程序内部的一条执行路径。
// 若一个进程同一时间并行执行多个线程，就是支持多线程的。
// 【4】协程(goroutine)
// 又称为微线程，纤程，协程是一种用户态的轻量级线程
// 作用:在执行A函数的时候，可以随时中断，去执行B函数，然后中断继续执行A函数(可以自动切换)，
// 注意这一切换过程并不是函数调用（没有调用语句)，过程很像多线程，然而协程中只有一个线程在执行（协程的本质是个单线程）
// 对于单线程下，我们不可避免程序中出现io操作，
// 但如果我们能在自己的程序中(即用户程序级别，而非操作系统级别)
// 控制单线程下的多个任务能在一个任务遇到io阻塞时就将寄存器上下文和栈保存到某个其他地方，
// 然后切换到另外一个任务去计算。在任务切回来的时候，恢复先前保存的寄存器上下文和栈，
// 这样就保证了该线程能够最大限度地处于就绪态，即随时都可以被cpu执行的状态，
// 相当于我们在用户程序级别将自己的io操作最大限度地隐藏起来，
// 从而可以迷惑操作系统，让其看到：该线程好像是一直在计算，io比较少，
// 从而会更多的将cpu的执行权限分配给我们的线程
// （注意:线程是CPU控制的，而协程是程序自身控制的，属于程序级别的切换，操作系统完全感知不到，因而更加轻量级)
// 【5】Go语言实现的协程模型：GMP模型结构：
// Goroutine 就是Go语言实现的协程模型。其核心结构有三个，称为GMP，也叫GMP模型。分别是：
// G，Goroutine，我们使用关键字go调用的函数。存储于P的本地队列或者是全局队列中。协程由OS负责调度交由具体的CPU核心中执行
// M，Machine，就是Work Thread，就是传统意义的线程，用于执行Goroutine，G。只有在M与具体的P绑定后，才能执行P中的G。
// P，Processor，处理器，主要用于协调G和M之间的关系，存储需要执行的G队列，与特定的M绑定后，执行Go程序，也就是G。P的本地队列当前最多存储256个G
// 【6】Golang为什么可以支持百万级并发
// Golang之所以支持百万级的goroutine并发，核心是因为每个goroutine的初始栈内存为2KB，
// 例如16G内存的服务器，栈中没有其它变量占用内存的情况下可以开启300多万个的协程，
// 栈内存用于保持goroutine中的执行数据，例如局部变量等。相对来说，线程线程的栈内存通常为2MB。
// 除了比较小的初始栈内存外，goroutine的栈内存可扩容的，也就是说支持按需增大或缩小，一个goroutine最大的栈内存当前限制为1GB。
func main() {
	// 编写一个程序，完成如下功能:
	//（1）在主线程中，开启一个goroutine，该goroutine每隔1秒输出"hello goroutine"
	//（2）在主线程中也每隔一秒输出"hello main"，输出10次后，退出程序
	//（3）要求主线程和goroutine同时执行
	// 注意：如果主线程退出了，即使协程还没有执行完毕，也会退出

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println(strconv.Itoa(i), ":hello golang ")
			// 阻塞一秒
			time.Sleep(time.Second)
		}
	}()

	for i := 1; i <= 10; i++ {
		fmt.Println(strconv.Itoa(i), ":hello main ")
		time.Sleep(time.Second)
	}

}

// GoroutineAnts 通过协程池使用协程
func GoroutineAnts() {
	// 1. 统计当前存在的goroutine的数量
	go func() {
		for {
			fmt.Println("NumGoroutine:", runtime.NumGoroutine())
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// 2. 初始化协程池，goroutine pool
	size := 1024
	pool, err := ants.NewPool(size)
	if err != nil {
		log.Fatalln(err)
	}
	// 保证pool被关闭
	defer pool.Release()

	// 3. 利用 pool，调度需要并发的大量goroutine
	for {
		// 向pool中提交一个执行的goroutine
		err := pool.Submit(func() {
			time.Sleep(100 * time.Second)
		})
		if err != nil {
			log.Fatalln(err)
		}
	}
}
