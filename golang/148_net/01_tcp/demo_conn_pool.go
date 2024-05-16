package _1_tcp

import (
	"errors"
	"log"
	"net"
	"sync"
	"time"
)

type Pool interface {
	// Get 获取连接
	Get() (net.Conn, error)
	// Put 放回连接
	Put(conn net.Conn) error
	// Release 释放池(全部连接)
	Release() error
	// Len 有效连接个数
	Len() int
}

// ConnFactory 连接池还应该有能力创建新连接. 在Get操作时,若没有空闲可用的连接, 在数量允许的情况下,会创造新的连接. 通过连接工厂来提供这样的能力
type ConnFactory interface {
	// NewConn 构造连接
	NewConn(addr string) (net.Conn, error)
	// Close 关闭连接的方法
	Close(net.Conn) error
	// Ping 检查连接是否有效的方法
	Ping(net.Conn) error
}

type PoolConfig struct {
	//初始连接数, 池初始化时的连接数
	InitConnNum int
	//最大连接数, 池中最多支持多少连接
	MaxConnNum int
	//最大空闲连接数, 池中最多有多少可用的连接
	MaxIdleNum int
	//空闲连接超时时间, 多久后空闲连接会被释放
	IdleTimeout time.Duration

	// 连接地址
	addr string

	// 连接工厂
	Factory ConnFactory
}

const defaultMaxConnNum = 10
const defaultInitConnNum = 10

func NewPoolConfig(initConnNum int, maxConnNum int, maxIdleNum int, idleTimeout time.Duration, connFactory ConnFactory) (*PoolConfig, error) {
	poolConfig := &PoolConfig{
		InitConnNum: initConnNum,
		MaxConnNum:  maxConnNum,
		MaxIdleNum:  maxIdleNum,
		IdleTimeout: idleTimeout,
		addr:        "",
		Factory:     connFactory,
	}

	// 最大连接数
	if poolConfig.MaxConnNum == 0 {
		// 1. return错误
		//return nil, errors.New("max conn num is zero")

		// 2. 人为修改一个合理的
		poolConfig.MaxConnNum = defaultMaxConnNum
	}
	// 初始化连接数
	if poolConfig.InitConnNum == 0 {
		poolConfig.InitConnNum = defaultInitConnNum
	} else if poolConfig.InitConnNum > poolConfig.MaxConnNum {
		poolConfig.InitConnNum = poolConfig.MaxConnNum
	}
	// 若最大空闲连接数为0或最大空闲连接数多于最大连接数时，更正配置，使得配置合理
	if poolConfig.MaxIdleNum == 0 {
		poolConfig.MaxIdleNum = poolConfig.InitConnNum
	} else if poolConfig.MaxIdleNum > poolConfig.MaxConnNum {
		poolConfig.MaxIdleNum = poolConfig.MaxConnNum
	}

	return poolConfig, nil
}

// IdleConn 由于需要判断连接的空闲时间,因此,需要记录连接被放入到连接池中的时间，因此,需要定义一个空闲连接结构
type IdleConn struct {
	// 连接本身
	conn net.Conn
	// 放回时间
	putTime time.Time
}

type TcpPool struct {
	// 相关配置
	config *PoolConfig

	// 开放使用的连接数量
	openingConnNum int
	// 空闲的连接队列
	idleList chan *IdleConn

	addr string

	// 并发安全锁
	mu sync.RWMutex
}

func (pool *TcpPool) Get() (net.Conn, error) {
	// 1锁定
	pool.mu.Lock()
	defer pool.mu.Unlock()

	// 2获取空闲连接，若没有则创建连接
	for {
		select {
		// 获取空闲连接
		case idleConn, ok := <-pool.idleList:
			// 判断channel是否被关闭
			if !ok {
				return nil, errors.New("idle list closed")
			}

			// 判断连接是否超时
			//pool.config.IdleTimeout, idleConn.putTime
			if pool.config.IdleTimeout > 0 { // 设置了超时时间
				// putTime + timeout 是否在 now 之前
				if idleConn.putTime.Add(pool.config.IdleTimeout).Before(time.Now()) {
					// 关闭连接，继续查找下一个连接
					_ = pool.config.Factory.Close(idleConn.conn)
					continue
				}
			}

			// 判断连接是否可用
			if err := pool.config.Factory.Ping(idleConn.conn); err != nil {
				// ping 失败，连接不可用
				// 关闭连接，继续查找
				_ = pool.config.Factory.Close(idleConn.conn)
				continue
			}

			// 找到了可用的空闲连接
			log.Println("get conn from Idle")
			// 使用的连接计数
			pool.openingConnNum++
			// 返回连接
			return idleConn.conn, nil

		// 创建连接
		default:
			// a判断是否还可以继续创建
			// 基于开放的连接是否已经达到了连接池最大的连接数
			if pool.openingConnNum >= pool.config.MaxConnNum {
				return nil, errors.New("max opening connection")
				// 另一种方案，就是阻塞
				//continue
			}

			// b创建连接
			conn, err := pool.config.Factory.NewConn(pool.addr)
			if err != nil {
				return nil, err
			}

			// c正确创建了可用的连接
			log.Println("get conn from Factory")
			// 使用的连接计数
			pool.openingConnNum++
			// 返回连接
			return conn, nil
		}
	}
}

