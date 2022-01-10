package main

/***
1.接收客户端 request，并将 request 中带的 header 写入 response header
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4.当访问 localhost/healthz 时，应返回 200


*/

import (
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	glog.V(2).Info("Starting http server...")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	os.Setenv("VERSION", "v0.0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("get os version: %s \n", version)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	clientip := getClientip(r)
	fmt.Printf("clientip: \n" + clientip)
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Printf("Header key:% s, Header value: %s \n", k, v)
			w.Header().Set(k, vv)
		}
	}
}

func getClientip(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, "：")[0]
	}
	return ip
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}
