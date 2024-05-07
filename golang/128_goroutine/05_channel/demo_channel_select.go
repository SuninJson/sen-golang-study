package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	demo_util "sen-golang-study/golang/util"
	"sync"
	"time"
)

// SelectStmt
// 【1】select功能：解决多个管道的选择问题，也可以叫做多路复用，可以从多个管道中随机公平地选择一个来执行
// PS：case后面必须进行的是io操作，不能是等值，随机去选择一个io操作
// PS：default防止select被阻塞住，加入default
// 【2】select语句的执行分为几个步骤：
// 1. 对于全部的 case，receive操作的channel操作数、send语句的channel和右表达式在进入select语句时只会基于源码顺序计算一次。计算结果是一组要从中接收或者发送到的channel，以及要发送的相应值。RecvStmt的左侧带有短变量声明或赋值的表达式尚未计算。
// 2. 如果一个或多个通信可以继续，通过伪随机数选择其中一个继续执行。否则，如果存在default case，选择该case。如果没有default case，select 语句会阻塞直到至少一个通信操作可以继续。
// 3. 除非选择了default case，那么相应的通信操作会被执行。
// 4. 如果选择的case是带有短变量声明或赋值的RecvStmt，左侧表达式会被计算，并分配接收的值（或多个值）。
// 5. 执行所选择的case的语句列表。
func SelectStmt() {
	// 声明需要的变量
	var a [4]int
	var c1, c2, c3, c4 = make(chan int), make(chan int), make(chan int), make(chan int)
	var i1, i2 int

	// 用于操作channel的goroutine
	go func() {
		c1 <- 10
	}()
	go func() {
		<-c2
	}()
	go func() {
		close(c3)
	}()
	go func() {
		c4 <- 40
	}()

	// 用于select的goroutine
	go func() {
		select {
		case i1 = <-c1:
			println("received ", i1, " from c1")
		case c2 <- i2:
			println("sent ", i2, " to c2")
		case i3, ok := <-c3:
			if ok {
				println("received ", i3, " from c3")
			} else {
				println("c3 is closed")
			}
		case a[f()] = <-c4:
			println("received ", a[f()], " from c4")
		default:
			println("no communication")
		}
	}()

	// 简单sleep测试
	time.Sleep(100 * time.Millisecond)
}

func f() int {
	print("f() was run")
	return 2
}

