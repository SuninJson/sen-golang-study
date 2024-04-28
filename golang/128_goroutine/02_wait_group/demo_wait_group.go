package main

import (
	"fmt"
	"sync"
)

// sync.WaitGroup 用于等待一组协程的结束。
var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		// 父线程调用Add方法来设定应等待的协程的数量
		wg.Add(1)
		go func(n int) {
			// 协程执行结束时，调用Done方法减少WaitGroup计数器的值
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}

	// Wait方法等待阻塞直到WaitGroup计数器为零，从而解决主线程退出导致即使协程还没有执行完毕，也会退出的问题
	wg.Wait()
}
