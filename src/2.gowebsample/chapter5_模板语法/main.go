package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("解析模板异常err= %v", err)
		return

	}

	// 渲染模板

	u1 := User{
		Name:   "李茶茶",
		Gender: "女",
		Age:    18,
	}

	m1 := map[string]interface{}{
		"name":   "张爱国",
		"gender": "男",
		"age":    22,
	}
	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}
	data := map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("渲染模版异常err= %v", err)
		return
	}

}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server  err= %v", err)
		return
	}
}
