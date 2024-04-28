package main

import (
	"fmt"
	"sync"
	"time"
)

// 【1】管道（channel）特质介绍：
// （1）管道本质就是一个数据结构-队列
// （2）数据是先进先出
// （3）自身线程安全，多协程访问时，不需要加锁，channel本身就是线程安全的
// （4）管道有类型的，一个string的管道只能存放string类型数据
func main() {
	// 管道的定义语法："var 变量名 chan 数据类型"
	var intChan chan int
	// 通过make初始化，第二个参数为管道的容量，存储的数据量不能超过管道的容量
	intChan = make(chan int, 3)
	fmt.Printf("验证管道是引用类型，intChan的值：%v\n", intChan)
	fmt.Printf("管道的实际长度：%v,管道的容量是：%v\n", len(intChan), cap(intChan))
	// 向管道存放数据
	intChan <- 1
	intChan <- 2
	intChan <- 3
	// 从管道读取数据
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan

	fmt.Println("从管道中取得的第一个数：", num1)
	fmt.Println("从管道中取得的第二个数：", num2)
	fmt.Println("从管道中取得的第三个数：", num3)

	// 注意：在没有使用协程的情况下，如果管道的数据已经全部取出，那么再取就会报错
	//num4 := <-intChan
	//fmt.Println("管道的数据已经全部取出的情况：", num4)

	println("\n可以使用内置函数close可以关闭管道，当管道关闭后，就不能再向管道写数据了，但是仍然可以从该管道读取数据")

	intChan <- 4

	close(intChan)

	// 若在关闭管道后再次写入数据会报错 panic: send on closed channel
	//intChan <- 5
	testCloseNum := <-intChan
	fmt.Println("关闭管道读取的数据：", testCloseNum)

	println("\n遍历管道")
	// 管道支持for-range的方式进行遍历
	testForRangeChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		testForRangeChan <- i
	}
	// 注意在遍历时，如果管道没有关闭，会报错 fatal error: all goroutines are asleep - deadlock!
	close(testForRangeChan)
	for v := range testForRangeChan {
		fmt.Println("遍历管道 value = ", v)
	}

	println("\n协程与管道协作")
	intChan = make(chan int, 10)
	wg.Add(2)

	// 开启一个协程向管道写入数据
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			intChan <- i
			fmt.Println("向管道写入数据为：", i)
			// 通过休眠，模拟协程阻塞的场景
			time.Sleep(time.Second)
		}
		close(intChan)
	}()

	// 开启一个协程从管道读取数据
	go func() {
		defer wg.Done()
		for v := range intChan {
			fmt.Println("从管道读取数据为：", v)
			// 通过休眠，模拟协程阻塞的场景
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}

var wg sync.WaitGroup
