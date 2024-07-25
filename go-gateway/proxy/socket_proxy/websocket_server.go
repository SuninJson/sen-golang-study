package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func main() {
	// 下面这行代码声明了一个名为 addr 的变量，类型为指向字符串的指针，并调用了 flag 包的 String 函数。
	// flag 包是 Golang 标准库中用于解析命令行参数的包。String 函数接受三个参数，分别是参数名称、默认值和描述。
	// 具体来说，这行代码的作用是定义了一个名为 addr 的命令行参数，当程序在命令行中启动时，可以通过 -addr 选项来指定某个地址。
	// 如果用户没有指定 -addr 选项，那么程序将会使用默认值 "localhost:2003"。
	var addr = flag.String("addr", "localhost:2003", "http_proxy service address")
	flag.Parse()

	// log.SetFlags(0)，可以清除掉所有的默认标志。
	// 通常情况下，log 包在输出日志信息时会默认添加时间、日期、文件名等额外的标志信息，以便更好地标识日志的来源和时间。
	log.SetFlags(0)

	http.HandleFunc("/handleWebSocket", handleWebSocket)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func handleWebSocket(writer http.ResponseWriter, request *http.Request) {
	// 将HTTP协议升级为WebSocket协议
	upGrader := websocket.Upgrader{}
	conn, err := upGrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println("upgrade err:", err)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("Failed to close connection:", err)
			return
		}
	}(conn)

	go func() {
		// 开启一个携程每隔一秒服务端检测客户端心跳
		for {
			if err := conn.WriteMessage(1, []byte("heartbeat")); err != nil {
				log.Println("The heartbeat detection failed. The WebSocket processing ended")
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// 接收并回复消息
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message.", err)
			break
		}
		log.Println("Received message:", msg)
		newMsg := string(msg) + " Hello,This is replay."
		msg = []byte(newMsg)
		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Println("Failed to write message.", err)
			break
		}
	}
}
