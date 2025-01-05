package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type IndexData struct {
	Title string    `json:"title"`
	Desc  string    `json:"desc"`
	Time  time.Time `json:"time"`
}

func indexData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	indexData := IndexData{
		Title: "Hello World",
		Desc:  "Hello World",
		Time:  time.Now(),
	}
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}

func indexHtml(w http.ResponseWriter, r *http.Request) {

	indexData := IndexData{
		Title: "Hello World",
		Desc:  "Hello World",
		Time:  time.Now(),
	}

	t := template.New("index.html")

	// 拿到当前路径
	path, _ := os.Getwd()
	t.ParseFiles("template/index.html", path)
	t.Execute(w, indexData)
}

func main() {

	//1. 程序入口,一个程序只有一个入口
	//2.构建一个server
	server := http.Server{
		Addr: ":9090",
	}
	http.HandleFunc("/", indexData)
	http.HandleFunc("/index.html", indexHtml)
	log.Printf("http://localhost:9090/  start %v", time.Now().Format("2006-01-02 15:04:05"))
	if err := server.ListenAndServe(); err != nil {
		log.Panicln(err)
	}

}
