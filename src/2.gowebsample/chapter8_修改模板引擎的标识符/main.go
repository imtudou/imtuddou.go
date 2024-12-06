package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	// 渲染模板
	name := "李茶茶是小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("execute template failed, err:%v\n", err)
		return
	}
}

func xxx(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	// 解析模板之前定义一个自定义的函数safe
	t, err := template.New("xxx.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xxx.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	// 渲染模板
	str1 := "<script>alert(123);</script>"
	str2 := "<a href='https://www.baidu.com'>百度一下</a>"
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xxx", xxx)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err:%v", err)
		return
	}
}
