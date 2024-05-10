package _6_sync_lock

import (
	"log"
	"sync"
)

// SyncCondition 示例代码：
// - 一个goroutine负责接收数据，完毕后，广播给处理数据的goroutine
// - 多个goroutine处理数据，在数据未处理完前，等待广播信号。信号来了，处理数据
func SyncCondition() {
	data := make([]int, 1024*1024)

	condition := sync.NewCond(&sync.Mutex{})

	// 开启一个协程来填充数据
	go func(cond *sync.Cond) {
		log.Println("开始准备数据")
		cond.L.Lock()
		defer cond.L.Unlock()
		for i := 0; i < 1024; i++ {
			data[i] = i
		}
		log.Println("数据准备完成，发送数据准备完成信号")
		cond.Broadcast() // 广播以唤醒等待的协程
	}(condition)

	// 开启多个协程来处理数据
	workerNum := 1024
	var wg sync.WaitGroup
	wg.Add(workerNum)
	for i := 0; i < workerNum; i++ {
		go func(partIndex int, cond *sync.Cond) {
			defer wg.Done()
			log.Println("协程 ", partIndex, "等待数据准备完成")
			// 注意Wait()操作，是会先解锁，等到广播信号后，再加锁。
			// 若解锁时，未加锁，会导致panic，因此，Wait()操作前，要加锁。
			cond.L.Lock()
			defer cond.L.Unlock()
			for len(data) < 1024 {
				// wait所在的goroutine要判定是否需要wait，所以wait要出现在条件中，
				// 因为goroutine调用的关系，不能保证wait在broadcast前面执行
				// wait要使用for进行条件判定，因为在wait返回后，条件不一定成立。
				// 因为Broadcast()操作可能其它协程被提前调用
				cond.Wait()
			}
			dataPart := data[partIndex : partIndex+1] // 分配数据部分
			log.Printf("协程 %v 开始处理数据, 数据：%v \n",
				partIndex, dataPart)
		}(i, condition)
	}

	wg.Wait()
}
