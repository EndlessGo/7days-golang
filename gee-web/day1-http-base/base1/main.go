package main

// $ curl http://localhost:9999/
// URL.Path = "/"
// $ curl http://localhost:9999/hello
// Header["Accept"] = ["*/*"]
// Header["User-Agent"] = ["curl/7.54.0"]

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 设置两个静态路由，根据不同的HTTP请求调用不同的处理函数
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	//启动 Web 服务
	// 第一个参数是地址，:9999表示在 9999 端口监听。
	// 第二个参数则代表处理所有的HTTP请求的实例，nil 代表使用标准库中的实例处理。第二个参数，则是我们基于net/http标准库实现Web框架的入口。
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
