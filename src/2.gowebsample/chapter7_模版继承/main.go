package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 渲染模板
	data := "李茶茶 index"
	t.Execute(w, data)
}

func home(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	data := "李茶茶 home"

	//渲染模板
	t.Execute(w, data)
}

func index2(w http.ResponseWriter, r *http.Request) {
	// 解析模板（模板继承的方式）
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 渲染模板
	data := "李茶茶 index2"
	t.Execute(w, data)
}

func home2(w http.ResponseWriter, r *http.Request) {
	//解析模板 （模板继承的方式）
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	data := "李茶茶 home2"

	//渲染模板
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
