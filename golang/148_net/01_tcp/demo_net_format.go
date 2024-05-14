package _1_tcp

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"net"
	"sync"
	"time"
)

// 在发送或接收消息时，需要对消息进行格式化处理，才能在应用程序中保证消息具有逻辑含义。
// 前面的例子，我们采用的是字符串传递消息，也是一种格式，但能够包含的数据字段有限。
// 典型编程时，我们会将两端处理好的数据，使用特定格式进行发送。典型的有两类：
//  - 文本编码，例如JSON，YAML，CSV等
//  - 二进制编码，例如GOB（Go Binary），Protocol Buffer等

const tcp = "tcp"

func TcpServerFormat() {
	// A. 基于某个地址建立监听
	// 服务端地址
	address := ":8888" // Any IP or version
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	// 关闭监听
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalln("Net listen err", err)
		}
	}(listener)
	log.Printf("%s server is listening on %s\n", tcp, listener.Addr())

	// B. 接受连接请求
	// 循环接受
	for {
		// 阻塞接受
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		// 处理连接，读写
		go handleConnFormat(conn)
	}
}

func handleConnFormat(conn net.Conn) {
	// 日志连接的远程地址（client addr）
	log.Printf("accept from %s\n", conn.RemoteAddr())
	// A.保证连接关闭
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)

	wg := sync.WaitGroup{}
	wg.Add(1)
	// 发送端，
	go serverWriteFormat(conn, &wg)
	wg.Wait()
}

func serverWriteFormat(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// 向客户端发送数据
		// 数据编码后发送

		// 创建需要传递的数据
		message := Message{
			ID:      uint(rand.Int()),
			Code:    "SERVER-STANDARD",
			Content: "message from server",
		}

		// 创建编码器
		// 1.JSON, 文本编码
		encoder := json.NewEncoder(conn)

		// 2.GOB, 二进制编码
		//encoder := gob.NewEncoder(conn)

		// 利用编码器进行编码
		// encode 成功后，会写入到conn，已经完成了conn.Write()
		if err := encoder.Encode(message); err != nil {
			log.Println(err)
			continue
		}
		log.Println("message was send")

		time.Sleep(1 * time.Second)
	}
}

func TcpClientFormat() {
	address := ":8888"
	conn, err := net.DialTimeout(tcp, address, 1*time.Second)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Printf("Client %v connect remote server %v success!", conn.LocalAddr(), conn.RemoteAddr())
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		clientReadFormat(conn)
	}()

	wg.Wait()
}

func clientReadFormat(conn net.Conn) {
	for {
		message := Message{}

		// 创建编码器
		// 1. JSON, 文本编码
		decoder := json.NewDecoder(conn)

		// 2. GOB, 二进制编码
		//decoder := gob.NewDecoder(conn)

		// 利用编码器进行编码
		// encode 成功后，会写入到conn，已经完成了conn.Write()
		if err := decoder.Decode(&message); err != nil {
			log.Println(err)
			continue
		}
		log.Printf("message was received from remote server %v,message : %v \n", conn.RemoteAddr(), message)
	}
}

// Message 自定义的消息结构类型
type Message struct {
	ID      uint   `json:"id,omitempty"`
	Code    string `json:"code,omitempty"`
	Content string `json:"content,omitempty"`
}
