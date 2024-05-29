package _2_udp

import (
	"fmt"
	"log"
	"net"
	"time"
)

// 【多播（Multicast）方式的数据传输】是基于 UDP 完成的。
// 与 UDP 服务器端/客户端的单播方式不同，区别是，单播数据传输以单一目标进行，而多播数据同时传递到加入（注册）特定组的大量主机。
// 换言之，采用多播方式时，可以同时向多个主机传递数据。
// 【多播的特点如下】：
// （1） 多播发送端针对特定多播组
// （2） 发送端发送 1 次数据，但该组内的所有接收端都会接收数据
// （3） 多播组数可以在 IP 地址范围内任意增加
// 【多播组是 D 类IP地址（224.0.0.0~239.255.255.255）】：
// （1） 224.0.0.0～224.0.0.255为预留的组播地址（永久组地址），地址224.0.0.0保留不做分配，其它地址供路由协议使用；
// （2） 224.0.1.0～224.0.1.255是公用组播地址，Internet work Control Block；
// （3） 224.0.2.0～238.255.255.255为用户可用的组播地址（临时组地址），全网范围内有效；
// （4） 239.0.0.0～239.255.255.255为本地管理组播地址，仅在特定的本地范围内有效

const address = "224.0.0.1:9999"

func UDPMulticastReceive() {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)

	for {
		buf := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(buf[:n]))
	}
}

func UDPMulticastSend() {
	// 1.建立UDP多播组连接
	remoteAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatalln(err)
	}
	udpConn, err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(udpConn *net.UDPConn) {
		err := udpConn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(udpConn)

	// 2.发送内容
	// 每隔1秒发送一次消息
	for {
		data := fmt.Sprintf("[%s]: %s", time.Now().Format("03:04:05.000"), "hello!")
		wn, err := udpConn.Write([]byte(data))
		if err != nil {
			log.Println(err)
		}
		log.Printf("send \"%s\"(%d) to %s\n", data, wn, remoteAddr.String())

		time.Sleep(time.Second)
	}
}
