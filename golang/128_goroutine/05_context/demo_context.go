package _5_context

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

func ContextCancelCall() {
	// 1. 创建cancelContext
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(4)
	// 2. 启动goroutine，携带cancelCtx
	for i := 0; i < 4; i++ {
		// 启动goroutine，携带ctx参数
		go func(c context.Context, n int) {
			defer wg.Done()
			// 监听context的取消完成channel，来确定是否执行了主动cancel操作
			for {
				select {
				// 等待接收c.Done()这个channel
				case <-c.Done():
					fmt.Println("Cancel")
					return
				default:

				}
				fmt.Println(strings.Repeat("  ", n), n)
				time.Sleep(300 * time.Millisecond)
			}
		}(ctx, i)
	}

	// 3. 主动取消 cancel()
	// 3s后取消
	select {
	case <-time.NewTimer(2 * time.Second).C:
		cancel() // ctx.Done() <- struct{}
	}

	select {
	case <-ctx.Done():
		fmt.Println("main Cancel")
	}

	wg.Wait()

}

// ContextCancelDeep 当父上下文被取消时，子上下文也会被取消。Context 结构如下：
// ctxOne
//
//	|    \
//
// ctxTwo    ctxThree
//
//	|
//
// ctxFour
func ContextCancelDeep() {
	ctxOne, cancel := context.WithCancel(context.Background())
	ctxTwo, _ := context.WithCancel(ctxOne)
	ctxThree, _ := context.WithCancel(ctxOne)
	ctxFour, _ := context.WithCancel(ctxTwo)

	// 带有timeout的cancel
	//ctxOne, _ := context.WithTimeout(context.Background(), 1*time.Second)
	//ctxTwo, cancel := context.WithTimeout(ctxOne, 1*time.Second)
	//ctxThree, _ := context.WithTimeout(ctxOne, 1*time.Second)
	//ctxFour, _ := context.WithTimeout(ctxTwo, 1*time.Second)

	cancel()
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		select {
		case <-ctxOne.Done():
			fmt.Println("one cancel")
		}
	}()
	go func() {
		defer wg.Done()
		select {
		case <-ctxTwo.Done():
			fmt.Println("two cancel")
		}
	}()
	go func() {
		defer wg.Done()
		select {
		case <-ctxThree.Done():
			fmt.Println("three cancel")
		}
	}()
	go func() {
		defer wg.Done()
		select {
		case <-ctxFour.Done():
			fmt.Println("four cancel")
		}
	}()
	wg.Wait()
}

type MyContextKey string

func ContextValue() {
	wg := sync.WaitGroup{}

	ctx := context.WithValue(context.Background(), MyContextKey("title"), "Go")

	wg.Add(1)
	go func(c context.Context) {
		defer wg.Done()
		if v := c.Value(MyContextKey("title")); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", MyContextKey("title"))
	}(ctx)

	wg.Wait()
}

// ContextValueDeep 如果 context.valueCtx.Value 方法查询的 key 不存在于当前 valueCtx 中，
// 就会从父上下文中查找该键对应的值直到某个父上下文中返回 `nil` 或者查找到对应的值。
// 例如
func ContextValueDeep() {
	wgOne := sync.WaitGroup{}

	ctxOne := context.WithValue(context.Background(), MyContextKey("title"), "One")
	//ctxOne := context.WithValue(context.Background(), MyContextKey("key"), "Value")
	//ctxTwo := context.WithValue(ctxOne, MyContextKey("title"), "Two")
	ctxTwo := context.WithValue(ctxOne, MyContextKey("key"), "Value")
	//ctxThree := context.WithValue(ctxTwo, MyContextKey("title"), "Three")
	ctxThree := context.WithValue(ctxTwo, MyContextKey("key"), "Value")

	wgOne.Add(1)
	go func(c context.Context) {
		defer wgOne.Done()
		if v := c.Value(MyContextKey("title")); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", MyContextKey("title"))
	}(ctxThree)

	wgOne.Wait()
}