// SelectFor select 匹配到可操作的case或者是defaultcase后，就执行完毕。
// 实操时，我们通常需要持续监听某些channel的操作，因此典型的select使用会配合for完成。
func SelectFor() {
	ch := make(chan int)
	// send to channel
	go func() {
		for {
			// 模拟演示数据来自随机数
			// 实操时，数据可以来自各种I/O，例如网络、缓存、数据库等
			ch <- rand.Intn(100)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	// select receive from channel
	go func() {
		for {
			select {
			case v := <-ch:
				println("received value: ", v)
			}
		}

	}()

	time.Sleep(3 * time.Second)
}

// SelectEmptyBlock 不存在任何case的select语句会阻塞
func SelectEmptyBlock() {
	// 空select阻塞
	println("before select")
	select {}
	println("after select")
}

// SelectNilChannelBlock 若select语句中只使用了nil的channel会阻塞
func SelectNilChannelBlock() {
	// nil select阻塞
	var ch chan int
	go func() {
		ch <- 1024
	}()
	println("before select")
	select {
	case <-ch:
	case ch <- 42:
	}
	println("after select")
}

func SelectNilChannel() {
	ch := make(chan int)
	// 写channel
	go func() {
		// 设置rand的随机种子
		rand.New(rand.NewSource(time.Now().Unix()))
		for {
			ch <- rand.Intn(10)
			time.Sleep(400 * time.Millisecond)
		}
	}()

	// 处理三秒后，停止处理
	go func() {
		sum := 0
		// 设置一个3秒的定时器
		timer := time.After(3 * time.Second)
		for {
			select {
			case v := <-ch:
				println("received value: ", v)
				sum += v
			case <-timer:
				// select会在定时器到时后，读取到定时器的内容，此时将channel设置为nil，不再读写
				ch = nil
				println("ch was set nil, sum is ", sum)
			}
		}

	}()

	// sleep 5 秒
	time.Sleep(5 * time.Second)
}

// SelectNonBlock 当select语句存在default case时：
// - 若没有可操作的channel，会执行default case
// - 若有可操作的channel，会执行对应的case
// 这样select语句不会进入block状态，称之为非阻塞（non-block）的收发（channel 的接收和发送）。
// 示例：多人猜数字游戏，我们在乎是否有人猜中数字：
func SelectNonBlock() {
	// 初始化数据
	counter := 10 // 参与人数
	maxNum := 20  // [0, 19] // 最大范围
	rand.New(rand.NewSource(time.Now().UnixMilli()))
	answer := rand.Intn(maxNum) // 随机答案
	println("The answer is ", answer)
	println("------------------------------")

	// 定义猜中了正确答案的channel
	bingoCh := make(chan int, counter)
	// wg
	wg := sync.WaitGroup{}
	wg.Add(counter)
	for i := 0; i < counter; i++ {
		// 每个goroutine代表一个猜数字的人
		go func() {
			defer wg.Done()
			result := rand.Intn(maxNum)
			println("someone guess ", result)
			// 若猜对了答案，将答案放入bingo channel中
			if result == answer {
				bingoCh <- result
			}
		}()
	}
	wg.Wait()

	println("------------------------------")
	// 是否有人发送了正确结果，可以是0或多个人
	select {
	case result := <-bingoCh:
		println("some one hint the answer ", result)
	default:
		println("no one hint the answer")
	}

	// 特别的情况是存在两个case，其中一个是default，
	// 另一个是channel case，那么go的优化器会优化内部这个select。
	// 内部会以if结构完成处理。因为这种情况，不用考虑随机性的问题
}

type Rows struct {
	Index int
}

// SelectRace Race模式，典型的并发执行模式之一，多路同时操作资源，哪路先操作成功，优先使用，同时放弃其他路的等待。
// 简而言之，从多个操作中选择一个最快的。核心工作：
// - 选择最快的
// - 停止其他未完成的
// 示例代码，示例从多个查询器同时读取数据，使用最先反返回结果的，其他查询器结束
func SelectRace() {
	const queryCoroutineNum = 8
	resultCh := make(chan Rows, 1)
	stopChs := [queryCoroutineNum]chan struct{}{}
	for i := range stopChs {
		stopChs[i] = make(chan struct{})
	}

	wg := sync.WaitGroup{}
	wg.Add(queryCoroutineNum)
	rand.New(rand.NewSource(time.Now().UnixMilli()))
	for i := 0; i < queryCoroutineNum; i++ {
		go func(index int) {
			defer wg.Done()
			//模拟查询的执行时间
			randTime := rand.Intn(1000)
			println("Query coroutine", index, "start query data,need duration is ", randTime, " ms")

			// 查询结果的channel
			tempResultCh := make(chan Rows, 1)

			// 执行查询工作
			go func() {
				// 模拟时长
				time.Sleep(time.Duration(randTime) * time.Millisecond)
				tempResultCh <- Rows{
					Index: index,
				}
			}()

			// 监听查询结果和停止信号channel
			select {
			// 查询结果
			case rows := <-tempResultCh:
				println("Query coroutine ", index, " get result.")
				// 保证没有其他结果写入，才写入结果
				if len(resultCh) == 0 {
					resultCh <- rows
				}
			// stop信号
			case <-stopChs[index]:
				println("Query coroutine ", index, " is stopping.")
				return
			}
		}(i)
	}

	// 等待第一个查询结果的反馈
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 等待ch中传递的结果
		select {
		// 等待第一个查询结果
		case rows := <-resultCh:
			println("get first result from ", rows.Index, ". stop other query coroutine.")
			// 通知全部查询协程结束
			for i := range stopChs {
				// 当前返回结果的goroutine不需要了，因为已经结束
				if i == rows.Index {
					continue
				}
				stopChs[i] <- struct{}{}
			}

		// 计划一个超时时间
		case <-time.After(5 * time.Second):
			println("all query coroutine timeout.")
			// 所有协程执行超时，通知全部查询协程结束
			for i := range stopChs {
				stopChs[i] <- struct{}{}
			}
		}
	}()

	wg.Wait()
}

// SelectAll Race模式是多个Goroutine获取相同的结果，优先使用快速响应的。
// 而All模式是多个Goroutine分别获取结果的各个部分，全部获取完毕后，组合成完整的数据，要保证全部的Goroutine都响应后，继续执行。
// 示例代码，核心逻辑：
// - 视频内容，分为三个goroutine分别处理subject、tags、file三个部分
// - 3个goroutine要全部执行完毕，数据才会整体获取
// - 不会一直等待，设置超时时间，进行超时处理。
func SelectAll() {
	type Video struct {
		Subject string
		File    string
		Tags    []string
	}

	type VideoPart struct {
		part string
		Video
	}

	const (
		QuerySubject = "QuerySubject"
		QueryFile    = "QueryFile"
		QueryTags    = "QueryTags"
	)

	goroutineNames := []string{QuerySubject, QueryFile, QueryTags}

	// 定义存储video部份数据的channel
	videoPartCh := make(chan VideoPart, len(goroutineNames))

	timeoutTimer := time.After(time.Millisecond * 1000)
	timeoutStopChMap := make(map[string]chan struct{}, 3)
	for _, name := range goroutineNames {
		timeoutStopChMap[name] = make(chan struct{}, 1)
	}

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(goroutineNames))

	// 并发获取Video的数据
	for _, name := range goroutineNames {
		go func(goroutineName string) {
			defer waitGroup.Done()

			println("Start ", goroutineName, " goroutine")
			resultCh := make(chan VideoPart, 1)

			go func() {
				rand.New(rand.NewSource(time.Now().UnixMilli()))
				demo_util.MockExecuteByRandTime(goroutineName, 1000)
				videoPart := VideoPart{part: goroutineName}
				switch goroutineName {
				case QuerySubject:
					videoPart.Subject = "Goroutine select all demo"
				case QueryTags:
					videoPart.Tags = []string{"Golang", "Channel", "select"}
				case QueryFile:
					videoPart.File = "https://..."
				}
				resultCh <- videoPart
			}()

			select {
			case result := <-resultCh:
				println(goroutineName, " already has obtained the result!")
				videoPartCh <- result
			case <-timeoutStopChMap[goroutineName]:
				// 若监听到因超时需要停止goroutine的信号，结束goroutine
				println(goroutineName, " already has been stopped!")
				return
			}

		}(name)
	}

	// 接收并整合Video每个部份的数据
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		video := Video{}
		received := make(map[string]struct{}, len(goroutineNames))
		// 循环等待接收数据以及超时处理
	loopReceive:
		for {
			select {
			case <-timeoutTimer:
				println("Query timeout,Video is incomplete!")
				// 通知超时但未完成的Goroutine结束
				for goroutineName := range timeoutStopChMap {
					if _, exists := received[goroutineName]; !exists {
						timeoutStopChMap[goroutineName] <- struct{}{}
					}
				}
				close(videoPartCh)
				// 超时处理完成后，结束监听
				break loopReceive
			case videoPart := <-videoPartCh:
				switch videoPart.part {
				case QuerySubject:
					video.Subject = videoPart.Subject
					received[QuerySubject] = struct{}{}
				case QueryTags:
					video.Tags = videoPart.Tags
					received[QueryTags] = struct{}{}
				case QueryFile:
					video.File = videoPart.File
					received[QueryFile] = struct{}{}
				}
			}

			// 若received长度与协程数量一致，则表示全部协程已处理完毕，结束监听
			if len(received) == len(goroutineNames) {
				println("All goroutine has obtained the result. Video is complete")
				close(videoPartCh)
				break loopReceive
			}
		}
		fmt.Println("Video:", video)
	}()

	waitGroup.Wait()
}

