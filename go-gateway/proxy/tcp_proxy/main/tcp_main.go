package main

import (
	"log"
	"os"
	"os/signal"
	"sen-golang-study/go-gateway/proxy/tcp_proxy"
	"syscall"
)

// TCP服务器
// 创建一个TCP服务器
// 1.监听端口
// 2.获取连接
// 3.封装新连接对象，设置服务参数（上下文、超时、连接关闭）
// 4.定义回调的 handler

func main() {
	// 启动TCP服务器
	go func() {
		var addr = "127.0.0.1:8000"

		// 创建TCP服务实例
		tcpServer := &tcp_proxy.TCPServer{
			Addr:    addr,
			Handler: &tcp_proxy.DefaultTCPHandler{},
		}

		// 开始监听并提供服务
		log.Println("Starting TCP server at " + addr)
		tcpServer.ListenAndServe()
	}()

	// 启动TCP代理服务器
	go func() {
		var proxyServerAddr = "192.168.110.11:8081"

		tcpServer := tcp_proxy.NewSingleHostReverseProxy(proxyServerAddr)

		// 开始监听并提供服务
		log.Println("Starting TCP server at " + proxyServerAddr)
		tcpServer.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
}
