## 使用sync.WaitGroup实现协同调度

WaitGroup用于等待一组goroutine完成。等待思路是计数器方案：

- 调用等待goroutine时，调用Add()增加等待的goroutine的数量
- 当具体的goroutine运行结束后，Done()用来减去计数。
- 主goroutine可以使用Wait来阻塞，直到所有goroutine都完成（计数器归零）。

示例代码：

```go
func GoroutineWG() {
    // 1. 初始化 WaitGroup
    wg := sync.WaitGroup{}
    // 定义输出奇数的函数
    printOdd := func() {
        // 3.并发执行结束后，计数器-1
        defer wg.Done()
        for i := 1; i <= 10; i += 2 {
            fmt.Println(i)
            time.Sleep(100 * time.Millisecond)
        }
    }

    // 定义输出偶数的函数
    printEven := func() {
        // 3.并发执行结束后，计数器-1
        defer wg.Done()
        for i := 2; i <= 10; i += 2 {
            fmt.Println(i)
            time.Sleep(100 * time.Millisecond)
        }
    }
    // 在 main goroutine 中，开启新的goroutine
    // 并发调用
    // 2, 累加WG的计数器
    wg.Add(2)
    go printOdd()
    go printEven()

    // main goroutine 运行结束
    // 内部调用的goroutine也就结束
    // 4. 主goroutine等待
    wg.Wait()
    fmt.Println("after main wait")
}
```

**WaitGroup() 适用于主goroutine需要等待其他goroutine全部运行结束后，才结束的情况。不适用于，主goroutine需要结束，而通知其他goroutine结束的情景**。

如图：

![image.png](https://fynotefile.oss-cn-zhangjiakou.aliyuncs.com/fynote/fyfile/13080/1672136903072/4f5508ffdc3c43178925e8c36ed7f55a.png)

注意，不得复制WaitGroup。因为内部维护的计数器不能被意外修改。

可以同时存在多个goroutine进行等待。

### WaitGroup的基本实现原理

WaitGroup 结构：

```go
type WaitGroup struct {
    // 用于保证不会被拷贝
    noCopy noCopy
    // 当前状态，存储计数器，存储等待的goroutine
    state1 uint64
    state2 uint32
}
```

状态 32bit和64bit的计算机不同，以64bit为例：

- 高32 bits是计数器
- 低32 bits是等待者

Add() 和 Done() 是用来操作计数器，操作计数器的操作是原子操作，保证并发安全性。

Wait()操作，在计数器为0时，结束阻塞状态。

核心代码示例：

```go
func (wg *WaitGroup) Add(delta int) {
    // 原子操作，累加计数器
    state := atomic.AddUint64(statep, uint64(delta)<<32)
}

func (wg *WaitGroup) Done() {
    wg.Add(-1)
}

func (wg *WaitGroup) Wait() {

    for {
        state := atomic.LoadUint64(statep)
        v := int32(state >> 32)
        w := uint32(state)
        // 如果计数器为0，则不需要等待
        if v == 0 {
            // Counter is 0, no need to wait.
            if race.Enabled {
                race.Enable()
                race.Acquire(unsafe.Pointer(wg))
            }
            return
        }
        // Increment waiters count.
}
```

## 