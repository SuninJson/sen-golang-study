package _1_tcp

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

// 在使用长连接时，通常需要使用规律性的发送数据包，以维持在线状态，称为心跳检测。
//一旦心跳检测不能正确响应，那么就意味着对方（或者己方）不在线，关闭连接。
//心跳检测用来解决半连接问题。
//测试：将连接建立后，关闭客户端或服务器，查看另一端的状态。
//发送心跳检测的发送端：
//- 可以是客户端
//- 也可以是服务端
//- 甚至是两端都发
//典型的有两种发送策略：
//1. 建立连接后，就使用固定的频率发送
//2. 一段时间没有接收到数据后，发送检测包。（TCP 层的KeepAlive就是该策略）
//心跳检测包的数据内容：
//- 可以无数据
//- 可以携带数据，例如做时钟同步，业务状态同步
//- 典型的 ping pong 结构
//心跳检测包是否需要响应？
//- 可以不响应，发送成功即可
//- 可以响应，通常用于同步数据
//总而言之，都是业务来决定。

//示例， ping-pong模式，在连接建立后持续心跳：
//* 定时心跳
//* 判断是否接收到正确心跳响应
//* 当N次心跳检测失败后，断开连接
//* Server端，发送ping包
//* Client端，接收到ping后，响应pong
//* Server端，要检测是否收到了正确的响应pong，进而判断是否要主动断开连接

const MsgCodePing = "Ping"
const MsgCodePong = "Pong"

type Ping struct {
	Message
	time time.Time
}

type Pong struct {
	Message
	time time.Time
}

func NewPing(id uint, content string) Ping {
	return Ping{
		Message: Message{
			ID:      id,
			Content: content,
			Code:    MsgCodePing,
		},
		time: time.Now(),
	}
}

func NewPong(id uint, content string) Ping {
	return Ping{
		Message: Message{
			ID:      id,
			Content: content,
			Code:    MsgCodePong,
		},
		time: time.Now(),
	}
}

func TcpServerHeartBeat() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalln("Net listen err", err)
		}
	}(listener)

	for {
		// 阻塞接受
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		// 处理连接，读写
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	log.Printf("Accept a connection from the remote client %v\n", conn.RemoteAddr())
	defer func(conn net.Conn) {
		log.Println("Close the connection from the remote client", conn.RemoteAddr())
		err := conn.Close()
		if err != nil {
			log.Fatalln("Net conn close err", err)
		}
	}(conn)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		serverPing(conn, &wg)
	}()

	wg.Wait()
}

func serverPing(conn net.Conn, wg *sync.WaitGroup) {
	const maxPingNum = 3
	pingErrCount := 0

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	wg.Add(1)
	go func() {
		defer wg.Done()
		serverAcceptPong(conn, ctx)
	}()

	ticker := time.NewTicker(3 * time.Second)
	for t := range ticker.C {
		ping := NewPing(uint(rand.Int()), "")
		encoder := gob.NewEncoder(conn)
		err := encoder.Encode(ping)
		if err != nil {
			log.Println(err)
			pingErrCount++
			if pingErrCount == maxPingNum {
				return
			}
		}
		log.Printf("Send ping to remote client %v on %v", conn.RemoteAddr(), t.Format("2006-01-02 15:04:05"))
	}
}

func serverAcceptPong(conn net.Conn, ctx context.Context) {
	for {
		select {
		// 若接收到ping结束的信号，则结束Pong的处理
		case <-ctx.Done():
			return
		default:
			message := Pong{}
			// GOB解码
			decoder := gob.NewDecoder(conn)
			// 解码操作，从conn中读取内容，成功会将解码后的结果，赋值到message变量
			err := decoder.Decode(&message)
			// 错误 io.EOF 时，表示连接被给关闭
			if err != nil && errors.Is(err, io.EOF) {
				log.Println(err)
				break
			}
			// 判断是为为 pong 类型消息
			if message.Code == MsgCodePong {
				log.Printf("receive pong from %s, %s\n", conn.RemoteAddr(), message.Content)
			}
		}
	}
}

func TcpClientHeartBeat() {
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
		clientReadPing(conn)
	}()

	wg.Wait()
}

func clientReadPing(conn net.Conn) {
	// 传递的消息类型
	ping := Ping{}
	for {
		// GOB解码
		decoder := gob.NewDecoder(conn)
		// 解码操作，从conn中读取内容，成功会将解码后的结果，赋值到message变量
		err := decoder.Decode(&ping)
		// 错误 io.EOF 时，表示连接被给关闭
		if err != nil && errors.Is(err, io.EOF) {
			log.Println(err)
			break
		}
		// 判断是为为 ping 类型消息
		if ping.Code == MsgCodePing {
			log.Println("receive ping from", conn.RemoteAddr())
			clientWritePong(conn, ping)
		}
	}
}

func clientWritePong(conn net.Conn, ping Ping) {
	pong := NewPong(uint(rand.Int()), fmt.Sprintf("pingID:%d", ping.ID))

	// GOB, 二进制编码
	// 创建编码器
	encoder := gob.NewEncoder(conn)
	// 利用编码器进行编码
	// encode 成功后，会写入到conn，已经完成了conn.Write()
	if err := encoder.Encode(pong); err != nil {
		log.Println(err)
		return
	}
	log.Println("pong was send to", conn.RemoteAddr())
	return
}
