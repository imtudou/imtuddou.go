package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义一个函数 要么一个返回值 要么两个返回值 第二个返回值必须是error类型
	f := func(name string) (string, error) {
		return name + "帅气逼人", nil
	}
	// 定义模板
	// 创建一个名字是f的模板对象， 名字一定要与模板的名字一样
	t := template.New("f.tmpl")
	// 告诉模板引擎 我现在多了一个自定义函数 customfunc99
	t.Funcs(template.FuncMap{
		"customfunc99": f,
	})

	//解析模板
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Println("解析模板异常err=%v\n", err)
		return
	}
	name := "李茶茶"
	// 渲染模板
	t.Execute(w, name)
}

func f2(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("解析模板异常err=%v\n", err)
		return
	}
	name := "李茶茶"
	// 渲染模板
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/f2", f2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server start err=%v\n", err)
		return
	}
}