// SelectChannelCloseSignal 无缓冲Channel+关闭作典型同步信号
// 基于无缓冲Channel是同步的以及closed 的channel是可以接收内容的
// 以上两点原因，经常使用关闭无缓冲channel的方案来作为信号传递使用。前提是，信号纯粹是信号，
// 没有其他含义，比如关闭时间等
func SelectChannelCloseSignal() {
	wg := sync.WaitGroup{}
	// 定义无缓冲channel
	// 作为一个终止信号使用（啥功能的信号都可以，信号本身不分功能）
	ch := make(chan struct{})

	// goroutine，用来close, 表示 发出信号
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("发出信号, close(ch)")
		close(ch)
	}()

	// goroutine，接收ch，表示接收信号
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 先正常处理，等待ch的信号到来
		for {
			select {
			case <-ch:
				fmt.Println("收到信号, <-ch")
				return
			default:

			}
			// 正常的业务逻辑
			fmt.Println("业务逻辑处理中....")
			time.Sleep(300 * time.Millisecond)
		}
	}()

	wg.Wait()
}

// SelectSignal signal.Notify 信号通知监控
// 系统信号也是通过channel与应用程序交互，
// 例如典型的 ctrl+c 中断程序， `os.Interrupt`，
// 若不监控系统信号，ctrl+c后程序会直接终止，
// 而如果监控了信号，那么可以在ctrl+c后，执行一系列的关闭处理，
// 例如：
func SelectSignal() {
	// 一：模拟一段长时间运行的goroutine
	go func() {
		for {
			fmt.Println(time.Now().Format(".15.04.05.000"))
			time.Sleep(300 * time.Millisecond)
		}
	}()

	// 要求主goroutine等待上面的goroutine，方案：
	// 1. wg.Wait()
	// 2. time.Sleep()
	// 3. select{}

	// 持久阻塞
	//select {}

	// 二，监控系统的中断信号,interrupt
	// 1 创建channel，用于传递信号
	chSignal := make(chan os.Signal, 1)
	// 2 设置该channel可以监控哪些信号
	signal.Notify(chSignal, os.Interrupt)
	//signal.Notify(chSignal, os.Interrupt, os.Kill)
	//signal.Notify(chSignal) // 全部类型的信号都可以使用该channel
	// 3 监控channel
	select {
	case <-chSignal:
		fmt.Println("received os signal: Interrupt")
	}
}
