package tcp_proxy

// TCP服务器
// 创建一个TCP服务器
// 1.监听端口
// 2.获取连接
// 3.封装新连接对象，设置服务参数（上下文、超时、连接关闭）
// 4.定义回调的 handler

func main() {
	var addr = "127.0.0.1:8000"

	// 创建TCP服务实例
	tcpServer := &TCPServer{
		Addr:    addr,
		Handler: &DefaultTCPHandler{},
	}

	tcpServer.ListenAndServe()
}

type TCPServer struct {
	Addr    string
	Handler TCPHandler
}

func (srv *TCPServer) ListenAndServe() error {

	return nil
}

func ListenAndServe(addr string, handler TCPHandler) error {
	server := &TCPServer{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
