package _3_poll

import (
	"log"
	"net"
	"sync"
	"time"
)

func BIONet() {
	// 模拟读
	addr := "127.0.0.1:5678"
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatalln(err)
		}
		defer func(conn net.Conn) {
			closeErr := conn.Close()
			if closeErr != nil {
				log.Fatalln(closeErr)
			}
		}(conn)

		log.Println("start read.", time.Now().Format("03:04:05.000"))
		buf := make([]byte, 1024)
		// 读操作会阻塞，直到接收到数据
		n, _ := conn.Read(buf)
		log.Println("content:", string(buf[:n]), time.Now().Format("03:04:05.000"))
	}()

	// 模拟写
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		listener, _ := net.Listen("tcp", addr)
		defer listener.Close()

		for {
			conn, _ := listener.Accept()
			go func(conn net.Conn) {
				defer conn.Close()
				log.Println("connected.")

				// 阻塞时长
				time.Sleep(3 * time.Second)
				conn.Write([]byte("Blocking I/O"))
			}(conn)
		}
	}(&wg)

	wg.Wait()
}
