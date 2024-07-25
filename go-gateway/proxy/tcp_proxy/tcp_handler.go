package tcp_proxy

type TCPHandler interface {
	TCPServe()
}

type DefaultTCPHandler struct {
}

func (handler DefaultTCPHandler) TCPServe() {

}
