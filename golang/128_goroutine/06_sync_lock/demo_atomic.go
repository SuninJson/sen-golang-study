package _6_sync_lock

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func SyncAtomicAdd() {
	// 并发的过程，没有加锁，Lock
	//var counter int32 = 0
	// type
	// atomic 原子的Int32, counter := 0
	counter := atomic.Int32{}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				//atomic.AddInt32(&counter, 1)
				// type
				// 原子累加操作 ， counter ++
				counter.Add(1)
			}
		}()
	}
	wg.Wait()
	//fmt.Println("counter:", atomic.LoadInt32(&counter))
	// type
	fmt.Println("counter:", counter.Load())
}

func SyncAtomicValue() {

	// 模拟加载配置，例如从文件中加载配置，返回解析后的配置信息
	var loadConfig = func() map[string]string {
		return map[string]string{
			// some config
			"title":   "Go并发编程",
			"varConf": fmt.Sprintf("%d", rand.Int31()),
		}
	}

	// config的操作应该是并发安全，选择原子操作进行时间，此外也可用锁的方式来实现
	var config atomic.Value

	// 每N秒加载一次配置文件
	go func() {
		for {
			config.Store(loadConfig())
			fmt.Println("latest config was loaded", time.Now().Format("15:04:05.99999999"))
			time.Sleep(time.Second)
		}
	}()

	// 使用配置
	// 不能在加载的过程中使用配置
	for {
		go func() {
			c := config.Load()
			fmt.Println(c, time.Now().Format("15:04:05.99999999"))
		}()

		time.Sleep(400 * time.Millisecond)
	}

	select {}
}
