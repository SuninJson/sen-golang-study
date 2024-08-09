package tcp_proxy

import (
	"context"
	"fmt"
	"net"
	"runtime"
)

type TCPConn struct {
	server        *TCPServer
	cancelCtx     context.CancelFunc
	readWriteConn net.Conn
	remoteAddr    string
}

func (c *TCPConn) close() {
	err := c.readWriteConn.Close()
	if err != nil {
		return
	}
}

func (c *TCPConn) serve(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil && err != ErrAbortHandler {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			fmt.Printf("tcp_proxy: panic serving %v: %v\n%s", c.remoteAddr, err, buf)
		}
		c.close()
	}()
	c.remoteAddr = c.readWriteConn.RemoteAddr().String()
	ctx = context.WithValue(ctx, LocalAddrContextKey, c.readWriteConn.LocalAddr())
	if c.server.Handler == nil {
		panic("handler empty")
	}
	c.server.Handler.Serve(ctx, c.readWriteConn)
}
