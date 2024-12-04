package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("http://localhost:8080/")

	// 1. 创建处理函数处理请求
	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)

	// 2. 使用处理器处理请求
	//customerHandler := CustmerHandler{}
	//http.Handle("/", &customerHandler)
	//http.ListenAndServe(":8080", nil)

	// 3.通过 Server 结构对服务器进行更详细的配置
	//customerHandler2 := CustmerHandler2{}
	//server := http.Server{
	//	Addr:        ":8080",
	//	Handler:     &customerHandler2,
	//	ReadTimeout: 10 * time.Second,
	//}
	//server.ListenAndServe()

	// 4. 使用自己创建的多路复用器

	mux := http.NewServeMux()
	mux.HandleFunc("/", serveMux)
	http.ListenAndServe(":8080", mux)
}

// 1. 创建处理函数处理请求
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "创建处理函数处理请求: hello world %v", time.Now().Format("2006-01-02 15:04:05"))
}

// 2. 使用处理器处理请求
type CustmerHandler struct{}

func (h *CustmerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "使用处理器处理请求:  hello world %v", time.Now().Format("2006-01-02 15:04:05"))
}

// 3.通过 Server 结构对服务器进行更详细的配置
type CustmerHandler2 struct{}

func (h *CustmerHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "通过 Server 结构对服务器进行更详细的配置:  hello world %v", time.Now().Format("2006-01-02 15:04:05"))
}

type ServeMux struct {
}

// 4. 使用自己创建的多路复用器
func serveMux(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "使用自己创建的多路复用器:  hello world %v", time.Now().Format("2006-01-02 15:04:05"))
}
