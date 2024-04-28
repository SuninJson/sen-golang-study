package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 读写锁
var lock sync.RWMutex

func main() {
	wg.Add(6)
	println("读多写少的场景下使用读写锁，锁在写时生效，不影响并发的读")
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			read()
		}()
	}

	go func() {
		defer wg.Done()
		write()
	}()
	wg.Wait()
}

func read() {
	lock.RLock()
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功")
	lock.RUnlock()
}

func write() {
	lock.Lock()
	fmt.Println("开始修改数据")
	time.Sleep(time.Second * 10)
	fmt.Println("修改数据成功")
	lock.Unlock()
}
