package _2_udp

import (
	"fmt"
	"log"
	"net"
	"time"
)

// 广播地址，Broadcast，指的是将消息发送到在同一广播网络上的每个主机
// 例如对于网络：
// # ip a
// ens33: <BROADCAST,MULTICAST,UP,LOWER_UP>
// inet 192.168.50.130/24 brd 192.168.50.255
// IP ADDR 就是 192.168.50.130/24， 广播地址就是 192.168.50.255。
// 意味着，只要发送数据包的目标地址（接收地址）为192.168.50.255时，那么该数据会发送给该网段上的所有计算机

func UDPReceiverBroadcast() {
	// 1.广播监听地址
	localAddr, err := net.ResolveUDPAddr("udp", ":6789")
	if err != nil {
		log.Fatalln(err)
	}

	// 2.广播监听
	udpConn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer udpConn.Close()
	log.Printf("%s server is listening on %s\n", "UDP", udpConn.LocalAddr().String())

	// 3.接收数据
	// 4.处理数据
	buf := make([]byte, 1024)
	for {
		rn, remoteAddr, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
		}
		log.Printf("received \"%s\" from %s\n", string(buf[:rn]), remoteAddr.String())
	}
}

func UDPSenderBroadcast() {
	// 1.监听地址
	// 2.建立连接
	localAddr, err := net.ResolveUDPAddr("udp", ":9876")
	if err != nil {
		log.Fatalln(err)
	}
	udpConn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer udpConn.Close()
	log.Printf("%s server is listening on %s\n", "UDP", udpConn.LocalAddr().String())

	// 3.发送数据
	// 广播地址
	rAddress := "192.168.50.255:6789"
	remoteAddr, err := net.ResolveUDPAddr("udp", rAddress)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		data := fmt.Sprintf("[%s]: %s", time.Now().Format("03:04:05.000"), "hello!")
		// 广播发送
		wn, err := udpConn.WriteToUDP([]byte(data), remoteAddr)
		if err != nil {
			log.Println(err)
		}
		log.Printf("send \"%s\"(%d) to %s\n", data, wn, remoteAddr.String())

		time.Sleep(time.Second)
	}
}
