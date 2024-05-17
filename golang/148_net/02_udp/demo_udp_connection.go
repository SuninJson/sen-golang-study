package _2_udp

import (
	"log"
	"net"
)

//UDP的连接分为：
// 1.已连接，connected, 使用方法 DialUDP建立的连接，称为已连接，pre-connected，`*UDPConn`是 `connected`,读写方法 `Read`和 `Write`。
// 2.未连接，unconnected，使用方法 ListenUDP 获得的连接，称为未连接，not connected，`*UDPConn`是 `unconnected`,读写方法 `ReadFromUDP`和 `WriteToUDP`

// UDPServerConnect 示例：获取远程地址，conn.RemoteAddr()**
func UDPServerConnect() {
	// 1.解析地址
	localAddr, err := net.ResolveUDPAddr("udp", ":9876")
	if err != nil {
		log.Fatalln(err)
	}

	// 2.监听
	udpConn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%s server is listening on %s\n", "UDP", udpConn.LocalAddr().String())
	defer func(udpConn *net.UDPConn) {
		err := udpConn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(udpConn)

	// 测试输出远程地址
	log.Println("测试输出远程地址：", udpConn.RemoteAddr())

	// 3.读
	buf := make([]byte, 1024)
	rn, remoteAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("received %s from %s\n", string(buf[:rn]), remoteAddr.String())

	// 测试输出远程地址
	log.Println("测试输出远程地址：", udpConn.RemoteAddr())

	// 4.写
	data := []byte("received:" + string(buf[:rn]))
	wn, err := udpConn.WriteToUDP(data, remoteAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("send %s(%d) to %s\n", string(data), wn, remoteAddr.String())

	// 测试输出远程地址
	log.Println("测试输出远程地址：", udpConn.RemoteAddr())
}

func UDPClientConnect() {
	// 1.建立连接
	remoteAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9876")
	if err != nil {
		log.Fatalln(err)
	}
	udpConn, err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		log.Fatalln(err)
	}

	// 测试输出远程地址
	log.Println(udpConn.RemoteAddr())

	// 2.写
	data := []byte("Go UDP program")
	wn, err := udpConn.Write(data) // WriteToUDP(data, remoteAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("send %s(%d) to %s\n", string(data), wn, remoteAddr.String())

	// 测试输出远程地址
	log.Println(udpConn.RemoteAddr())

	// 3.读
	buf := make([]byte, 1024)
	rn, remoteAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("received %s from %s\n", string(buf[:rn]), remoteAddr.String())

	// 测试输出远程地址
	log.Println(udpConn.RemoteAddr())
}
