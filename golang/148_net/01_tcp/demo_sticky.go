package _1_tcp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"math/rand"
	"net"
	"time"
)

// [粘包现象]
// 指TCP协议中，发送方发送的若干包数据到接收方接收时粘成一包，从接收缓冲区看，后一包数据的头紧接着前一包数据的尾
// 其实TCP是面向字节流的协议，就是没有界限的一串数据，本没有“包”的概念， 包可以当作一个逻辑上的数据单元。
// “粘包”和“拆包”是逻辑上的概念。

// [粘包原因]
//- 发送端：TCP使用Nagle算法来减少传输的报文数量，下面两个原因引发发送粘包问题
//  1. 前一个分组确认，发送下一个分组
//  2. 收集多个分组，收到确认后一起发送
//- 接收端：TCP将接收到的数据包保存在接收缓存里，然后应用程序主动从缓存读取收到的分组。应用程序不能及时接收到发送的数据。
//
// 当发送的多个数据包之间需要逻辑隔离，那么就需要处理粘包问题。反之若发送的数据本身就是一个连续的整体，那么不需要处理粘包问题。

// [解决方案]
//数据包连着发送这个是不能改变的。我们需要在数据包层面标注包与包的分离方案，来解决粘包现象带来的问题。
//典型的方案有：
// 1. 每个包都封装成固定的长度。读取到内容时，依据长度进行分割即可
// 2.数据包使用特定分隔符。读取到内容时，依据分隔符分割数据即可，例如HTTP,FTP协议的\r\n。
// 3.将消息封装为Header+Body形式，Header通常时固定长度，Header中包含Body的长度信息。读取到期待长度时，才表示成功。
//不论何种方案，在编码实现时，通常采用定义编解码器的方案来实现。就类似JSON和GOB编码。

// 示例编码，以方案三，Header+Body的模式为例：
// 约定Header的长度为4个字节，用于表示Body的长度。

type Encoder struct {
	// 编码结束后，写入的目标
	w io.Writer
}

// NewEncoder 创建编码器
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: w,
	}
}

// Encode 编码，将编码的结果，写入到w io.Writer
func (enc *Encoder) Encode(message string) error {
	// 1.获取message的长度
	l := int32(len(message))

	// 构建一个数据包缓存
	buf := new(bytes.Buffer)

	// 2.在数据包中写入长度
	// 需要二进制的写入操作，需要将数据以bit的形式写入
	if err := binary.Write(buf, binary.LittleEndian, l); err != nil {
		return err
	}

	// 3.将数据主体Body写入
	if _, err := buf.Write([]byte(message)); err != nil {
		return err
	}

	// 4.利用io.Writer发送数据
	if n, err := enc.w.Write(buf.Bytes()); err != nil {
		log.Println(n, err)
		return err
	}

	return nil
}

// Decoder 解码器
type Decoder struct {
	r io.Reader
}

// NewDecoder 创建Decoder
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r: r,
	}
}

// Decode 从Reader中读取内容，解码
func (dec *Decoder) Decode(message *string) error {
	// 1.读取前4个字节，读取header
	header := make([]byte, 4)
	hn, err := dec.r.Read(header)
	if err != nil {
		return err
	}
	if hn != 4 {
		return errors.New("header is not enough")
	}

	// 2.将前4个字节转换为int32类型，确定了body的长度
	var l int32
	headerBuf := bytes.NewBuffer(header)
	if err := binary.Read(headerBuf, binary.LittleEndian, &l); err != nil {
		return err
	}

	// 3.读取body
	body := make([]byte, l)
	bn, err := dec.r.Read(body)
	if err != nil {
		return err
	}
	if bn != int(l) {
		return errors.New("body is not enough")
	}

	// 4.设置message
	*message = string(body)

	return nil
}

func TcpServerEncoder() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalln("Net listen err", err)
		}
	}(listener)

	for {
		// 阻塞接受
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		// 处理连接，读写
		go handleConnCoder(conn)
	}

}

func handleConnCoder(conn net.Conn) {
	// 日志连接的远程地址（client addr）
	log.Printf("accept from %s\n", conn.RemoteAddr())
	// A.保证连接关闭
	defer func() {
		conn.Close()
		log.Println("connection be closed")
	}()

	// 连续发送数据
	data := []string{
		"package data.",
		"package.",
		"package data data",
		"pack",
	}
	encoder := NewEncoder(conn)
	// 模拟发送50次消息
	for i := 0; i < 50; i++ {
		// 创建编解码器
		// 利用编码器进行编码
		// encode 成功后，会写入到conn，已经完成了conn.Write()
		if err := encoder.Encode(data[rand.Intn(len(data))]); err != nil {
			log.Println(err)
		}
	}
}

func TcpClientDecoder() {
	// tcp服务端地址
	serverAddress := "127.0.0.1:8888"

	// A. 建立连接
	conn, err := net.DialTimeout(tcp, serverAddress, time.Second)
	if err != nil {
		log.Println(err)
		return
	}
	// 保证关闭
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)
	log.Printf("connection is establish, client addr is %s\n", conn.LocalAddr())

	// 从服务端接收数据，SerRead
	// 创建解码器
	decoder := NewDecoder(conn)
	data := ""
	i := 0
	for {
		// 错误 io.EOF 时，表示连接被给关闭
		if err := decoder.Decode(&data); err != nil {
			log.Println(err)
			break
		}

		log.Println(i, "received data:", data)
		i++
	}
}
