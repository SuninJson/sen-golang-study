package _6_sync_lock

import (
	"fmt"
	"sync"
	"time"
)

var totalNum int
var lock sync.Mutex

// 读写锁
var rwLock sync.RWMutex

func DemoMutex() {
	wg := sync.WaitGroup{}
	println("\n未加锁时，多个协程操作同一数据")
	wg.Add(2)
	go func() {
		defer wg.Done()
		add()
	}()
	go func() {
		defer wg.Done()
		sub()
	}()
	wg.Wait()
	// 在理论上，这个totalNum结果应该是0
	fmt.Println("未加锁时，多个协程操作同一数据的结果：", totalNum)

	totalNum = 0
	println("\n通过加入互斥锁来确保一个协程在执行时，另一个协程不执行")
	wg.Add(2)
	go func() {
		defer wg.Done()
		syncAdd()
	}()
	go func() {
		defer wg.Done()
		syncSub()
	}()
	wg.Wait()
	fmt.Println("加锁时，多个协程操作同一数据的结果：", totalNum)
}

func DemoRWMutex() {
	wg := sync.WaitGroup{}
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

func add() {
	for i := 0; i < 100000; i++ {
		totalNum = totalNum + 1
	}
}

func sub() {
	for i := 0; i < 100000; i++ {
		totalNum = totalNum - 1
	}
}

func syncAdd() {
	for i := 0; i < 100000; i++ {
		lock.Lock()
		totalNum = totalNum + 1
		lock.Unlock()
	}
}

func syncSub() {
	for i := 0; i < 100000; i++ {
		lock.Lock()
		totalNum = totalNum - 1
		lock.Unlock()
	}
}

func read() {
	rwLock.RLock()
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功")
	rwLock.RUnlock()
}

func write() {
	rwLock.Lock()
	fmt.Println("开始修改数据")
	time.Sleep(time.Second * 10)
	fmt.Println("修改数据成功")
	rwLock.Unlock()
}
