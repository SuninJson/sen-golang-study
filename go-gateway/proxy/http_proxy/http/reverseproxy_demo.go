package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// HTTP反向代理简单版：用ReverseProxy实现
func main() {
	// 下游真实服务器地址
	realServer := "http://127.0.0.1:8001"
	// parse 解析url
	// 从 http://127.0.0.1:8001/base?param=0100&p2=haha&p3=world#container
	// 到 http://127.0.0.1:8001/base
	serverURL, err1 := url.Parse(realServer)
	if err1 != nil {
		log.Println(err1)
	}
	proxy := httputil.NewSingleHostReverseProxy(serverURL)
	// 代理服务器: 8081
	var addr = "127.0.0.1:8081"
	log.Println("Starting proxy http server at: " + addr)
	// 下面这两行代码与最后一行等价
	// http.Handle("/", proxy)
	// http.ListenAndServe(addr, nil)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
