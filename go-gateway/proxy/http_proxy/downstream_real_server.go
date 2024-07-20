package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sen-golang-study/go-gateway/load_balance/service_discovery/zookeeper"
	"syscall"
	"time"
)

/*
证书签名生成方式:

//CA私钥
openssl genrsa -out ca.key 2048
//CA数据证书
openssl req -x509 -new -nodes -key ca.key -subj "/CN=example1.com" -days 5000 -out ca.crt

//服务器私钥（默认由CA签发）
openssl genrsa -out server.key 2048
//服务器证书签名请求：Certificate Sign Request，简称csr（example1.com代表你的域名）
openssl req -new -key server.key -subj "/CN=example1.com" -out server.csr
//上面2个文件生成服务器证书（days代表有效期）
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000
*/

func main() {
	server1 := &RealServer{Addr: "127.0.0.1:8001"}
	server1.Run()
	//server2 := &RealServer{Addr: "127.0.0.1:8002"}
	//server2.Run()

	// 监听系统关闭信号. 否则main协程终止后,将直接导致守护协程终止
	// 相当于 ctrl + c, kill
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

// RealServer 下游真实服务器
type RealServer struct {
	Addr string // 服务器主机地址: {host:port}
}

// Run 新建协程启动服务器
func (r *RealServer) Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/realserver", r.HelloHandler)
	mux.HandleFunc("/realserver/error", r.ErrorHandler)
	server := &http.Server{
		Addr:         r.Addr,
		Handler:      mux,
		WriteTimeout: time.Second * 3,
	}
	// 以新的协程的方式启动服务
	go func() {
		// 下游服务器启动时，注册zk节点：临时节点
		zkManager := zookeeper.NewZkManager([]string{"192.168.154.132:2181"})
		err := zkManager.GetConnect()
		if err != nil {
			fmt.Printf(" connect zk error: %s ", err)
		}
		defer zkManager.Close()
		err = zkManager.RegisterServerPath("/realserver", r.Addr)
		if err != nil {
			fmt.Printf(" regist node error: %s ", err)
		}
		//zlist, err := zkManager.GetServerListByPath("/realserver")
		//fmt.Println(zlist)
		log.Fatal(server.ListenAndServe())
	}()
}

// HelloHandler 路由处理器
func (r *RealServer) HelloHandler(w http.ResponseWriter, req *http.Request) {
	// 重新拼接好url后返回
	newPath := fmt.Sprintf("here is realserver: http://%s%s", r.Addr, req.URL.Path)
	//fmt.Println(newPath)
	w.Write([]byte(newPath))
}

// ErrorHandler 错误处理器
func (r *RealServer) ErrorHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusInternalServerError) // 服务器内部错误
	w.Write([]byte("error: 服务器内部错误"))
}
