package main

import (
	"fmt"
	"sync"
)

var totalNum int
var wg sync.WaitGroup
var lock sync.Mutex

func main() {
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
