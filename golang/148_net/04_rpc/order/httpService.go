package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"order/protobufs/compiles"
	"time"
)

var (
	// 目标 grpc 服务器地址
	gRPCAddr = flag.String("grpc", "localhost:5051", "the address to connect to")
	// http 命令行参数
	addr = flag.String("addr", "127.0.0.1", "The Address for listen. Default is 127.0.0.1")
	port = flag.Int("port", 8080, "The Port for listen. Default is 8080.")
)

func main() {
	flag.Parse()
	// 连接 grpc 服务器
	conn, err := grpc.NewClient(*gRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)
	// 实例化 grpc 客户端
	productClient := compiles.NewProductClient(conn)

	// 订单服务对外提供HTTP接口
	service := http.NewServeMux()
	service.HandleFunc("/orders", func(writer http.ResponseWriter, request *http.Request) {

		// RPC的超时处理
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// 调用产品服务提供的gRPC接口，获取产品信息
		productResponse, err := productClient.ProductInfo(ctx, &compiles.ProductRequest{
			Id: 1,
		})
		if err != nil {
			log.Fatalln(err)
		}

		// 构建HTTP接口的响应信息
		resp := struct {
			ID       int                         `json:"id"`
			Quantity int                         `json:"quantity"`
			Products []*compiles.ProductResponse `json:"products"`
		}{
			9527, 1,
			[]*compiles.ProductResponse{
				productResponse,
			},
		}
		respJson, err := json.Marshal(resp)
		if err != nil {
			log.Fatalln(err)
		}
		writer.Header().Set("Content-Type", "application/json")
		_, err = fmt.Fprintf(writer, "%s", string(respJson))
		if err != nil {
			log.Fatalln(err)
		}
	})

	// 启动监听
	address := fmt.Sprintf("%s:%d", *addr, *port)
	fmt.Printf("Order service is listening on %s.\n", address)
	log.Fatalln(http.ListenAndServe(address, service))
}
