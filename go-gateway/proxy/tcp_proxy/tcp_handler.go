package tcp_proxy

import (
	"context"
	"net"
)

type TCPHandler interface {
	Serve(context.Context, net.Conn)
}

type DefaultTCPHandler struct {
}

func (handler DefaultTCPHandler) Serve(ctx context.Context, conn net.Conn) {
	conn.Write([]byte("Pong!TCP handler here.\n"))
}
