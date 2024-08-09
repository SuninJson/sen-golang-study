package tcp_proxy

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrServerClosed     = errors.New("http: Server closed")
	ServerContextKey    = &contextKey{"tcp_proxy-server"}
	ErrAbortHandler     = errors.New("tcp_proxy: abort TCPHandler")
	LocalAddrContextKey = &contextKey{"local-addr"}
)

type contextKey struct {
	name string
}

type TCPServer struct {
	Addr    string     //主机地址
	Handler TCPHandler //TCP请求的处理器

	BaseContext      context.Context //上下文:用来收集取消、终止、错误等信息
	err              error
	ReadTimeout      time.Duration //读操作的超时时间
	WriteTimeout     time.Duration //写操作的超时时间
	KeepAliveTimeout time.Duration //长连接断开后的超时时间

	mu         sync.Mutex         //连接初始化、关闭等动作需要加锁
	doneChan   chan struct{}      //服务已完成时，doneChan监听信号
	isShutdown int32              //服务终止状态：0-未关闭，1-已关闭
	listener   *onceCloseListener //服务器监听器，使用完成后将进程关闭
}

func (srv *TCPServer) ListenAndServe() error {
	if srv.shuttingDown() {
		return ErrServerClosed
	}
	addr := srv.Addr
	if addr == "" {
		return errors.New("TCP Server address is empty")
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve(listener)
}

func ListenAndServe(addr string, handler TCPHandler) error {
	server := &TCPServer{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}

type onceCloseListener struct {
	net.Listener
	once     sync.Once
	closeErr error
}

func (oc *onceCloseListener) Close() error {
	oc.once.Do(oc.close)
	return oc.closeErr
}

func (oc *onceCloseListener) close() { oc.closeErr = oc.Listener.Close() }

// Close 关闭TCP Server
func (srv *TCPServer) Close() error {
	// 可能存在其它协程同时去关闭TCP Server
	atomic.StoreInt32(&srv.isShutdown, 1) // 用原子操作修改服务的状态
	close(srv.doneChan)                   // 关闭监听服务完成的信号Channel
	srv.listener.close()

	return nil
}

func (srv *TCPServer) shuttingDown() bool {
	return atomic.LoadInt32(&srv.isShutdown) == 1

}

func (srv *TCPServer) Serve(l net.Listener) error {
	srv.listener = &onceCloseListener{Listener: l}
	defer func(listener *onceCloseListener) {
		err := listener.Close()
		if err != nil {
			log.Fatalln("Close listener failed")
		}
	}(srv.listener) //执行listener关闭
	if srv.BaseContext == nil {
		srv.BaseContext = context.Background()
	}
	baseCtx := srv.BaseContext
	ctx := context.WithValue(baseCtx, ServerContextKey, srv)
	for {
		rwConn, e := l.Accept()
		if e != nil {
			select {
			case <-srv.getDoneChan():
				// 若获得到了终止信息，则返回服务器已经被关闭的错误信息
				return ErrServerClosed
			default:
			}
			fmt.Printf("accept fail, err: %v\n", e)
			continue
		}
		tcpConn := srv.newConn(rwConn)
		go tcpConn.serve(ctx)
	}
}

func (srv *TCPServer) getDoneChan() <-chan struct{} {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	if srv.doneChan == nil {
		// 若doneChan没初始化过，则将其初始化
		srv.doneChan = make(chan struct{})
	}
	return srv.doneChan
}

// 对监听到的连接进行再次封装，来支持我们的额外设置的参数比如ReadTimeout、WriteTimeout和KeepAliveTimeout等
func (srv *TCPServer) newConn(readWriteConn net.Conn) *TCPConn {
	tcpConn := &TCPConn{
		server:        srv,
		readWriteConn: readWriteConn,
		remoteAddr:    readWriteConn.RemoteAddr().String(),
	}
	// 设置参数
	if d := tcpConn.server.ReadTimeout; d != 0 {
		tcpConn.readWriteConn.SetReadDeadline(time.Now().Add(d))
	}
	if d := tcpConn.server.WriteTimeout; d != 0 {
		tcpConn.readWriteConn.SetWriteDeadline(time.Now().Add(d))
	}
	if d := tcpConn.server.KeepAliveTimeout; d != 0 {
		if tcpConn, ok := tcpConn.readWriteConn.(*net.TCPConn); ok {
			tcpConn.SetKeepAlive(true)
			tcpConn.SetKeepAlivePeriod(d)
		}
	}
	return tcpConn
}
