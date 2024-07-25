package main

import (
	"bytes"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// HTTP反向代理完整版：用ReverseProxy实现
//
// 支持功能：
//	URL重写、更改请求或响应内容、错误信息回调、连接池

func main() {
	// 真实服务地址的URL [protocol]://[IP]:[port]/?[query parameter1]&[query parameter2]#[Fragment Identifier]
	realServer := "http://127.0.0.1:8001/?a=1&b=2#hello"
	// 将服务的URL字符串解析为 Golang net包中的URL结构体
	serverURL, err := url.Parse(realServer)
	if err != nil {
		log.Println(err)
		return
	}

	proxy := NewSingleHostReverseProxy(serverURL)

	var addr = "127.0.0.1:8081"
	log.Println("Starting http proxy server at " + addr)
	log.Fatalln(http.ListenAndServe(addr, proxy))
}

var transport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second, // 拨号超时时间
		KeepAlive: 30 * time.Second, // 长连接超时时间
	}).DialContext,
	MaxIdleConns:          100,              //最大空闲连接数
	IdleConnTimeout:       90 * time.Second, //空闲连接超时时间
	TLSHandshakeTimeout:   10 * time.Second, //TLS握手超时时间
	ExpectContinueTimeout: time.Second,      //
}

func NewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	// 重写请求的URL
	director := func(req *http.Request) {
		rewriteRequestURL(req, target)
	}

	return &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifyResponse,
	}
}

func rewriteRequestURL(req *http.Request, target *url.URL) {
	targetQuery := target.RawQuery
	req.URL.Scheme = target.Scheme                               //保存目标 URL 的查询参数部分
	req.URL.Host = target.Host                                   //将请求的 URL 的协议和主机部分设置为目标 URL 的协议和主机部分
	req.URL.Path, req.URL.RawPath = joinURLPath(target, req.URL) //将请求的 URL 的路径和原始路径设置为目标 URL 路径和请求的 URL 路径的拼接结果。
	if targetQuery == "" || req.URL.RawQuery == "" {
		req.URL.RawQuery = targetQuery + req.URL.RawQuery
	} else {
		req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
	}
}

// joinURLPath 将两个URL路径进行拼接
// 如果两个URL的RawPath都为空，则直接通过singleJoiningSlash方法拼接路径，并返回结果
// 否则根据是否以"/"结尾来判断是否需要添加"/"来拼接路径，保证拼接后的路径格式正确
//
// 假设用户在浏览器中输入：http://127.0.0.1:8081/realserver，这个请求将被发送到您的代理服务器。
// 首先，用户在浏览器输入的地址被解析成一个 http.Request 对象，其中包含了用户请求的 URL 信息。
// 接着，我们将代理服务器的地址解析成一个 target *url.URL 对象，目标 URL 的路径是 "/?a=1&b=2#af"。
// 当用户请求到达代理服务器时，rewriteRequestURL 函数会被调用，其中会用到 joinURLPath 函数来重新构造请求的目标地址。
// 最终会被构造为http://127.0.0.1:8001/realserver/?a=1&b=2#af
// 从而确保最终的请求 URL 是正确的，并且能够正确地代理用户的请求到目标服务器上。
func joinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}

// singleJoiningSlash 将两个路径进行拼接
// 根据路径a和b的最后一个字符和第一个字符是否为'/'，来确定是否需要添加'/'
func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func modifyResponse(res *http.Response) error {
	log.Println("Start modify response")
	if res.StatusCode == http.StatusOK {
		srcBody, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		newBody := []byte(string(srcBody) + " modify response demo")
		res.Body = io.NopCloser(bytes.NewBuffer(newBody))

		// 注意修改了HTTP的Body后要同时修改表示内容长度的 Content-Length 头
		length := int64(len(newBody))
		res.ContentLength = length
		res.Header.Set("Content-Length", strconv.FormatInt(length, 10))
	}
	return nil
}
