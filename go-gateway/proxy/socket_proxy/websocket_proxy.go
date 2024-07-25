package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	proxyAddr = "127.0.0.1:2002"
	serverURL = "http://127.0.0.1:2003"
)

func main() {
	URL, err := url.Parse(serverURL)
	if err != nil {
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(URL)
	log.Println("Starting HTTP server at " + proxyAddr)
	log.Fatal(http.ListenAndServe(proxyAddr, proxy))
}
