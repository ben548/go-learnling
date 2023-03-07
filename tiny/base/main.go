package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func newProxy(targetHost string) (*httputil.ReverseProxy, error) {
	u, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	return proxy, nil
}

func proxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func main() {
	newProxy, err := newProxy("http://127.0.0.1:9090")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", proxyRequestHandler(newProxy))
	log.Fatal(http.ListenAndServe(":9091", nil))
}
