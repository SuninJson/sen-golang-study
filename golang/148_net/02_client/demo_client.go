package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	println("客户端启动。。。")
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端连接失败：", err)
		return
	}
	fmt.Println("连接成功，连接信息：", conn)

	// 通过客户端发送单行数据
	// os.Stdin代表终端标准输入
	reader := bufio.NewReader(os.Stdin)
	terminalInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("终端输入失败，err:", err)
	}

	//将终端输入的数据发送给服务器
	n, err := conn.Write([]byte(terminalInput))
	if err != nil {
		fmt.Println("连接失败，err:", err)
	}
	fmt.Printf("终端数据通过客户端发送成功，一共发送了%d字节的数据,并退出", n)
}