func (pool *TcpPool) Put(conn net.Conn) error {
	// 1锁
	pool.mu.Lock()
	defer pool.mu.Unlock()

	// 2做一些校验
	if conn == nil {
		return errors.New("connection is not exists")
	}
	// 判断空闲连接列表是否存在
	if pool.idleList == nil {
		// 关闭连接
		_ = pool.config.Factory.Close(conn)
		return errors.New("idle list is not exists")
	}

	// 3放回连接
	select {
	// 放回连接
	case pool.idleList <- &IdleConn{
		conn:    conn,
		putTime: time.Now(),
	}:
		// 只要可以发送成功，任务完成
		// 更新开放的连接数量
		pool.openingConnNum--
		return nil
	// 关闭连接
	default:
		_ = pool.config.Factory.Close(conn)
		return nil
	}
}

// Release 释放连接池
func (pool *TcpPool) Release() error {
	// 1并发安全锁
	pool.mu.Lock()
	defer pool.mu.Unlock()

	// 2确定连接池是否被释放
	if pool.idleList == nil {
		return nil
	}

	// 3关闭IdleList
	close(pool.idleList)

	// 4释放全部空闲连接
	// 继续接收已关闭channel中的元素
	for idleConn := range pool.idleList {
		// 关闭连接
		_ = pool.config.Factory.Close(idleConn.conn)
	}

	return nil
}

func (pool *TcpPool) Len() int {
	return len(pool.idleList)
}

// TcpConnFactory Tcp连接工厂类型，使用ConnFactory接口的所有方法
type TcpConnFactory struct{}

// NewConn 产生连接方法
func (*TcpConnFactory) NewConn(addr string) (net.Conn, error) {
	// 校验参数的合理性
	if addr == "" {
		return nil, errors.New("addr is empty")
	}

	// 建立连接
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		return nil, err
	}

	// return
	return conn, nil
}

func (*TcpConnFactory) Close(conn net.Conn) error {
	return conn.Close()
}

func (*TcpConnFactory) Ping(conn net.Conn) error {
	return nil
}

// NewTcpPool 创建TcpPool对象
func NewTcpPool(addr string, poolConfig PoolConfig) (*TcpPool, error) {
	// 1.校验参数
	if addr == "" {
		return nil, errors.New("addr is empty")
	}

	// 校验工厂的存在
	if poolConfig.Factory == nil {
		return nil, errors.New("factory is not exists")
	}

	// 2.初始化TcpPool对象
	pool := TcpPool{
		config:         &poolConfig,
		openingConnNum: 0,
		idleList:       make(chan *IdleConn, poolConfig.MaxIdleNum),
		addr:           addr,
		mu:             sync.RWMutex{},
	}

	// 3.初始化连接
	// 根据InitConnNum的配置来创建
	for i := 0; i < poolConfig.InitConnNum; i++ {
		conn, err := pool.config.Factory.NewConn(addr)
		if err != nil {
			// 通常意味着，连接池初始化失败
			// 释放可能已经存在的连接
			err := pool.Release()
			if err != nil {
				return nil, err
			}
			return nil, err
		}
		// 连接创建成功
		// 加入到空闲连接队列中
		pool.idleList <- &IdleConn{
			conn:    conn,
			putTime: time.Now(),
		}
	}

	// 4返回
	return &pool, nil
}

func TcpServerPool() {
	// A. 基于某个地址建立监听
	// 服务端地址
	address := ":8888" // Any IP or version
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	// 关闭监听
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(listener)
	log.Printf("%s server is listening on %s\n", tcp, listener.Addr())

	// B. 接受连接请求
	// 循环接受
	for {
		// 阻塞接受
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		// 处理连接，读写
		go handleConnPool(conn)
	}
}

func handleConnPool(conn net.Conn) {
	// 日志连接的远程地址（client addr）
	log.Printf("accept from %s\n", conn.RemoteAddr())
	// A.保证连接关闭
	defer func() {
		err := conn.Close()
		if err != nil {
			return
		}
		log.Println("connection be closed")
	}()

	select {}
}

func TcpClientPool() {
	// tcp服务端地址
	serverAddress := "127.0.0.1:8888" // IPv6 4
	// A，建立连接池
	pool, err := NewTcpPool(serverAddress, PoolConfig{
		Factory:     &TcpConnFactory{},
		InitConnNum: 4,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(pool, len(pool.idleList))

	wg := sync.WaitGroup{}
	clientNum := 50
	wg.Add(clientNum)
	// B, 复用连接池中的连接
	for i := 0; i < clientNum; i++ {
		// goroutine 模拟独立的客户端
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 获取连接
			conn, err := pool.Get()
			if err != nil {
				log.Println(err)
				return
			}
			//log.Println(conn)
			// 回收连接
			putErr := pool.Put(conn)
			if putErr != nil {
				log.Println(putErr)
				return
			}
		}(&wg)
	}
	wg.Wait()
}
