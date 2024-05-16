package _2_udp

import (
	"log"
	"net"
)

// UDP，User Datagram Protocol，用户数据报协议，是一个简单的面向数据报(package-oriented)的传输层协议，
// 规范为：[RFC 768](https://datatracker.ietf.org/doc/html/rfc768)。
// UDP提供数据的不可靠传递，它一旦把应用程序发给网络层的数据发送出去，就不保留数据备份。缺乏可靠性，缺乏拥塞控制

func UDPServerBasic() {
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

	// 3.读
	buf := make([]byte, 1024)
	rn, remoteAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("received %s from %s\n", string(buf[:rn]), remoteAddr.String())

	// 4.写
	data := []byte("received:" + string(buf[:rn]))
	wn, err := udpConn.WriteToUDP(data, remoteAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("send %s(%d) to %s\n", string(data), wn, remoteAddr.String())
}

func UDPClientBasic() {
	// 1.建立连接
	remoteAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9876")
	if err != nil {
		log.Fatalln(err)
	}
	udpConn, err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(udpConn)

	// 2.写
	data := []byte("Go UDP program")
	wn, err := udpConn.Write(data) // WriteToUDP(data, remoteAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("send %s(%d) to %s\n", string(data), wn, remoteAddr.String())

	// 3.读
	buf := make([]byte, 1024)
	rn, remoteAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("received %s from %s\n", string(buf[:rn]), remoteAddr.String())
}
