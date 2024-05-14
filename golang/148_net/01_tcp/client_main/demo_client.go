package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	StartClient()
}

func StartClient() {
	println("客户端启动。。。")
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8888", time.Second)
	if err != nil {
		fmt.Println("客户端连接失败：", err)
		return
	}
	fmt.Printf("连接成功，客户端地址：%v,服务端地址：%v\n", conn.LocalAddr(), conn.RemoteAddr())

	// 通过客户端发送单行数据
	// os.Stdin代表终端标准输入
	reader := bufio.NewReader(os.Stdin)
	terminalInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("终端输入失败，err:", err)
	}

	// 将终端输入的数据发送给服务器
	// （1）conn.Write函数特点
	//- 写成功, err ==nil && wn == len(data) 表示写入成功
	//- 写阻塞，当无法继续写时，Write会进入阻塞状态. 无法继续写,通常意味着TCP的窗口已满.
	//- 已关闭的连接不能继续写入
	//- 可以使用如下方法控制Write的超时时长
	//  - `SetDeadline(t time.Time) error`
	//  - `SetWriteDeadline(t time.Time) error`
	// （2）并发读写，指的是两方面：
	//- 读操作和写操作是并发执行的
	//- 可能出现多个Goroutine同时写或读
	// 因此在Go中，要使用Goroutine完成。
	// 同一个连接的并发读或写操作是Goroutine并发安全的。
	// 指的是同时存在多个Goroutine并发的读写，之间是不会相互影响的，这个在实操中，主要针对Write操作。
	// conn.Write()是通过锁来实现的。
	n, err := conn.Write([]byte(terminalInput))
	if err != nil && n == len(terminalInput) {
		fmt.Println("写入失败，err:", err)
	}
	fmt.Printf("终端数据通过客户端发送成功，一共发送了%d字节的数据,并退出", n)
}
