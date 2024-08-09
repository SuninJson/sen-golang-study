package tcp_proxy

type TCPProxy struct {
}

func (p TCPProxy) ListenAndServe() {

}

func NewSingleHostReverseProxy(addr string) TCPProxy {
	return TCPProxy{}
}
