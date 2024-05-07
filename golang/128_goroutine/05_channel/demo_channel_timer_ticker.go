package main

import (
	"log"
	"math/rand"
	"time"
)

// Timer&Ticker是Go标准包time中定义的类型，通过Channel与程序进行通信。
// 定时器time.Timer 类似于一次性闹钟
// 断续器time.Ticker类似于重复性闹钟，也成循环定时器
// 无论是一次性还是重复性计时器，都是通过Channel与应用程序交互的。我们通过监控Timer和Ticker返回的Channel，来确定是否到时的需求

func TimerA() {
	t := time.NewTimer(time.Second)
	println("Set the timer, \ttime is ", time.Now().String())

	now := <-t.C
	println("The time is up, time is ", now.String())
}

// TimerB 示例代码，简单的猜数字游戏
func TimerB() {
	guessChan := make(chan int)
	// 写channel
	go func() {
		defer close(guessChan)                          // 确保channel在结束时关闭，避免资源泄漏
		rand.New(rand.NewSource(time.Now().UnixNano())) // 正确初始化随机数生成器
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop() // 确保ticker在结束时停止，避免goroutine泄漏
		i := 1
		for now := range ticker.C {
			guessNum := rand.Intn(10)
			log.Printf("now is %s ,Guess %d: %d\n", now.String(), i, guessNum)
			guessChan <- guessNum
			i++
		}
	}()

	// 每局时间
	duration := 5 * time.Second
	var timer = time.NewTimer(duration)
	var hint, miss int
	// 共猜5局，每局持续5秒钟
	for i := 0; i < 5; i++ {
	receiveGuess:
		for {
			select {
			case v := <-guessChan:
				log.Printf("Receive receiveGuess value: %d\n", v)
				if v == 6 {
					log.Println("Bingo! some one hint the answer.")
					// 猜中后开始新的一局游戏，重置定时器
					timer.Reset(duration)
					hint++
					break receiveGuess
				}
				miss++
			case <-timer.C:
				log.Println("The time is up, no one hint.")
				// 重新创建定时器
				timer = time.NewTimer(duration)
				break receiveGuess
			}
		}
	}
	log.Printf("Game Over! Hint %d, Miss %d\n", hint, miss)
}

// TickerA 模拟了一个心跳程序，间隔1秒，发送ping操作。整体到时，运行结束
func TickerA() {
	// 断续器
	ticker := time.NewTicker(time.Second)

	// 定时器
	timer := time.After(5 * time.Second)
loop: // 持续心跳
	for now := range ticker.C {
		println("now is ", now.String())
		// heart beat
		println("http.Get(\"/ping\")")

		// 非阻塞读timer，到时结束断续器
		select {
		case <-timer:
			ticker.Stop()
			break loop
		default:
		}
	}
}
