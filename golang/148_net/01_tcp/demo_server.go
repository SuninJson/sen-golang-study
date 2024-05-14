package _1_tcp

import (
	"fmt"
	"net"
)

func StartServer() {
	println("启动服务端。。。")
	// net.Listen函数中 network参数 表示网络类型, 支持的TCP类型字符串:
	//- tcp, 使用IPv4或IPv6
	//- tcp4, 仅使用IPv4
	//- tcp6, 仅使用IPv6
	//- 省略IP部分, 绑定可用的全部IP, 包括IPv4和IPv6
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("监听失败：", err)
		return
	}

	// 若监听成功，则循环等待客户端的连接
	fmt.Println("等待客户端的连接")
	for {
		conn, acceptErr := listener.Accept()
		if acceptErr != nil {
			fmt.Println("客户端接受出现错误：", acceptErr)
			continue
		}

		fmt.Printf("等待链接成功，con=%v ，接收到的客户端信息：%v \n", conn, conn.RemoteAddr().String())

		// 通过一个协程处理接收到的客户端的连接
		go handleAccept(conn)
	}

}

func handleAccept(conn net.Conn) {
	// 用完连接后，一定要关闭
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭连接出现错误：", err)
		}
	}(conn)

	for {
		// 创建一个切片，后续需要将读取的数据放入切片
		buf := make([]byte, 1024)
		// 从conn中读取数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("从连接中读取数据出现错误：", err)
			return
		}
		// 将读取的内容在终端上显示
		fmt.Println("从客户端接受到的内容：", string(buf[0:n]))
	}
}
