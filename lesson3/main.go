package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"runtime"
)

//9.25课后作业
//内容：编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把1都做完
//1.接收客户端 request，并将 request 中带的 header 写入 response header
//2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
//3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
//4.当访问 localhost/healthz 时，应返回200
//提交链接🔗：https://jinshuju.net/f/PlZ3xg

func main() {
	// 创建一个bar路由和处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//1.接收客户端 request，并将 request 中带的 header 写入 response header
		header := w.Header()
		requestHeader := r.Header
		for key, value := range requestHeader {
			//fmt.Printf("HTTP header为 %s : %s\n", key, value)
			header.Add(key, value[0])
		}
		//2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
		header.Add("VERSION", runtime.Version())
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		w.WriteHeader(200)
		//3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
		fmt.Println("客户端IP地址为：" +  r.RemoteAddr)
		fmt.Println("HTTP返回码为：200")
	})

	//4.当访问 localhost/healthz 时，应返回200
	// 创建一个healthz路由和处理函数
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		//io.WriteString(w, "ok\n")
		w.WriteHeader(200)
		fmt.Fprintf(w, "healthz: OK")
	})
	// 监听8080端口
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("服务启动成功，监听端口8080")
}