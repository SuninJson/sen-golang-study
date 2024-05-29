package _2_udp

import (
	"errors"
	"io"
	"log"
	"net"
	"os"
	"time"
)

// UDP协议在传输数据时，由于不能保证稳定性传输，因此比较适合多媒体通信领域，例如直播、视频、音频即时播放，即时通信等领域。
// 本案例使用文件传输为例。

// UDPFileClient 客户端：
// 1. 发送文件mp3（任意类型都ok）
// 2. 发送文件名
// 3. 发送文件内容
func UDPFileClient() {
	// 1.获取文件信息
	filename := "./data/Beyond.mp3"
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	// 关闭文件
	defer file.Close()
	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("send file size:", fileInfo.Size())

	// 2.连接服务器
	remoteAddress := "192.168.50.131:5678"
	remoteUDPAddr, err := net.ResolveUDPAddr("udp", remoteAddress)
	if err != nil {
		log.Fatalln(err)
	}
	udpConn, err := net.DialUDP("udp", nil, remoteUDPAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer udpConn.Close()

	// 3.发送文件名
	if _, err := udpConn.Write([]byte(fileInfo.Name())); err != nil {
		log.Fatalln(err)
	}

	// 4.服务端确认
	buf := make([]byte, 4*1024)
	rn, err := udpConn.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	// 判断是否为文件名正确接收响应
	if "filename ok" != string(buf[:rn]) {
		log.Fatalln(errors.New("server not ready"))
	}

	// 5.发送文件内容
	// 读取文件内容，利用连接发送到服务端
	// file.Read()
	i := 0
	for {
		// 读取文件内容
		rn, err := file.Read(buf)
		if err != nil {
			// io.EOF 错误表示文件读取完毕
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		// 发送到服务端
		if _, err := udpConn.Write(buf[:rn]); err != nil {
			log.Fatalln(err)
		}
		i++
	}
	log.Println(i)
	// 文件发送完成。
	log.Println("file send complete.")

	// 等待的测试
	time.Sleep(2 * time.Second)
}

// UDPFileServer 服务端：
// 1. 接收文件
// 2. 存储为客户端发送的名字
// 3. 接收文件内容
// 4. 写入到具体文件中
func UDPFileServer() {
	// 1.建立UDP连接
	localAddress := ":5678"
	localAddr, err := net.ResolveUDPAddr("udp", localAddress)
	if err != nil {
		log.Fatalln(err)
	}
	udpConn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer udpConn.Close()
	log.Printf("%s server is listening on %s\n", "UDP", udpConn.LocalAddr().String())

	// 2.接收文件名，并确认
	buf := make([]byte, 4*1024)
	rn, remoteAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Fatalln(err)
	}
	filename := string(buf[:rn])
	if _, err := udpConn.WriteToUDP([]byte("filename ok"), remoteAddr); err != nil {
		log.Fatalln(err)
	}

	// 3.接收文件内容，并写入文件
	// 打开文件（创建）
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// 网络读取
	i := 0
	for {
		// 一次读取
		rn, _, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			log.Fatalln(err)
		}

		// 写入文件
		if _, err := file.Write(buf[:rn]); err != nil {
			log.Fatalln(err)
		}
		i++
		log.Println("file write some content", i)
	}
}
