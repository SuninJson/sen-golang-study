package _6_sync_lock

import (
	"fmt"
	"math/rand"
	"sync"
)

// SyncOnce 若需要保证多个并发goroutine中，某段代码仅仅执行一次，就可以使用 sync.Once 结构实现。
// 例如，在获取配置的时候，往往仅仅需要获取一次，然后去使用。
// 在多个goroutine并发时，要保证能够获取到配置，同时仅获取一次配置，就可以使用sync.Once结构：
func SyncOnce() {
	// 初始化config变量
	config := make(map[string]string)

	// 1. 初始化 sync.Once
	once := sync.Once{}

	// 加载配置的函数
	loadConfig := func() {
		// 2. 利用 once.Do() 来执行
		once.Do(func() {
			// 保证执行一次
			config = map[string]string{
				"varInt": fmt.Sprintf("%d", rand.Int31()),
			}
			fmt.Println("config loaded")
		})
	}

	// 模拟多个goroutine，多次调用加载配置
	// 测试加载配置操作，执行了几次
	workers := 10
	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			// 并发的多次加载配置
			loadConfig()
			// 使用配置
			_ = config

		}()
	}
	wg.Wait()
}
