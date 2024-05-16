package _1_tcp

import (
	"testing"
)

func TestStartServer(t *testing.T) {
	StartServer()
}

func TestTcpServerFormat(t *testing.T) {
	TcpServerFormat()
}

func TestTcpClientFormat(t *testing.T) {
	TcpClientFormat()
}

func TestTcpServerHeartBeat(t *testing.T) {
	TcpServerHeartBeat()
}

func TestTcpClientHeartBeat(t *testing.T) {
	TcpClientHeartBeat()
}

func TestTcpServerPool(t *testing.T) {
	TcpServerPool()
}

func TestTcpClientPool(t *testing.T) {
	TcpClientPool()
}

func TestTcpServerEncoder(t *testing.T) {
	TcpServerEncoder()
}

func TestTcpClientDecoder(t *testing.T) {
	TcpClientDecoder()
}
