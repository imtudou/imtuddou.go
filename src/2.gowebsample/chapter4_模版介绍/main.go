package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {

	// 1. 解析模板
	t, err := template.ParseFiles("./hello.tmpl")

	if err != nil {
		fmt.Println("解析模板异常err= %v", err)
		return
	}

	// 2. 渲染模版
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println("渲染模版异常err= %v", err)
		return
	}

}

func main() {
	fmt.Println("http server start on:http://localhost:9090/")
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server err= %v", err)
		return
	}

}
