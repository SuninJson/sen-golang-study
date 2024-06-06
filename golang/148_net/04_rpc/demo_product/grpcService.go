package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"sen-golang-study/golang/148_net/04_rpc/demo_product/protobufs/compiles"
)

var (
	port = flag.Int("port", 5051, "The gRPC Server's port")
)

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()

	// 将产品服务注册到gRPC服务中
	compiles.RegisterProductServer(server, &ProductServer{})

	// 启动TCP监听
	log.Printf("gRPC Server is listening on %s.\n", listener.Addr())
	serverErr := server.Serve(listener)
	if serverErr != nil {
		log.Fatalln(serverErr)
	}
}

type ProductServer struct {
	compiles.UnimplementedProductServer
}

func (ProductServer) ProductInfo(ctx context.Context, request *compiles.ProductRequest) (*compiles.ProductResponse, error) {
	resp := compiles.ProductResponse{
		Name:   "gRPC Product Demo",
		Id:     1,
		IsSale: true,
	}
	return &resp, nil
}
