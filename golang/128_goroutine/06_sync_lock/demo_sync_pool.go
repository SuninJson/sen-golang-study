package _6_sync_lock

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Connection struct {
	IP   string
	Port string
}

// SyncPool 池是一组可以单独保存和检索的可以复用的临时对象。
// 存储在池中的任何项目可随时自动删除，无需通知。
// 一个池可以安全地同时被多个goroutine使用。
// 典型特征：
// - sync.Pool 是并发安全的
// - 池中的对象由Go负责删除，内存由Go自己回收
// - 池中元素的数量由Go负责管理，用户无法干预
// - 池中元素应该是临时的，不应该是持久的。例如长连接不适合放入 sync.Pool 中
// 池的目的是缓存已分配但未使用的项目以供以后重用，从而减轻垃圾收集器的压力。
// 也就是说，它使构建高效、线程安全的自由元素变得容易。
// 池的一个适当用途是管理一组临时项，这些临时项在包的并发独立客户端之间默默共享，并可能被其重用。
// 池提供了一种在许多客户机上分摊分配开销的方法。
// 一个很好地使用池的例子是fmt包，它维护了临时输出缓冲区的动态大小存储。
// 池由 sync.Pool类型实现，具体三个操作：
// - 初始化Pool实例，需要提供池中缓存元素的New方法。
// - 申请元素，func (p *Pool) Get() any
// - 交回对象，func (p *Pool) Put(x any)
// 代码示例：创建一个连接池，让1024个goroutine通过连接池使用连接
func SyncPool() {
	wg := sync.WaitGroup{}

	const IP = "127.0.0.1"
	var port int32 = 8080

	var connCount int32
	connectionNewer := func() any {
		atomic.AddInt32(&connCount, 1)
		conn := new(Connection)
		conn.IP = IP
		conn.Port = fmt.Sprintf("%d", port+connCount)
		return conn
	}

	pool := sync.Pool{
		New: connectionNewer,
	}

	for i := 0; i < 1024; i++ {
		go func() {
			defer wg.Done()
			wg.Add(1)
			conn := pool.Get().(*Connection)
			// 注意获取到的连接，在使用完成后一定要放回池子中
			defer pool.Put(conn)
			println("使用连接:", conn.IP, " ", conn.Port)

		}()
	}

	wg.Wait()

	println("使用的连接数:", atomic.LoadInt32(&connCount))
}
